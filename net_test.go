/**
 * @Author:      leafney
 * @Date:        2022-11-05 00:37
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import (
	"testing"
)

func TestLocalIP(t *testing.T) {
	//ip, err := LocalIP()
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf(ip)

	t.Log(IsLocalIPAddr("::1"))
	t.Log(IsLocalIPAddr("127.0.0.1"))

}
