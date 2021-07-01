// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the association between an OpsItem and a related resource. For example,
// this API action can delete an Incident Manager incident from an OpsItem.
// Incident Manager is a capability of AWS Systems Manager.
func (c *Client) DisassociateOpsItemRelatedItem(ctx context.Context, params *DisassociateOpsItemRelatedItemInput, optFns ...func(*Options)) (*DisassociateOpsItemRelatedItemOutput, error) {
	if params == nil {
		params = &DisassociateOpsItemRelatedItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateOpsItemRelatedItem", params, optFns, c.addOperationDisassociateOpsItemRelatedItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateOpsItemRelatedItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateOpsItemRelatedItemInput struct {

	// The ID of the association for which you want to delete an association between
	// the OpsItem and a related resource.
	//
	// This member is required.
	AssociationId *string

	// The ID of the OpsItem for which you want to delete an association between the
	// OpsItem and a related resource.
	//
	// This member is required.
	OpsItemId *string
}

type DisassociateOpsItemRelatedItemOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDisassociateOpsItemRelatedItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateOpsItemRelatedItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateOpsItemRelatedItem{}, middleware.After)
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
	if err = addOpDisassociateOpsItemRelatedItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateOpsItemRelatedItem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateOpsItemRelatedItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DisassociateOpsItemRelatedItem",
	}
}
