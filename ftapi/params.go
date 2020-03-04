package ftapi

import (
	"net/url"
)

type (
	Params struct {
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

func (p *Params) Add(key, value string) RequestParams {
	if p.Values == nil {
		p.Values = url.Values{}
	}
	p.Values.Add(key, value)
	return p
}

func (p *Params) Clear() RequestParams {
	if p.Values != nil {
		p.Values = url.Values{}
	}
	return p
}

func (p *Params) Del(key string) RequestParams {
	p.Values.Del(key)
	return p
}

func (p *Params) Encode() string {
	return p.Values.Encode()
}

func (p *Params) Get(key string) string {
	return p.Values.Get(key)
}

func (p *Params) Has(key string) bool {
	if p.Values == nil {
		return false
	}
	_, present := p.Values[key]
	return present
}

func (p *Params) Set(key, value string) RequestParams {
	if p.Values == nil {
		p.Values = url.Values{}
	}
	p.Values.Set(key, value)
	return p
}
