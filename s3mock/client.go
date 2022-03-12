package s3mock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/mock"
)

// Client provides the API client to mock operations call for Amazon S3 Service.
type Client struct {
	mock.Mock
}

// AbortMultipartUpload mock.
func (c *Client) AbortMultipartUpload(ctx context.Context, params *s3.AbortMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.AbortMultipartUploadOutput), args.Error(1)
}

// CompleteMultipartUpload mock.
func (c *Client) CompleteMultipartUpload(ctx context.Context, params *s3.CompleteMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.CompleteMultipartUploadOutput), args.Error(1)
}

// CopyObject mock.
func (c *Client) CopyObject(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.CopyObjectOutput), args.Error(1)
}

// CreateBucket mock.
func (c *Client) CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.CreateBucketOutput), args.Error(1)
}

// CreateMultipartUpload mock.
func (c *Client) CreateMultipartUpload(ctx context.Context, params *s3.CreateMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.CreateMultipartUploadOutput), args.Error(1)
}

// DeleteBucket mock.
func (c *Client) DeleteBucket(ctx context.Context, params *s3.DeleteBucketInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketOutput), args.Error(1)
}

// DeleteBucketAnalyticsConfiguration mock.
func (c *Client) DeleteBucketAnalyticsConfiguration(ctx context.Context, params *s3.DeleteBucketAnalyticsConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketAnalyticsConfigurationOutput), args.Error(1)
}

// DeleteBucketCors mock.
func (c *Client) DeleteBucketCors(ctx context.Context, params *s3.DeleteBucketCorsInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketCorsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketCorsOutput), args.Error(1)
}

// DeleteBucketEncryption mock.
func (c *Client) DeleteBucketEncryption(ctx context.Context, params *s3.DeleteBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketEncryptionOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketEncryptionOutput), args.Error(1)
}

// DeleteBucketIntelligentTieringConfiguration mock.
func (c *Client) DeleteBucketIntelligentTieringConfiguration(ctx context.Context, params *s3.DeleteBucketIntelligentTieringConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketIntelligentTieringConfigurationOutput), args.Error(1)
}

// DeleteBucketInventoryConfiguration mock.
func (c *Client) DeleteBucketInventoryConfiguration(ctx context.Context, params *s3.DeleteBucketInventoryConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketInventoryConfigurationOutput), args.Error(1)
}

// DeleteBucketLifecycle mock.
func (c *Client) DeleteBucketLifecycle(ctx context.Context, params *s3.DeleteBucketLifecycleInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketLifecycleOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketLifecycleOutput), args.Error(1)
}

// DeleteBucketMetricsConfiguration mock.
func (c *Client) DeleteBucketMetricsConfiguration(ctx context.Context, params *s3.DeleteBucketMetricsConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketMetricsConfigurationOutput), args.Error(1)
}

// DeleteBucketOwnershipControls mock.
func (c *Client) DeleteBucketOwnershipControls(ctx context.Context, params *s3.DeleteBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketOwnershipControlsOutput), args.Error(1)
}

// DeleteBucketPolicy mock.
func (c *Client) DeleteBucketPolicy(ctx context.Context, params *s3.DeleteBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketPolicyOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketPolicyOutput), args.Error(1)
}

// DeleteBucketReplication mock.
func (c *Client) DeleteBucketReplication(ctx context.Context, params *s3.DeleteBucketReplicationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketReplicationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketReplicationOutput), args.Error(1)
}

// DeleteBucketTagging mock.
func (c *Client) DeleteBucketTagging(ctx context.Context, params *s3.DeleteBucketTaggingInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketTaggingOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketTaggingOutput), args.Error(1)
}

// DeleteBucketWebsite mock.
func (c *Client) DeleteBucketWebsite(ctx context.Context, params *s3.DeleteBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketWebsiteOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteBucketWebsiteOutput), args.Error(1)
}

// DeleteObject mock.
func (c *Client) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteObjectOutput), args.Error(1)
}

// DeleteObjectTagging mock.
func (c *Client) DeleteObjectTagging(ctx context.Context, params *s3.DeleteObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectTaggingOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteObjectTaggingOutput), args.Error(1)
}

// DeleteObjects mock.
func (c *Client) DeleteObjects(ctx context.Context, params *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteObjectsOutput), args.Error(1)
}

// DeletePublicAccessBlock mock.
func (c *Client) DeletePublicAccessBlock(ctx context.Context, params *s3.DeletePublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.DeletePublicAccessBlockOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeletePublicAccessBlockOutput), args.Error(1)
}

// GetBucketAccelerateConfiguration mock.
func (c *Client) GetBucketAccelerateConfiguration(ctx context.Context, params *s3.GetBucketAccelerateConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketAccelerateConfigurationOutput), args.Error(1)
}

// GetBucketAcl mock.
func (c *Client) GetBucketAcl(ctx context.Context, params *s3.GetBucketAclInput, optFns ...func(*s3.Options)) (*s3.GetBucketAclOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketAclOutput), args.Error(1)
}

// GetBucketAnalyticsConfiguration mock.
func (c *Client) GetBucketAnalyticsConfiguration(ctx context.Context, params *s3.GetBucketAnalyticsConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketAnalyticsConfigurationOutput), args.Error(1)
}

// GetBucketCors mock.
func (c *Client) GetBucketCors(ctx context.Context, params *s3.GetBucketCorsInput, optFns ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketCorsOutput), args.Error(1)
}

// GetBucketEncryption mock.
func (c *Client) GetBucketEncryption(ctx context.Context, params *s3.GetBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketEncryptionOutput), args.Error(1)
}

// GetBucketIntelligentTieringConfiguration mock.
func (c *Client) GetBucketIntelligentTieringConfiguration(ctx context.Context, params *s3.GetBucketIntelligentTieringConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketIntelligentTieringConfigurationOutput), args.Error(1)
}

// GetBucketInventoryConfiguration mock.
func (c *Client) GetBucketInventoryConfiguration(ctx context.Context, params *s3.GetBucketInventoryConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketInventoryConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketInventoryConfigurationOutput), args.Error(1)
}

// GetBucketLifecycleConfiguration mock.
func (c *Client) GetBucketLifecycleConfiguration(ctx context.Context, params *s3.GetBucketLifecycleConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketLifecycleConfigurationOutput), args.Error(1)
}

// GetBucketLocation mock.
func (c *Client) GetBucketLocation(ctx context.Context, params *s3.GetBucketLocationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketLocationOutput), args.Error(1)
}

// GetBucketLogging mock.
func (c *Client) GetBucketLogging(ctx context.Context, params *s3.GetBucketLoggingInput, optFns ...func(*s3.Options)) (*s3.GetBucketLoggingOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketLoggingOutput), args.Error(1)
}

// GetBucketMetricsConfiguration mock.
func (c *Client) GetBucketMetricsConfiguration(ctx context.Context, params *s3.GetBucketMetricsConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketMetricsConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketMetricsConfigurationOutput), args.Error(1)
}

// GetBucketNotificationConfiguration mock.
func (c *Client) GetBucketNotificationConfiguration(ctx context.Context, params *s3.GetBucketNotificationConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketNotificationConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketNotificationConfigurationOutput), args.Error(1)
}

// GetBucketOwnershipControls mock.
func (c *Client) GetBucketOwnershipControls(ctx context.Context, params *s3.GetBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.GetBucketOwnershipControlsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketOwnershipControlsOutput), args.Error(1)
}

// GetBucketPolicy mock.
func (c *Client) GetBucketPolicy(ctx context.Context, params *s3.GetBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketPolicyOutput), args.Error(1)
}

// GetBucketPolicyStatus mock.
func (c *Client) GetBucketPolicyStatus(ctx context.Context, params *s3.GetBucketPolicyStatusInput, optFns ...func(*s3.Options)) (*s3.GetBucketPolicyStatusOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketPolicyStatusOutput), args.Error(1)
}

// GetBucketReplication mock.
func (c *Client) GetBucketReplication(ctx context.Context, params *s3.GetBucketReplicationInput, optFns ...func(*s3.Options)) (*s3.GetBucketReplicationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketReplicationOutput), args.Error(1)
}

// GetBucketRequestPayment mock.
func (c *Client) GetBucketRequestPayment(ctx context.Context, params *s3.GetBucketRequestPaymentInput, optFns ...func(*s3.Options)) (*s3.GetBucketRequestPaymentOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketRequestPaymentOutput), args.Error(1)
}

// GetBucketTagging mock.
func (c *Client) GetBucketTagging(ctx context.Context, params *s3.GetBucketTaggingInput, optFns ...func(*s3.Options)) (*s3.GetBucketTaggingOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketTaggingOutput), args.Error(1)
}

// GetBucketVersioning mock.
func (c *Client) GetBucketVersioning(ctx context.Context, params *s3.GetBucketVersioningInput, optFns ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketVersioningOutput), args.Error(1)
}

// GetBucketWebsite mock.
func (c *Client) GetBucketWebsite(ctx context.Context, params *s3.GetBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.GetBucketWebsiteOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetBucketWebsiteOutput), args.Error(1)
}

// GetObject mock.
func (c *Client) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetObjectOutput), args.Error(1)
}

// GetObjectAcl mock.
func (c *Client) GetObjectAcl(ctx context.Context, params *s3.GetObjectAclInput, optFns ...func(*s3.Options)) (*s3.GetObjectAclOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetObjectAclOutput), args.Error(1)
}

// GetObjectAttributes mock.
func (c *Client) GetObjectAttributes(ctx context.Context, params *s3.GetObjectAttributesInput, optFns ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetObjectAttributesOutput), args.Error(1)
}

// GetObjectLegalHold mock.
func (c *Client) GetObjectLegalHold(ctx context.Context, params *s3.GetObjectLegalHoldInput, optFns ...func(*s3.Options)) (*s3.GetObjectLegalHoldOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetObjectLegalHoldOutput), args.Error(1)
}

// GetObjectLockConfiguration mock.
func (c *Client) GetObjectLockConfiguration(ctx context.Context, params *s3.GetObjectLockConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetObjectLockConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetObjectLockConfigurationOutput), args.Error(1)
}

// GetObjectRetention mock.
func (c *Client) GetObjectRetention(ctx context.Context, params *s3.GetObjectRetentionInput, optFns ...func(*s3.Options)) (*s3.GetObjectRetentionOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetObjectRetentionOutput), args.Error(1)
}

// GetObjectTagging mock.
func (c *Client) GetObjectTagging(ctx context.Context, params *s3.GetObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.GetObjectTaggingOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetObjectTaggingOutput), args.Error(1)
}

// GetObjectTorrent mock.
func (c *Client) GetObjectTorrent(ctx context.Context, params *s3.GetObjectTorrentInput, optFns ...func(*s3.Options)) (*s3.GetObjectTorrentOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetObjectTorrentOutput), args.Error(1)
}

// GetPublicAccessBlock mock.
func (c *Client) GetPublicAccessBlock(ctx context.Context, params *s3.GetPublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.GetPublicAccessBlockOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.GetPublicAccessBlockOutput), args.Error(1)
}

// HeadBucket mock.
func (c *Client) HeadBucket(ctx context.Context, params *s3.HeadBucketInput, optFns ...func(*s3.Options)) (*s3.HeadBucketOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.HeadBucketOutput), args.Error(1)
}

// HeadObject mock.
func (c *Client) HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.HeadObjectOutput), args.Error(1)
}

// ListBucketAnalyticsConfigurations mock.
func (c *Client) ListBucketAnalyticsConfigurations(ctx context.Context, params *s3.ListBucketAnalyticsConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.ListBucketAnalyticsConfigurationsOutput), args.Error(1)
}

// ListBucketIntelligentTieringConfigurations mock.
func (c *Client) ListBucketIntelligentTieringConfigurations(ctx context.Context, params *s3.ListBucketIntelligentTieringConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.ListBucketIntelligentTieringConfigurationsOutput), args.Error(1)
}

// ListBucketInventoryConfigurations mock.
func (c *Client) ListBucketInventoryConfigurations(ctx context.Context, params *s3.ListBucketInventoryConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.ListBucketInventoryConfigurationsOutput), args.Error(1)
}

// ListBucketMetricsConfigurations mock.
func (c *Client) ListBucketMetricsConfigurations(ctx context.Context, params *s3.ListBucketMetricsConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.ListBucketMetricsConfigurationsOutput), args.Error(1)
}

// ListBuckets mock.
func (c *Client) ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.ListBucketsOutput), args.Error(1)
}

// ListMultipartUploads mock.
func (c *Client) ListMultipartUploads(ctx context.Context, params *s3.ListMultipartUploadsInput, optFns ...func(*s3.Options)) (*s3.ListMultipartUploadsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.ListMultipartUploadsOutput), args.Error(1)
}

// ListObjectVersions mock.
func (c *Client) ListObjectVersions(ctx context.Context, params *s3.ListObjectVersionsInput, optFns ...func(*s3.Options)) (*s3.ListObjectVersionsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.ListObjectVersionsOutput), args.Error(1)
}

// ListObjects mock.
func (c *Client) ListObjects(ctx context.Context, params *s3.ListObjectsInput, optFns ...func(*s3.Options)) (*s3.ListObjectsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.ListObjectsOutput), args.Error(1)
}

// ListObjectsV2 mock.
func (c *Client) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.ListObjectsV2Output), args.Error(1)
}

// ListParts mock.
func (c *Client) ListParts(ctx context.Context, params *s3.ListPartsInput, optFns ...func(*s3.Options)) (*s3.ListPartsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.ListPartsOutput), args.Error(1)
}

// PutBucketAccelerateConfiguration mock.
func (c *Client) PutBucketAccelerateConfiguration(ctx context.Context, params *s3.PutBucketAccelerateConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketAccelerateConfigurationOutput), args.Error(1)
}

// PutBucketAcl mock.
func (c *Client) PutBucketAcl(ctx context.Context, params *s3.PutBucketAclInput, optFns ...func(*s3.Options)) (*s3.PutBucketAclOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketAclOutput), args.Error(1)
}

// PutBucketAnalyticsConfiguration mock.
func (c *Client) PutBucketAnalyticsConfiguration(ctx context.Context, params *s3.PutBucketAnalyticsConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketAnalyticsConfigurationOutput), args.Error(1)
}

// PutBucketCors mock.
func (c *Client) PutBucketCors(ctx context.Context, params *s3.PutBucketCorsInput, optFns ...func(*s3.Options)) (*s3.PutBucketCorsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketCorsOutput), args.Error(1)
}

// PutBucketEncryption mock.
func (c *Client) PutBucketEncryption(ctx context.Context, params *s3.PutBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.PutBucketEncryptionOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketEncryptionOutput), args.Error(1)
}

// PutBucketIntelligentTieringConfiguration mock.
func (c *Client) PutBucketIntelligentTieringConfiguration(ctx context.Context, params *s3.PutBucketIntelligentTieringConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketIntelligentTieringConfigurationOutput), args.Error(1)
}

// PutBucketInventoryConfiguration mock.
func (c *Client) PutBucketInventoryConfiguration(ctx context.Context, params *s3.PutBucketInventoryConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketInventoryConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketInventoryConfigurationOutput), args.Error(1)
}

// PutBucketLifecycleConfiguration mock.
func (c *Client) PutBucketLifecycleConfiguration(ctx context.Context, params *s3.PutBucketLifecycleConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketLifecycleConfigurationOutput), args.Error(1)
}

// PutBucketLogging mock.
func (c *Client) PutBucketLogging(ctx context.Context, params *s3.PutBucketLoggingInput, optFns ...func(*s3.Options)) (*s3.PutBucketLoggingOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketLoggingOutput), args.Error(1)
}

// PutBucketMetricsConfiguration mock.
func (c *Client) PutBucketMetricsConfiguration(ctx context.Context, params *s3.PutBucketMetricsConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketMetricsConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketMetricsConfigurationOutput), args.Error(1)
}

// PutBucketNotificationConfiguration mock.
func (c *Client) PutBucketNotificationConfiguration(ctx context.Context, params *s3.PutBucketNotificationConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketNotificationConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketNotificationConfigurationOutput), args.Error(1)
}

// PutBucketOwnershipControls mock.
func (c *Client) PutBucketOwnershipControls(ctx context.Context, params *s3.PutBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.PutBucketOwnershipControlsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketOwnershipControlsOutput), args.Error(1)
}

// PutBucketPolicy mock.
func (c *Client) PutBucketPolicy(ctx context.Context, params *s3.PutBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.PutBucketPolicyOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketPolicyOutput), args.Error(1)
}

// PutBucketReplication mock.
func (c *Client) PutBucketReplication(ctx context.Context, params *s3.PutBucketReplicationInput, optFns ...func(*s3.Options)) (*s3.PutBucketReplicationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketReplicationOutput), args.Error(1)
}

// PutBucketRequestPayment mock.
func (c *Client) PutBucketRequestPayment(ctx context.Context, params *s3.PutBucketRequestPaymentInput, optFns ...func(*s3.Options)) (*s3.PutBucketRequestPaymentOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketRequestPaymentOutput), args.Error(1)
}

// PutBucketTagging mock.
func (c *Client) PutBucketTagging(ctx context.Context, params *s3.PutBucketTaggingInput, optFns ...func(*s3.Options)) (*s3.PutBucketTaggingOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketTaggingOutput), args.Error(1)
}

// PutBucketVersioning mock.
func (c *Client) PutBucketVersioning(ctx context.Context, params *s3.PutBucketVersioningInput, optFns ...func(*s3.Options)) (*s3.PutBucketVersioningOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketVersioningOutput), args.Error(1)
}

// PutBucketWebsite mock.
func (c *Client) PutBucketWebsite(ctx context.Context, params *s3.PutBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.PutBucketWebsiteOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutBucketWebsiteOutput), args.Error(1)
}

// PutObject mock.
func (c *Client) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutObjectOutput), args.Error(1)
}

// PutObjectAcl mock.
func (c *Client) PutObjectAcl(ctx context.Context, params *s3.PutObjectAclInput, optFns ...func(*s3.Options)) (*s3.PutObjectAclOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutObjectAclOutput), args.Error(1)
}

// PutObjectLegalHold mock.
func (c *Client) PutObjectLegalHold(ctx context.Context, params *s3.PutObjectLegalHoldInput, optFns ...func(*s3.Options)) (*s3.PutObjectLegalHoldOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutObjectLegalHoldOutput), args.Error(1)
}

// PutObjectLockConfiguration mock.
func (c *Client) PutObjectLockConfiguration(ctx context.Context, params *s3.PutObjectLockConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutObjectLockConfigurationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutObjectLockConfigurationOutput), args.Error(1)
}

// PutObjectRetention mock.
func (c *Client) PutObjectRetention(ctx context.Context, params *s3.PutObjectRetentionInput, optFns ...func(*s3.Options)) (*s3.PutObjectRetentionOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutObjectRetentionOutput), args.Error(1)
}

// PutObjectTagging mock.
func (c *Client) PutObjectTagging(ctx context.Context, params *s3.PutObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.PutObjectTaggingOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutObjectTaggingOutput), args.Error(1)
}

// PutPublicAccessBlock mock.
func (c *Client) PutPublicAccessBlock(ctx context.Context, params *s3.PutPublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.PutPublicAccessBlockOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutPublicAccessBlockOutput), args.Error(1)
}

// RestoreObject mock.
func (c *Client) RestoreObject(ctx context.Context, params *s3.RestoreObjectInput, optFns ...func(*s3.Options)) (*s3.RestoreObjectOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.RestoreObjectOutput), args.Error(1)
}

// SelectObjectContent mock.
func (c *Client) SelectObjectContent(ctx context.Context, params *s3.SelectObjectContentInput, optFns ...func(*s3.Options)) (*s3.SelectObjectContentOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.SelectObjectContentOutput), args.Error(1)
}

// UploadPart mock.
func (c *Client) UploadPart(ctx context.Context, params *s3.UploadPartInput, optFns ...func(*s3.Options)) (*s3.UploadPartOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.UploadPartOutput), args.Error(1)
}

// UploadPartCopy mock.
func (c *Client) UploadPartCopy(ctx context.Context, params *s3.UploadPartCopyInput, optFns ...func(*s3.Options)) (*s3.UploadPartCopyOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.UploadPartCopyOutput), args.Error(1)
}

// WriteGetObjectResponse mock.
func (c *Client) WriteGetObjectResponse(ctx context.Context, params *s3.WriteGetObjectResponseInput, optFns ...func(*s3.Options)) (*s3.WriteGetObjectResponseOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.WriteGetObjectResponseOutput), args.Error(1)
}
