# S3 Client

## Usage

```go
// main.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Client interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

func main() {
	cfg, _ := config.LoadDefaultConfig(context.Background())
	client := s3.NewFromConfig(cfg)

	buckets, _ := getBuckets(context.Background(), client)
	fmt.Printf("%#v", buckets)
}

func getBuckets(ctx context.Context, c s3Client) ([]string, error) {
	buckets, err := c.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	var bucketNames []string
	for _, b := range buckets.Buckets {
		bucketNames = append(bucketNames, *b.Name)
	}
	return bucketNames, nil
}

```

```go
// main_test.go
package main

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gvre/awsmock-v2/s3mock"
	"github.com/stretchr/testify/mock"
)

func TestGetBuckets(t *testing.T) {
	bucketNames := []types.Bucket{
		{
			Name: new(string),
		},
	}
	client := new(s3mock.Client)
	client.On("ListBuckets", mock.Anything, mock.Anything, mock.Anything).Return(&s3.ListBucketsOutput{
		Buckets: bucketNames,
	}, nil)

	buckets, _ := getBuckets(context.Background(), client)

	if nbuckets, nBucketNames := len(buckets), len(bucketNames); nbuckets != nBucketNames {
		t.Errorf("Invalid number of buckets, got: %v, want: %v", nbuckets, nBucketNames)
	}
}

```