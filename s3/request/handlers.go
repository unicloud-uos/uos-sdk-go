package request

import (
	"fmt"
	"io"
	"runtime"
	"time"

	"github.com/uos-sdk-go/s3/credential"
	. "github.com/uos-sdk-go/s3/error"
	"github.com/uos-sdk-go/s3/helper"
)

type Handlers struct {
	Validate  HandlerList
	Set       HandlerList
	Send      HandlerList
	Sign      HandlerList
	Marshal   HandlerList
	Unmarshal HandlerList
}

func NewHandlers() Handlers {
	var handlers Handlers

	handlers.Validate.PushBackHandlerItem(ValidateEndpointHandler)
	handlers.Validate.PushBackHandlerItem(ValidateParametersHandler)
	handlers.Validate.AfterEachFn = StopHandlerListOnErr
	handlers.Set.PushBackHandlerItem(AgentInfoAndSDKVersionHandler)
	handlers.Set.PushBackHandlerItem(ContentLengthHandler)
	handlers.Set.AfterEachFn = StopHandlerListOnErr
	handlers.Send.PushBackHandlerItem(ValidateReqSigHandler)
	handlers.Send.PushBackHandlerItem(SendHandler)
	handlers.Send.AfterEachFn = StopHandlerListOnErr

	return handlers
}

func (h *Handlers) Copy() Handlers {
	return Handlers{
		Validate:  h.Validate.copy(),
		Set:       h.Set.copy(),
		Send:      h.Send.copy(),
		Sign:      h.Sign.copy(),
		Marshal:   h.Marshal.copy(),
		Unmarshal: h.Unmarshal.copy(),
	}
}

type HandlerList struct {
	list []HandlerItem

	// If Handler func return false, stop iterating
	AfterEachFn func(item HandlerRunItem) bool
}

// A running HandlerList entry
type HandlerRunItem struct {
	Index   int
	Handler HandlerItem
	Request *Request
}

// An entry used to load request func
type HandlerItem struct {
	Name string
	Fn   func(*Request)
}

func (l *HandlerList) copy() HandlerList {
	n := HandlerList{
		AfterEachFn: l.AfterEachFn,
	}
	if len(l.list) == 0 {
		return n
	}

	n.list = append(make([]HandlerItem, 0, len(l.list)), l.list...)
	return n
}

func (l *HandlerList) PushBackHandlerItem(n HandlerItem) {
	if cap(l.list) == 0 {
		l.list = make([]HandlerItem, 0, 5)
	}
	l.list = append(l.list, n)
}

func (l *HandlerList) PushFrontHandlerItem(n HandlerItem) {
	if cap(l.list) == len(l.list) {
		l.list = append([]HandlerItem{n}, l.list...)
	} else {
		l.list = append(l.list, HandlerItem{})
		copy(l.list[1:], l.list)
		l.list[0] = n
	}
}

func (l *HandlerList) RemoveHandlerItem(name string) {
	for i := 0; i < len(l.list); i++ {
		m := l.list[i]
		if m.Name == name {
			// Shift array preventing creating new arrays
			copy(l.list[i:], l.list[i+1:])
			l.list[len(l.list)-1] = HandlerItem{}
			l.list = l.list[:len(l.list)-1]

			// decrement list so next check to length is correct
			i--
		}
	}
}

func (l *HandlerList) Run(r *Request) {
	for i, h := range l.list {
		h.Fn(r)
		item := HandlerRunItem{
			Index: i, Handler: h, Request: r,
		}
		if l.AfterEachFn != nil && !l.AfterEachFn(item) {
			return
		}
	}
}

// Validate the request had endpoint or not
var ValidateEndpointHandler = HandlerItem{
	Name: "core.validate.endpoint.request",
	Fn: func(request *Request) {
		if request.Metadata.SigningRegion == "" && request.Config.Region == "" {
			request.Config.Logger.Debug("Validate Endpoint err: ", ErrMissingRegion)
			request.Error = ErrMissingRegion
		} else if len(request.Config.Endpoint) == 0 {
			request.Config.Logger.Debug("Validate Endpoint err: ", ErrMissingEndpoint)
			request.Error = ErrMissingEndpoint
		}
	},
}

// Validate the input parameters
var ValidateParametersHandler = HandlerItem{
	Name: "core.validate.parameters.request",
	Fn: func(request *Request) {
		if !request.IsParamsValid() {
			return
		}
		if v, ok := request.Params.(Validator); ok {
			err := v.Validate()
			if err != nil {
				request.Config.Logger.Debug("Validate parameter err: ", err)
				request.Error = err
			}
		}
	},
}

// Add agent info to request
var AgentInfoAndSDKVersionHandler = HandlerItem{
	Name: "core.set.agentInfo.request",
	Fn: func(request *Request) {
		agent := fmt.Sprintf("%s-%s; %s; %s; %s", helper.SDKName, helper.SDKVersion,
			runtime.Version(), runtime.GOOS, runtime.GOARCH)
		request.AddAgentInfo(agent)
	},
}

var ContentLengthHandler = HandlerItem{
	Name: "core.set.contentLength.request",
	Fn: func(request *Request) {
		var length int64
		switch body := request.Body.(type) {
		case nil:
			length = 0
		case io.Seeker:
			var err error
			request.BodyStart, err = body.Seek(0, io.SeekCurrent)
			if err != nil {
				request.Config.Logger.Debug("failed to determine start of the request body:", err)
				request.Error = NewBaseError("GetLengthERR", "failed to determine start of the request body:", err)
			}
			end, err := body.Seek(0, io.SeekEnd)
			if err != nil {
				request.Config.Logger.Debug("failed to determine end of the request body:", err)
				request.Error = NewBaseError("GetLengthERR", "failed to determine end of the request body:", err)
			}
			_, err = body.Seek(request.BodyStart, io.SeekStart) // make sure to seek back to original location
			if err != nil {
				request.Config.Logger.Debug("failed to seek back to the original location", err)
				request.Error = NewBaseError("GetLengthERR", "failed to seek back to the original location:", err)
			}
			length = end - request.BodyStart
		default:
			panic("Cannot get length of body, must provide `ContentLength`")
		}
		if length > 0 {
			request.HTTPRequest.ContentLength = length
			request.HTTPRequest.Header.Set("Content-Length", fmt.Sprintf("%d", length))
		} else {
			request.HTTPRequest.ContentLength = 0
			request.HTTPRequest.Header.Del("Content-Length")
		}
	},
}

var ContentMd5Handler = HandlerItem{
	Name: "core.set.contentMd5.request",
	Fn:   AddBodyContentMD5Handler,
}

var ValidateReqSigHandler = HandlerItem{
	Name: "core.send.sign.validate.request",
	Fn: func(request *Request) {
		if request.Config.Credentials == credential.DefaultCredentials {
			return
		}

		signedTime := request.Time
		if request.LastSignedTime.IsZero() {
			signedTime = request.LastSignedTime
		}

		if signedTime.Add(10 * time.Minute).After(time.Now()) {
			return
		}
		err := request.Sign()
		if err != nil {
			request.Config.Logger.Error("core.send.validateSign.request err: ", err)
		}
	},
}

var SendHandler = HandlerItem{
	Name: "core.send.request",
	Fn: func(request *Request) {
		request.HTTPResponse, request.Error = request.HTTPClient.Do(request.HTTPRequest)
		request.Config.Logger.Debug("Http response: ", request.HTTPResponse)

	},
}

func StopHandlerListOnErr(item HandlerRunItem) bool {
	return item.Request.Error == nil
}
