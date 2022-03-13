package snsmock_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/stretchr/testify/mock"

	"github.com/gvre/awsmock-v2/snsmock"
)

type snsClient interface {
	AddPermission(ctx context.Context, params *sns.AddPermissionInput, optFns ...func(*sns.Options)) (*sns.AddPermissionOutput, error)
	CheckIfPhoneNumberIsOptedOut(ctx context.Context, params *sns.CheckIfPhoneNumberIsOptedOutInput, optFns ...func(*sns.Options)) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error)
	ConfirmSubscription(ctx context.Context, params *sns.ConfirmSubscriptionInput, optFns ...func(*sns.Options)) (*sns.ConfirmSubscriptionOutput, error)
	CreatePlatformApplication(ctx context.Context, params *sns.CreatePlatformApplicationInput, optFns ...func(*sns.Options)) (*sns.CreatePlatformApplicationOutput, error)
	CreatePlatformEndpoint(ctx context.Context, params *sns.CreatePlatformEndpointInput, optFns ...func(*sns.Options)) (*sns.CreatePlatformEndpointOutput, error)
	CreateSMSSandboxPhoneNumber(ctx context.Context, params *sns.CreateSMSSandboxPhoneNumberInput, optFns ...func(*sns.Options)) (*sns.CreateSMSSandboxPhoneNumberOutput, error)
	CreateTopic(ctx context.Context, params *sns.CreateTopicInput, optFns ...func(*sns.Options)) (*sns.CreateTopicOutput, error)
	DeleteEndpoint(ctx context.Context, params *sns.DeleteEndpointInput, optFns ...func(*sns.Options)) (*sns.DeleteEndpointOutput, error)
	DeletePlatformApplication(ctx context.Context, params *sns.DeletePlatformApplicationInput, optFns ...func(*sns.Options)) (*sns.DeletePlatformApplicationOutput, error)
	DeleteSMSSandboxPhoneNumber(ctx context.Context, params *sns.DeleteSMSSandboxPhoneNumberInput, optFns ...func(*sns.Options)) (*sns.DeleteSMSSandboxPhoneNumberOutput, error)
	DeleteTopic(ctx context.Context, params *sns.DeleteTopicInput, optFns ...func(*sns.Options)) (*sns.DeleteTopicOutput, error)
	GetEndpointAttributes(ctx context.Context, params *sns.GetEndpointAttributesInput, optFns ...func(*sns.Options)) (*sns.GetEndpointAttributesOutput, error)
	GetPlatformApplicationAttributes(ctx context.Context, params *sns.GetPlatformApplicationAttributesInput, optFns ...func(*sns.Options)) (*sns.GetPlatformApplicationAttributesOutput, error)
	GetSMSAttributes(ctx context.Context, params *sns.GetSMSAttributesInput, optFns ...func(*sns.Options)) (*sns.GetSMSAttributesOutput, error)
	GetSMSSandboxAccountStatus(ctx context.Context, params *sns.GetSMSSandboxAccountStatusInput, optFns ...func(*sns.Options)) (*sns.GetSMSSandboxAccountStatusOutput, error)
	GetSubscriptionAttributes(ctx context.Context, params *sns.GetSubscriptionAttributesInput, optFns ...func(*sns.Options)) (*sns.GetSubscriptionAttributesOutput, error)
	GetTopicAttributes(ctx context.Context, params *sns.GetTopicAttributesInput, optFns ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error)
	ListEndpointsByPlatformApplication(ctx context.Context, params *sns.ListEndpointsByPlatformApplicationInput, optFns ...func(*sns.Options)) (*sns.ListEndpointsByPlatformApplicationOutput, error)
	ListOriginationNumbers(ctx context.Context, params *sns.ListOriginationNumbersInput, optFns ...func(*sns.Options)) (*sns.ListOriginationNumbersOutput, error)
	ListPhoneNumbersOptedOut(ctx context.Context, params *sns.ListPhoneNumbersOptedOutInput, optFns ...func(*sns.Options)) (*sns.ListPhoneNumbersOptedOutOutput, error)
	ListPlatformApplications(ctx context.Context, params *sns.ListPlatformApplicationsInput, optFns ...func(*sns.Options)) (*sns.ListPlatformApplicationsOutput, error)
	ListSMSSandboxPhoneNumbers(ctx context.Context, params *sns.ListSMSSandboxPhoneNumbersInput, optFns ...func(*sns.Options)) (*sns.ListSMSSandboxPhoneNumbersOutput, error)
	ListSubscriptions(ctx context.Context, params *sns.ListSubscriptionsInput, optFns ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error)
	ListSubscriptionsByTopic(ctx context.Context, params *sns.ListSubscriptionsByTopicInput, optFns ...func(*sns.Options)) (*sns.ListSubscriptionsByTopicOutput, error)
	ListTagsForResource(ctx context.Context, params *sns.ListTagsForResourceInput, optFns ...func(*sns.Options)) (*sns.ListTagsForResourceOutput, error)
	ListTopics(ctx context.Context, params *sns.ListTopicsInput, optFns ...func(*sns.Options)) (*sns.ListTopicsOutput, error)
	OptInPhoneNumber(ctx context.Context, params *sns.OptInPhoneNumberInput, optFns ...func(*sns.Options)) (*sns.OptInPhoneNumberOutput, error)
	Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
	PublishBatch(ctx context.Context, params *sns.PublishBatchInput, optFns ...func(*sns.Options)) (*sns.PublishBatchOutput, error)
	RemovePermission(ctx context.Context, params *sns.RemovePermissionInput, optFns ...func(*sns.Options)) (*sns.RemovePermissionOutput, error)
	SetEndpointAttributes(ctx context.Context, params *sns.SetEndpointAttributesInput, optFns ...func(*sns.Options)) (*sns.SetEndpointAttributesOutput, error)
	SetPlatformApplicationAttributes(ctx context.Context, params *sns.SetPlatformApplicationAttributesInput, optFns ...func(*sns.Options)) (*sns.SetPlatformApplicationAttributesOutput, error)
	SetSMSAttributes(ctx context.Context, params *sns.SetSMSAttributesInput, optFns ...func(*sns.Options)) (*sns.SetSMSAttributesOutput, error)
	SetSubscriptionAttributes(ctx context.Context, params *sns.SetSubscriptionAttributesInput, optFns ...func(*sns.Options)) (*sns.SetSubscriptionAttributesOutput, error)
	SetTopicAttributes(ctx context.Context, params *sns.SetTopicAttributesInput, optFns ...func(*sns.Options)) (*sns.SetTopicAttributesOutput, error)
	Subscribe(ctx context.Context, params *sns.SubscribeInput, optFns ...func(*sns.Options)) (*sns.SubscribeOutput, error)
	TagResource(ctx context.Context, params *sns.TagResourceInput, optFns ...func(*sns.Options)) (*sns.TagResourceOutput, error)
	Unsubscribe(ctx context.Context, params *sns.UnsubscribeInput, optFns ...func(*sns.Options)) (*sns.UnsubscribeOutput, error)
	UntagResource(ctx context.Context, params *sns.UntagResourceInput, optFns ...func(*sns.Options)) (*sns.UntagResourceOutput, error)
	VerifySMSSandboxPhoneNumber(ctx context.Context, params *sns.VerifySMSSandboxPhoneNumberInput, optFns ...func(*sns.Options)) (*sns.VerifySMSSandboxPhoneNumberOutput, error)
}

func TestAddPermission(t *testing.T) {
	client := new(snsmock.Client)
	client.On("AddPermission", mock.Anything, mock.Anything, mock.Anything).Return(&sns.AddPermissionOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.AddPermission(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCheckIfPhoneNumberIsOptedOut(t *testing.T) {
	client := new(snsmock.Client)
	client.On("CheckIfPhoneNumberIsOptedOut", mock.Anything, mock.Anything, mock.Anything).Return(&sns.CheckIfPhoneNumberIsOptedOutOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.CheckIfPhoneNumberIsOptedOut(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestConfirmSubscription(t *testing.T) {
	client := new(snsmock.Client)
	client.On("ConfirmSubscription", mock.Anything, mock.Anything, mock.Anything).Return(&sns.ConfirmSubscriptionOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.ConfirmSubscription(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCreatePlatformApplication(t *testing.T) {
	client := new(snsmock.Client)
	client.On("CreatePlatformApplication", mock.Anything, mock.Anything, mock.Anything).Return(&sns.CreatePlatformApplicationOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.CreatePlatformApplication(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCreatePlatformEndpoint(t *testing.T) {
	client := new(snsmock.Client)
	client.On("CreatePlatformEndpoint", mock.Anything, mock.Anything, mock.Anything).Return(&sns.CreatePlatformEndpointOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.CreatePlatformEndpoint(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCreateSMSSandboxPhoneNumber(t *testing.T) {
	client := new(snsmock.Client)
	client.On("CreateSMSSandboxPhoneNumber", mock.Anything, mock.Anything, mock.Anything).Return(&sns.CreateSMSSandboxPhoneNumberOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.CreateSMSSandboxPhoneNumber(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCreateTopic(t *testing.T) {
	client := new(snsmock.Client)
	client.On("CreateTopic", mock.Anything, mock.Anything, mock.Anything).Return(&sns.CreateTopicOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.CreateTopic(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteEndpoint(t *testing.T) {
	client := new(snsmock.Client)
	client.On("DeleteEndpoint", mock.Anything, mock.Anything, mock.Anything).Return(&sns.DeleteEndpointOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.DeleteEndpoint(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeletePlatformApplication(t *testing.T) {
	client := new(snsmock.Client)
	client.On("DeletePlatformApplication", mock.Anything, mock.Anything, mock.Anything).Return(&sns.DeletePlatformApplicationOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.DeletePlatformApplication(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteSMSSandboxPhoneNumber(t *testing.T) {
	client := new(snsmock.Client)
	client.On("DeleteSMSSandboxPhoneNumber", mock.Anything, mock.Anything, mock.Anything).Return(&sns.DeleteSMSSandboxPhoneNumberOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.DeleteSMSSandboxPhoneNumber(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteTopic(t *testing.T) {
	client := new(snsmock.Client)
	client.On("DeleteTopic", mock.Anything, mock.Anything, mock.Anything).Return(&sns.DeleteTopicOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.DeleteTopic(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetEndpointAttributes(t *testing.T) {
	client := new(snsmock.Client)
	client.On("GetEndpointAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sns.GetEndpointAttributesOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.GetEndpointAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetPlatformApplicationAttributes(t *testing.T) {
	client := new(snsmock.Client)
	client.On("GetPlatformApplicationAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sns.GetPlatformApplicationAttributesOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.GetPlatformApplicationAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetSMSAttributes(t *testing.T) {
	client := new(snsmock.Client)
	client.On("GetSMSAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sns.GetSMSAttributesOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.GetSMSAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetSMSSandboxAccountStatus(t *testing.T) {
	client := new(snsmock.Client)
	client.On("GetSMSSandboxAccountStatus", mock.Anything, mock.Anything, mock.Anything).Return(&sns.GetSMSSandboxAccountStatusOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.GetSMSSandboxAccountStatus(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetSubscriptionAttributes(t *testing.T) {
	client := new(snsmock.Client)
	client.On("GetSubscriptionAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sns.GetSubscriptionAttributesOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.GetSubscriptionAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetTopicAttributes(t *testing.T) {
	client := new(snsmock.Client)
	client.On("GetTopicAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sns.GetTopicAttributesOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.GetTopicAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListEndpointsByPlatformApplication(t *testing.T) {
	client := new(snsmock.Client)
	client.On("ListEndpointsByPlatformApplication", mock.Anything, mock.Anything, mock.Anything).Return(&sns.ListEndpointsByPlatformApplicationOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.ListEndpointsByPlatformApplication(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListOriginationNumbers(t *testing.T) {
	client := new(snsmock.Client)
	client.On("ListOriginationNumbers", mock.Anything, mock.Anything, mock.Anything).Return(&sns.ListOriginationNumbersOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.ListOriginationNumbers(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListPhoneNumbersOptedOut(t *testing.T) {
	client := new(snsmock.Client)
	client.On("ListPhoneNumbersOptedOut", mock.Anything, mock.Anything, mock.Anything).Return(&sns.ListPhoneNumbersOptedOutOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.ListPhoneNumbersOptedOut(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListPlatformApplications(t *testing.T) {
	client := new(snsmock.Client)
	client.On("ListPlatformApplications", mock.Anything, mock.Anything, mock.Anything).Return(&sns.ListPlatformApplicationsOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.ListPlatformApplications(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListSMSSandboxPhoneNumbers(t *testing.T) {
	client := new(snsmock.Client)
	client.On("ListSMSSandboxPhoneNumbers", mock.Anything, mock.Anything, mock.Anything).Return(&sns.ListSMSSandboxPhoneNumbersOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.ListSMSSandboxPhoneNumbers(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListSubscriptions(t *testing.T) {
	client := new(snsmock.Client)
	client.On("ListSubscriptions", mock.Anything, mock.Anything, mock.Anything).Return(&sns.ListSubscriptionsOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.ListSubscriptions(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListSubscriptionsByTopic(t *testing.T) {
	client := new(snsmock.Client)
	client.On("ListSubscriptionsByTopic", mock.Anything, mock.Anything, mock.Anything).Return(&sns.ListSubscriptionsByTopicOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.ListSubscriptionsByTopic(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListTagsForResource(t *testing.T) {
	client := new(snsmock.Client)
	client.On("ListTagsForResource", mock.Anything, mock.Anything, mock.Anything).Return(&sns.ListTagsForResourceOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.ListTagsForResource(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListTopics(t *testing.T) {
	client := new(snsmock.Client)
	client.On("ListTopics", mock.Anything, mock.Anything, mock.Anything).Return(&sns.ListTopicsOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.ListTopics(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestOptInPhoneNumber(t *testing.T) {
	client := new(snsmock.Client)
	client.On("OptInPhoneNumber", mock.Anything, mock.Anything, mock.Anything).Return(&sns.OptInPhoneNumberOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.OptInPhoneNumber(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPublish(t *testing.T) {
	client := new(snsmock.Client)
	client.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(&sns.PublishOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.Publish(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPublishBatch(t *testing.T) {
	client := new(snsmock.Client)
	client.On("PublishBatch", mock.Anything, mock.Anything, mock.Anything).Return(&sns.PublishBatchOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.PublishBatch(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestRemovePermission(t *testing.T) {
	client := new(snsmock.Client)
	client.On("RemovePermission", mock.Anything, mock.Anything, mock.Anything).Return(&sns.RemovePermissionOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.RemovePermission(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestSetEndpointAttributes(t *testing.T) {
	client := new(snsmock.Client)
	client.On("SetEndpointAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sns.SetEndpointAttributesOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.SetEndpointAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestSetPlatformApplicationAttributes(t *testing.T) {
	client := new(snsmock.Client)
	client.On("SetPlatformApplicationAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sns.SetPlatformApplicationAttributesOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.SetPlatformApplicationAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestSetSMSAttributes(t *testing.T) {
	client := new(snsmock.Client)
	client.On("SetSMSAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sns.SetSMSAttributesOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.SetSMSAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestSetSubscriptionAttributes(t *testing.T) {
	client := new(snsmock.Client)
	client.On("SetSubscriptionAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sns.SetSubscriptionAttributesOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.SetSubscriptionAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestSetTopicAttributes(t *testing.T) {
	client := new(snsmock.Client)
	client.On("SetTopicAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sns.SetTopicAttributesOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.SetTopicAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestSubscribe(t *testing.T) {
	client := new(snsmock.Client)
	client.On("Subscribe", mock.Anything, mock.Anything, mock.Anything).Return(&sns.SubscribeOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.Subscribe(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestTagResource(t *testing.T) {
	client := new(snsmock.Client)
	client.On("TagResource", mock.Anything, mock.Anything, mock.Anything).Return(&sns.TagResourceOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.TagResource(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUnsubscribe(t *testing.T) {
	client := new(snsmock.Client)
	client.On("Unsubscribe", mock.Anything, mock.Anything, mock.Anything).Return(&sns.UnsubscribeOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.Unsubscribe(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUntagResource(t *testing.T) {
	client := new(snsmock.Client)
	client.On("UntagResource", mock.Anything, mock.Anything, mock.Anything).Return(&sns.UntagResourceOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.UntagResource(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestVerifySMSSandboxPhoneNumber(t *testing.T) {
	client := new(snsmock.Client)
	client.On("VerifySMSSandboxPhoneNumber", mock.Anything, mock.Anything, mock.Anything).Return(&sns.VerifySMSSandboxPhoneNumberOutput{}, nil)
	fn := func(c snsClient) {
		_, _ = c.VerifySMSSandboxPhoneNumber(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}
