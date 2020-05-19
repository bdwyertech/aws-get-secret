package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/mkideal/cli"
)

var GitCommit, ReleaseVer, ReleaseDate string

func showVersion() {
	if GitCommit == "" {
		GitCommit = "DEVELOPMENT"
	}
	if ReleaseVer == "" {
		ReleaseVer = "DEVELOPMENT"
	}
	fmt.Println("version:", ReleaseVer)
	fmt.Println("commit:", GitCommit)
	fmt.Println("date:", ReleaseDate)
}

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
	sess_opts := session.Options{
		Config: *aws.NewConfig().WithRegion(region).
			WithCredentialsChainVerboseErrors(true),
	}
	if argv.Profile != "" {
		sess_opts.Profile = argv.Profile
	}
	sess := session.Must(session.NewSessionWithOptions(sess_opts))

	var creds *credentials.Credentials
	if argv.RoleARN != "" {
		// Get AssumeRole Credentials
		creds = stscreds.NewCredentials(sess, argv.RoleARN, func(p *stscreds.AssumeRoleProvider) {
			if argv.RoleSessionName != "" {
				p.RoleSessionName = argv.RoleSessionName
			} else {
				p.RoleSessionName = fmt.Sprintf("aws-get-secret-%d", time.Now().UTC().UnixNano())
			}
			if argv.ExternalID != "" {
				p.ExternalID = aws.String(argv.ExternalID)
			}
		})
	} else {
		creds = sess.Config.Credentials
	}

	// SSM Client
	ssmcfg := &aws.Config{
		Credentials: creds,
	}
	ssmclient := ssm.New(sess, ssmcfg)
	resp, err := ssmclient.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(prm),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatalf("ERROR: ssm.GetParameter:: %s\n%s", prm, err)
	}
	val := *resp.Parameter.Value

	fmt.Println(val)

	return
}
