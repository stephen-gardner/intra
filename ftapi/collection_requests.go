package ftapi

import (
	"strconv"
	"strings"
)

type CollectionRequest interface {
	CachedRequest
	FilterBy(field string, values ...string) CollectionRequest
	GetParams() RequestParams
	RangeBy(field, minValue, maxValue string) CollectionRequest
	SetPageNumber(page int) CollectionRequest
	SetPageSize(size int) CollectionRequest
	SortBy(field string, desc bool) CollectionRequest
}

func (req *RequestData) FilterBy(field string, values ...string) CollectionRequest {
	key := "filter[" + field + "]"
	req.Params.Set(key, strings.Join(values, ","))
	return req
}

func (req *RequestData) GetParams() RequestParams {
	return &req.Params
}

func (req *RequestData) RangeBy(field, minValue, maxValue string) CollectionRequest {
	key := "range[" + field + "]"
	req.Params.Set(key, minValue+","+maxValue)
	return req
}

func (req *RequestData) SetPageNumber(page int) CollectionRequest {
	req.Params.Set("page[number]", strconv.Itoa(page))
	return req
}

func (req *RequestData) SetPageSize(size int) CollectionRequest {
	req.Params.Set("page[size]", strconv.Itoa(size))
	return req
}

func (req *RequestData) SortBy(field string, desc bool) CollectionRequest {
	if desc {
		field = "-" + field
	}
	if req.Params.Has("sort") {
		field = req.Params.Get("sort") + "," + field
	}
	req.Params.Set("sort", field)
	return req
}
