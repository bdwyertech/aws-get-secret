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

// Provides details about all active and terminated Automation executions.
func (c *Client) DescribeAutomationExecutions(ctx context.Context, params *DescribeAutomationExecutionsInput, optFns ...func(*Options)) (*DescribeAutomationExecutionsOutput, error) {
	if params == nil {
		params = &DescribeAutomationExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAutomationExecutions", params, optFns, addOperationDescribeAutomationExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAutomationExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutomationExecutionsInput struct {

	// Filters used to limit the scope of executions that are requested.
	Filters []types.AutomationExecutionFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeAutomationExecutionsOutput struct {

	// The list of details about each automation execution which has occurred which
	// matches the filter specification, if any.
	AutomationExecutionMetadataList []types.AutomationExecutionMetadata

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAutomationExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAutomationExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAutomationExecutions{}, middleware.After)
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
	if err = addOpDescribeAutomationExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutomationExecutions(options.Region), middleware.Before); err != nil {
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

// DescribeAutomationExecutionsAPIClient is a client that implements the
// DescribeAutomationExecutions operation.
type DescribeAutomationExecutionsAPIClient interface {
	DescribeAutomationExecutions(context.Context, *DescribeAutomationExecutionsInput, ...func(*Options)) (*DescribeAutomationExecutionsOutput, error)
}

var _ DescribeAutomationExecutionsAPIClient = (*Client)(nil)

// DescribeAutomationExecutionsPaginatorOptions is the paginator options for
// DescribeAutomationExecutions
type DescribeAutomationExecutionsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAutomationExecutionsPaginator is a paginator for
// DescribeAutomationExecutions
type DescribeAutomationExecutionsPaginator struct {
	options   DescribeAutomationExecutionsPaginatorOptions
	client    DescribeAutomationExecutionsAPIClient
	params    *DescribeAutomationExecutionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAutomationExecutionsPaginator returns a new
// DescribeAutomationExecutionsPaginator
func NewDescribeAutomationExecutionsPaginator(client DescribeAutomationExecutionsAPIClient, params *DescribeAutomationExecutionsInput, optFns ...func(*DescribeAutomationExecutionsPaginatorOptions)) *DescribeAutomationExecutionsPaginator {
	if params == nil {
		params = &DescribeAutomationExecutionsInput{}
	}

	options := DescribeAutomationExecutionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAutomationExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAutomationExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeAutomationExecutions page.
func (p *DescribeAutomationExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAutomationExecutionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeAutomationExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAutomationExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeAutomationExecutions",
	}
}
