/**
 * @Author:      leafney
 * @Date:        2022-08-08 11:08
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package reqx

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (r *Request) Do() (resp *Response, err error) {

	if err := parseRequestURL(r); err != nil {
		r.log.Printf("[ERROR] failed to get params %v", err)
		return nil, err
	}

	if err := parseRequestBody(r); err != nil {
		r.log.Printf("[ERROR] failed to get body %v", err)
		return nil, err
	}

	if r.Debug {
		r.log.Printf("%s %s \r\n", r.Method, r.URL.String())
	}

	var host string
	if h := r.Headers.Get("Host"); h != "" {
		host = h
	} else {
		host = r.URL.Host
	}

	var header http.Header
	if r.Headers == nil {
		header = make(http.Header)
	} else {
		header = r.Headers
	}

	if r.Debug {
		for k, vv := range header {
			for _, v := range vv {
				r.log.Printf("%s: %s \r\n", k, v)
			}
		}
	}

	contentLength := int64(len(r.body))

	var reqBody io.ReadCloser
	if r.getBody != nil {
		reqBody, err = r.getBody()
		if err != nil {
			return
		}
	}

	req := &http.Request{
		Method:        r.Method,
		Header:        header,
		URL:           r.URL,
		Host:          host,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: contentLength,
		Body:          reqBody,
		GetBody:       r.getBody,
	}

	for _, cookie := range r.Cookies {
		req.AddCookie(cookie)
	}

	client := &http.Client{}

	if r.Proxy != nil {
		client.Transport = &http.Transport{
			Proxy:           r.Proxy,
			TLSClientConfig: r.TLSConfig,
		}
	}

	if r.TimeOut.Nanoseconds() > 0 {
		client.Timeout = r.TimeOut
	}

	httpResp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[ERROR] failed to request %v", err)
		return nil, err
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		fmt.Printf("[ERROR] failed to read body %v", err)
		return nil, err
	}
	fmt.Println("返回数据：")
	fmt.Println(string(body))
	//return &Response{
	//	resp,
	//	req,
	//}, err
	return nil, nil
}

func (r *Request) Get(url string) *Request {
	r.Method = http.MethodGet
	r.RawURL = url
	return r
}

func (r *Request) Post(url string) *Request {
	r.Method = http.MethodPost
	r.RawURL = url
	return r
}

func (r *Request) Put(url string) *Request {
	r.Method = http.MethodPut
	r.RawURL = url
	return r
}

func (r *Request) Patch(url string) *Request {
	r.Method = http.MethodPatch
	r.RawURL = url
	return r
}

func (r *Request) Delete(url string) *Request {
	r.Method = http.MethodDelete
	r.RawURL = url
	return r
}
