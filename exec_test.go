/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-08-12 21:47
 * @Description:
 */

package rose

import "testing"

func TestExecCmd(t *testing.T) {

	//o1, err := ExecCmdBashOut("pwd")
	o1, err := ExecCmdBashOutDir("./reqx", "pwd")
	if err != nil {
		t.Log(err)
	}
	t.Log(o1)
}
