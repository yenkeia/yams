package mircodec

import (
	"encoding/binary"
	"errors"
	"reflect"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"github.com/davyxu/golog"
	"github.com/yenkeia/yams/game/proto/server"
)

var log = golog.New("codec.mirCodec")

var (
	ErrInvalidType = errors.New("invalid type")
	ErrOutOfData   = errors.New("out of data")
)

func init() {
	codec.RegisterCodec(new(mirCodec))
}

// mirCodec 编码解码
type mirCodec struct{}

// Name 返回名字
func (m *mirCodec) Name() string {
	return "mirCodec"
}

// MimeType 我也不知道是干嘛的
func (m *mirCodec) MimeType() string {
	return "application/binary"
}

// Decode 将字节数组转换为结构体
// 作用是将客户端发来的字节转换为 client.Packet
func (*mirCodec) Decode(bytes interface{}, msgObj interface{}) error {
	data := bytes.([]byte)

	if len(data) == 0 {
		return nil
	}

	v := reflect.ValueOf(msgObj)

	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
	}

	size := dataSize(v, nil)
	if size < 0 {
		return ErrInvalidType
	}

	if len(data) < size {
		return ErrOutOfData
	}

	d := &decoder{order: binary.LittleEndian, buf: data}
	d.value(v)

	return nil
}

// Encode 将结构体转换为字节数组
// 作用是将 server.Packet 转换为字节发送给客户端
func (*mirCodec) Encode(msgObj interface{}, ctx cellnet.ContextSet) (data interface{}, err error) {
	switch res := msgObj.(type) {
	case *server.UserInformation:
		return encodeUserInformation(res)
	case *server.SplitItem:
		return encodeSplitItem(res)
	case *server.PlayerInspect:
		return encodePlayerInspect(res)
	case *server.ObjectPlayer:
		return encodeObjectPlayer(res)
	case *server.ObjectNPC:
		return encodeObjectNPC(res)
	case *server.NPCResponse:
		return encodeNPCResponse(res)
	case *server.TradeItem:
		return encodeTradeItem(res)
	default:
		return encode(msgObj)
	}
}

func encode(msgObj interface{}) (data interface{}, err error) {
	v := reflect.Indirect(reflect.ValueOf(msgObj))
	size := dataSize(v, nil)
	if size < 0 {
		return nil, ErrInvalidType
	}
	buf := make([]byte, size)
	e := &encoder{order: binary.LittleEndian, buf: buf}
	e.value(v)
	return buf, nil
}

type wrapper struct {
	buf []byte
}

func (w *wrapper) Write(obj interface{}) {
	data, err := encode(obj)
	if err != nil {
		panic(err)
	}
	bytes := data.([]byte)
	w.buf = append(w.buf, bytes...)
}

func encodeUserInformation(msg *server.UserInformation) (data interface{}, err error) {
	ui := msg

	writer := &wrapper{buf: make([]byte, 0)}
	writer.Write(ui.ObjectID)
	writer.Write(ui.RealID)
	writer.Write(ui.Name)
	writer.Write(ui.GuildName)
	writer.Write(ui.GuildRank)
	writer.Write(ui.NameColor)
	writer.Write(ui.Class)
	writer.Write(ui.Gender)
	writer.Write(ui.Level)
	writer.Write(ui.Location.X)
	writer.Write(ui.Location.Y)
	writer.Write(ui.Direction)
	writer.Write(ui.Hair)
	writer.Write(ui.HP)
	writer.Write(ui.MP)
	writer.Write(ui.Experience)
	writer.Write(ui.MaxExperience)
	writer.Write(ui.LevelEffect)

	// Inventory
	hasInventory := true
	if ui.Inventory == nil || len(ui.Inventory) == 0 {
		hasInventory = false
	}
	writer.Write(hasInventory)
	if hasInventory {
		l := len(ui.Inventory)
		//l := 46
		writer.Write(int32(l))
		for i := 0; i < l; i++ {
			hasUserItem := !isNull(ui.Inventory[i])
			writer.Write(hasUserItem)
			if !hasUserItem {
				continue
			}
			writer.Write(ui.Inventory[i])
		}
	}

	// Equipment
	hasEquipment := true
	if ui.Equipment == nil || len(ui.Equipment) == 0 {
		hasEquipment = false
	}
	writer.Write(hasEquipment)
	if hasEquipment {
		l := len(ui.Equipment)
		//l := 14
		writer.Write(int32(l))
		for i := 0; i < l; i++ {
			hasUserItem := !isNull(ui.Equipment[i])
			writer.Write(hasUserItem)
			if !hasUserItem {
				continue
			}
			writer.Write(ui.Equipment[i])
		}
	}

	// QuestInventory
	hasQuestInventory := true
	if ui.QuestInventory == nil || len(ui.QuestInventory) == 0 {
		hasQuestInventory = false
	}
	writer.Write(hasQuestInventory)
	if hasQuestInventory {
		l := len(ui.QuestInventory)
		//l := 40
		writer.Write(int32(l))
		for i := 0; i < l; i++ {
			hasUserItem := !isNull(ui.QuestInventory[i])
			writer.Write(hasUserItem)
			if !hasUserItem {
				continue
			}
			writer.Write(ui.QuestInventory[i])
		}
	}
	writer.Write(ui.Gold)
	writer.Write(ui.Credit)
	writer.Write(ui.HasExpandedStorage)
	writer.Write(ui.ExpandedStorageExpiryTime)

	count := len(ui.ClientMagics)
	writer.Write(int32(count))
	for i := range ui.ClientMagics {
		writer.Write(ui.ClientMagics[i])
	}

	// IntelligentCreature 直接填充字节数组, 功能不做
	writer.Write([]byte{0, 0, 0, 0, 99, 0})

	return writer.buf, nil
}

func encodeSplitItem(msg *server.SplitItem) (data interface{}, err error) {
	writer := &wrapper{buf: make([]byte, 0)}
	if msg.Item != nil {
		writer.Write(true)
		writer.Write(msg.Item)
	} else {
		writer.Write(false)
	}
	writer.Write(msg.Grid)
	return writer.buf, nil
}

func encodePlayerInspect(msg *server.PlayerInspect) (data interface{}, err error) {
	pi := msg
	writer := &wrapper{buf: make([]byte, 0)}
	writer.Write(pi.Name)
	writer.Write(pi.GuildName)
	writer.Write(pi.GuildRank)
	// Equipment
	l := len(pi.Equipment)
	if l != 14 {
		panic("equipment != 14")
	}
	//l := 14
	writer.Write(int32(l))
	for i := 0; i < l; i++ {
		hasUserItem := !isNull(pi.Equipment[i])
		writer.Write(hasUserItem)
		if !hasUserItem {
			continue
		}
		writer.Write(pi.Equipment[i])
	}

	writer.Write(pi.Class)
	writer.Write(pi.Gender)
	writer.Write(pi.Hair)
	writer.Write(pi.Level)
	writer.Write(pi.LoverName)
	return writer.buf, nil
}

func encodeObjectPlayer(msg *server.ObjectPlayer) (data interface{}, err error) {
	op := msg
	writer := &wrapper{buf: make([]byte, 0)}
	writer.Write(op.ObjectID)
	writer.Write(op.Name)
	writer.Write(op.GuildName)
	writer.Write(op.GuildRankName)
	writer.Write(op.NameColor)
	writer.Write(op.Class)
	writer.Write(op.Gender)
	writer.Write(op.Level)
	writer.Write(op.Location.X)
	writer.Write(op.Location.Y)
	writer.Write(op.Direction)
	writer.Write(op.Hair)
	writer.Write(op.Light)
	writer.Write(op.Weapon)
	writer.Write(op.WeaponEffect)
	writer.Write(op.Armour)
	writer.Write(op.Poison)
	writer.Write(op.Dead)
	writer.Write(op.Hidden)
	writer.Write(op.Effect)
	writer.Write(op.WingEffect)
	writer.Write(op.Extra)
	writer.Write(op.MountType)
	writer.Write(op.RidingMount)
	writer.Write(op.Fishing)
	writer.Write(op.TransformType)
	writer.Write(op.ElementOrbEffect)
	writer.Write(op.ElementOrbLvl)
	writer.Write(op.ElementOrbMax)
	bc := len(op.Buffs)
	writer.Write(int32(bc))
	for i := range op.Buffs {
		b := op.Buffs[i]
		writer.Write(b)
	}
	writer.Write(op.LevelEffects)
	return writer.buf, nil
}

func encodeObjectNPC(msg *server.ObjectNPC) (data interface{}, err error) {
	on := msg
	writer := &wrapper{buf: make([]byte, 0)}
	writer.Write(on.ObjectID)
	writer.Write(on.Name)
	writer.Write(on.NameColor)
	writer.Write(on.Image)
	writer.Write(on.Color)
	writer.Write(on.Location.X)
	writer.Write(on.Location.Y)
	writer.Write(uint8(on.Direction))
	qc := len(on.QuestIDs)
	writer.Write(int32(qc))
	for i := range on.QuestIDs {
		writer.Write(on.QuestIDs[i])
	}
	return writer.buf, nil
}

func encodeNPCResponse(msg *server.NPCResponse) (data interface{}, err error) {
	res := msg
	writer := &wrapper{buf: make([]byte, 0)}
	count := len(res.Page)
	writer.Write(int32(count))
	for i := 0; i < count; i++ {
		writer.Write(res.Page[i])
	}
	return writer.buf, nil
}

func encodeTradeItem(msg *server.TradeItem) (data interface{}, err error) {
	writer := &wrapper{buf: make([]byte, 0)}
	length := len(msg.TradeItems)
	writer.Write(int32(length))
	for i := 0; i < length; i++ {
		ui := msg.TradeItems[i]
		if ui == nil {
			writer.Write(false)
		} else {
			writer.Write(true)
			writer.Write(ui)
		}
	}
	return writer.buf, nil
}
