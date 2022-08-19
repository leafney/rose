/**
 * @Author:      leafney
 * @Date:        2022-06-30 16:41
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import "math"

// GetDistanceByLocation 获取两个坐标点之间的距离（返回距离单位km
func GetDistanceByLocation(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6371000.0
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius / 1000
}
