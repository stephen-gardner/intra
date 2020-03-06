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

type RequestData struct {
	Error              error
	ExecuteMethod      func()
	Endpoint           string
	ContentType        string
	Body               []byte
	Params             Params
	CacheReadBypassed  bool
	CacheWriteBypassed bool
}

const (
	ContentTypeForm = "application/x-www-form-urlencoded"
	ContentTypeJson = "application/json"
	ScopeELearning  = "elearning"
	ScopeForum      = "forum"
	ScopeProfile    = "profile"
	ScopeProjects   = "projects"
	ScopePublic     = "public"
	ScopeTig        = "tig"
)

var (
	oauth     clientcredentials.Config
	rateLimit = time.NewTicker(time.Second / 2)
	ScopeAll  = []string{ScopeELearning, ScopeForum, ScopeProfile, ScopeProjects, ScopePublic, ScopeTig}
)

func init() {
	SetScope(ScopePublic)
}

func GetClient(ctx context.Context) *http.Client {
	if ctx == nil {
		ctx = context.Background()
	}
	client := oauth.Client(ctx)
	client.Timeout = 2 * time.Minute
	return client
}

func SetScope(scopes ...string) {
	oauth = clientcredentials.Config{
		ClientID:     os.Getenv("INTRA_CLIENT_ID"),
		ClientSecret: os.Getenv("INTRA_CLIENT_SECRET"),
		TokenURL:     "https://api.intra.42.fr/oauth/token",
		Scopes:       scopes,
	}
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

func EncapsulatedMarshal(container string, params interface{}) json.RawMessage {
	encapsulated := map[string]interface{}{container: params}
	data, _ := json.Marshal(encapsulated)
	return data
}

func (req *RequestData) Clean() {
	if req.Error == nil {
		req.ExecuteMethod = nil
		req.Params.Clear()
		req.Body = nil
	}
}

func (req *RequestData) Make(ctx context.Context, method string) (status int, body []byte) {
	<-rateLimit.C
	var requestBody io.Reader
	switch req.ContentType {
	case ContentTypeForm:
		requestBody = strings.NewReader(req.Params.Encode())
	case ContentTypeJson:
		requestBody = bytes.NewReader(req.Body)
	}
	var finalReq *http.Request
	if finalReq, req.Error = http.NewRequest(method, req.Endpoint, requestBody); req.Error == nil {
		finalReq.Header.Add("Content-Type", req.ContentType)
		var resp *http.Response
		if resp, req.Error = GetClient(ctx).Do(finalReq); req.Error == nil {
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

func (req *RequestData) Create(ctx context.Context, objPtr interface{}, params json.RawMessage) *RequestData {
	defer req.Clean()
	req.ContentType = ContentTypeJson
	req.Body = params
	_, body := req.Make(ctx, http.MethodPost)
	if req.Error == nil {
		req.Error = json.Unmarshal(body, objPtr)
	}
	if req.Error == nil && req.cacheWriteEnabled() {
		intraCache.put(objPtr)
	}
	return req
}

func (req *RequestData) Delete(ctx context.Context, objPtr interface{}) *RequestData {
	defer req.Clean()
	req.ContentType = ContentTypeForm
	_, body := req.Make(ctx, http.MethodDelete)
	if req.Error != nil {
		value := reflect.Indirect(reflect.ValueOf(objPtr))
		ID := value.FieldByName("ID").Interface().(int)
		req.Error = fmt.Errorf("%s: (%d) %s: %s", value.Type().String(), ID, req.Error, string(body))
		return req
	}
	if intraCache.isEnabled() {
		intraCache.delete(objPtr)
	}
	return req
}

func (req *RequestData) Get(ctx context.Context, objPtr interface{}) *RequestData {
	defer req.Clean()
	req.ContentType = ContentTypeForm
	if req.cacheReadEnabled() && intraCache.get(objPtr) {
		return req
	}
	_, body := req.Make(ctx, http.MethodGet)
	if req.Error == nil {
		req.Error = json.Unmarshal(body, objPtr)
	}
	if req.Error != nil {
		value := reflect.Indirect(reflect.ValueOf(objPtr))
		ID := value.FieldByName("ID").Interface().(int)
		req.Error = fmt.Errorf("%s (%d): %s: %s", value.Type().String(), ID, req.Error, string(body))
	} else if req.cacheWriteEnabled() {
		intraCache.put(objPtr)
	}
	return req
}

func (req *RequestData) GetAll(ctx context.Context, objPtr interface{}) *RequestData {
	defer req.Clean()
	req.ContentType = ContentTypeForm
	value := reflect.Indirect(reflect.ValueOf(objPtr))
	pageNumber := 1
	singlePage := false
	if _, present := req.Params.Values["page[number]"]; present {
		pageNumber, _ = strconv.Atoi(req.Params.Get("page[number]"))
		singlePage = true
	}
	data := bytes.Buffer{}
	data.WriteByte('[')
	for {
		req.Params.Set("page[number]", strconv.Itoa(pageNumber))
		_, page := req.Make(ctx, http.MethodGet)
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
	} else if req.cacheWriteEnabled() {
		for i := 0; i < value.Len(); i++ {
			ptr := reflect.New(value.Index(i).Type())
			ptr.Elem().Set(value.Index(i))
			intraCache.put(ptr.Interface())
		}
	}
	return req
}

// Objects will have to be manually mutated after patching, as the 42 Intra API uses a different format for patching
// Thus CacheObject() should be called on these objects after mutation
func (req *RequestData) Patch(ctx context.Context, params json.RawMessage) *RequestData {
	defer req.Clean()
	req.ContentType = ContentTypeJson
	req.Body = params
	req.Make(ctx, http.MethodPatch)
	return req
}

func (req *RequestData) cacheWriteEnabled() bool {
	return intraCache.isEnabled() && !req.CacheWriteBypassed
}

func (req *RequestData) cacheReadEnabled() bool {
	return intraCache.isEnabled() && !req.CacheReadBypassed
}
