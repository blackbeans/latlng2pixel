package core

import (
	"container/list"
	// "fmt"
	"math"
	"strconv"
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

func groundResolution(lat float64, level int) float64 {
	lat = clip(lat, minLat, maxLat)
	return math.Cos(lat*math.Pi/180) * 2 * math.Pi * radius / float64(mapsize(level))
}

func mapsize(level int) uint {
	return uint(256) << uint(level)
}

/**
 * 这里将经纬度转换成格子的xy坐标
 */
func LatLng2Tile(lat float64, lng float64, level int) (tileX int, tileY int) {
	lat = clip(lat, minLat, maxLat)
	lng = clip(lng, minLng, maxLng)
	//投影为平面坐标
	x := (lng + 180) / 360
	sval := math.Sin(lat * math.Pi / 180)
	y := 0.5 - math.Log((1+sval)/(1-sval))/(4*math.Pi)

	//算出总共该的像总素数
	blockSize := mapsize(level)

	//求取块中的像素坐标
	pixelX := int(clip(x*float64(blockSize)+0.5, 0, float64(blockSize)-1))
	pixelY := int(clip(y*float64(blockSize)+0.5, 0, float64(blockSize)-1))

	tileX = int(pixelX / 256)
	tileY = int(pixelY / 256)
	return

}

/**
 *
 * 10进制两个数交叉形成新数
 * 并以4进制的形式输出
 * 用于生成tile索引前缀
 */
func GenBlockKey(a int, b int, level int) string {

	stack := list.New()
	for i := level; i > 0; i-- {

		// var bita uint32 = a & 0x01
		// var bitb uint32 = b & 0x01

		digit := 0
		mask := int(1 << uint(i-1))

		if (a & mask) != 0 {
			digit++
		}

		if (b & mask) != 0 {
			digit++
			digit++
		}

		stack.PushBack(digit)
	}

	quadkey := ""

	//翻转显示
	for stack.Len() > 0 {
		head := stack.Front()
		val := head.Value.(int)
		quadkey += strconv.Itoa(val)
		stack.Remove(head)
	}

	return quadkey
}
