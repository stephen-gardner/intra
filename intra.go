package intra

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
)

var (
	clientID     string
	clientSecret string
)

func init() {
	clientID = os.Getenv("INTRA_CLIENT_ID")
	clientSecret = os.Getenv("INTRA_CLIENT_SECRET")
}

func GetClient(ctx context.Context, scopes ...string) *http.Client {
	oauth := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://api.intra.42.fr/oauth/token",
		Scopes:       scopes,
	}
	return oauth.Client(ctx)
}

func GetEndpoint(path string, params url.Values) string {
	baseURL, _ := url.Parse("https://api.intra.42.fr/v2/")
	baseURL.Path += path
	baseURL.RawQuery = params.Encode()
	return baseURL.String()
}

func RunRequest(client *http.Client, method, endpoint string, formData url.Values) (int, []byte, error) {
	req, err := http.NewRequest(method, endpoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return 0, nil, err
	}
	if formData != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		err := fmt.Errorf("intra error [response: %d] %s", resp.StatusCode, string(data))
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, data, err
}

func GetAll(client *http.Client, endpoint string, params url.Values) ([][]byte, error) {
	var res [][]byte
	pageNumber := 1
	singlePage := false
	if _, ok := params["page[number]"]; ok {
		pageNumber, _ = strconv.Atoi(params.Get("page[number]"))
		singlePage = true
	}
	for {
		params.Set("page[number]", strconv.Itoa(pageNumber))
		endpoint := GetEndpoint(endpoint, params)
		_, page, err := RunRequest(client, http.MethodGet, endpoint, nil)
		if err != nil {
			return res, err
		}
		if singlePage {
			return [][]byte{page}, nil
		}
		if string(page) == "[]" {
			break
		}
		res = append(res, page)
		pageNumber++
	}
	return res, nil
}
