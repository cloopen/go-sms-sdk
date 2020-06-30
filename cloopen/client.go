package cloopen

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	config *Config
}

func NewJsonClient(config *Config) *Client {
	config.withContentType(CONTENT_JSON)
	return &Client{config}
}

func NewXmlClient(config *Config) *Client {
	config.withContentType(CONTENT_XML)
	return &Client{config}
}

var DefHeaders = map[string]string{
	"Api-Lang":   "go",
	"Connection": "keep-alive",
}

type request struct {
	config *Config
	method string
	url    *url.URL
	body   *bytes.Buffer
	header http.Header
	params url.Values
}

func (c *Client) newRequest(method, endpoint, path string) *request {

	r := &request{
		config: c.config,
		method: method,
		params: make(map[string][]string),
		header: make(http.Header),
	}

	u := &url.URL{
		Host: endpoint,
		Path: path,
	}

	if c.config.UseSSL {
		u.Scheme = SCHEME_HTTPS
	} else {
		u.Scheme = SCHEME_HTTP
	}

	r.url = u
	return r
}

func (c *Client) doRequest(r *request) (*http.Response, error) {
	req, err := r.toHTTP()
	if err != nil {
		return nil, err
	}

	return c.config.HttpClient.Do(req)
}

func (r *request) toHTTP() (*http.Request, error) {
	r.url.RawQuery = r.params.Encode()
	req, err := http.NewRequest(r.method, r.url.RequestURI(), r.body)
	if err != nil {
		return nil, err
	}

	req.URL.Host = r.url.Host
	req.URL.Scheme = r.url.Scheme
	req.Header = r.header

	for key, val := range DefHeaders {
		req.Header.Set(key, val)
	}

	log.Printf("Request URL: %s \n", r.url)
	log.Printf("Http Header: %s \n", req.Header)
	log.Printf("Request Body: %s \n", r.body)
	return req, nil
}

func (c *Client) handleResponse(resp *http.Response, out interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("请求失败：%s \n", string(body))
	}
	log.Printf("Response Body: %s \n", string(body))
	contentType := c.config.ContentType
	if contentType == CONTENT_JSON {
		return json.NewDecoder(bytes.NewBuffer(body)).Decode(out)
	} else {
		return xml.NewDecoder(bytes.NewBuffer(body)).Decode(out)
	}
}
