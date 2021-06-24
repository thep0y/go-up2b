/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: http.go
 * @Created: 2021-06-21 09:52:54
 * @Modified: 2021-06-24 11:51:45
 */

package request

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/mailru/easyjson"
	"github.com/thep0y/go-up2b/models"
)

type Request struct {
	client  *http.Client
	headers http.Header
}

func NewRequest(headers http.Header) *Request {
	req := new(Request)
	req.client = &http.Client{}

	req.headers = headers

	return req
}

func (r *Request) parseResponse(resp *http.Response) (*Response, error) {
	response := new(Response)
	response.StatusCode = resp.StatusCode
	response.Request = resp.Request

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response.Body = body
	response.Header = resp.Header

	return response, nil
}

func (r *Request) request(req *http.Request) (*Response, error) {
	if len(req.Header) == 0 {
		req.Header = r.headers
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return r.parseResponse(resp)
}

func (r *Request) Get(url string) (*Response, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return r.request(req)
}

func (r *Request) requestWithBody(url, method string, data models.FileData) (*Response, error) {
	formdata, err := easyjson.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		method,
		url,
		bytes.NewReader(formdata),
	)
	if err != nil {
		return nil, err
	}

	return r.request(req)

}

func (r *Request) Put(url string, data models.FileData) (*Response, error) {
	return r.requestWithBody(url, http.MethodPut, data)
}

func (r *Request) Post(url string, data models.FileData) (*Response, error) {
	return r.requestWithBody(url, http.MethodPost, data)
}

func (r *Request) PostWithoutHeader(u string, data models.FileData) (*Response, error) {
	values := make(url.Values)
	for k, v := range data {
		values.Add(k, v)
	}

	fmt.Println(values)

	req, err := http.NewRequest(
		http.MethodPost,
		u,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return r.parseResponse(resp)
}

func (r *Request) PostMultipartForm(url, contentType string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}

	if len(contentType) > 0 {
		req.Header = r.headers.Clone()
		req.Header.Add("Content-Type", contentType)
	}

	return r.request(req)
}

func PostWithHeadersWithNoRedirect(u string, headers http.Header, data models.FileData) (*http.Response, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// data 不是 json 格式，改成 url 格式
	formdata := make(url.Values)
	for k, v := range data {
		formdata.Add(k, v)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		u,
		strings.NewReader(formdata.Encode()),
	)
	if err != nil {
		return nil, err
	}
	req.Header = headers

	return client.Do(req)
}