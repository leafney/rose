/**
 * @Author:      leafney
 * @Date:        2022-06-30 13:46
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import "testing"

func TestFloat64Round(t *testing.T) {
	t.Log(Float64Round(3.1415926, 4))
	t.Log(Float64Round(3.1415926, 2))
	t.Log(Float64Round(5.678, 0))
}

func TestFloat64RoundInt64(t *testing.T) {
	t.Log(Float64RoundInt64(3.1415926))
	t.Log(Float64RoundInt64(5.678))
}
