package ftapi

import (
	"errors"
)

type (
	Request interface {
		Execute() error
	}
	BasicRequest interface {
		Request
		BypassCache(bypass bool) BasicRequest
		BypassCacheRead(bypass bool) BasicRequest
		BypassCacheWrite(bypass bool) BasicRequest
	}
)

func (req *RequestData) BypassCache(bypass bool) BasicRequest {
	req.bypassCacheRead = bypass
	req.bypassCacheWrite = bypass
	return req
}

func (req *RequestData) BypassCacheRead(bypass bool) BasicRequest {
	req.bypassCacheRead = bypass
	return req
}

func (req *RequestData) BypassCacheWrite(bypass bool) BasicRequest {
	req.bypassCacheWrite = bypass
	return req
}

func (req *RequestData) Execute() error {
	if req.ExecuteMethod == nil {
		req.Error = errors.New("request method not specified")
		return req.Error
	}
	req.ExecuteMethod()
	return req.Error
}
