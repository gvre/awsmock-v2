package dynamodbmock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/mock"
)

// Client provides the API client to mock operations call for Amazon Simple Queue Service.
type Client struct {
	mock.Mock
}

// BatchExecuteStatement mock.
func (c *Client) BatchExecuteStatement(ctx context.Context, params *dynamodb.BatchExecuteStatementInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchExecuteStatementOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.BatchExecuteStatementOutput), args.Error(1)
}

// BatchGetItem mock.
func (c *Client) BatchGetItem(ctx context.Context, params *dynamodb.BatchGetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.BatchGetItemOutput), args.Error(1)
}

// BatchWriteItem mock.
func (c *Client) BatchWriteItem(ctx context.Context, params *dynamodb.BatchWriteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.BatchWriteItemOutput), args.Error(1)
}

// CreateBackup mock.
func (c *Client) CreateBackup(ctx context.Context, params *dynamodb.CreateBackupInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateBackupOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.CreateBackupOutput), args.Error(1)
}

// CreateGlobalTable mock.
func (c *Client) CreateGlobalTable(ctx context.Context, params *dynamodb.CreateGlobalTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateGlobalTableOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.CreateGlobalTableOutput), args.Error(1)
}

// CreateTable mock.
func (c *Client) CreateTable(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.CreateTableOutput), args.Error(1)
}

// DeleteBackup mock.
func (c *Client) DeleteBackup(ctx context.Context, params *dynamodb.DeleteBackupInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteBackupOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DeleteBackupOutput), args.Error(1)
}

// DeleteItem mock.
func (c *Client) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DeleteItemOutput), args.Error(1)
}

// DeleteTable mock.
func (c *Client) DeleteTable(ctx context.Context, params *dynamodb.DeleteTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DeleteTableOutput), args.Error(1)
}

// DescribeBackup mock.
func (c *Client) DescribeBackup(ctx context.Context, params *dynamodb.DescribeBackupInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeBackupOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeBackupOutput), args.Error(1)
}

// DescribeContinuousBackups mock.
func (c *Client) DescribeContinuousBackups(ctx context.Context, params *dynamodb.DescribeContinuousBackupsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeContinuousBackupsOutput), args.Error(1)
}

// DescribeContributorInsights mock.
func (c *Client) DescribeContributorInsights(ctx context.Context, params *dynamodb.DescribeContributorInsightsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeContributorInsightsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeContributorInsightsOutput), args.Error(1)
}

// DescribeEndpoints mock.
func (c *Client) DescribeEndpoints(ctx context.Context, params *dynamodb.DescribeEndpointsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeEndpointsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeEndpointsOutput), args.Error(1)
}

// DescribeExport mock.
func (c *Client) DescribeExport(ctx context.Context, params *dynamodb.DescribeExportInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeExportOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeExportOutput), args.Error(1)
}

// DescribeGlobalTable mock.
func (c *Client) DescribeGlobalTable(ctx context.Context, params *dynamodb.DescribeGlobalTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeGlobalTableOutput), args.Error(1)
}

// DescribeGlobalTableSettings mock.
func (c *Client) DescribeGlobalTableSettings(ctx context.Context, params *dynamodb.DescribeGlobalTableSettingsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeGlobalTableSettingsOutput), args.Error(1)
}

// DescribeKinesisStreamingDestination mock.
func (c *Client) DescribeKinesisStreamingDestination(ctx context.Context, params *dynamodb.DescribeKinesisStreamingDestinationInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeKinesisStreamingDestinationOutput), args.Error(1)
}

// DescribeLimits mock.
func (c *Client) DescribeLimits(ctx context.Context, params *dynamodb.DescribeLimitsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeLimitsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeLimitsOutput), args.Error(1)
}

// DescribeTable mock.
func (c *Client) DescribeTable(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeTableOutput), args.Error(1)
}

// DescribeTableReplicaAutoScaling mock.
func (c *Client) DescribeTableReplicaAutoScaling(ctx context.Context, params *dynamodb.DescribeTableReplicaAutoScalingInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeTableReplicaAutoScalingOutput), args.Error(1)
}

// DescribeTimeToLive mock.
func (c *Client) DescribeTimeToLive(ctx context.Context, params *dynamodb.DescribeTimeToLiveInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTimeToLiveOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DescribeTimeToLiveOutput), args.Error(1)
}

// DisableKinesisStreamingDestination mock.
func (c *Client) DisableKinesisStreamingDestination(ctx context.Context, params *dynamodb.DisableKinesisStreamingDestinationInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DisableKinesisStreamingDestinationOutput), args.Error(1)
}

// EnableKinesisStreamingDestination mock.
func (c *Client) EnableKinesisStreamingDestination(ctx context.Context, params *dynamodb.EnableKinesisStreamingDestinationInput, optFns ...func(*dynamodb.Options)) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.EnableKinesisStreamingDestinationOutput), args.Error(1)
}

// ExecuteStatement mock.
func (c *Client) ExecuteStatement(ctx context.Context, params *dynamodb.ExecuteStatementInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.ExecuteStatementOutput), args.Error(1)
}

// ExecuteTransaction mock.
func (c *Client) ExecuteTransaction(ctx context.Context, params *dynamodb.ExecuteTransactionInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExecuteTransactionOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.ExecuteTransactionOutput), args.Error(1)
}

// ExportTableToPointInTime mock.
func (c *Client) ExportTableToPointInTime(ctx context.Context, params *dynamodb.ExportTableToPointInTimeInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.ExportTableToPointInTimeOutput), args.Error(1)
}

// GetItem mock.
func (c *Client) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.GetItemOutput), args.Error(1)
}

// ListBackups mock.
func (c *Client) ListBackups(ctx context.Context, params *dynamodb.ListBackupsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListBackupsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.ListBackupsOutput), args.Error(1)
}

// ListContributorInsights mock.
func (c *Client) ListContributorInsights(ctx context.Context, params *dynamodb.ListContributorInsightsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListContributorInsightsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.ListContributorInsightsOutput), args.Error(1)
}

// ListExports mock.
func (c *Client) ListExports(ctx context.Context, params *dynamodb.ListExportsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListExportsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.ListExportsOutput), args.Error(1)
}

// ListGlobalTables mock.
func (c *Client) ListGlobalTables(ctx context.Context, params *dynamodb.ListGlobalTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListGlobalTablesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.ListGlobalTablesOutput), args.Error(1)
}

// ListTables mock.
func (c *Client) ListTables(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.ListTablesOutput), args.Error(1)
}

// ListTagsOfResource mock.
func (c *Client) ListTagsOfResource(ctx context.Context, params *dynamodb.ListTagsOfResourceInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTagsOfResourceOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.ListTagsOfResourceOutput), args.Error(1)
}

// PutItem mock.
func (c *Client) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
}

// Query mock.
func (c *Client) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.QueryOutput), args.Error(1)
}

// RestoreTableFromBackup mock.
func (c *Client) RestoreTableFromBackup(ctx context.Context, params *dynamodb.RestoreTableFromBackupInput, optFns ...func(*dynamodb.Options)) (*dynamodb.RestoreTableFromBackupOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.RestoreTableFromBackupOutput), args.Error(1)
}

// RestoreTableToPointInTime mock.
func (c *Client) RestoreTableToPointInTime(ctx context.Context, params *dynamodb.RestoreTableToPointInTimeInput, optFns ...func(*dynamodb.Options)) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.RestoreTableToPointInTimeOutput), args.Error(1)
}

// Scan mock.
func (c *Client) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.ScanOutput), args.Error(1)
}

// TagResource mock.
func (c *Client) TagResource(ctx context.Context, params *dynamodb.TagResourceInput, optFns ...func(*dynamodb.Options)) (*dynamodb.TagResourceOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.TagResourceOutput), args.Error(1)
}

// TransactGetItems mock.
func (c *Client) TransactGetItems(ctx context.Context, params *dynamodb.TransactGetItemsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.TransactGetItemsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.TransactGetItemsOutput), args.Error(1)
}

// TransactWriteItems mock.
func (c *Client) TransactWriteItems(ctx context.Context, params *dynamodb.TransactWriteItemsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.TransactWriteItemsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.TransactWriteItemsOutput), args.Error(1)
}

// UntagResource mock.
func (c *Client) UntagResource(ctx context.Context, params *dynamodb.UntagResourceInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UntagResourceOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.UntagResourceOutput), args.Error(1)
}

// UpdateContinuousBackups mock.
func (c *Client) UpdateContinuousBackups(ctx context.Context, params *dynamodb.UpdateContinuousBackupsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.UpdateContinuousBackupsOutput), args.Error(1)
}

// UpdateContributorInsights mock.
func (c *Client) UpdateContributorInsights(ctx context.Context, params *dynamodb.UpdateContributorInsightsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateContributorInsightsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.UpdateContributorInsightsOutput), args.Error(1)
}

// UpdateGlobalTable mock.
func (c *Client) UpdateGlobalTable(ctx context.Context, params *dynamodb.UpdateGlobalTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateGlobalTableOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.UpdateGlobalTableOutput), args.Error(1)
}

// UpdateGlobalTableSettings mock.
func (c *Client) UpdateGlobalTableSettings(ctx context.Context, params *dynamodb.UpdateGlobalTableSettingsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.UpdateGlobalTableSettingsOutput), args.Error(1)
}

// UpdateItem mock.
func (c *Client) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.UpdateItemOutput), args.Error(1)
}

// UpdateTable mock.
func (c *Client) UpdateTable(ctx context.Context, params *dynamodb.UpdateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateTableOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.UpdateTableOutput), args.Error(1)
}

// UpdateTableReplicaAutoScaling mock.
func (c *Client) UpdateTableReplicaAutoScaling(ctx context.Context, params *dynamodb.UpdateTableReplicaAutoScalingInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.UpdateTableReplicaAutoScalingOutput), args.Error(1)
}

// UpdateTimeToLive mock.
func (c *Client) UpdateTimeToLive(ctx context.Context, params *dynamodb.UpdateTimeToLiveInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateTimeToLiveOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.UpdateTimeToLiveOutput), args.Error(1)
}
