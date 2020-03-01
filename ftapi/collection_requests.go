package ftapi

import (
	"strconv"
	"strings"
)

type CollectionRequest interface {
	CachedRequest
	FilterBy(field string, values ...string) CollectionRequest
	Params() RequestParams
	RangeBy(field, minValue, maxValue string) CollectionRequest
	SetPageNumber(page int) CollectionRequest
	SetPageSize(size int) CollectionRequest
	SortBy(field string, desc bool) CollectionRequest
}

func (req *RequestData) FilterBy(field string, values ...string) CollectionRequest {
	key := "filter[" + field + "]"
	req.params.Set(key, strings.Join(values, ","))
	return req
}

func (req *RequestData) Params() RequestParams {
	return &req.params
}

func (req *RequestData) RangeBy(field, minValue, maxValue string) CollectionRequest {
	key := "range[" + field + "]"
	req.params.Set(key, minValue+","+maxValue)
	return req
}

func (req *RequestData) SetPageNumber(page int) CollectionRequest {
	req.params.Set("page[number]", strconv.Itoa(page))
	return req
}

func (req *RequestData) SetPageSize(size int) CollectionRequest {
	req.params.Set("page[size]", strconv.Itoa(size))
	return req
}

func (req *RequestData) SortBy(field string, desc bool) CollectionRequest {
	if desc {
		field = "-" + field
	}
	if req.params.Has("sort") {
		field = req.params.Get("sort") + "," + field
	}
	req.params.Set("sort", field)
	return req
}
