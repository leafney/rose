/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2024-03-22 22:28
 * @Description:
 */

package rose

import "testing"

func TestGeoDemo(t *testing.T) {
	//t.Log(GeoInChina(116.404, 39.915))
	//
	//t.Log(GeoBD09toWGS84(116.404, 39.915))
	//t.Log(GeoBD09toGCJ02(116.404, 39.915))

	//t.Log(GeoWGS84toGCJ02(116.404, 39.915))
	//t.Log(GeoWGS84toBD09(116.404, 39.915))

	//t.Log(Wgs84ToGcj02(116.404, 39.915))
	//t.Log(Wgs84ToBd09(116.404, 39.915))

	// 验证两种实现方式的结果是否一致
	/*
		//t.Log(GeoWgs84ToBd09(104.69313101, 31.494952300000001))
		t.Log(GeoWGS84toBD09(104.69313101, 31.494952300000001))

		//t.Log(GeoWgs84ToGcj02(104.69313101, 31.494952300000001))
		t.Log(GeoWGS84toGCJ02(104.69313101, 31.494952300000001))

		t.Log(GeoBD09toGCJ02(104.69313101, 31.494952300000001))
		//t.Log(GeoBd09ToGcj02(104.69313101, 31.494952300000001))

		t.Log(GeoGCJ02toBD09(104.69313101, 31.494952300000001))
		//t.Log(GeoGcj02ToBd09(104.69313101, 31.494952300000001))

		t.Log(GeoBD09toWGS84(104.69313101, 31.494952300000001))
		//t.Log(GeoBd09ToWgs84(104.69313101, 31.494952300000001))

		t.Log(GeoGCJ02toWGS84(104.69313101, 31.494952300000001))
		//t.Log(GeoGcj02ToWps84(104.69313101, 31.494952300000001))
	*/

	// 解决 wgs转 bd 偏移问题

	t.Log(GeoWGS84toBD09(107.8003647, 29.3347596))
	t.Log(GeoWGS84toGCJ02(107.8003647, 29.3347596))
	t.Log(GeoGCJ02toBD09(107.80473041151717, 29.33190938671697))
}
