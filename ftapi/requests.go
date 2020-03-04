package ftapi

import (
	"errors"
)

type (
	Request interface {
		Execute() error
	}
	CachedRequest interface {
		Request
		BypassCache(bypass bool) CachedRequest
		BypassCacheRead(bypass bool) CachedRequest
		BypassCacheWrite(bypass bool) CachedRequest
	}
)

func (req *RequestData) BypassCache(bypass bool) CachedRequest {
	req.CacheReadBypassed = bypass
	req.CacheWriteBypassed = bypass
	return req
}

func (req *RequestData) BypassCacheRead(bypass bool) CachedRequest {
	req.CacheReadBypassed = bypass
	return req
}

func (req *RequestData) BypassCacheWrite(bypass bool) CachedRequest {
	req.CacheWriteBypassed = bypass
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
