# AWS Go SDK v2 Mock Libraries

## Usage

### SQS Client

```go
// main.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type sqsClient interface {
	ListQueues(ctx context.Context, params *sqs.ListQueuesInput, optFns ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error)
}

func main() {
	cfg, _ := config.LoadDefaultConfig(context.Background())
	client := sqs.NewFromConfig(cfg)

	queues, _ := getQueues(context.Background(), client)
	fmt.Printf("%#v", queues)
}

func getQueues(ctx context.Context, c sqsClient) ([]string, error) {
	queues, err := c.ListQueues(ctx, &sqs.ListQueuesInput{})
	if err != nil {
		return nil, err
	}
	return queues.QueueUrls, nil
}
```

```go
// main_test.go
package main

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/gvre/awsmock-v2/sqsmock"
	"github.com/stretchr/testify/mock"
)

func TestGetQueues(t *testing.T) {
	queueURLs := []string{
		"https://sqs.us-west-2.amazonaws.com/1234567890/queue1",
		"https://sqs.us-west-2.amazonaws.com/1234567890/queue2",
	}
	client := new(sqsmock.Client)
	client.On("ListQueues", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.ListQueuesOutput{
		QueueUrls: queueURLs,
	}, nil)

	queues, _ := getQueues(context.Background(), client)

	if nqueues, nQueueURLs := len(queues), len(queueURLs); nqueues != nQueueURLs {
		t.Errorf("Invalid number of queues, got: %v, want: %v", nqueues, nQueueURLs)
	}
}

```

## License
- MIT