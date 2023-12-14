/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-12-14 23:05
 * @Description:
 */

package rose

import (
	"testing"
	"time"
)

func TestCookieToJsonStrNVDPE(t *testing.T) {
	cookieStr := "aaa=hello; bbb=world; ccc=welcome"

	cookies := CookieFromStrWithDPE(cookieStr, ".baidu.com", "/", time.Hour)
	res := CookieToJsonStrNVDPE(cookies)
	t.Log(res)
}
