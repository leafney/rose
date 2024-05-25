/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2024-05-25 02:15
 * @Description:
 */

package rose

import (
	"testing"
)

func TestFlowKBtoMB(t *testing.T) {

	//t.Log(FlowKBtoMB(1500))
	//t.Log(FlowKBtoGB(2500000))

	//t.Log(FlowKBtoGMB(2500000))
	//t.Log(FlowKBtoGMB(1500))
	//t.Log(FlowKBtoGMB(500000))

	//t.Log(FlowMBtoKB(952))
	//t.Log(FlowMBtoGB(2048))
	//t.Log(FlowMBtoGMB(3000))
	//t.Log(FlowMBtoGMB(512))

	//t.Log(FlowMBtoGMB(2560))
	//t.Log(FlowGBtoGMB(2.5))

	//t.Log(FlowGBtoGMB(2.01))
	//t.Log(FlowGBtoKB(2.01))

	//a, _ := FlowParseToKB("2GB 10MB")
	//b, _ := FlowParseToKB("2107637.76KB")
	//c, _ := FlowParseToKB("2058.24mb")
	//t.Logf("%.2fKB-%.2fKB-%.2fKB", a, b, c)

	//d, _ := FlowParseToKB("2.3gb")
	//t.Logf("%.2fKB", d)
	//t.Log(FlowGBtoKB(2.3))

	//d, _ := FlowParseToKB("1.1MB1.5kb")
	//t.Logf("%v-%.2fKB", d, d)
	//t.Log(FlowMBtoKB(1.1))
	//
	//t.Log(FlowKBtoMB(d))
	//t.Log(FlowKBtoMB(1127.90))

	d, _ := FlowParseToKB("2GB 393.41MB")
	t.Logf("%.2fKB", d) // 2500003.84KB

}
