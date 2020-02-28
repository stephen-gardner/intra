package ftapi

import (
	"net/url"
)

type (
	params struct {
		url.Values
	}
	RequestParams interface {
		Add(key, value string) RequestParams
		Clear() RequestParams
		Del(key string) RequestParams
		Encode() string
		Get(key string) string
		Has(key string) bool
		Set(key, value string) RequestParams
	}
)

func (p *params) Add(key, value string) RequestParams {
	if p.Values == nil {
		p.Values = url.Values{}
	}
	p.Values.Add(key, value)
	return p
}

func (p *params) Clear() RequestParams {
	if p.Values != nil {
		p.Values = url.Values{}
	}
	return p
}

func (p *params) Del(key string) RequestParams {
	p.Values.Del(key)
	return p
}

func (p *params) Encode() string {
	return p.Values.Encode()
}

func (p *params) Get(key string) string {
	return p.Values.Get(key)
}

func (p *params) Has(key string) bool {
	if p.Values == nil {
		return false
	}
	_, present := p.Values[key]
	return present
}

func (p *params) Set(key, value string) RequestParams {
	if p.Values == nil {
		p.Values = url.Values{}
	}
	p.Values.Set(key, value)
	return p
}
