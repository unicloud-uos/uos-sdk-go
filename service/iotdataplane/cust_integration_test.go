// +build integration

package iotdataplane_test

import (
	"testing"

	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/awserr"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/awstesting/integration"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/service/iot"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/service/iotdataplane"
)

func TestInteg_DescribeEndpoint(t *testing.T) {
	sess := integration.Session.Copy()
	if v := aws.StringValue(sess.Config.Region); len(v) == 0 {
		sess.Config.Region = aws.String("us-east-1")
	}

	ctrlSvc := iot.New(sess)
	descResp, err := ctrlSvc.DescribeEndpoint(&iot.DescribeEndpointInput{})
	if err != nil {
		t.Fatalf("failed to get dataplane endpoint, %v", err)
	}

	dataSvc := iotdataplane.New(sess, &aws.Config{
		Endpoint: descResp.EndpointAddress,
	})
	_, err = dataSvc.GetThingShadow(&iotdataplane.GetThingShadowInput{
		ThingName: aws.String("fake-thing"),
	})
	if err == nil {
		t.Fatalf("expect error")
	}

	aerr, ok := err.(awserr.Error)
	if !ok {
		t.Fatalf("expect awserr.Error, got %T, %v", err, err)
	}
	if e, a := "ResourceNotFoundException", aerr.Code(); e != a {
		t.Errorf("expect %v error, got %v", e, aerr)
	}
}
