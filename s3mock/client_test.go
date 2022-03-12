package s3mock_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/mock"

	"github.com/gvre/awsmock-v2/s3mock"
)

type s3Client interface {
	AbortMultipartUpload(ctx context.Context, params *s3.AbortMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error)
	CompleteMultipartUpload(ctx context.Context, params *s3.CompleteMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error)
	CopyObject(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error)
	CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
	CreateMultipartUpload(ctx context.Context, params *s3.CreateMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error)
	DeleteBucket(ctx context.Context, params *s3.DeleteBucketInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketOutput, error)
	DeleteBucketAnalyticsConfiguration(ctx context.Context, params *s3.DeleteBucketAnalyticsConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketAnalyticsConfigurationOutput, error)
	DeleteBucketCors(ctx context.Context, params *s3.DeleteBucketCorsInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketCorsOutput, error)
	DeleteBucketEncryption(ctx context.Context, params *s3.DeleteBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketEncryptionOutput, error)
	DeleteBucketIntelligentTieringConfiguration(ctx context.Context, params *s3.DeleteBucketIntelligentTieringConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error)
	DeleteBucketInventoryConfiguration(ctx context.Context, params *s3.DeleteBucketInventoryConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketInventoryConfigurationOutput, error)
	DeleteBucketLifecycle(ctx context.Context, params *s3.DeleteBucketLifecycleInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketLifecycleOutput, error)
	DeleteBucketMetricsConfiguration(ctx context.Context, params *s3.DeleteBucketMetricsConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketMetricsConfigurationOutput, error)
	DeleteBucketOwnershipControls(ctx context.Context, params *s3.DeleteBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketOwnershipControlsOutput, error)
	DeleteBucketPolicy(ctx context.Context, params *s3.DeleteBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketPolicyOutput, error)
	DeleteBucketReplication(ctx context.Context, params *s3.DeleteBucketReplicationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketReplicationOutput, error)
	DeleteBucketTagging(ctx context.Context, params *s3.DeleteBucketTaggingInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketTaggingOutput, error)
	DeleteBucketWebsite(ctx context.Context, params *s3.DeleteBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketWebsiteOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
	DeleteObjectTagging(ctx context.Context, params *s3.DeleteObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectTaggingOutput, error)
	DeleteObjects(ctx context.Context, params *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error)
	DeletePublicAccessBlock(ctx context.Context, params *s3.DeletePublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.DeletePublicAccessBlockOutput, error)
	GetBucketAccelerateConfiguration(ctx context.Context, params *s3.GetBucketAccelerateConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketAccelerateConfigurationOutput, error)
	GetBucketAcl(ctx context.Context, params *s3.GetBucketAclInput, optFns ...func(*s3.Options)) (*s3.GetBucketAclOutput, error)
	GetBucketAnalyticsConfiguration(ctx context.Context, params *s3.GetBucketAnalyticsConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketAnalyticsConfigurationOutput, error)
	GetBucketCors(ctx context.Context, params *s3.GetBucketCorsInput, optFns ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error)
	GetBucketEncryption(ctx context.Context, params *s3.GetBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error)
	GetBucketIntelligentTieringConfiguration(ctx context.Context, params *s3.GetBucketIntelligentTieringConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketIntelligentTieringConfigurationOutput, error)
	GetBucketInventoryConfiguration(ctx context.Context, params *s3.GetBucketInventoryConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketInventoryConfigurationOutput, error)
	GetBucketLifecycleConfiguration(ctx context.Context, params *s3.GetBucketLifecycleConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLocation(ctx context.Context, params *s3.GetBucketLocationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error)
	GetBucketLogging(ctx context.Context, params *s3.GetBucketLoggingInput, optFns ...func(*s3.Options)) (*s3.GetBucketLoggingOutput, error)
	GetBucketMetricsConfiguration(ctx context.Context, params *s3.GetBucketMetricsConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketMetricsConfigurationOutput, error)
	GetBucketNotificationConfiguration(ctx context.Context, params *s3.GetBucketNotificationConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketNotificationConfigurationOutput, error)
	GetBucketOwnershipControls(ctx context.Context, params *s3.GetBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.GetBucketOwnershipControlsOutput, error)
	GetBucketPolicy(ctx context.Context, params *s3.GetBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error)
	GetBucketPolicyStatus(ctx context.Context, params *s3.GetBucketPolicyStatusInput, optFns ...func(*s3.Options)) (*s3.GetBucketPolicyStatusOutput, error)
	GetBucketReplication(ctx context.Context, params *s3.GetBucketReplicationInput, optFns ...func(*s3.Options)) (*s3.GetBucketReplicationOutput, error)
	GetBucketRequestPayment(ctx context.Context, params *s3.GetBucketRequestPaymentInput, optFns ...func(*s3.Options)) (*s3.GetBucketRequestPaymentOutput, error)
	GetBucketTagging(ctx context.Context, params *s3.GetBucketTaggingInput, optFns ...func(*s3.Options)) (*s3.GetBucketTaggingOutput, error)
	GetBucketVersioning(ctx context.Context, params *s3.GetBucketVersioningInput, optFns ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error)
	GetBucketWebsite(ctx context.Context, params *s3.GetBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.GetBucketWebsiteOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	GetObjectAcl(ctx context.Context, params *s3.GetObjectAclInput, optFns ...func(*s3.Options)) (*s3.GetObjectAclOutput, error)
	GetObjectAttributes(ctx context.Context, params *s3.GetObjectAttributesInput, optFns ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error)
	GetObjectLegalHold(ctx context.Context, params *s3.GetObjectLegalHoldInput, optFns ...func(*s3.Options)) (*s3.GetObjectLegalHoldOutput, error)
	GetObjectLockConfiguration(ctx context.Context, params *s3.GetObjectLockConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetObjectLockConfigurationOutput, error)
	GetObjectRetention(ctx context.Context, params *s3.GetObjectRetentionInput, optFns ...func(*s3.Options)) (*s3.GetObjectRetentionOutput, error)
	GetObjectTagging(ctx context.Context, params *s3.GetObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.GetObjectTaggingOutput, error)
	GetObjectTorrent(ctx context.Context, params *s3.GetObjectTorrentInput, optFns ...func(*s3.Options)) (*s3.GetObjectTorrentOutput, error)
	GetPublicAccessBlock(ctx context.Context, params *s3.GetPublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.GetPublicAccessBlockOutput, error)
	HeadBucket(ctx context.Context, params *s3.HeadBucketInput, optFns ...func(*s3.Options)) (*s3.HeadBucketOutput, error)
	HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error)
	ListBucketAnalyticsConfigurations(ctx context.Context, params *s3.ListBucketAnalyticsConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketAnalyticsConfigurationsOutput, error)
	ListBucketIntelligentTieringConfigurations(ctx context.Context, params *s3.ListBucketIntelligentTieringConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error)
	ListBucketInventoryConfigurations(ctx context.Context, params *s3.ListBucketInventoryConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketInventoryConfigurationsOutput, error)
	ListBucketMetricsConfigurations(ctx context.Context, params *s3.ListBucketMetricsConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketMetricsConfigurationsOutput, error)
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	ListMultipartUploads(ctx context.Context, params *s3.ListMultipartUploadsInput, optFns ...func(*s3.Options)) (*s3.ListMultipartUploadsOutput, error)
	ListObjectVersions(ctx context.Context, params *s3.ListObjectVersionsInput, optFns ...func(*s3.Options)) (*s3.ListObjectVersionsOutput, error)
	ListObjects(ctx context.Context, params *s3.ListObjectsInput, optFns ...func(*s3.Options)) (*s3.ListObjectsOutput, error)
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	ListParts(ctx context.Context, params *s3.ListPartsInput, optFns ...func(*s3.Options)) (*s3.ListPartsOutput, error)
	PutBucketAccelerateConfiguration(ctx context.Context, params *s3.PutBucketAccelerateConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketAccelerateConfigurationOutput, error)
	PutBucketAcl(ctx context.Context, params *s3.PutBucketAclInput, optFns ...func(*s3.Options)) (*s3.PutBucketAclOutput, error)
	PutBucketAnalyticsConfiguration(ctx context.Context, params *s3.PutBucketAnalyticsConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketAnalyticsConfigurationOutput, error)
	PutBucketCors(ctx context.Context, params *s3.PutBucketCorsInput, optFns ...func(*s3.Options)) (*s3.PutBucketCorsOutput, error)
	PutBucketEncryption(ctx context.Context, params *s3.PutBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.PutBucketEncryptionOutput, error)
	PutBucketIntelligentTieringConfiguration(ctx context.Context, params *s3.PutBucketIntelligentTieringConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketIntelligentTieringConfigurationOutput, error)
	PutBucketInventoryConfiguration(ctx context.Context, params *s3.PutBucketInventoryConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketInventoryConfigurationOutput, error)
	PutBucketLifecycleConfiguration(ctx context.Context, params *s3.PutBucketLifecycleConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLogging(ctx context.Context, params *s3.PutBucketLoggingInput, optFns ...func(*s3.Options)) (*s3.PutBucketLoggingOutput, error)
	PutBucketMetricsConfiguration(ctx context.Context, params *s3.PutBucketMetricsConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketMetricsConfigurationOutput, error)
	PutBucketNotificationConfiguration(ctx context.Context, params *s3.PutBucketNotificationConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketNotificationConfigurationOutput, error)
	PutBucketOwnershipControls(ctx context.Context, params *s3.PutBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.PutBucketOwnershipControlsOutput, error)
	PutBucketPolicy(ctx context.Context, params *s3.PutBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.PutBucketPolicyOutput, error)
	PutBucketReplication(ctx context.Context, params *s3.PutBucketReplicationInput, optFns ...func(*s3.Options)) (*s3.PutBucketReplicationOutput, error)
	PutBucketRequestPayment(ctx context.Context, params *s3.PutBucketRequestPaymentInput, optFns ...func(*s3.Options)) (*s3.PutBucketRequestPaymentOutput, error)
	PutBucketTagging(ctx context.Context, params *s3.PutBucketTaggingInput, optFns ...func(*s3.Options)) (*s3.PutBucketTaggingOutput, error)
	PutBucketVersioning(ctx context.Context, params *s3.PutBucketVersioningInput, optFns ...func(*s3.Options)) (*s3.PutBucketVersioningOutput, error)
	PutBucketWebsite(ctx context.Context, params *s3.PutBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.PutBucketWebsiteOutput, error)
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	PutObjectAcl(ctx context.Context, params *s3.PutObjectAclInput, optFns ...func(*s3.Options)) (*s3.PutObjectAclOutput, error)
	PutObjectLegalHold(ctx context.Context, params *s3.PutObjectLegalHoldInput, optFns ...func(*s3.Options)) (*s3.PutObjectLegalHoldOutput, error)
	PutObjectLockConfiguration(ctx context.Context, params *s3.PutObjectLockConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutObjectLockConfigurationOutput, error)
	PutObjectRetention(ctx context.Context, params *s3.PutObjectRetentionInput, optFns ...func(*s3.Options)) (*s3.PutObjectRetentionOutput, error)
	PutObjectTagging(ctx context.Context, params *s3.PutObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.PutObjectTaggingOutput, error)
	PutPublicAccessBlock(ctx context.Context, params *s3.PutPublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.PutPublicAccessBlockOutput, error)
	RestoreObject(ctx context.Context, params *s3.RestoreObjectInput, optFns ...func(*s3.Options)) (*s3.RestoreObjectOutput, error)
	SelectObjectContent(ctx context.Context, params *s3.SelectObjectContentInput, optFns ...func(*s3.Options)) (*s3.SelectObjectContentOutput, error)
	UploadPart(ctx context.Context, params *s3.UploadPartInput, optFns ...func(*s3.Options)) (*s3.UploadPartOutput, error)
	UploadPartCopy(ctx context.Context, params *s3.UploadPartCopyInput, optFns ...func(*s3.Options)) (*s3.UploadPartCopyOutput, error)
	WriteGetObjectResponse(ctx context.Context, params *s3.WriteGetObjectResponseInput, optFns ...func(*s3.Options)) (*s3.WriteGetObjectResponseOutput, error)
}

func TestAbortMultipartUpload(t *testing.T) {
	client := new(s3mock.Client)
	client.On("AbortMultipartUpload", mock.Anything, mock.Anything, mock.Anything).Return(&s3.AbortMultipartUploadOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.AbortMultipartUpload(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCompleteMultipartUpload(t *testing.T) {
	client := new(s3mock.Client)
	client.On("CompleteMultipartUpload", mock.Anything, mock.Anything, mock.Anything).Return(&s3.CompleteMultipartUploadOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.CompleteMultipartUpload(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCopyObject(t *testing.T) {
	client := new(s3mock.Client)
	client.On("CopyObject", mock.Anything, mock.Anything, mock.Anything).Return(&s3.CopyObjectOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.CopyObject(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCreateBucket(t *testing.T) {
	client := new(s3mock.Client)
	client.On("CreateBucket", mock.Anything, mock.Anything, mock.Anything).Return(&s3.CreateBucketOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.CreateBucket(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCreateMultipartUpload(t *testing.T) {
	client := new(s3mock.Client)
	client.On("CreateMultipartUpload", mock.Anything, mock.Anything, mock.Anything).Return(&s3.CreateMultipartUploadOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.CreateMultipartUpload(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucket(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucket", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucket(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketAnalyticsConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketAnalyticsConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketAnalyticsConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketAnalyticsConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketCors(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketCors", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketCorsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketCors(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketEncryption(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketEncryption", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketEncryptionOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketEncryption(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketIntelligentTieringConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketIntelligentTieringConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketIntelligentTieringConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketIntelligentTieringConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketInventoryConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketInventoryConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketInventoryConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketInventoryConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketLifecycle(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketLifecycle", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketLifecycleOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketLifecycle(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketMetricsConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketMetricsConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketMetricsConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketMetricsConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketOwnershipControls(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketOwnershipControls", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketOwnershipControlsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketOwnershipControls(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketPolicy(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketPolicy", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketPolicyOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketPolicy(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketReplication(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketReplication", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketReplicationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketReplication(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketTagging(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketTagging", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketTaggingOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketTagging(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteBucketWebsite(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteBucketWebsite", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteBucketWebsiteOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteBucketWebsite(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteObject(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteObject", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteObjectOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteObject(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteObjectTagging(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteObjectTagging", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteObjectTaggingOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteObjectTagging(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteObjects(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeleteObjects", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeleteObjectsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeleteObjects(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeletePublicAccessBlock(t *testing.T) {
	client := new(s3mock.Client)
	client.On("DeletePublicAccessBlock", mock.Anything, mock.Anything, mock.Anything).Return(&s3.DeletePublicAccessBlockOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.DeletePublicAccessBlock(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketAccelerateConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketAccelerateConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketAccelerateConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketAccelerateConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketAcl(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketAcl", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketAclOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketAcl(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketAnalyticsConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketAnalyticsConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketAnalyticsConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketAnalyticsConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketCors(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketCors", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketCorsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketCors(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketEncryption(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketEncryption", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketEncryptionOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketEncryption(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketIntelligentTieringConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketIntelligentTieringConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketIntelligentTieringConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketIntelligentTieringConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketInventoryConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketInventoryConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketInventoryConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketInventoryConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketLifecycleConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketLifecycleConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketLifecycleConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketLifecycleConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketLocation(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketLocation", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketLocationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketLocation(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketLogging(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketLogging", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketLoggingOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketLogging(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketMetricsConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketMetricsConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketMetricsConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketMetricsConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketNotificationConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketNotificationConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketNotificationConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketNotificationConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketOwnershipControls(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketOwnershipControls", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketOwnershipControlsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketOwnershipControls(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketPolicy(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketPolicy", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketPolicyOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketPolicy(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketPolicyStatus(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketPolicyStatus", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketPolicyStatusOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketPolicyStatus(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketReplication(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketReplication", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketReplicationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketReplication(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketRequestPayment(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketRequestPayment", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketRequestPaymentOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketRequestPayment(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketTagging(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketTagging", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketTaggingOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketTagging(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketVersioning(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketVersioning", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketVersioningOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketVersioning(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetBucketWebsite(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetBucketWebsite", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetBucketWebsiteOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetBucketWebsite(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetObject(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetObject", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetObjectOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetObject(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetObjectAcl(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetObjectAcl", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetObjectAclOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetObjectAcl(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetObjectAttributes(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetObjectAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetObjectAttributesOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetObjectAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetObjectLegalHold(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetObjectLegalHold", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetObjectLegalHoldOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetObjectLegalHold(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetObjectLockConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetObjectLockConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetObjectLockConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetObjectLockConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetObjectRetention(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetObjectRetention", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetObjectRetentionOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetObjectRetention(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetObjectTagging(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetObjectTagging", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetObjectTaggingOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetObjectTagging(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetObjectTorrent(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetObjectTorrent", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetObjectTorrentOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetObjectTorrent(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetPublicAccessBlock(t *testing.T) {
	client := new(s3mock.Client)
	client.On("GetPublicAccessBlock", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetPublicAccessBlockOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.GetPublicAccessBlock(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestHeadBucket(t *testing.T) {
	client := new(s3mock.Client)
	client.On("HeadBucket", mock.Anything, mock.Anything, mock.Anything).Return(&s3.HeadBucketOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.HeadBucket(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestHeadObject(t *testing.T) {
	client := new(s3mock.Client)
	client.On("HeadObject", mock.Anything, mock.Anything, mock.Anything).Return(&s3.HeadObjectOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.HeadObject(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListBucketAnalyticsConfigurations(t *testing.T) {
	client := new(s3mock.Client)
	client.On("ListBucketAnalyticsConfigurations", mock.Anything, mock.Anything, mock.Anything).Return(&s3.ListBucketAnalyticsConfigurationsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.ListBucketAnalyticsConfigurations(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListBucketIntelligentTieringConfigurations(t *testing.T) {
	client := new(s3mock.Client)
	client.On("ListBucketIntelligentTieringConfigurations", mock.Anything, mock.Anything, mock.Anything).Return(&s3.ListBucketIntelligentTieringConfigurationsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.ListBucketIntelligentTieringConfigurations(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListBucketInventoryConfigurations(t *testing.T) {
	client := new(s3mock.Client)
	client.On("ListBucketInventoryConfigurations", mock.Anything, mock.Anything, mock.Anything).Return(&s3.ListBucketInventoryConfigurationsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.ListBucketInventoryConfigurations(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListBucketMetricsConfigurations(t *testing.T) {
	client := new(s3mock.Client)
	client.On("ListBucketMetricsConfigurations", mock.Anything, mock.Anything, mock.Anything).Return(&s3.ListBucketMetricsConfigurationsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.ListBucketMetricsConfigurations(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListBuckets(t *testing.T) {
	client := new(s3mock.Client)
	client.On("ListBuckets", mock.Anything, mock.Anything, mock.Anything).Return(&s3.ListBucketsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.ListBuckets(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListMultipartUploads(t *testing.T) {
	client := new(s3mock.Client)
	client.On("ListMultipartUploads", mock.Anything, mock.Anything, mock.Anything).Return(&s3.ListMultipartUploadsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.ListMultipartUploads(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListObjectVersions(t *testing.T) {
	client := new(s3mock.Client)
	client.On("ListObjectVersions", mock.Anything, mock.Anything, mock.Anything).Return(&s3.ListObjectVersionsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.ListObjectVersions(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListObjects(t *testing.T) {
	client := new(s3mock.Client)
	client.On("ListObjects", mock.Anything, mock.Anything, mock.Anything).Return(&s3.ListObjectsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.ListObjects(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListObjectsV2(t *testing.T) {
	client := new(s3mock.Client)
	client.On("ListObjectsV2", mock.Anything, mock.Anything, mock.Anything).Return(&s3.ListObjectsV2Output{}, nil)
	fn := func(c s3Client) {
		_, _ = c.ListObjectsV2(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListParts(t *testing.T) {
	client := new(s3mock.Client)
	client.On("ListParts", mock.Anything, mock.Anything, mock.Anything).Return(&s3.ListPartsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.ListParts(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketAccelerateConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketAccelerateConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketAccelerateConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketAccelerateConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketAcl(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketAcl", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketAclOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketAcl(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketAnalyticsConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketAnalyticsConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketAnalyticsConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketAnalyticsConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketCors(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketCors", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketCorsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketCors(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketEncryption(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketEncryption", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketEncryptionOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketEncryption(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketIntelligentTieringConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketIntelligentTieringConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketIntelligentTieringConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketIntelligentTieringConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketInventoryConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketInventoryConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketInventoryConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketInventoryConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketLifecycleConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketLifecycleConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketLifecycleConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketLifecycleConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketLogging(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketLogging", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketLoggingOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketLogging(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketMetricsConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketMetricsConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketMetricsConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketMetricsConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketNotificationConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketNotificationConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketNotificationConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketNotificationConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketOwnershipControls(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketOwnershipControls", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketOwnershipControlsOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketOwnershipControls(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketPolicy(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketPolicy", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketPolicyOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketPolicy(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketReplication(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketReplication", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketReplicationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketReplication(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketRequestPayment(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketRequestPayment", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketRequestPaymentOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketRequestPayment(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketTagging(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketTagging", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketTaggingOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketTagging(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketVersioning(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketVersioning", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketVersioningOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketVersioning(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutBucketWebsite(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutBucketWebsite", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutBucketWebsiteOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutBucketWebsite(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutObject(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutObject", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutObjectOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutObject(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutObjectAcl(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutObjectAcl", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutObjectAclOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutObjectAcl(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutObjectLegalHold(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutObjectLegalHold", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutObjectLegalHoldOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutObjectLegalHold(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutObjectLockConfiguration(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutObjectLockConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutObjectLockConfigurationOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutObjectLockConfiguration(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutObjectRetention(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutObjectRetention", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutObjectRetentionOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutObjectRetention(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutObjectTagging(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutObjectTagging", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutObjectTaggingOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutObjectTagging(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPutPublicAccessBlock(t *testing.T) {
	client := new(s3mock.Client)
	client.On("PutPublicAccessBlock", mock.Anything, mock.Anything, mock.Anything).Return(&s3.PutPublicAccessBlockOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.PutPublicAccessBlock(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestRestoreObject(t *testing.T) {
	client := new(s3mock.Client)
	client.On("RestoreObject", mock.Anything, mock.Anything, mock.Anything).Return(&s3.RestoreObjectOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.RestoreObject(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestSelectObjectContent(t *testing.T) {
	client := new(s3mock.Client)
	client.On("SelectObjectContent", mock.Anything, mock.Anything, mock.Anything).Return(&s3.SelectObjectContentOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.SelectObjectContent(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUploadPart(t *testing.T) {
	client := new(s3mock.Client)
	client.On("UploadPart", mock.Anything, mock.Anything, mock.Anything).Return(&s3.UploadPartOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.UploadPart(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUploadPartCopy(t *testing.T) {
	client := new(s3mock.Client)
	client.On("UploadPartCopy", mock.Anything, mock.Anything, mock.Anything).Return(&s3.UploadPartCopyOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.UploadPartCopy(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestWriteGetObjectResponse(t *testing.T) {
	client := new(s3mock.Client)
	client.On("WriteGetObjectResponse", mock.Anything, mock.Anything, mock.Anything).Return(&s3.WriteGetObjectResponseOutput{}, nil)
	fn := func(c s3Client) {
		_, _ = c.WriteGetObjectResponse(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}
