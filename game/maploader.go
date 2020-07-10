package game

import (
	"fmt"
	"io/ioutil"

	"github.com/yenkeia/yams/game/cm"
)

func loadMap(filepath string) *mirMap {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	v := detectMapVersion(fileBytes)
	switch v {
	case 0:
		return getMapV0(fileBytes)
	case 1:
		return getMapV1(fileBytes)
	case 3:
		return getMapV3(fileBytes)
	case 5:
		return getMapV5(fileBytes)
	default:
		panic(fmt.Sprintf("map version not support! %d", int(v)))
	}
}

func detectMapVersion(input []byte) byte {
	//c# custom map format
	if (input[2] == 0x43) && (input[3] == 0x23) {
		return 100
	}
	//wemade mir3 maps have no title they just start with blank bytes
	if input[0] == 0 {
		return 5
	}
	//shanda mir3 maps start with title: (C) SNDA, MIR3.
	if (input[0] == 0x0F) && (input[5] == 0x53) && (input[14] == 0x33) {
		return 6
	}
	//wemades antihack map (laby maps) title start with: Mir2 AntiHack
	if (input[0] == 0x15) && (input[4] == 0x32) && (input[6] == 0x41) && (input[19] == 0x31) {
		return 4
	}
	//wemades 2010 map format i guess title starts with: Map 2010 Ver 1.0
	if (input[0] == 0x10) && (input[2] == 0x61) && (input[7] == 0x31) && (input[14] == 0x31) {
		return 1
	}
	//shanda's 2012 format and one of shandas(wemades) older formats share same header info, only difference is the filesize
	if (input[4] == 0x0F) && (input[18] == 0x0D) && (input[19] == 0x0A) {
		/*
			W := int(input[0] + (input[1] << 8))
			H := int(input[2] + (input[3] << 8))
			if len(input) > (52 + (W * H * 14)) {
				return 3
			}
			return 2
		*/
		panic("not support shanda's 2012 map format")
	}
	//3/4 heroes map format (myth/lifcos i guess)
	if (input[0] == 0x0D) && (input[1] == 0x4C) && (input[7] == 0x20) && (input[11] == 0x6D) {
		return 7
	}
	return 0
}

func getMapV0(bytes []byte) *mirMap {
	offset := 0
	w := cm.BytesToUint16(bytes[offset:])
	offset += 2
	h := cm.BytesToUint16(bytes[offset:])
	width := int(w)
	height := int(h)

	m := newMirMap(width, height, 0)

	offset = 52
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			cellAttr := cm.CellAttributeWalk

			if (cm.BytesToUint16(bytes[offset:]) & 0x8000) != 0 {
				// cell = HighWallCell //Can Fire Over.
				cellAttr = cm.CellAttributeHighWall
			}

			offset += 2
			if (cm.BytesToUint16(bytes[offset:]) & 0x8000) != 0 {
				// cell = LowWallCell //Can't Fire Over.
				cellAttr = cm.CellAttributeLowWall
			}

			offset += 2
			if (cm.BytesToUint16(bytes[offset:]) & 0x8000) != 0 {
				// cell = HighWallCell //No Floor Tile.
				cellAttr = cm.CellAttributeHighWall
			}

			m.setCellAttribute(x, y, cellAttr)

			offset += 4
			// TODO 地图连接门
			// if bytes[offset] > 0 {
			// 	m.AddDoor(bytes[offset], point)
			// }

			offset += 3 + 1

			// byte light = fileBytes[offSet++];

			// if (light >= 100 && light <= 119)
			// 	Cells[x, y].FishingAttribute = (sbyte)(light - 100);
		}
	}
	return m
}

func getMapV1(bytes []byte) *mirMap {
	offset := 21
	w := cm.BytesToUint16(bytes[offset:])
	offset += 2
	xor := cm.BytesToUint16(bytes[offset:])
	offset += 2
	h := cm.BytesToUint16(bytes[offset:])
	width := int(w ^ xor)
	height := int(h ^ xor)

	m := newMirMap(width, height, 1)

	offset = 54
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			cellAttr := cm.CellAttributeWalk

			if (cm.BytesToUint32(bytes[offset:])^0xAA38AA38)&0x20000000 != 0 {
				cellAttr = cm.CellAttributeHighWall
			}

			offset += 6
			if ((cm.BytesToUint16(bytes[offset:]) ^ xor) & 0x8000) != 0 {
				cellAttr = cm.CellAttributeLowWall
			}

			m.setCellAttribute(x, y, cellAttr)

			offset += 2
			// TODO 地图连接门
			// if bytes[offset] > 0 {
			// 	m.AddDoor(bytes[offset], point)
			// }

			offset += 5

			// byte light = fileBytes[offSet++];
			// if (light >= 100 && light <= 119)
			// 	Cells[x, y].FishingAttribute = (sbyte)(light - 100);
			offset += 1 + 1
		}
	}
	return m
}

func getMapV3(bytes []byte) *mirMap {
	offset := 0
	w := cm.BytesToUint16(bytes[offset:])
	offset += 2
	h := cm.BytesToUint16(bytes[offset:])
	width := int(w)
	height := int(h)

	m := newMirMap(width, height, 3)

	offset = 52
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			cellAttr := cm.CellAttributeWalk

			if (cm.BytesToUint16(bytes[offset:]) & 0x8000) != 0 {
				cellAttr = cm.CellAttributeHighWall
			}

			offset += 2
			if (cm.BytesToUint16(bytes[offset:]) & 0x8000) != 0 {
				cellAttr = cm.CellAttributeLowWall
			}

			offset += 2
			if (cm.BytesToUint16(bytes[offset:]) & 0x8000) != 0 {
				cellAttr = cm.CellAttributeHighWall
			}

			m.setCellAttribute(x, y, cellAttr)

			offset += 2
			// TODO 地图连接门
			// if bytes[offset] > 0 {
			// 	m.AddDoor(bytes[offset], point)
			// }

			offset += 12

			// byte light = fileBytes[offSet++];

			// if (light >= 100 && light <= 119)
			// 	Cells[x, y].FishingAttribute = (sbyte)(light - 100);

			offset += 17 + 1
		}
	}
	return m
}

func getMapV5(bytes []byte) *mirMap {
	offset := 22
	w := cm.BytesToUint16(bytes[offset:])
	offset += 2
	h := cm.BytesToUint16(bytes[offset:])
	width := int(w)
	height := int(h)

	m := newMirMap(width, height, 5)

	offset = 28 + (3 * ((width / 2) + (width % 2)) * (height / 2))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			cellAttr := cm.CellAttributeWalk

			if (bytes[offset] & 0x01) != 1 {
				cellAttr = cm.CellAttributeHighWall
			} else if (bytes[offset] & 0x02) != 2 {
				cellAttr = cm.CellAttributeLowWall
			}

			m.setCellAttribute(x, y, cellAttr)

			offset += 13

			// byte light = fileBytes[offSet++];

			// if (light >= 100 && light <= 119)
			// 	Cells[x, y].FishingAttribute = (sbyte)(light - 100);

			offset++

		}
	}
	return m
}
