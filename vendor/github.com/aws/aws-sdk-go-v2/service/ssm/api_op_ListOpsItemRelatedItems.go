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

// Lists all related-item resources associated with an OpsItem.
func (c *Client) ListOpsItemRelatedItems(ctx context.Context, params *ListOpsItemRelatedItemsInput, optFns ...func(*Options)) (*ListOpsItemRelatedItemsOutput, error) {
	if params == nil {
		params = &ListOpsItemRelatedItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOpsItemRelatedItems", params, optFns, c.addOperationListOpsItemRelatedItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOpsItemRelatedItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOpsItemRelatedItemsInput struct {

	// One or more OpsItem filters. Use a filter to return a more specific list of
	// results.
	Filters []types.OpsItemRelatedItemsFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The ID of the OpsItem for which you want to list all related-item resources.
	OpsItemId *string
}

type ListOpsItemRelatedItemsOutput struct {

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// A list of related-item resources for the specified OpsItem.
	Summaries []types.OpsItemRelatedItemSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListOpsItemRelatedItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOpsItemRelatedItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOpsItemRelatedItems{}, middleware.After)
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
	if err = addOpListOpsItemRelatedItemsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOpsItemRelatedItems(options.Region), middleware.Before); err != nil {
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

// ListOpsItemRelatedItemsAPIClient is a client that implements the
// ListOpsItemRelatedItems operation.
type ListOpsItemRelatedItemsAPIClient interface {
	ListOpsItemRelatedItems(context.Context, *ListOpsItemRelatedItemsInput, ...func(*Options)) (*ListOpsItemRelatedItemsOutput, error)
}

var _ ListOpsItemRelatedItemsAPIClient = (*Client)(nil)

// ListOpsItemRelatedItemsPaginatorOptions is the paginator options for
// ListOpsItemRelatedItems
type ListOpsItemRelatedItemsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOpsItemRelatedItemsPaginator is a paginator for ListOpsItemRelatedItems
type ListOpsItemRelatedItemsPaginator struct {
	options   ListOpsItemRelatedItemsPaginatorOptions
	client    ListOpsItemRelatedItemsAPIClient
	params    *ListOpsItemRelatedItemsInput
	nextToken *string
	firstPage bool
}

// NewListOpsItemRelatedItemsPaginator returns a new
// ListOpsItemRelatedItemsPaginator
func NewListOpsItemRelatedItemsPaginator(client ListOpsItemRelatedItemsAPIClient, params *ListOpsItemRelatedItemsInput, optFns ...func(*ListOpsItemRelatedItemsPaginatorOptions)) *ListOpsItemRelatedItemsPaginator {
	if params == nil {
		params = &ListOpsItemRelatedItemsInput{}
	}

	options := ListOpsItemRelatedItemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOpsItemRelatedItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOpsItemRelatedItemsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListOpsItemRelatedItems page.
func (p *ListOpsItemRelatedItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOpsItemRelatedItemsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListOpsItemRelatedItems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOpsItemRelatedItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListOpsItemRelatedItems",
	}
}
