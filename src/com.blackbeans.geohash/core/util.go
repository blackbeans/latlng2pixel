package core

import (
	"container/list"
	"math"
)

const (
	radius float64 = 6378137
	minLat float64 = -85.05112878
	maxLat float64 = 85.05112878
	minLng float64 = -180
	maxLng float64 = 180
)

/**
 * 修改location的经纬度以适合墨托投影
 */
func clip(val float64, minVal float64, maxVal float64) float64 {
	return math.Min(maxVal, math.Max(val, minVal))
}

/**
 * 这里将经纬度转换成格子的xy坐标
 */
func LatLng2Tile(lat float64, lng float64, level int) (tileX uint32, tileY uint32) {
	lat = math.Min(maxLat, math.Max(lat, minLat))
	lng = math.Min(maxLng, math.Max(lng, minLng))
	//投影为平面坐标
	x := (lng + 180) / 360
	sval := math.Sin(lat * math.Pi / 180)
	y := 0.5 - math.Log((1+sval)/(1-sval))/(4*math.Pi)

	//算出总共该的像总素数
	blockSize := float64(uint32(256 << uint(level)))

	//求取块中的像素坐标
	pixelX := clip(x*blockSize+0.5, 0, blockSize-1)
	pixelY := clip(y*blockSize+0.5, 0, blockSize-1)

	tileX = uint32(pixelX / 256)
	tileY = uint32(pixelY / 256)
	return

}

/**
 *
 * 10进制两个数交叉形成新数
 * 并以4进制的形式输出
 * 用于生成tile索引前缀
 */
func GenBlockKey(a uint32, b uint32, level int) uint32 {

	stack := list.New()
	for i := 0; i < level; i++ {

		var bita uint32 = a & 0x01
		var bitb uint32 = b & 0x01

		a = a >> 1
		b = b >> 1
		c := (bita<<1 | bitb) % 4

		stack.PushFront(c)
	}

	quadkey := uint32(0)

	//翻转显示
	for stack.Len() > 0 {
		head := stack.Front()
		val := head.Value.(uint32)
		quadkey = quadkey*10 + val
		stack.Remove(head)
	}

	return quadkey
}
