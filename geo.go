/**
 * @Author:      leafney
 * @Date:        2022-06-30 16:41
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import "math"

/*
  WGS84坐标系：即地球坐标系，国际上通用的坐标系。
  GCJ02坐标系：即火星坐标系，WGS84坐标系经加密后的坐标系。如：Google Maps中国版，高德地图、腾讯地图、苹果地图中国版
  BD09坐标系：即百度坐标系，GCJ02坐标系经加密后的坐标系。如：百度地图
*/

const (
	xPi     = math.Pi * 3000.0 / 180.0
	xOffset = 0.00669342162296594323 // 扁率
	xAxis   = 6378245.0              // 长半轴
)

// GeoWGS84toGCJ02 WGS84坐标系->火星坐标系
func GeoWGS84toGCJ02(lng, lat float64) (nLng, nLat float64) {
	if !GeoInChina(lng, lat) {
		return lng, lat
	}
	nLng, nLat = delta(lng, lat)
	return
}

// GeoGCJ02toWGS84 火星坐标系->WGS84坐标系
func GeoGCJ02toWGS84(lng, lat float64) (nLng, nLat float64) {
	if !GeoInChina(lng, lat) {
		return lng, lat
	}

	nLng, nLat = delta(lng, lat)
	return lng*2 - nLng, lat*2 - nLat
}

// GeoWGS84toBD09 WGS84坐标系->百度坐标系
func GeoWGS84toBD09(lng, lat float64) (nLng, nLat float64) {
	tLng, tLat := GeoWGS84toGCJ02(lng, lat)
	return GeoGCJ02toBD09(tLng, tLat)
}

// GeoBD09toWGS84 百度坐标系->WGS84坐标系
func GeoBD09toWGS84(lng, lat float64) (nLng, nLat float64) {
	tLng, tLat := GeoBD09toGCJ02(lng, lat)
	return GeoGCJ02toWGS84(tLng, tLat)
}

// GeoBD09toGCJ02 百度坐标系->火星坐标系
func GeoBD09toGCJ02(lng, lat float64) (nLng, nLat float64) {
	x := lng - 0.0065
	y := lat - 0.006

	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*xPi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*xPi)

	nLng = z * math.Cos(theta)
	nLat = z * math.Sin(theta)
	return
}

// GeoGCJ02toBD09 火星坐标系->百度坐标系
func GeoGCJ02toBD09(lng, lat float64) (nLng, nLat float64) {
	z := math.Sqrt(lng*lng+lat*lat) + 0.00002*math.Sin(lat*xPi)
	theta := math.Atan2(lat, lng) + 0.000003*math.Cos(lng*xPi)

	nLng = z*math.Cos(theta) + 0.0065
	nLat = z*math.Sin(theta) + 0.006
	return
}

func delta(lng, lat float64) (float64, float64) {
	dLat, dLng := transform(lng-105.0, lat-35.0)
	radLat := lat / 180.0 * math.Pi
	magic := math.Sin(radLat)
	magic = 1 - xOffset*magic*magic
	sqrtMagic := math.Sqrt(magic)

	dLat = (dLat * 180.0) / ((xAxis * (1 - xOffset)) / (magic * sqrtMagic) * math.Pi)
	dLng = (dLng * 180.0) / (xAxis / sqrtMagic * math.Cos(radLat) * math.Pi)

	mgLat := lat + dLat
	mgLng := lng + dLng

	return mgLng, mgLat
}

func transform(lng, lat float64) (x, y float64) {
	var lngLat = lng * lat
	var absX = math.Sqrt(math.Abs(lng))
	var lngPi, latPi = lng * math.Pi, lat * math.Pi
	var d = 20.0*math.Sin(6.0*lngPi) + 20.0*math.Sin(2.0*lngPi)
	x, y = d, d
	x += 20.0*math.Sin(latPi) + 40.0*math.Sin(latPi/3.0)
	y += 20.0*math.Sin(lngPi) + 40.0*math.Sin(lngPi/3.0)
	x += 160.0*math.Sin(latPi/12.0) + 320*math.Sin(latPi/30.0)
	y += 150.0*math.Sin(lngPi/12.0) + 300.0*math.Sin(lngPi/30.0)
	x *= 2.0 / 3.0
	y *= 2.0 / 3.0
	x += -100.0 + 2.0*lng + 3.0*lat + 0.2*lat*lat + 0.1*lngLat + 0.2*absX
	y += 300.0 + lng + 2.0*lat + 0.1*lng*lng + 0.1*lngLat + 0.1*absX
	return
}

// GeoInChina 判断坐标是否在中国境内
func GeoInChina(lng, lat float64) bool {
	//return lng > 72.004 && lng < 135.05 && lat > 3.86 && lat < 53.55
	return lng >= 72.004 && lng <= 137.8347 && lat >= 0.8293 && lat <= 55.8271
}

// GeoLocationDistance 获取两个坐标点之间的距离（距离单位(米)
func GeoLocationDistance(lng1, lat1, lng2, lat2 float64) float64 {
	radius := 6371000.0
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}

/*
// 第二种实现方式

// WGS84转GCJ02(火星坐标系)；
// 一般情况下，google地图、高德地图、腾讯地图，使用这个坐标；
func GeoWgs84ToGcj02(lng, lat float64) (float64, float64) {
	if !GeoInChina(lng, lat) {
		return lng, lat
	}
	dLat := transformLat(lng-105.0, lat-35.0)
	dLng := transformLng(lng-105.0, lat-35.0)
	radLat := lat / 180.0 * math.Pi
	magic := math.Sin(radLat)
	magic = 1 - xOffset*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((xAxis * (1 - xOffset)) / (magic * sqrtMagic) * math.Pi)
	dLng = (dLng * 180.0) / (xAxis / sqrtMagic * math.Cos(radLat) * math.Pi)
	return lng + dLng, lat + dLat
}

// GCJ02(火星坐标系)转WGS84
func GeoGcj02ToWps84(lng, lat float64) (float64, float64) {
	if !GeoInChina(lng, lat) {
		return lng, lat
	}
	dLat := transformLat(lng-105.0, lat-35.0)
	dLng := transformLng(lng-105.0, lat-35.0)
	radLat := lat / 180.0 * math.Pi
	magic := math.Sin(radLat)
	magic = 1 - xOffset*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((xAxis * (1 - xOffset)) / (magic * sqrtMagic) * math.Pi)
	dLng = (dLng * 180.0) / (xAxis / sqrtMagic * math.Cos(radLat) * math.Pi)
	return lng*2 - (lng + dLng), lat*2 - (lat + dLat)
}

// 火星坐标转百度坐标
func GeoGcj02ToBd09(lng, lat float64) (float64, float64) {
	z := math.Sqrt(lng*lng+lat*lat) + 0.00002*math.Sin(lat*xPi)
	theta := math.Atan2(lat, lng) + 0.000003*math.Cos(lng*xPi)
	return z*math.Cos(theta) + 0.0065, z*math.Sin(theta) + 0.006
}

// 百度坐标转火星
func GeoBd09ToGcj02(lng, lat float64) (float64, float64) {
	x := lng - 0.0065
	y := lat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*xPi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*xPi)
	return z * math.Cos(theta), z * math.Sin(theta)
}

// wgs84转百度
func GeoWgs84ToBd09(lng, lat float64) (float64, float64) {
	return GeoGcj02ToBd09(GeoWgs84ToGcj02(lng, lat))
}

// 百度转wgs84
func GeoBd09ToWgs84(lng, lat float64) (float64, float64) {
	return GeoGcj02ToWps84(GeoBd09ToGcj02(lng, lat))
}

func transformLat(lng, lat float64) float64 {
	ret := -100.0 + 2.0*lng + 3.0*lat + 0.2*lat*lat + 0.1*lng*lat + 0.2*math.Sqrt(math.Abs(lng))
	ret += (20.0*math.Sin(6.0*lng*math.Pi) + 20.0*math.Sin(2.0*lng*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lat*math.Pi) + 40.0*math.Sin(lat/3.0*math.Pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(lat/12.0*math.Pi) + 320*math.Sin(lat*math.Pi/30.0)) * 2.0 / 3.0
	return ret
}

func transformLng(lng, lat float64) float64 {
	ret := 300.0 + lng + 2.0*lat + 0.1*lng*lng + 0.1*lng*lat + 0.1*math.Sqrt(math.Abs(lng))
	ret += (20.0*math.Sin(6.0*lng*math.Pi) + 20.0*math.Sin(2.0*lng*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lng*math.Pi) + 40.0*math.Sin(lng/3.0*math.Pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(lng/12.0*math.Pi) + 300.0*math.Sin(lng/30.0*math.Pi)) * 2.0 / 3.0
	return ret
}

*/
