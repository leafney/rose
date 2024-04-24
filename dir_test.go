/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2024-04-24 08:33
 * @Description:
 */

package rose

import "testing"

func TestDIsExist(t *testing.T) {
	path := "./aaa/bbb/ccc.txt"
	t.Log(DIsExist(path))
	//t.Log(DIsExistE(path))
	//t.Log(DEnsurePathExist(path))
}
