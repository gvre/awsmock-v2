package snsmock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/stretchr/testify/mock"
)

// Client provides the API client to mock operations call for Amazon SNS Service.
type Client struct {
	mock.Mock
}

// AddPermission mock.
func (c *Client) AddPermission(ctx context.Context, params *sns.AddPermissionInput, optFns ...func(*sns.Options)) (*sns.AddPermissionOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.AddPermissionOutput), args.Error(1)
}

// CheckIfPhoneNumberIsOptedOut mock.
func (c *Client) CheckIfPhoneNumberIsOptedOut(ctx context.Context, params *sns.CheckIfPhoneNumberIsOptedOutInput, optFns ...func(*sns.Options)) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.CheckIfPhoneNumberIsOptedOutOutput), args.Error(1)
}

// ConfirmSubscription mock.
func (c *Client) ConfirmSubscription(ctx context.Context, params *sns.ConfirmSubscriptionInput, optFns ...func(*sns.Options)) (*sns.ConfirmSubscriptionOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.ConfirmSubscriptionOutput), args.Error(1)
}

// CreatePlatformApplication mock.
func (c *Client) CreatePlatformApplication(ctx context.Context, params *sns.CreatePlatformApplicationInput, optFns ...func(*sns.Options)) (*sns.CreatePlatformApplicationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.CreatePlatformApplicationOutput), args.Error(1)
}

// CreatePlatformEndpoint mock.
func (c *Client) CreatePlatformEndpoint(ctx context.Context, params *sns.CreatePlatformEndpointInput, optFns ...func(*sns.Options)) (*sns.CreatePlatformEndpointOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.CreatePlatformEndpointOutput), args.Error(1)
}

// CreateSMSSandboxPhoneNumber mock.
func (c *Client) CreateSMSSandboxPhoneNumber(ctx context.Context, params *sns.CreateSMSSandboxPhoneNumberInput, optFns ...func(*sns.Options)) (*sns.CreateSMSSandboxPhoneNumberOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.CreateSMSSandboxPhoneNumberOutput), args.Error(1)
}

// CreateTopic mock.
func (c *Client) CreateTopic(ctx context.Context, params *sns.CreateTopicInput, optFns ...func(*sns.Options)) (*sns.CreateTopicOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.CreateTopicOutput), args.Error(1)
}

// DeleteEndpoint mock.
func (c *Client) DeleteEndpoint(ctx context.Context, params *sns.DeleteEndpointInput, optFns ...func(*sns.Options)) (*sns.DeleteEndpointOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.DeleteEndpointOutput), args.Error(1)
}

// DeletePlatformApplication mock.
func (c *Client) DeletePlatformApplication(ctx context.Context, params *sns.DeletePlatformApplicationInput, optFns ...func(*sns.Options)) (*sns.DeletePlatformApplicationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.DeletePlatformApplicationOutput), args.Error(1)
}

// DeleteSMSSandboxPhoneNumber mock.
func (c *Client) DeleteSMSSandboxPhoneNumber(ctx context.Context, params *sns.DeleteSMSSandboxPhoneNumberInput, optFns ...func(*sns.Options)) (*sns.DeleteSMSSandboxPhoneNumberOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.DeleteSMSSandboxPhoneNumberOutput), args.Error(1)
}

// DeleteTopic mock.
func (c *Client) DeleteTopic(ctx context.Context, params *sns.DeleteTopicInput, optFns ...func(*sns.Options)) (*sns.DeleteTopicOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.DeleteTopicOutput), args.Error(1)
}

// GetEndpointAttributes mock.
func (c *Client) GetEndpointAttributes(ctx context.Context, params *sns.GetEndpointAttributesInput, optFns ...func(*sns.Options)) (*sns.GetEndpointAttributesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.GetEndpointAttributesOutput), args.Error(1)
}

// GetPlatformApplicationAttributes mock.
func (c *Client) GetPlatformApplicationAttributes(ctx context.Context, params *sns.GetPlatformApplicationAttributesInput, optFns ...func(*sns.Options)) (*sns.GetPlatformApplicationAttributesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.GetPlatformApplicationAttributesOutput), args.Error(1)
}

// GetSMSAttributes mock.
func (c *Client) GetSMSAttributes(ctx context.Context, params *sns.GetSMSAttributesInput, optFns ...func(*sns.Options)) (*sns.GetSMSAttributesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.GetSMSAttributesOutput), args.Error(1)
}

// GetSMSSandboxAccountStatus mock.
func (c *Client) GetSMSSandboxAccountStatus(ctx context.Context, params *sns.GetSMSSandboxAccountStatusInput, optFns ...func(*sns.Options)) (*sns.GetSMSSandboxAccountStatusOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.GetSMSSandboxAccountStatusOutput), args.Error(1)
}

// GetSubscriptionAttributes mock.
func (c *Client) GetSubscriptionAttributes(ctx context.Context, params *sns.GetSubscriptionAttributesInput, optFns ...func(*sns.Options)) (*sns.GetSubscriptionAttributesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.GetSubscriptionAttributesOutput), args.Error(1)
}

// GetTopicAttributes mock.
func (c *Client) GetTopicAttributes(ctx context.Context, params *sns.GetTopicAttributesInput, optFns ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.GetTopicAttributesOutput), args.Error(1)
}

// ListEndpointsByPlatformApplication mock.
func (c *Client) ListEndpointsByPlatformApplication(ctx context.Context, params *sns.ListEndpointsByPlatformApplicationInput, optFns ...func(*sns.Options)) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.ListEndpointsByPlatformApplicationOutput), args.Error(1)
}

// ListOriginationNumbers mock.
func (c *Client) ListOriginationNumbers(ctx context.Context, params *sns.ListOriginationNumbersInput, optFns ...func(*sns.Options)) (*sns.ListOriginationNumbersOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.ListOriginationNumbersOutput), args.Error(1)
}

// ListPhoneNumbersOptedOut mock.
func (c *Client) ListPhoneNumbersOptedOut(ctx context.Context, params *sns.ListPhoneNumbersOptedOutInput, optFns ...func(*sns.Options)) (*sns.ListPhoneNumbersOptedOutOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.ListPhoneNumbersOptedOutOutput), args.Error(1)
}

// ListPlatformApplications mock.
func (c *Client) ListPlatformApplications(ctx context.Context, params *sns.ListPlatformApplicationsInput, optFns ...func(*sns.Options)) (*sns.ListPlatformApplicationsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.ListPlatformApplicationsOutput), args.Error(1)
}

// ListSMSSandboxPhoneNumbers mock.
func (c *Client) ListSMSSandboxPhoneNumbers(ctx context.Context, params *sns.ListSMSSandboxPhoneNumbersInput, optFns ...func(*sns.Options)) (*sns.ListSMSSandboxPhoneNumbersOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.ListSMSSandboxPhoneNumbersOutput), args.Error(1)
}

// ListSubscriptions mock.
func (c *Client) ListSubscriptions(ctx context.Context, params *sns.ListSubscriptionsInput, optFns ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.ListSubscriptionsOutput), args.Error(1)
}

// ListSubscriptionsByTopic mock.
func (c *Client) ListSubscriptionsByTopic(ctx context.Context, params *sns.ListSubscriptionsByTopicInput, optFns ...func(*sns.Options)) (*sns.ListSubscriptionsByTopicOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.ListSubscriptionsByTopicOutput), args.Error(1)
}

// ListTagsForResource mock.
func (c *Client) ListTagsForResource(ctx context.Context, params *sns.ListTagsForResourceInput, optFns ...func(*sns.Options)) (*sns.ListTagsForResourceOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.ListTagsForResourceOutput), args.Error(1)
}

// ListTopics mock.
func (c *Client) ListTopics(ctx context.Context, params *sns.ListTopicsInput, optFns ...func(*sns.Options)) (*sns.ListTopicsOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.ListTopicsOutput), args.Error(1)
}

// OptInPhoneNumber mock.
func (c *Client) OptInPhoneNumber(ctx context.Context, params *sns.OptInPhoneNumberInput, optFns ...func(*sns.Options)) (*sns.OptInPhoneNumberOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.OptInPhoneNumberOutput), args.Error(1)
}

// Publish mock.
func (c *Client) Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.PublishOutput), args.Error(1)
}

// PublishBatch mock.
func (c *Client) PublishBatch(ctx context.Context, params *sns.PublishBatchInput, optFns ...func(*sns.Options)) (*sns.PublishBatchOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.PublishBatchOutput), args.Error(1)
}

// RemovePermission mock.
func (c *Client) RemovePermission(ctx context.Context, params *sns.RemovePermissionInput, optFns ...func(*sns.Options)) (*sns.RemovePermissionOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.RemovePermissionOutput), args.Error(1)
}

// SetEndpointAttributes mock.
func (c *Client) SetEndpointAttributes(ctx context.Context, params *sns.SetEndpointAttributesInput, optFns ...func(*sns.Options)) (*sns.SetEndpointAttributesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.SetEndpointAttributesOutput), args.Error(1)
}

// SetPlatformApplicationAttributes mock.
func (c *Client) SetPlatformApplicationAttributes(ctx context.Context, params *sns.SetPlatformApplicationAttributesInput, optFns ...func(*sns.Options)) (*sns.SetPlatformApplicationAttributesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.SetPlatformApplicationAttributesOutput), args.Error(1)
}

// SetSMSAttributes mock.
func (c *Client) SetSMSAttributes(ctx context.Context, params *sns.SetSMSAttributesInput, optFns ...func(*sns.Options)) (*sns.SetSMSAttributesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.SetSMSAttributesOutput), args.Error(1)
}

// SetSubscriptionAttributes mock.
func (c *Client) SetSubscriptionAttributes(ctx context.Context, params *sns.SetSubscriptionAttributesInput, optFns ...func(*sns.Options)) (*sns.SetSubscriptionAttributesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.SetSubscriptionAttributesOutput), args.Error(1)
}

// SetTopicAttributes mock.
func (c *Client) SetTopicAttributes(ctx context.Context, params *sns.SetTopicAttributesInput, optFns ...func(*sns.Options)) (*sns.SetTopicAttributesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.SetTopicAttributesOutput), args.Error(1)
}

// Subscribe mock.
func (c *Client) Subscribe(ctx context.Context, params *sns.SubscribeInput, optFns ...func(*sns.Options)) (*sns.SubscribeOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.SubscribeOutput), args.Error(1)
}

// TagResource mock.
func (c *Client) TagResource(ctx context.Context, params *sns.TagResourceInput, optFns ...func(*sns.Options)) (*sns.TagResourceOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.TagResourceOutput), args.Error(1)
}

// Unsubscribe mock.
func (c *Client) Unsubscribe(ctx context.Context, params *sns.UnsubscribeInput, optFns ...func(*sns.Options)) (*sns.UnsubscribeOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.UnsubscribeOutput), args.Error(1)
}

// UntagResource mock.
func (c *Client) UntagResource(ctx context.Context, params *sns.UntagResourceInput, optFns ...func(*sns.Options)) (*sns.UntagResourceOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.UntagResourceOutput), args.Error(1)
}

// VerifySMSSandboxPhoneNumber mock.
func (c *Client) VerifySMSSandboxPhoneNumber(ctx context.Context, params *sns.VerifySMSSandboxPhoneNumberInput, optFns ...func(*sns.Options)) (*sns.VerifySMSSandboxPhoneNumberOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.VerifySMSSandboxPhoneNumberOutput), args.Error(1)
}