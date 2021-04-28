package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/mkideal/cli"
)

type ArgT struct {
	cli.Helper
	Parameter       string `cli:"p,parameter" usage:"The path to the parameter in AWS Parameter Store"`
	Profile         string `cli:"profile" usage:"AWS Profile" dft:"$AWS_PROFILE"`
	Region          string `cli:"region" usage:"AWS Region" dft:"$AWS_REGION"`
	RoleARN         string `cli:"r,role-arn" usage:"The Role ARN to assume to retrieve the parameter"`
	RoleSessionName string `cli:"session-name" usage:"The Session Name to use during the AssumeRole operation"`
	ExternalID      string `cli:"e,external-id" usage:"The External ID to use during the AssumeRole operation"`
	Version         bool   `cli:"v, version" usage:"Display the version of this utility"`
}

// Validate & Coalesce CLI Arguments
func (argv *ArgT) Validate(ctx *cli.Context) (err error) {

	if argv.Version {
		showVersion()
		os.Exit(0)
	}

	if argv.Parameter == "" {
		if len(ctx.Args()) > 0 {
			argv.Parameter = ctx.Args()[0]
		} else {
			err = fmt.Errorf("%s is a required argument", ctx.Color().Yellow("Parameter"))
			return
		}
	}

	if argv.Region == "" {
		argv.Region = "us-east-1"
	}

	return
}

func main() {
	os.Exit(cli.Run(new(ArgT), func(ctx *cli.Context) (err error) {
		err = Run(ctx)
		return
	}))
}

func Run(ctx *cli.Context) (err error) {
	argv := ctx.Argv().(*ArgT)
	// Marshal Parameter
	prm := argv.Parameter
	region := argv.Region
	switch len(strings.Split(prm, ":")) {
	case 6:
		prm_arn, err := arn.Parse(prm)
		if err != nil {
			log.Fatal(err)
		}
		if prm_arn.Region != "" {
			region = prm_arn.Region
		}
		resource := strings.Split(prm_arn.Resource, "parameter")
		if len(resource) == 2 {
			prm = resource[1]
		} else {
			return fmt.Errorf("ERROR: Bad Parameter Resource Path: %s", resource)
		}
	case 1:
		break
	default:
		return fmt.Errorf("ERROR: Invalid Parameter ARN or Path")
	}

	// AWS Session
	cfgOpts := []func(*config.LoadOptions) error{
		config.WithRegion(region),
	}
	if argv.Profile != "" {
		cfgOpts = append(cfgOpts, config.WithSharedConfigProfile(argv.Profile))
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), cfgOpts...)
	if err != nil {
		panic(err)
	}

	if argv.RoleARN != "" {
		// Get AssumeRole Credentials
		client := sts.NewFromConfig(cfg)

		creds := stscreds.NewAssumeRoleProvider(client, argv.RoleARN, func(p *stscreds.AssumeRoleOptions) {
			if argv.RoleSessionName != "" {
				p.RoleSessionName = argv.RoleSessionName
			} else {
				p.RoleSessionName = fmt.Sprintf("aws-get-secret-%d", time.Now().UTC().UnixNano())
			}
			if argv.ExternalID != "" {
				p.ExternalID = aws.String(argv.ExternalID)
			}
		})
		cfg.Credentials = aws.NewCredentialsCache(creds)
	}

	// SSM Client
	ssmclient := ssm.NewFromConfig(cfg)
	resp, err := ssmclient.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String(prm),
		WithDecryption: true,
	})
	if err != nil {
		log.Fatalf("ERROR: ssm.GetParameter:: %s\n%s", prm, err)
	}
	val := *resp.Parameter.Value

	fmt.Println(val)

	return
}
