/**
 * @Author:      leafney
 * @Date:        2022-08-08 11:05
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package reqx

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/leafney/rose"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	urlpkg "net/url"
	"strings"
	"time"
)

type Request struct {
	PathParams       map[string]string
	QueryParams      urlpkg.Values
	FormData         urlpkg.Values
	Headers          http.Header
	Cookies          []*http.Cookie
	RawRequest       *http.Request
	RawURL           string
	Method           string
	URL              *urlpkg.URL
	TimeOut          time.Duration
	Proxy            func(*http.Request) (*urlpkg.URL, error)
	TLSConfig        *tls.Config
	unReplayableBody io.ReadCloser
	getBody          GetContentFunc
	body             []byte
	marshalBody      interface{}

	t     *http.Transport
	Debug bool
	log   *log.Logger
}

type GetContentFunc func() (io.ReadCloser, error)

func (r *Request) Do() (resp *Response, err error) {

	if err = parseRequestURL(r); err != nil {
		r.log.Printf("[ERROR] failed to get params %v", err)
		return nil, err
	}

	if err = parseRequestBody(r); err != nil {
		r.log.Printf("[ERROR] failed to get body %v", err)
		return nil, err
	}

	if r.Debug {
		r.log.Printf("URL: %s %s\n", r.Method, r.URL.String())
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
				r.log.Printf("Headers: \n%s: %s\n", k, v)
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

	//defer httpResp.Body.Close()
	//body, err := ioutil.ReadAll(httpResp.Body)
	//if err != nil {
	//	fmt.Printf("[ERROR] failed to read body %v", err)
	//	return nil, err
	//}
	//
	//if r.Debug {
	//	r.log.Printf("Response: \n %s \r\n", string(body))
	//}

	return &Response{
		r:    r,
		resp: httpResp,
		req:  r.RawRequest,
	}, nil

}

// SetFormData set the form data from a map
func (r *Request) SetFormData(data map[string]string) *Request {
	if r.FormData == nil {
		r.FormData = urlpkg.Values{}
	}
	for k, v := range data {
		r.FormData.Set(k, v)
	}
	return r
}

// SetCookies set http cookies for the request.
func (r *Request) SetCookies(cookies ...*http.Cookie) *Request {
	r.Cookies = append(r.Cookies, cookies...)
	return r
}

// SetQueryString set URL query parameters for the request using
// raw query string.
func (r *Request) SetQueryString(query string) *Request {
	params, err := urlpkg.ParseQuery(strings.TrimSpace(query))
	if err != nil {
		//r.client.log.Warnf("failed to parse query string (%s): %v", query, err)
		return r
	}
	if r.QueryParams == nil {
		r.QueryParams = make(urlpkg.Values)
	}
	for p, v := range params {
		for _, pv := range v {
			r.QueryParams.Add(p, pv)
		}
	}
	return r
}

// SetBearerAuthToken set bearer auth token for the request.
func (r *Request) SetBearerAuthToken(token string) *Request {
	return r.SetHeader("Authorization", "Bearer "+token)
}

// SetBasicAuth set basic auth for the request.
func (r *Request) SetBasicAuth(username, password string) *Request {
	return r.SetHeader("Authorization", BasicAuthHeaderValue(username, password))
}

// SetHeaders set headers from a map for the request.
func (r *Request) SetHeaders(hdrs map[string]string) *Request {
	for k, v := range hdrs {
		r.SetHeader(k, v)
	}
	return r
}

// SetHeader set a header for the request.
func (r *Request) SetHeader(key, value string) *Request {
	if r.Headers == nil {
		r.Headers = make(http.Header)
	}
	r.Headers.Set(key, value)

	return r
}

// SetQueryParams set URL query parameters from a map for the request.
func (r *Request) SetQueryParams(params map[string]string) *Request {
	for k, v := range params {
		r.SetQueryParam(k, v)
	}
	return r
}

// SetQueryParam set an URL query parameter for the request.
func (r *Request) SetQueryParam(key, value string) *Request {
	if r.QueryParams == nil {
		r.QueryParams = make(urlpkg.Values)
	}
	r.QueryParams.Set(key, value)
	return r
}

// AddQueryParam add a URL query parameter for the request.
func (r *Request) AddQueryParam(key, value string) *Request {
	if r.QueryParams == nil {
		r.QueryParams = make(urlpkg.Values)
	}
	r.QueryParams.Add(key, value)
	return r
}

// SetPathParams set URL path parameters from a map for the request.
// https://www.xxx.com/user/{name}
func (r *Request) SetPathParams(params map[string]string) *Request {
	for key, value := range params {
		r.SetPathParam(key, value)
	}
	return r
}

// SetPathParam set a URL path parameter for the request.
func (r *Request) SetPathParam(key, value string) *Request {
	if r.PathParams == nil {
		r.PathParams = make(map[string]string)
	}
	r.PathParams[key] = value
	return r
}

// SetContentType set the `Content-Type` for the request.
func (r *Request) SetContentType(contentType string) *Request {
	return r.SetHeader("Content-Type", contentType)
}

//
func (r *Request) SetUserAgent(userAgent string) *Request {
	return r.SetHeader("User-Agent", userAgent)
}

func (r *Request) SetDefUserAgent() *Request {
	agentList := []string{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.28",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:90.0) Gecko/20100101 Firefox/90.0",
	}
	agent := rose.SliceRandomItemStr(agentList)
	return r.SetHeader("User-Agent", agent)
}

func (r *Request) SetProxy(proxyUrl string, skipVerify bool) *Request {
	u, err := urlpkg.Parse(proxyUrl)
	if err != nil {
		return r
	}
	r.Proxy = http.ProxyURL(u)
	r.TLSConfig = &tls.Config{InsecureSkipVerify: skipVerify}
	return r
}

func (r *Request) SetTimeout(d time.Duration) *Request {
	r.TimeOut = d
	return r
}

func (r *Request) SetDebug(debug bool) *Request {
	r.Debug = debug
	return r
}

// SetBody set the request body, accepts string, []byte, io.Reader, map and struct.
func (r *Request) SetBody(body interface{}) *Request {
	if body == nil {
		return r
	}
	switch b := body.(type) {
	case io.ReadCloser:
		r.unReplayableBody = b
		r.getBody = func() (io.ReadCloser, error) {
			return r.unReplayableBody, nil
		}
	case io.Reader:
		r.unReplayableBody = ioutil.NopCloser(b)
		r.getBody = func() (io.ReadCloser, error) {
			return r.unReplayableBody, nil
		}
	case []byte:
		r.SetBodyBytes(b)
	case string:
		r.SetBodyString(b)
	case func() (io.ReadCloser, error):
		r.getBody = b
	case GetContentFunc:
		r.getBody = b
	default:
		r.marshalBody = body
	}
	return r
}

// SetBodyBytes set the request body as []byte.
func (r *Request) SetBodyBytes(body []byte) *Request {
	r.body = body
	r.getBody = func() (io.ReadCloser, error) {
		return ioutil.NopCloser(bytes.NewReader(body)), nil
	}
	return r
}

// SetBodyString set the request body as string.
func (r *Request) SetBodyString(body string) *Request {
	return r.SetBodyBytes([]byte(body))
}

func (r *Request) SetJsonContentType() *Request {
	return r.SetContentType(JsonContentType)
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
