package csm

import (
	"net/http"
	"testing"

	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/client"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/client/metadata"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/credentials"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/defaults"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
)

func TestMaxRetriesExceeded(t *testing.T) {
	md := metadata.ClientInfo{
		Endpoint: "http://127.0.0.1",
	}

	cfg := aws.Config{
		Region:      aws.String("foo"),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	}

	op := &request.Operation{}
	cases := []struct {
		name                    string
		httpStatusCode          int
		expectedMaxRetriesValue int
		expectedMetrics         int
	}{
		{
			name:                    "max retry reached",
			httpStatusCode:          http.StatusBadGateway,
			expectedMaxRetriesValue: 1,
		},
		{
			name:                    "status ok",
			httpStatusCode:          http.StatusOK,
			expectedMaxRetriesValue: 0,
		},
	}

	for _, c := range cases {
		r := request.New(cfg, md, defaults.Handlers(), client.DefaultRetryer{NumMaxRetries: 2}, op, nil, nil)
		reporter := newReporter("", "")
		r.Handlers.Send.Clear()
		reporter.InjectHandlers(&r.Handlers)

		r.Handlers.Send.PushBack(func(r *request.Request) {
			r.HTTPResponse = &http.Response{
				StatusCode: c.httpStatusCode,
			}
		})
		r.Send()

		for {
			m := <-reporter.metricsCh.ch

			if *m.Type != "ApiCall" {
				// ignore non-ApiCall metrics since MaxRetriesExceeded is only on ApiCall events
				continue
			}

			if val := *m.MaxRetriesExceeded; val != c.expectedMaxRetriesValue {
				t.Errorf("%s: expected %d, but received %d", c.name, c.expectedMaxRetriesValue, val)
			}

			break
		}
	}
}
