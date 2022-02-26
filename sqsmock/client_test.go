package sqsmock_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/stretchr/testify/mock"

	"github.com/gvre/awsmock-v2/sqsmock"
)

type sqsClient interface {
	AddPermission(ctx context.Context, params *sqs.AddPermissionInput, optFns ...func(*sqs.Options)) (*sqs.AddPermissionOutput, error)
	ChangeMessageVisibility(ctx context.Context, params *sqs.ChangeMessageVisibilityInput, optFns ...func(*sqs.Options)) (*sqs.ChangeMessageVisibilityOutput, error)
	ChangeMessageVisibilityBatch(ctx context.Context, params *sqs.ChangeMessageVisibilityBatchInput, optFns ...func(*sqs.Options)) (*sqs.ChangeMessageVisibilityBatchOutput, error)
	CreateQueue(ctx context.Context, params *sqs.CreateQueueInput, optFns ...func(*sqs.Options)) (*sqs.CreateQueueOutput, error)
	DeleteMessage(ctx context.Context, params *sqs.DeleteMessageInput, optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error)
	DeleteMessageBatch(ctx context.Context, params *sqs.DeleteMessageBatchInput, optFns ...func(*sqs.Options)) (*sqs.DeleteMessageBatchOutput, error)
	DeleteQueue(ctx context.Context, params *sqs.DeleteQueueInput, optFns ...func(*sqs.Options)) (*sqs.DeleteQueueOutput, error)
	GetQueueAttributes(ctx context.Context, params *sqs.GetQueueAttributesInput, optFns ...func(*sqs.Options)) (*sqs.GetQueueAttributesOutput, error)
	GetQueueUrl(ctx context.Context, params *sqs.GetQueueUrlInput, optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)
	ListDeadLetterSourceQueues(ctx context.Context, params *sqs.ListDeadLetterSourceQueuesInput, optFns ...func(*sqs.Options)) (*sqs.ListDeadLetterSourceQueuesOutput, error)
	ListQueueTags(ctx context.Context, params *sqs.ListQueueTagsInput, optFns ...func(*sqs.Options)) (*sqs.ListQueueTagsOutput, error)
	ListQueues(ctx context.Context, params *sqs.ListQueuesInput, optFns ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error)
	PurgeQueue(ctx context.Context, params *sqs.PurgeQueueInput, optFns ...func(*sqs.Options)) (*sqs.PurgeQueueOutput, error)
	ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)
	RemovePermission(ctx context.Context, params *sqs.RemovePermissionInput, optFns ...func(*sqs.Options)) (*sqs.RemovePermissionOutput, error)
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
	SendMessageBatch(ctx context.Context, params *sqs.SendMessageBatchInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageBatchOutput, error)
	SetQueueAttributes(ctx context.Context, params *sqs.SetQueueAttributesInput, optFns ...func(*sqs.Options)) (*sqs.SetQueueAttributesOutput, error)
	TagQueue(ctx context.Context, params *sqs.TagQueueInput, optFns ...func(*sqs.Options)) (*sqs.TagQueueOutput, error)
	UntagQueue(ctx context.Context, params *sqs.UntagQueueInput, optFns ...func(*sqs.Options)) (*sqs.UntagQueueOutput, error)
}

func TestAddPermission(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("AddPermission", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.AddPermissionOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.AddPermission(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestChangeMessageVisibility(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("ChangeMessageVisibility", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.ChangeMessageVisibilityOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.ChangeMessageVisibility(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestChangeMessageVisibilityBatch(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("ChangeMessageVisibilityBatch", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.ChangeMessageVisibilityBatchOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.ChangeMessageVisibilityBatch(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestCreateQueue(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("CreateQueue", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.CreateQueueOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.CreateQueue(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteMessage(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("DeleteMessage", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.DeleteMessageOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.DeleteMessage(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteMessageBatch(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("DeleteMessageBatch", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.DeleteMessageBatchOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.DeleteMessageBatch(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestDeleteQueue(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("DeleteQueue", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.DeleteQueueOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.DeleteQueue(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetQueueAttributes(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("GetQueueAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.GetQueueAttributesOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.GetQueueAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestGetQueueUrl(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("GetQueueUrl", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.GetQueueUrlOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.GetQueueUrl(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListDeadLetterSourceQueues(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("ListDeadLetterSourceQueues", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.ListDeadLetterSourceQueuesOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.ListDeadLetterSourceQueues(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListQueueTags(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("ListQueueTags", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.ListQueueTagsOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.ListQueueTags(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestListQueues(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("ListQueues", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.ListQueuesOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.ListQueues(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestPurgeQueue(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("PurgeQueue", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.PurgeQueueOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.PurgeQueue(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestReceiveMessage(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("ReceiveMessage", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.ReceiveMessageOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.ReceiveMessage(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestRemovePermission(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("RemovePermission", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.RemovePermissionOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.RemovePermission(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestSendMessage(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("SendMessage", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.SendMessageOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.SendMessage(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestSendMessageBatch(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("SendMessageBatch", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.SendMessageBatchOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.SendMessageBatch(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestSetQueueAttributes(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("SetQueueAttributes", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.SetQueueAttributesOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.SetQueueAttributes(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestTagQueue(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("TagQueue", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.TagQueueOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.TagQueue(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}

func TestUntagQueue(t *testing.T) {
	client := new(sqsmock.Client)
	client.On("UntagQueue", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.UntagQueueOutput{}, nil)
	fn := func(c sqsClient) {
		_, _ = c.UntagQueue(context.Background(), nil, nil)
	}

	fn(client)

	client.AssertExpectations(t)
}
