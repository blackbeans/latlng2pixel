package core

import (
	"time"
)

const (
	F = iota
	M
)

/**
 * 地点
 */
type Location struct {
	Lat float64
	Lng float64
}

/**
 * location单个实体
 */
type LocationEntry struct {
	MomoId     string
	ActiveTime time.Time
	Sex        int32
	GeoCode    string
	GeoLoc     Location
}

/**
 * 本块的属性
 */
type BlockRange struct {
	TileX   uint32
	TileY   uint32
	Peoples map[string]*LocationEntry
	IdxKey  uint32
}

type BlockTreeNode struct {
	ParentRange *BlockTreeNode
	LeftChild   *BlockTreeNode
	RightChild  *BlockTreeNode
	BRang       *BlockRange
}

// /**
//  * 只有当前的经纬度在该块的区间则认为是这个block中的
//  */
// func (brange *BlockRange) Contains(lat float64, lng float64, incBound bool) bool {
// 	flag := false
// 	if incBound {
// 		flag = brange.StartLat <= lat && brange.EndLat > lat &&
// 			brange.StartLng <= lng && brange.EndLng > lng
// 	} else {
// 		flag = brange.StartLat < lat && brange.EndLat > lat &&
// 			brange.StartLng < lng && brange.EndLng > lng
// 	}
// 	return flag
// }

// func (brange *BlockRange) ToString() string {
// 	return strconv.FormatFloat(brange.StartLat, 'f', -1, 64) +
// 		"," + strconv.FormatFloat(brange.StartLng, 'f', -1, 64) + "->" +
// 		strconv.FormatFloat(brange.EndLat, 'f', -1, 64) + "," + strconv.FormatFloat(brange.EndLng, 'f', -1, 64)
// }
