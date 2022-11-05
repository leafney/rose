/**
 * @Author:      leafney
 * @Date:        2022-08-08 11:05
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package reqx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	r        *Request
	resp     *http.Response
	req      *http.Request
	respBody []byte
}

func (r *Response) Request() *http.Request {
	return r.req
}

func (r *Response) Response() *http.Response {
	return r.resp
}

func (r *Response) StatusCode() int {
	return r.resp.StatusCode
}

func (r *Response) ToBytes() ([]byte, error) {

	defer r.resp.Body.Close()
	body, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		fmt.Printf("[ERROR] failed to read body %v", err)
		return nil, err
	}

	if r.r.Debug {
		r.r.log.Printf("Response: \r\n %s \r\n", string(body))
	}

	r.respBody = body
	return r.respBody, nil
}

func (r *Response) Bytes() []byte {
	data, _ := r.ToBytes()
	return data
}

func (r *Response) ToString() (string, error) {
	data, err := r.ToBytes()
	return string(data), err
}

func (r *Response) String() string {
	data, _ := r.ToBytes()
	return string(data)
}

func (r *Response) ToJson(v interface{}) error {
	data, err := r.ToBytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (r *Response) IsSuccess() bool {
	return r.resp.StatusCode == http.StatusOK
}
