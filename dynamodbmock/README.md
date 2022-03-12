# DynamoDB Client

## Usage

```go
// main.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamodbClient interface {
	ListTables(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error)
}

func main() {
	cfg, _ := config.LoadDefaultConfig(context.Background())
	client := dynamodb.NewFromConfig(cfg)

	tables, _ := getTables(context.Background(), client)
	fmt.Printf("%#v", tables)
}

func getTables(ctx context.Context, c dynamodbClient) ([]string, error) {
	tables, err := c.ListTables(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		return nil, err
	}
	return tables.TableNames, nil
}

```

```go
// main_test.go
package main

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gvre/awsmock-v2/dynamodbmock"
	"github.com/stretchr/testify/mock"
)

func TestGetTables(t *testing.T) {
	tableNames := []string{
		"t1",
		"t2",
	}
	client := new(dynamodbmock.Client)
	client.On("ListTables", mock.Anything, mock.Anything, mock.Anything).Return(&dynamodb.ListTablesOutput{
		TableNames: tableNames,
	}, nil)

	tables, _ := getTables(context.Background(), client)

	if ntables, nTableNames := len(tables), len(tableNames); ntables != nTableNames {
		t.Errorf("Invalid number of tables, got: %v, want: %v", ntables, nTableNames)
	}
}


```