package dynamodbmock_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/mock"

	"github.com/gvre/awsmock-v2/dynamodbmock"
)

type dynamodbClient interface {
	BatchExecuteStatement(ctx context.Context, params *dynamodb.BatchExecuteStatementInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchExecuteStatementOutput, error)
	BatchGetItem(ctx context.Context, params *dynamodb.BatchGetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error)
	BatchWriteItem(ctx context.Context, params *dynamodb.BatchWriteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error)
	CreateBackup(ctx context.Context, params *dynamodb.CreateBackupInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateBackupOutput, error)
	CreateGlobalTable(ctx context.Context, params *dynamodb.CreateGlobalTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateGlobalTableOutput, error)
	CreateTable(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)
	DeleteBackup(ctx context.Context, params *dynamodb.DeleteBackupInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteBackupOutput, error)
	DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
	DeleteTable(ctx context.Context, params *dynamodb.DeleteTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error)
	DescribeBackup(ctx context.Context, params *dynamodb.DescribeBackupInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeBackupOutput, error)
	DescribeContinuousBackups(ctx context.Context, params *dynamodb.DescribeContinuousBackupsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeContinuousBackupsOutput, error)
	DescribeContributorInsights(ctx context.Context, params *dynamodb.DescribeContributorInsightsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeContributorInsightsOutput, error)
	DescribeEndpoints(ctx context.Context, params *dynamodb.DescribeEndpointsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeEndpointsOutput, error)
	DescribeExport(ctx context.Context, params *dynamodb.DescribeExportInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeExportOutput, error)
	DescribeGlobalTable(ctx context.Context, params *dynamodb.DescribeGlobalTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableOutput, error)
	DescribeGlobalTableSettings(ctx context.Context, params *dynamodb.DescribeGlobalTableSettingsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
	DescribeKinesisStreamingDestination(ctx context.Context, params *dynamodb.DescribeKinesisStreamingDestinationInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error)
	DescribeLimits(ctx context.Context, params *dynamodb.DescribeLimitsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeLimitsOutput, error)
	DescribeTable(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)
	DescribeTableReplicaAutoScaling(ctx context.Context, params *dynamodb.DescribeTableReplicaAutoScalingInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
	DescribeTimeToLive(ctx context.Context, params *dynamodb.DescribeTimeToLiveInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTimeToLiveOutput, error)
	DisableKinesisStreamingDestination(ctx context.Context, params *dynamodb.DisableKinesisStreamingDestinationInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DisableKinesisStreamingDestinationOutput, error)
	EnableKinesisStreamingDestination(ctx context.Context, params *dynamodb.EnableKinesisStreamingDestinationInput, optFns ...func(*dynamodb.Options)) (*dynamodb.EnableKinesisStreamingDestinationOutput, error)
	ExecuteStatement(ctx context.Context, params *dynamodb.ExecuteStatementInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error)
	ExecuteTransaction(ctx context.Context, params *dynamodb.ExecuteTransactionInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExecuteTransactionOutput, error)
	ExportTableToPointInTime(ctx context.Context, params *dynamodb.ExportTableToPointInTimeInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExportTableToPointInTimeOutput, error)
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	ListBackups(ctx context.Context, params *dynamodb.ListBackupsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListBackupsOutput, error)
	ListContributorInsights(ctx context.Context, params *dynamodb.ListContributorInsightsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListContributorInsightsOutput, error)
	ListExports(ctx context.Context, params *dynamodb.ListExportsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListExportsOutput, error)
	ListGlobalTables(ctx context.Context, params *dynamodb.ListGlobalTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListGlobalTablesOutput, error)
	ListTables(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error)
	ListTagsOfResource(ctx context.Context, params *dynamodb.ListTagsOfResourceInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTagsOfResourceOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
	RestoreTableFromBackup(ctx context.Context, params *dynamodb.RestoreTableFromBackupInput, optFns ...func(*dynamodb.Options)) (*dynamodb.RestoreTableFromBackupOutput, error)
	RestoreTableToPointInTime(ctx context.Context, params *dynamodb.RestoreTableToPointInTimeInput, optFns ...func(*dynamodb.Options)) (*dynamodb.RestoreTableToPointInTimeOutput, error)
	Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
	TagResource(ctx context.Context, params *dynamodb.TagResourceInput, optFns ...func(*dynamodb.Options)) (*dynamodb.TagResourceOutput, error)
	TransactGetItems(ctx context.Context, params *dynamodb.TransactGetItemsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.TransactGetItemsOutput, error)
	TransactWriteItems(ctx context.Context, params *dynamodb.TransactWriteItemsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.TransactWriteItemsOutput, error)
	UntagResource(ctx context.Context, params *dynamodb.UntagResourceInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UntagResourceOutput, error)
	UpdateContinuousBackups(ctx context.Context, params *dynamodb.UpdateContinuousBackupsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateContinuousBackupsOutput, error)
	UpdateContributorInsights(ctx context.Context, params *dynamodb.UpdateContributorInsightsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateContributorInsightsOutput, error)
	UpdateGlobalTable(ctx context.Context, params *dynamodb.UpdateGlobalTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateGlobalTableOutput, error)
	UpdateGlobalTableSettings(ctx context.Context, params *dynamodb.UpdateGlobalTableSettingsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	UpdateTable(ctx context.Context, params *dynamodb.UpdateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateTableOutput, error)
	UpdateTableReplicaAutoScaling(ctx context.Context, params *dynamodb.UpdateTableReplicaAutoScalingInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error)
	UpdateTimeToLive(ctx context.Context, params *dynamodb.UpdateTimeToLiveInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateTimeToLiveOutput, error)
}

func TestBatchExecuteStatement(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("BatchExecuteStatement", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.BatchExecuteStatementOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.BatchExecuteStatement(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestBatchGetItem(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("BatchGetItem", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.BatchGetItemOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.BatchGetItem(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestBatchWriteItem(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("BatchWriteItem", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.BatchWriteItemOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.BatchWriteItem(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCreateBackup(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("CreateBackup", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.CreateBackupOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.CreateBackup(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCreateGlobalTable(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("CreateGlobalTable", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.CreateGlobalTableOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.CreateGlobalTable(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCreateTable(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("CreateTable", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.CreateTableOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.CreateTable(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBackup(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DeleteBackup", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DeleteBackupOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DeleteBackup(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteItem(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DeleteItem", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DeleteItemOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DeleteItem(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteTable(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DeleteTable", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DeleteTableOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DeleteTable(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeBackup(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeBackup", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeBackupOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeBackup(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeContinuousBackups(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeContinuousBackups", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeContinuousBackupsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeContinuousBackups(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeContributorInsights(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeContributorInsights", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeContributorInsightsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeContributorInsights(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeEndpoints(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeEndpoints", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeEndpointsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeEndpoints(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeExport(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeExport", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeExportOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeExport(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeGlobalTable(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeGlobalTable", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeGlobalTableOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeGlobalTable(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeGlobalTableSettings(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeGlobalTableSettings", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeGlobalTableSettingsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeGlobalTableSettings(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeKinesisStreamingDestination(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeKinesisStreamingDestination", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeKinesisStreamingDestinationOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeKinesisStreamingDestination(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeLimits(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeLimits", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeLimitsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeLimits(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeTable(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeTable", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeTableOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeTable(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeTableReplicaAutoScaling(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeTableReplicaAutoScaling", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeTableReplicaAutoScalingOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeTableReplicaAutoScaling(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDescribeTimeToLive(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DescribeTimeToLive", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DescribeTimeToLiveOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DescribeTimeToLive(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDisableKinesisStreamingDestination(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("DisableKinesisStreamingDestination", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.DisableKinesisStreamingDestinationOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.DisableKinesisStreamingDestination(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestEnableKinesisStreamingDestination(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("EnableKinesisStreamingDestination", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.EnableKinesisStreamingDestinationOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.EnableKinesisStreamingDestination(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestExecuteStatement(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("ExecuteStatement", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.ExecuteStatementOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.ExecuteStatement(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestExecuteTransaction(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("ExecuteTransaction", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.ExecuteTransactionOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.ExecuteTransaction(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestExportTableToPointInTime(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("ExportTableToPointInTime", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.ExportTableToPointInTimeOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.ExportTableToPointInTime(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetItem(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("GetItem", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.GetItemOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.GetItem(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListBackups(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("ListBackups", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.ListBackupsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.ListBackups(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListContributorInsights(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("ListContributorInsights", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.ListContributorInsightsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.ListContributorInsights(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListExports(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("ListExports", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.ListExportsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.ListExports(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListGlobalTables(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("ListGlobalTables", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.ListGlobalTablesOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.ListGlobalTables(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListTables(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("ListTables", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.ListTablesOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.ListTables(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListTagsOfResource(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("ListTagsOfResource", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.ListTagsOfResourceOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.ListTagsOfResource(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutItem(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("PutItem", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.PutItemOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.PutItem(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestQuery(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.QueryOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.Query(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestRestoreTableFromBackup(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("RestoreTableFromBackup", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.RestoreTableFromBackupOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.RestoreTableFromBackup(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestRestoreTableToPointInTime(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("RestoreTableToPointInTime", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.RestoreTableToPointInTimeOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.RestoreTableToPointInTime(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestScan(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("Scan", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.ScanOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.Scan(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestTagResource(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("TagResource", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.TagResourceOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.TagResource(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestTransactGetItems(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("TransactGetItems", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.TransactGetItemsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.TransactGetItems(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestTransactWriteItems(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("TransactWriteItems", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.TransactWriteItemsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.TransactWriteItems(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUntagResource(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("UntagResource", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.UntagResourceOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.UntagResource(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUpdateContinuousBackups(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("UpdateContinuousBackups", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.UpdateContinuousBackupsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.UpdateContinuousBackups(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUpdateContributorInsights(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("UpdateContributorInsights", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.UpdateContributorInsightsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.UpdateContributorInsights(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUpdateGlobalTable(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("UpdateGlobalTable", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.UpdateGlobalTableOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.UpdateGlobalTable(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUpdateGlobalTableSettings(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("UpdateGlobalTableSettings", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.UpdateGlobalTableSettingsOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.UpdateGlobalTableSettings(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUpdateItem(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("UpdateItem", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.UpdateItemOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.UpdateItem(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUpdateTable(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("UpdateTable", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.UpdateTableOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.UpdateTable(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUpdateTableReplicaAutoScaling(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("UpdateTableReplicaAutoScaling", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.UpdateTableReplicaAutoScalingOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.UpdateTableReplicaAutoScaling(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUpdateTimeToLive(t *testing.T) {
	client := new(dynamodbmock.Client)
	client.On("UpdateTimeToLive", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.UpdateTimeToLiveOutput{}, nil)
	fn := func(c dynamodbClient) {
		_, _ = c.UpdateTimeToLive(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}
