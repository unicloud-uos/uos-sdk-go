package handler

import (
	"fmt"
	"io"
	"runtime"

	"github.com/uos-sdk-go/s3/auxiliary"
	. "github.com/uos-sdk-go/s3/error"
	. "github.com/uos-sdk-go/s3/request"
)

type Handlers struct {
	Build            HandlerList
	Validate         HandlerList
	ValidateResponse HandlerList
	Set              HandlerList
	Send             HandlerList
	Sign             HandlerList
	Unmarshal        HandlerList
}

func NewHandlers() Handlers {
	var handlers Handlers

	handlers.Validate.PushBackHandlerItem(ValidateEndpointHandler)
	handlers.Validate.PushBackHandlerItem(ValidateParametersHandler)
	handlers.Validate.AfterFn = StopHandlerListOnErr
	handlers.Set.PushBackHandlerItem(RequestIDForTrackingHandler)
	handlers.Set.PushBackHandlerItem(AgentInfoAndSDKVersionHandler)
	handlers.Set.AfterFn = StopHandlerListOnErr

}

func (h *Handlers) Copy() Handlers {
	return Handlers{
		Build:            h.Build.copy(),
		Validate:         h.Validate.copy(),
		ValidateResponse: h.ValidateResponse.copy(),
		Set:              h.Set.copy(),
		Send:             h.Send.copy(),
		Sign:             h.Sign.copy(),
		Unmarshal:        h.Unmarshal.copy(),
	}
}

type HandlerList struct {
	list []HandlerItem

	// If Handler func return err, stop iterating
	AfterFn func(item HandlerRunItem) bool
}

// A running HandlerList entry
type HandlerRunItem struct {
	Index   int
	Handler HandlerItem
	Request *Request
}

// An entry used to load handler func
type HandlerItem struct {
	Name string
	Fn   func(*Request)
}

func (l *HandlerList) copy() HandlerList {
	n := HandlerList{
		AfterFn: l.AfterFn,
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

func (l *HandlerList) Run(r *Request) {
	for i, h := range l.list {
		h.Fn(r)
		item := HandlerRunItem{
			Index: i, Handler: h, Request: r,
		}
		if l.AfterFn != nil && !l.AfterFn(item) {
			return
		}
	}
}

// Validate the request had endpoint or not
var ValidateEndpointHandler = HandlerItem{
	Name: "core.validate.endpoint.handler",
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
	Name: "core.validate.parameters.handler",
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
	Name: "core.set.agentInfo.handler",
	Fn: func(request *Request) {
		agent := fmt.Sprintf("%s-%s; %s; %s; %s", auxiliary.SDKName, auxiliary.SDKVersion,
			runtime.Version(), runtime.GOOS, runtime.GOARCH)
		request.AddAgentInfo(agent)
	},
}

// Add ID for tracking
var RequestIDForTrackingHandler = HandlerItem{
	Name: "core.set.requestID.handler",
	Fn: func(request *Request) {
		request.HTTPRequest.Header.Set(auxiliary.UOSRequestID, request.ID)
	},
}

var ContentLengthHandler = HandlerItem{
	Name: "core.set.contentLength.handler",
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
				request.Error = err
			}
			end, err := body.Seek(0, io.SeekEnd)
			if err != nil {
				request.Config.Logger.Debug("failed to determine end of the request body:", err)
				request.Error = err
			}
			_, err = body.Seek(request.BodyStart, io.SeekStart) // make sure to seek back to original location
			if err != nil {
				request.Config.Logger.Debug("failed to seek back to the original location", err)
				request.Error = err
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

func StopHandlerListOnErr(item HandlerRunItem) bool {
	return item.Request.Error == nil
}
