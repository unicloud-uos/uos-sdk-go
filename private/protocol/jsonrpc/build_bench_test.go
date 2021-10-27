// +build bench

package jsonrpc_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/awstesting"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/private/protocol/json/jsonutil"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/private/protocol/jsonrpc"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/service/dynamodb"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/service/dynamodb/dynamodbattribute"
)

func BenchmarkJSONRPCBuild_Simple_dynamodbPutItem(b *testing.B) {
	svc := awstesting.NewClient()

	params := getDynamodbPutItemParams(b)

	for i := 0; i < b.N; i++ {
		r := svc.NewRequest(&request.Operation{Name: "Operation"}, params, nil)
		jsonrpc.Build(r)
		if r.Error != nil {
			b.Fatal("Unexpected error", r.Error)
		}
	}
}

func BenchmarkJSONUtilBuild_Simple_dynamodbPutItem(b *testing.B) {
	svc := awstesting.NewClient()

	params := getDynamodbPutItemParams(b)

	for i := 0; i < b.N; i++ {
		r := svc.NewRequest(&request.Operation{Name: "Operation"}, params, nil)
		_, err := jsonutil.BuildJSON(r.Params)
		if err != nil {
			b.Fatal("Unexpected error", err)
		}
	}
}

func BenchmarkEncodingJSONMarshal_Simple_dynamodbPutItem(b *testing.B) {
	params := getDynamodbPutItemParams(b)

	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		encoder := json.NewEncoder(buf)
		if err := encoder.Encode(params); err != nil {
			b.Fatal("Unexpected error", err)
		}
	}
}

func getDynamodbPutItemParams(b *testing.B) *dynamodb.PutItemInput {
	av, err := dynamodbattribute.ConvertToMap(struct {
		Key  string
		Data string
	}{Key: "MyKey", Data: "MyData"})
	if err != nil {
		b.Fatal("benchPutItem, expect no ConvertToMap errors", err)
	}
	return &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("tablename"),
	}
}
