/**
 * @Author:      leafney
 * @Date:        2022-08-08 11:08
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package reqx

import (
	"log"
	"net/http"
	"os"
)

var std = New()

type Req struct {
	req *Request
	//Timeout time.Duration
}

func New() *Req {
	return &Req{
		req: &Request{
			//TimeOut:
			log:     log.New(os.Stdout, "", log.LstdFlags),
			Cookies: make([]*http.Cookie, 0),
		},
		//Timeout: 30 * time.Second,
	}
}

func Get(url string) *Request {
	return std.req.Get(url)
}

func Post(url string) *Request {
	return std.req.Post(url)
}

func Put(url string) *Request {
	return std.req.Put(url)
}

func Patch(url string) *Request {
	return std.req.Patch(url)
}

func Delete(url string) *Request {
	return std.req.Delete(url)
}
