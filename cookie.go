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

// CookieFromNV 通过键值设置 Cookie，默认对值作编码处理
func CookieFromNV(name, value string) *http.Cookie {
	return CookieFromNVEscape(name, value, true)
}

// CookieFromNVEscape 通过键值设置 Cookie，可选是否对值作编码处理
func CookieFromNVEscape(name, value string, escape bool) *http.Cookie {
	val := value
	if escape {
		// cookie值可能包含特殊字符
		val = url.QueryEscape(value)
	}

	return &http.Cookie{
		Name:  name,
		Value: val,
	}
}

// CookieFromNVs 通过 map 集合设置 Cookie，默认对值作编码处理
func CookieFromNVs(cookies map[string]string) []*http.Cookie {
	return CookieFromNVsEscape(cookies, true)
}

// CookieFromNVsEscape 通过 map 集合设置 Cookie，可选是否对值作编码处理
func CookieFromNVsEscape(cookies map[string]string, escape bool) []*http.Cookie {
	newCookies := make([]*http.Cookie, 0)

	for k, v := range cookies {
		val := v
		if escape {
			// cookie值可能包含特殊字符
			val = url.QueryEscape(v)
		}

		newCookies = append(newCookies, &http.Cookie{
			Name:  k,
			Value: val,
		})
	}
	return newCookies
}

// CookieFromStr 将 Cookie 字符串转换为 []*http.Cookie，默认对值作编码处理
// Cookie 字符串格式：key1=value1; key2=value2; key3=value3;
func CookieFromStr(cookieStr string) []*http.Cookie {
	return CookieFromStrEscape(cookieStr, true)
}

// CookieFromStrEscape 将 Cookie 字符串转换为 []*http.Cookie，可选是否对值作编码处理
// Cookie 字符串格式：key1=value1; key2=value2; key3=value3;
func CookieFromStrEscape(cookieStr string, escape bool) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", cookieStr)
	request := http.Request{Header: header}

	newCookies := make([]*http.Cookie, 0)

	for _, c := range request.Cookies() {
		val := c.Value
		if escape {
			// cookie值可能包含特殊字符
			val = url.QueryEscape(c.Value)
		}

		newCookies = append(newCookies, &http.Cookie{
			Name:  c.Name,
			Value: val,
		})
	}

	return newCookies
}

// CookieFromStrWithDPE 将 Cookie 字符串转换为 []*http.Cookie，同时设置 Domain、Path、Expires 信息，默认对值作编码处理
// duration: 从当前时间开始的持续时间
// Cookie 字符串格式：key1=value1; key2=value2; key3=value3;
func CookieFromStrWithDPE(cookieStr string, domain string, path string, duration time.Duration) []*http.Cookie {
	return CookieFromStrWithDPEEscape(cookieStr, domain, path, duration, true)
}

// CookieFromStrWithDPEEscape 将 Cookie 字符串转换为 []*http.Cookie，同时设置 Domain、Path、Expires 信息，可选是否对值作编码处理
// duration: 从当前时间开始的持续时间
// Cookie 字符串格式：key1=value1; key2=value2; key3=value3;
func CookieFromStrWithDPEEscape(cookieStr string, domain string, path string, duration time.Duration, escape bool) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", cookieStr)
	request := http.Request{Header: header}

	expr := time.Now().Add(duration)

	newCookies := make([]*http.Cookie, 0)

	for _, c := range request.Cookies() {
		val := c.Value
		if escape {
			// cookie值可能包含特殊字符
			val = url.QueryEscape(c.Value)
		}
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

// CookieFromFile 从文件中获取 Cookie 字符串并转换为 []*http.Cookie，默认对值作编码处理
func CookieFromFile(cookieFilePath string) ([]*http.Cookie, error) {
	return CookieFromFileEscape(cookieFilePath, true)
}

// CookieFromFileEscape 从文件中获取 Cookie 字符串并转换为 []*http.Cookie，可选是否对值作编码处理
func CookieFromFileEscape(cookieFilePath string, escape bool) ([]*http.Cookie, error) {
	cookieBytes, err := os.ReadFile(cookieFilePath)
	if err != nil {
		return nil, err
	}

	return CookieFromStrEscape(string(cookieBytes), escape), nil
}

// CookieFromFileWithDPE 从文件中获取 Cookie 字符串并转换为 []*http.Cookie，同时设置 Domain、Path、Expires 信息，默认对值作编码处理
// duration: 从当前时间开始的持续时间
func CookieFromFileWithDPE(cookieFilePath string, domain string, path string, duration time.Duration) ([]*http.Cookie, error) {
	return CookieFromFileWithDPEEscape(cookieFilePath, domain, path, duration, true)
}

// CookieFromFileWithDPEEscape 从文件中获取 Cookie 字符串并转换为 []*http.Cookie，同时设置 Domain、Path、Expires 信息，可选是否对值作编码处理
// duration: 从当前时间开始的持续时间
func CookieFromFileWithDPEEscape(cookieFilePath string, domain string, path string, duration time.Duration, escape bool) ([]*http.Cookie, error) {
	cookieBytes, err := os.ReadFile(cookieFilePath)
	if err != nil {
		return nil, err
	}

	return CookieFromStrWithDPEEscape(string(cookieBytes), domain, path, duration, escape), nil
}

// CookieToStrNV 将 []*http.Cookie 转换成 Cookie 字符串，格式为 name1=value1; name2=value2; name3=value3;
func CookieToStrNV(cookies []*http.Cookie) string {
	return CookieToStrNVUnEscape(cookies, true)
}

// CookieToStrNVUnEscape 将 []*http.Cookie 转换成 Cookie 字符串，格式为 name1=value1; name2=value2; name3=value3;
func CookieToStrNVUnEscape(cookies []*http.Cookie, unEscape bool) string {
	res := make([]string, 0)
	for _, c := range cookies {
		val := c.Value
		if unEscape {
			// 对特殊字符转义
			val, _ = url.QueryUnescape(c.Value)
		}
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
	return CookieToJsonStrNVDPEUnEscape(cookies, true)
}

// CookieToJsonStrNVDPEUnEscape 将 []*http.Cookie 转换成 Json 格式的 Cookie 字符串，带有 Name、Value、Path、Domain、Expires 参数
//
// 转换后格式为：[{"name":"","value":"","path":"","domain":"","expires":0},{"name":"","value":"","path":"","domain":"","expires":0}]
func CookieToJsonStrNVDPEUnEscape(cookies []*http.Cookie, unEscape bool) string {
	res := make([]*CookieModel, 0)
	for _, c := range cookies {
		val := c.Value
		if unEscape {
			// 对特殊字符转义
			val, _ = url.QueryUnescape(c.Value)
		}
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
