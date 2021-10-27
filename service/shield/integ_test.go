// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package shield_test

import (
	"context"
	"testing"
	"time"

	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/awserr"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/awstesting/integration"
	"github.com/unicloud-uos/uos-sdk-go/service/shield"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_ListAttacks(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-east-1")
	svc := shield.New(sess)
	params := &shield.ListAttacksInput{}
	_, err := svc.ListAttacksWithContext(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
