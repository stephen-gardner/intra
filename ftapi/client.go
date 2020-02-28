package ftapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2/clientcredentials"
)

type (
	RequestData struct {
		ExecuteMethod    func()
		Endpoint         string
		Error            error
		contentType      string
		body             []byte
		params           params
		bypassCacheRead  bool
		bypassCacheWrite bool
	}
	Time struct {
		time.Time
	}
)

const (
	ContentTypeForm = "application/x-www-form-urlencoded"
	ContentTypeJson = "application/json"
	TimeFormat      = "2006-01-02T15:04:05.000Z"
)

var (
	clientID     = os.Getenv("INTRA_CLIENT_ID")
	clientSecret = os.Getenv("INTRA_CLIENT_SECRET")
	rateLimit    = time.NewTicker(time.Second / 2)
)

func (it *Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", it.UTC().Format(TimeFormat))), nil
}

func (it *Time) UnmarshalJSON(data []byte) error {
	raw := strings.Trim(string(data), "\"")
	if raw == "null" {
		it.Time = time.Time{}
		return nil
	}
	date, err := time.Parse(TimeFormat, raw)
	if err == nil {
		it.Time = date
	}
	return err
}

func GetClient(ctx context.Context, scopes ...string) *http.Client {
	oauth := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://api.intra.42.fr/oauth/token",
		Scopes:       scopes,
	}
	if ctx == nil {
		ctx = context.Background()
	}
	client := oauth.Client(ctx)
	client.Timeout = 2 * time.Minute
	return client
}

func GetEndpoint(path string, params url.Values) string {
	baseURL, _ := url.Parse("https://api.intra.42.fr/v2/")
	baseURL.Path += path
	baseURL.RawQuery = params.Encode()
	return baseURL.String()
}

func SetRateLimit(requestsPerSecond int) {
	rateLimit.Stop()
	rateLimit = time.NewTicker(time.Second / time.Duration(requestsPerSecond))
}

func (req *RequestData) clean() {
	if req.Error == nil {
		req.ExecuteMethod = nil
		req.params.Clear()
		req.body = nil
	}
}

func (req *RequestData) make(client *http.Client, method string) (status int, body []byte) {
	<-rateLimit.C
	var requestBody io.Reader
	switch req.contentType {
	case ContentTypeForm:
		requestBody = strings.NewReader(req.params.Encode())
	case ContentTypeJson:
		requestBody = bytes.NewReader(req.body)
	}
	var finalReq *http.Request
	if finalReq, req.Error = http.NewRequest(method, req.Endpoint, requestBody); req.Error == nil {
		finalReq.Header.Add("Content-Type", req.contentType)
		var resp *http.Response
		if resp, req.Error = client.Do(finalReq); req.Error == nil {
			defer resp.Body.Close()
			status = resp.StatusCode
			body, req.Error = ioutil.ReadAll(resp.Body)
		}
	}
	if status >= 400 {
		req.Error = fmt.Errorf("intra error [response: %d]", status)
	}
	return
}

func (req *RequestData) Create(client *http.Client, objPtr, params interface{}) *RequestData {
	defer req.clean()
	req.contentType = ContentTypeJson
	req.body, _ = json.Marshal(params)
	_, body := req.make(client, http.MethodPost)
	if req.Error == nil {
		req.Error = json.Unmarshal(body, objPtr)
	}
	if req.Error == nil && !req.bypassCacheWrite {
		intraCache.put(objPtr)
	}
	return req
}

func (req *RequestData) Delete(client *http.Client, objPtr interface{}) *RequestData {
	defer req.clean()
	req.contentType = ContentTypeForm
	_, body := req.make(client, http.MethodDelete)
	if req.Error != nil {
		value := reflect.Indirect(reflect.ValueOf(objPtr))
		ID := value.FieldByName("ID").Interface().(int)
		req.Error = fmt.Errorf("%s: (%d) %s: %s", value.Type().String(), ID, req.Error, string(body))
		return req
	}
	intraCache.delete(objPtr)
	return req
}

func (req *RequestData) Get(client *http.Client, objPtr interface{}) *RequestData {
	defer req.clean()
	req.contentType = ContentTypeForm
	if !req.bypassCacheRead && intraCache.get(objPtr) {
		return req
	}
	_, body := req.make(client, http.MethodGet)
	if req.Error == nil {
		req.Error = json.Unmarshal(body, objPtr)
	}
	if req.Error != nil {
		value := reflect.Indirect(reflect.ValueOf(objPtr))
		ID := value.FieldByName("ID").Interface().(int)
		req.Error = fmt.Errorf("%s (%d): %s: %s", value.Type().String(), ID, req.Error, string(body))
	} else if !req.bypassCacheWrite {
		intraCache.put(objPtr)
	}
	return req
}

func (req *RequestData) GetAll(client *http.Client, objPtr interface{}) *RequestData {
	defer req.clean()
	req.contentType = ContentTypeForm
	value := reflect.Indirect(reflect.ValueOf(objPtr))
	pageNumber := 1
	singlePage := false
	if _, present := req.params.Values["page[number]"]; present {
		pageNumber, _ = strconv.Atoi(req.params.Get("page[number]"))
		singlePage = true
	}
	data := bytes.Buffer{}
	data.WriteByte('[')
	for {
		req.params.Set("page[number]", strconv.Itoa(pageNumber))
		_, page := req.make(client, http.MethodGet)
		if req.Error != nil {
			req.Error = fmt.Errorf("%s: %s: %s", value.Type().String(), req.Error, string(page))
			return req
		}
		if string(page) != "[]" {
			if data.Len() > 1 {
				data.WriteByte(',')
			}
			data.Write(page[1 : len(page)-1])
		}
		if singlePage || string(page) == "[]" {
			break
		}
		pageNumber++
	}
	data.WriteByte(']')
	if data.String() == "[]" {
		return req
	}
	req.Error = json.Unmarshal(data.Bytes(), objPtr)
	if req.Error != nil {
		req.Error = fmt.Errorf("%s: %s", value.Type().String(), req.Error)
	} else if !req.bypassCacheWrite {
		for i := 0; i < value.Len(); i++ {
			intraCache.put(value.Index(i).Interface())
		}
	}
	return req
}

// Objects will have to be manually mutated after patching, as the 42 Intra API uses a different format for patching
func (req *RequestData) Patch(client *http.Client, objPtr interface{}, params interface{}) *RequestData {
	defer req.clean()
	req.contentType = ContentTypeJson
	req.body, _ = json.Marshal(params)
	req.make(client, http.MethodPatch)
	if req.Error == nil {
		_ = json.Unmarshal(req.body, objPtr)
		if !req.bypassCacheWrite {
			intraCache.put(objPtr)
		}
	}
	return req
}
