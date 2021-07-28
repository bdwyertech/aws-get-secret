// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A parameter label is a user-defined alias to help you manage different versions
// of a parameter. When you modify a parameter, Amazon Web Services Systems Manager
// automatically saves a new version and increments the version number by one. A
// label can help you remember the purpose of a parameter when there are multiple
// versions. Parameter labels have the following requirements and restrictions.
//
// *
// A version of a parameter can have a maximum of 10 labels.
//
// * You can't attach
// the same label to different versions of the same parameter. For example, if
// version 1 has the label Production, then you can't attach Production to version
// 2.
//
// * You can move a label from one version of a parameter to another.
//
// * You
// can't create a label when you create a new parameter. You must attach a label to
// a specific version of a parameter.
//
// * If you no longer want to use a parameter
// label, then you can either delete it or move it to a different version of a
// parameter.
//
// * A label can have a maximum of 100 characters.
//
// * Labels can
// contain letters (case sensitive), numbers, periods (.), hyphens (-), or
// underscores (_).
//
// * Labels can't begin with a number, "aws" or "ssm" (not case
// sensitive). If a label fails to meet these requirements, then the label isn't
// associated with a parameter and the system displays it in the list of
// InvalidLabels.
func (c *Client) LabelParameterVersion(ctx context.Context, params *LabelParameterVersionInput, optFns ...func(*Options)) (*LabelParameterVersionOutput, error) {
	if params == nil {
		params = &LabelParameterVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "LabelParameterVersion", params, optFns, c.addOperationLabelParameterVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*LabelParameterVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type LabelParameterVersionInput struct {

	// One or more labels to attach to the specified parameter version.
	//
	// This member is required.
	Labels []string

	// The parameter name on which you want to attach one or more labels.
	//
	// This member is required.
	Name *string

	// The specific version of the parameter on which you want to attach one or more
	// labels. If no version is specified, the system attaches the label to the latest
	// version.
	ParameterVersion int64
}

type LabelParameterVersionOutput struct {

	// The label doesn't meet the requirements. For information about parameter label
	// requirements, see Labeling parameters
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-labels.html)
	// in the Amazon Web Services Systems Manager User Guide.
	InvalidLabels []string

	// The version of the parameter that has been labeled.
	ParameterVersion int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationLabelParameterVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpLabelParameterVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpLabelParameterVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpLabelParameterVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opLabelParameterVersion(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opLabelParameterVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "LabelParameterVersion",
	}
}