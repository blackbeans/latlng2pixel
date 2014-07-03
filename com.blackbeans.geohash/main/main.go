package main

import (
	// "fmt"
	"latlng2pixel/com.blackbeans.geohash/core"
	"net/http"
	"strconv"
)

var blcokMap map[uint32]*core.BlockRange = make(map[uint32]*core.BlockRange)
var initLevel int = 12

var lowestLevel int = initLevel

const (
	maxBlockPeople = 60 * 1000
)

func main() {

	// UpdateLocation(&core.LocationEntry{MomoId: "a", GeoLoc: *&core.Location{0, 90}})
	// valB := UpdateLocation(&core.LocationEntry{MomoId: "b", GeoLoc: *&core.Location{2.4123456, 90.323}})
	// UpdateLocation(&core.LocationEntry{MomoId: "c", GeoLoc: *&core.Location{2.4123453, 90.32300001}})
	// UpdateLocation(&core.LocationEntry{MomoId: "d", GeoLoc: *&core.Location{33, 121.12}})
	// block := UpdateLocation(&core.LocationEntry{MomoId: "e", GeoLoc: *&core.Location{2.4123455, 90.3230000011}})
	// fmt.Println(valB.IdxKey)
	// for _, p := range block.Peoples {
	// 	fmt.Println(p.MomoId)
	// }

	// x, y := core.LatLng2Tile(40.17887331434695, 116.3671875, initLevel)
	// idkey := core.GenBlockKey(x, y, int(initLevel))

	http.HandleFunc("/q", func(resp http.ResponseWriter, req *http.Request) {

		latV := req.FormValue("lat")
		lngV := req.FormValue("lng")

		lat, _ := strconv.ParseFloat(latV, 9)

		lng, _ := strconv.ParseFloat(lngV, 9)

		x, y := core.LatLng2Tile(lat, lng, initLevel)
		idkey := core.GenBlockKey(x, y, initLevel)

		val := "{\"key\":\"" + idkey + "\"}"

		resp.Header().Set("content-type", "text/json")

		resp.Write([]byte(val))

		defer req.Body.Close()

		// fmt.Println(val)

	})

	http.ListenAndServe(":8999", nil)

}

// func UpdateLocation(entry *core.LocationEntry) (tile *core.BlockRange) {

// 	startLevel := initLevel
// 	//这里根据这个块中的数据的${maxBlockPeople}进行下次分裂
// 	x, y := core.LatLng2Tile(entry.GeoLoc.Lat, entry.GeoLoc.Lng, startLevel)
// 	//获取坐标块的key
// 	idxKey := core.GenBlockKey(y, x, int(startLevel))

// 	//判断是否存在该块，存在判断是否该用户的location存在

// 	val, ok := blcokMap[idxKey]
// 	if !ok {
// 		//该块刚建立，则需要将用户写入即可
// 		val = &core.BlockRange{TileX: x, TileY: y, IdxKey: idxKey, Peoples: make(map[string]*core.LocationEntry)}
// 		blcokMap[idxKey] = val
// 		val.Peoples[entry.MomoId] = entry
// 	} else {
// 		_, exists := val.Peoples[entry.MomoId]
// 		//如果这个用户不存在并且当前块中的人数大于等于块中最大人数
// 		//则分裂
// 		if !exists && len(val.Peoples) >= maxBlockPeople {

// 		} else {
// 			//更新人的位置
// 			val.Peoples[entry.MomoId] = entry
// 		}
// 	}

// 	return val
// }
