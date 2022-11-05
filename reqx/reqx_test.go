/**
 * @Author:      leafney
 * @Date:        2022-09-13 09:56
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package reqx

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	resp, err :=
		Get("http://jsonplaceholder.typicode.com/posts/1").
			SetDebug(true).
			SetTimeout(1 * time.Second).
			Do()
	if err != nil {
		t.Error(err)
	}
	if resp.IsSuccess() {
		t.Log(resp.String())
	}
	t.Log(resp.StatusCode())
}
