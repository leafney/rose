/**
 * @Author:      leafney
 * @Date:        2022-08-08 11:06
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package reqx

import (
	"net/url"
	"strings"
)

func parseRequestURL(r *Request) error {
	tempURL := r.RawURL
	if len(r.PathParams) > 0 {
		for p, v := range r.PathParams {
			tempURL = strings.Replace(tempURL, "{"+p+"}", url.PathEscape(v), -1)
		}
	}

	// Parsing request URL
	reqURL, err := url.Parse(tempURL)
	if err != nil {
		return err
	}

	// Adding Query Param
	query := make(url.Values)

	for k, v := range r.QueryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	if len(query) > 0 {
		if IsStringEmpty(reqURL.RawQuery) {
			reqURL.RawQuery = query.Encode()
		} else {
			reqURL.RawQuery = reqURL.RawQuery + "&" + query.Encode()
		}
	}

	reqURL.Host = removeEmptyPort(reqURL.Host)
	r.URL = reqURL
	return nil
}

func parseRequestBody(r *Request) (err error) {
	if len(r.FormData) > 0 {
		r.SetContentType(FormContentType)
		r.SetBodyBytes([]byte(r.FormData.Encode()))
		return
	}
	return
}
