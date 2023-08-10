/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-08-10 15:10
 * @Description:
 */

package rose

import "testing"

func TestFormatString(t *testing.T) {
	t.Log(FmtString(`replace {msg} with {data}`, map[string]interface{}{"msg": "hello", "data": 123}))
	t.Log(FmtStringWith("${", "}", `replace ${msg} with ${data}`, FmtItems{"msg": "hello", "data": 123}))
}
