/**
 * @Author:      leafney
 * @Date:        2022-08-08 11:05
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package reqx

import "net/http"

type Response struct {
	*http.Response
	Request *Request
	Url     string
	Body    []byte
}
