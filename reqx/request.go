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
