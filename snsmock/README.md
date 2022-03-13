# SNS Client

## Usage

```go
// main.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type snsClient interface {
	ListTopics(ctx context.Context, params *sns.ListTopicsInput, optFns ...func(*sns.Options)) (*sns.ListTopicsOutput, error)
}

func main() {
	cfg, _ := config.LoadDefaultConfig(context.Background())
	client := sns.NewFromConfig(cfg)

	topics, _ := getTopics(context.Background(), client)
	fmt.Printf("%#v", topics)
}

func getTopics(ctx context.Context, c snsClient) ([]string, error) {
	topics, err := c.ListTopics(ctx, &sns.ListTopicsInput{})
	if err != nil {
		return nil, err
	}

	var topicsARNs []string
	for _, t := range topics.Topics {
		topicsARNs = append(topicsARNs, *t.TopicArn)
	}
	return topicsARNs, nil
}
```

```go
// main_test.go
package main

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/gvre/awsmock-v2/snsmock"
	"github.com/stretchr/testify/mock"
)

func TestGetTopics(t *testing.T) {
	topicsARNs := []types.Topic{
		{
			TopicArn: new(string),
		},
	}
	client := new(snsmock.Client)
	client.On("ListTopics", mock.Anything, mock.Anything, mock.Anything).Return(&sns.ListTopicsOutput{
		Topics: topicsARNs,
	}, nil)

	topics, _ := getTopics(context.Background(), client)

	if ntopics, nTopicsARNs := len(topics), len(topicsARNs); ntopics != nTopicsARNs {
		t.Errorf("Invalid number of topics, got: %v, want: %v", ntopics, nTopicsARNs)
	}
}
```