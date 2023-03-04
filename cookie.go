/**
 * @Author:      leafney
 * @Date:        2023-03-04 17:06
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description: cookies
 */

package rose

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// CookieFromStr 将 Cookie 字符串转换为 []*http.Cookie
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

// CookieFromFile 从文件中获取 Cookie 字符串并转换为 []*http.Cookie
func CookieFromFile(cookieFilePath string) ([]*http.Cookie, error) {
	fCookieStr, err := ioutil.ReadFile(cookieFilePath)
	if err != nil {
		return nil, err
	}

	return CookieFromStr(string(fCookieStr)), nil
}
