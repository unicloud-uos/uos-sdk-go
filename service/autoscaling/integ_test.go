// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package autoscaling_test

import (
	"context"
	"testing"
	"time"

	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/awserr"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/awstesting/integration"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/service/autoscaling"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_DescribeScalingProcessTypes(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := autoscaling.New(sess)
	params := &autoscaling.DescribeScalingProcessTypesInput{}
	_, err := svc.DescribeScalingProcessTypesWithContext(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_CreateLaunchConfiguration(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := autoscaling.New(sess)
	params := &autoscaling.CreateLaunchConfigurationInput{
		ImageId:                 aws.String("ami-12345678"),
		InstanceType:            aws.String("m1.small"),
		LaunchConfigurationName: aws.String("hello, world"),
	}
	_, err := svc.CreateLaunchConfigurationWithContext(ctx, params)
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if len(aerr.Code()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if v := aerr.Code(); v == request.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}