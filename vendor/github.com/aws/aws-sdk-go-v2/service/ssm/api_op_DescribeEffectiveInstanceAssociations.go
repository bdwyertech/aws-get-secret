// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// All associations for the instance(s).
func (c *Client) DescribeEffectiveInstanceAssociations(ctx context.Context, params *DescribeEffectiveInstanceAssociationsInput, optFns ...func(*Options)) (*DescribeEffectiveInstanceAssociationsOutput, error) {
	if params == nil {
		params = &DescribeEffectiveInstanceAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEffectiveInstanceAssociations", params, optFns, c.addOperationDescribeEffectiveInstanceAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEffectiveInstanceAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEffectiveInstanceAssociationsInput struct {

	// The instance ID for which you want to view all associations.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeEffectiveInstanceAssociationsOutput struct {

	// The associations for the requested instance.
	Associations []types.InstanceAssociation

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEffectiveInstanceAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEffectiveInstanceAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEffectiveInstanceAssociations{}, middleware.After)
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
	if err = addOpDescribeEffectiveInstanceAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEffectiveInstanceAssociations(options.Region), middleware.Before); err != nil {
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

// DescribeEffectiveInstanceAssociationsAPIClient is a client that implements the
// DescribeEffectiveInstanceAssociations operation.
type DescribeEffectiveInstanceAssociationsAPIClient interface {
	DescribeEffectiveInstanceAssociations(context.Context, *DescribeEffectiveInstanceAssociationsInput, ...func(*Options)) (*DescribeEffectiveInstanceAssociationsOutput, error)
}

var _ DescribeEffectiveInstanceAssociationsAPIClient = (*Client)(nil)

// DescribeEffectiveInstanceAssociationsPaginatorOptions is the paginator options
// for DescribeEffectiveInstanceAssociations
type DescribeEffectiveInstanceAssociationsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEffectiveInstanceAssociationsPaginator is a paginator for
// DescribeEffectiveInstanceAssociations
type DescribeEffectiveInstanceAssociationsPaginator struct {
	options   DescribeEffectiveInstanceAssociationsPaginatorOptions
	client    DescribeEffectiveInstanceAssociationsAPIClient
	params    *DescribeEffectiveInstanceAssociationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEffectiveInstanceAssociationsPaginator returns a new
// DescribeEffectiveInstanceAssociationsPaginator
func NewDescribeEffectiveInstanceAssociationsPaginator(client DescribeEffectiveInstanceAssociationsAPIClient, params *DescribeEffectiveInstanceAssociationsInput, optFns ...func(*DescribeEffectiveInstanceAssociationsPaginatorOptions)) *DescribeEffectiveInstanceAssociationsPaginator {
	if params == nil {
		params = &DescribeEffectiveInstanceAssociationsInput{}
	}

	options := DescribeEffectiveInstanceAssociationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEffectiveInstanceAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEffectiveInstanceAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeEffectiveInstanceAssociations page.
func (p *DescribeEffectiveInstanceAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEffectiveInstanceAssociationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeEffectiveInstanceAssociations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeEffectiveInstanceAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeEffectiveInstanceAssociations",
	}
}
