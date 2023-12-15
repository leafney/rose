/**
 * @Author:      leafney
 * @Date:        2023-03-04 17:06
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description: cookies
 */

package rose

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// CookieFromStr 将 Cookie 字符串转换为 []*http.Cookie
//
// Cookie 字符串格式：key1=value1; key2=value2; key3=value3;
func CookieFromStr(cookieStr string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", cookieStr)
	request := http.Request{Header: header}

	newCookies := make([]*http.Cookie, 0)

	for _, c := range request.Cookies() {
		// cookie值可能包含特殊字符
		val := url.QueryEscape(c.Value)
		newCookies = append(newCookies, &http.Cookie{
			Name:  c.Name,
			Value: val,
		})
	}

	return newCookies
}

// CookieFromStrWithDPE 将 Cookie 字符串转换为 []*http.Cookie，同时设置 Domain、Path、Expires 信息
// duration: 从当前时间开始的持续时间
//
// Cookie 字符串格式：key1=value1; key2=value2; key3=value3;
func CookieFromStrWithDPE(cookieStr string, domain string, path string, duration time.Duration) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", cookieStr)
	request := http.Request{Header: header}

	expr := time.Now().Add(duration)

	newCookies := make([]*http.Cookie, 0)

	for _, c := range request.Cookies() {
		// cookie值可能包含特殊字符
		val := url.QueryEscape(c.Value)
		newCookies = append(newCookies, &http.Cookie{
			Name:    c.Name,
			Value:   val,
			Expires: expr,
			Domain:  domain,
			Path:    path,
		})
	}

	return newCookies
}

// CookieFromFile 从文件中获取 Cookie 字符串并转换为 []*http.Cookie
func CookieFromFile(cookieFilePath string) ([]*http.Cookie, error) {
	cookieBytes, err := os.ReadFile(cookieFilePath)
	if err != nil {
		return nil, err
	}

	return CookieFromStr(string(cookieBytes)), nil
}

// CookieFromFileWithDPE 从文件中获取 Cookie 字符串并转换为 []*http.Cookie，同时设置 Domain、Path、Expires 信息
// duration: 从当前时间开始的持续时间
func CookieFromFileWithDPE(cookieFilePath string, domain string, path string, duration time.Duration) ([]*http.Cookie, error) {
	cookieBytes, err := os.ReadFile(cookieFilePath)
	if err != nil {
		return nil, err
	}

	return CookieFromStrWithDPE(string(cookieBytes), domain, path, duration), nil
}

// CookieToStrNV 将 []*http.Cookie 转换成 Cookie 字符串，格式为 name1=value1; name2=value2; name3=value3;
func CookieToStrNV(cookies []*http.Cookie) string {
	res := make([]string, 0)
	for _, c := range cookies {
		// 对特殊字符转义
		val, _ := url.QueryUnescape(c.Value)
		res = append(res, fmt.Sprintf("%s=%s", c.Name, val))
	}

	return strings.Join(res, "; ")
}

type CookieModel struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Path    string `json:"path"`
	Domain  string `json:"domain"`
	Expires int64  `json:"expires"` // unix时间戳
}

// CookieToJsonStrNVDPE 将 []*http.Cookie 转换成 Json 格式的 Cookie 字符串，带有 Name、Value、Path、Domain、Expires 参数
//
// 转换后格式为：[{"name":"","value":"","path":"","domain":"","expires":0},{"name":"","value":"","path":"","domain":"","expires":0}]
func CookieToJsonStrNVDPE(cookies []*http.Cookie) string {
	res := make([]*CookieModel, 0)
	for _, c := range cookies {
		// 对特殊字符转义
		val, _ := url.QueryUnescape(c.Value)

		res = append(res, &CookieModel{
			Name:    c.Name,
			Value:   val,
			Path:    c.Path,
			Domain:  c.Domain,
			Expires: c.Expires.Unix(),
		})
	}

	return JsonMarshalStr(res)
}

// Deprecated: Use CookieFromStrWithDPE instead.
func CookieFromStrWithDomain(cookieStr string, domain string, path string, expires time.Duration) []*http.Cookie {
	return CookieFromStrWithDPE(cookieStr, domain, path, expires)
}

// Deprecated: Use CookieToStrNV instead.
func CookieToStr(cookies []*http.Cookie) string {
	return CookieToStrNV(cookies)
}
