package machinelearning_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/awstesting/unit"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/service/machinelearning"
)

func TestPredictEndpoint(t *testing.T) {
	ml := machinelearning.New(unit.Session)
	ml.Handlers.Send.Clear()
	ml.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
		}
	})

	req, _ := ml.PredictRequest(&machinelearning.PredictInput{
		PredictEndpoint: aws.String("https://localhost/endpoint"),
		MLModelId:       aws.String("id"),
		Record:          map[string]*string{},
	})
	err := req.Send()

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "https://localhost/endpoint", req.HTTPRequest.URL.String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}
