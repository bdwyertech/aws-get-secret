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

// Lists the executions of a maintenance window. This includes information about
// when the maintenance window was scheduled to be active, and information about
// tasks registered and run with the maintenance window.
func (c *Client) DescribeMaintenanceWindowExecutions(ctx context.Context, params *DescribeMaintenanceWindowExecutionsInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowExecutionsOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceWindowExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceWindowExecutions", params, optFns, c.addOperationDescribeMaintenanceWindowExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceWindowExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowExecutionsInput struct {

	// The ID of the maintenance window whose executions should be retrieved.
	//
	// This member is required.
	WindowId *string

	// Each entry in the array is a structure containing:
	//
	// * Key. A string between 1
	// and 128 characters. Supported keys include ExecutedBefore and ExecutedAfter.
	//
	// *
	// Values. An array of strings, each between 1 and 256 characters. Supported values
	// are date/time strings in a valid ISO 8601 date/time format, such as
	// 2021-11-04T05:00:00Z.
	Filters []types.MaintenanceWindowFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeMaintenanceWindowExecutionsOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the maintenance window executions.
	WindowExecutions []types.MaintenanceWindowExecution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMaintenanceWindowExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowExecutions{}, middleware.After)
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
	if err = addOpDescribeMaintenanceWindowExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowExecutions(options.Region), middleware.Before); err != nil {
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

// DescribeMaintenanceWindowExecutionsAPIClient is a client that implements the
// DescribeMaintenanceWindowExecutions operation.
type DescribeMaintenanceWindowExecutionsAPIClient interface {
	DescribeMaintenanceWindowExecutions(context.Context, *DescribeMaintenanceWindowExecutionsInput, ...func(*Options)) (*DescribeMaintenanceWindowExecutionsOutput, error)
}

var _ DescribeMaintenanceWindowExecutionsAPIClient = (*Client)(nil)

// DescribeMaintenanceWindowExecutionsPaginatorOptions is the paginator options for
// DescribeMaintenanceWindowExecutions
type DescribeMaintenanceWindowExecutionsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMaintenanceWindowExecutionsPaginator is a paginator for
// DescribeMaintenanceWindowExecutions
type DescribeMaintenanceWindowExecutionsPaginator struct {
	options   DescribeMaintenanceWindowExecutionsPaginatorOptions
	client    DescribeMaintenanceWindowExecutionsAPIClient
	params    *DescribeMaintenanceWindowExecutionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeMaintenanceWindowExecutionsPaginator returns a new
// DescribeMaintenanceWindowExecutionsPaginator
func NewDescribeMaintenanceWindowExecutionsPaginator(client DescribeMaintenanceWindowExecutionsAPIClient, params *DescribeMaintenanceWindowExecutionsInput, optFns ...func(*DescribeMaintenanceWindowExecutionsPaginatorOptions)) *DescribeMaintenanceWindowExecutionsPaginator {
	if params == nil {
		params = &DescribeMaintenanceWindowExecutionsInput{}
	}

	options := DescribeMaintenanceWindowExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMaintenanceWindowExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMaintenanceWindowExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMaintenanceWindowExecutions page.
func (p *DescribeMaintenanceWindowExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMaintenanceWindowExecutionsOutput, error) {
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

	result, err := p.client.DescribeMaintenanceWindowExecutions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowExecutions",
	}
}
