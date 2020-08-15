package orm

import (
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/proto/server"
)

// MapInfo 记录地图信息
type MapInfo struct {
	ID              int    `gorm:"primary_key"`
	Filename        string `gorm:"Column:file_name"`
	Title           string
	MiniMap         int
	BigMap          int
	Music           int
	Light           int
	MapDarkLight    int
	MineIndex       int
	NoTeleport      bool
	NoReconnect     bool
	NoRandom        bool
	NoEscape        bool
	NoRecall        bool
	NoDrug          bool
	NoPosition      bool
	NoFight         bool
	NoThrowItem     bool
	NoDropPlayer    bool
	NoDropMonster   bool
	NoNames         bool
	NoMount         bool
	NeedBridle      bool
	Fight           bool
	Fire            bool
	Lightning       bool
	NoTownTeleport  int
	NoReincarnation int
	NoReconnectMap  string
	FireDamage      int
	LightningDamage int
}

// ItemInfo ..
type ItemInfo struct {
	ID             int // int32 `gorm:"primary_key"`
	Name           string
	Type           cm.ItemType
	Grade          cm.ItemGrade
	RequiredType   cm.RequiredType
	RequiredClass  cm.RequiredClass
	RequiredGender cm.RequiredGender
	ItemSet        cm.ItemSet
	Shape          int // int16
	Weight         int // uint8
	Light          int // uint8
	RequiredAmount int // uint8
	Image          int // uint16
	Durability     int // uint16
	StackSize      int // uint32
	Price          int // uint32
	MinAC          int // uint8
	MaxAC          int // uint8
	MinMAC         int // uint8
	MaxMAC         int // uint8
	MinDC          int // uint8
	MaxDC          int // uint8
	MinMC          int // uint8
	MaxMC          int // uint8
	MinSC          int // uint8
	MaxSC          int // uint8
	HP             int // uint16
	MP             int // uint16
	Accuracy       int // uint8
	Agility        int // uint8
	Luck           int // int8
	AttackSpeed    int // int8
	StartItem      bool
	BagWeight      int // uint8
	HandWeight     int // uint8
	WearWeight     int // uint8
	Effect         int // uint8
	Strong         int // uint8
	MagicResist    int // uint8
	PoisonResist   int // uint8
	HealthRecovery int // uint8
	ManaRecovery   int // uint8
	PoisonRecovery int // uint8
	HpRate         int // uint8 // C# HRate
	MpRate         int // uint8 // C# MRate
	CriticalRate   int // uint8
	CriticalDamage int // uint8
	Bools          int // uint8
	MaxAcRate      int // uint8
	MaxMacRate     int // uint8
	Holy           int // uint8
	Freezing       int // uint8
	PoisonAttack   int // uint8
	Bind           int // uint16
	Reflect        int // uint8
	HpDrainRate    int // uint8
	UniqueItem     int // int16
	RandomStatsID  int // uint8
	CanFastRun     bool
	CanAwakening   bool
	IsToolTip      bool
	ToolTip        string
	ClassBased     bool `gorm:"-"`
	LevelBased     bool `gorm:"-"`
}

// ToServerItemInfo 直接从数据库使用的 orm.ItemInfo 转换成和客户端沟通使用的 server.ItemInfo
func (info *ItemInfo) ToServerItemInfo() *server.ItemInfo {
	return &server.ItemInfo{
		ID:             int32(info.ID),             // int32 `gorm:"primary_key"`
		Name:           info.Name,                  // string
		Type:           info.Type,                  // cm.ItemType
		Grade:          info.Grade,                 // cm.ItemGrade
		RequiredType:   info.RequiredType,          // cm.RequiredType
		RequiredClass:  info.RequiredClass,         // cm.RequiredClass
		RequiredGender: info.RequiredGender,        // cm.RequiredGender
		ItemSet:        info.ItemSet,               // cm.ItemSet
		Shape:          int16(info.Shape),          // int16
		Weight:         uint8(info.Weight),         // uint8
		Light:          uint8(info.Light),          // uint8
		RequiredAmount: uint8(info.RequiredAmount), // uint8
		Image:          uint16(info.Image),         // uint16
		Durability:     uint16(info.Durability),    // uint16
		StackSize:      uint32(info.StackSize),     // uint32
		Price:          uint32(info.Price),         // uint32
		MinAC:          uint8(info.MinAC),          // uint8
		MaxAC:          uint8(info.MaxAC),          // uint8
		MinMAC:         uint8(info.MinMAC),         // uint8
		MaxMAC:         uint8(info.MaxMAC),         // uint8
		MinDC:          uint8(info.MinDC),          // uint8
		MaxDC:          uint8(info.MaxDC),          // uint8
		MinMC:          uint8(info.MinMC),          // uint8
		MaxMC:          uint8(info.MaxMC),          // uint8
		MinSC:          uint8(info.MinSC),          // uint8
		MaxSC:          uint8(info.MaxSC),          // uint8
		HP:             uint16(info.HP),            // uint16
		MP:             uint16(info.MP),            // uint16
		Accuracy:       uint8(info.Accuracy),       // uint8
		Agility:        uint8(info.Agility),        // uint8
		Luck:           int8(info.Luck),            // int8
		AttackSpeed:    int8(info.AttackSpeed),     // int8
		StartItem:      info.StartItem,             // bool
		BagWeight:      uint8(info.BagWeight),      // uint8
		HandWeight:     uint8(info.HandWeight),     // uint8
		WearWeight:     uint8(info.WearWeight),     // uint8
		Effect:         uint8(info.Effect),         // uint8
		Strong:         uint8(info.Strong),         // uint8
		MagicResist:    uint8(info.MagicResist),    // uint8
		PoisonResist:   uint8(info.PoisonResist),   // uint8
		HealthRecovery: uint8(info.HealthRecovery), // uint8
		ManaRecovery:   uint8(info.ManaRecovery),   // uint8
		PoisonRecovery: uint8(info.PoisonRecovery), // uint8
		HpRate:         uint8(info.HpRate),         // uint8 // C# HRate
		MpRate:         uint8(info.MpRate),         // uint8 // C# MRate
		CriticalRate:   uint8(info.CriticalRate),   // uint8
		CriticalDamage: uint8(info.CriticalDamage), // uint8
		Bools:          uint8(info.Bools),          // uint8
		MaxAcRate:      uint8(info.MaxAcRate),      // uint8
		MaxMacRate:     uint8(info.MaxMacRate),     // uint8
		Holy:           uint8(info.Holy),           // uint8
		Freezing:       uint8(info.Freezing),       // uint8
		PoisonAttack:   uint8(info.PoisonAttack),   // uint8
		Bind:           uint16(info.Bind),          // uint16
		Reflect:        uint8(info.Reflect),        // uint8
		HpDrainRate:    uint8(info.HpDrainRate),    // uint8
		UniqueItem:     int16(info.UniqueItem),     // int16
		RandomStatsId:  uint8(info.RandomStatsID),  // uint8
		CanFastRun:     info.CanFastRun,            // bool
		CanAwakening:   info.CanAwakening,          // bool
		IsToolTip:      info.IsToolTip,             // bool
		ToolTip:        info.ToolTip,               // string
	}
}

// NPCInfo ...
type NPCInfo struct {
	ID        int `gorm:"primary_key"`
	MapID     int
	Filename  string `gorm:"Column:file_name"`
	Name      string
	Image     int
	LocationX int `gorm:"Column:location_x"`
	LocationY int `gorm:"Column:location_y"`
}

// MonsterInfo ..
type MonsterInfo struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Image       int
	AI          int `gorm:"Column:ai"`
	Effect      int
	Level       int
	ViewRange   int
	CoolEye     int
	HP          int `gorm:"Column:hp"`
	MinAC       int
	MaxAC       int
	MinMAC      int
	MaxMAC      int
	MinDC       int
	MaxDC       int
	MinMC       int
	MaxMC       int
	MinSC       int
	MaxSC       int
	Accuracy    int
	Agility     int
	Light       int
	AttackSpeed int
	MoveSpeed   int
	Experience  int
	CanPush     int
	CanTame     int
	AutoRev     int
	Undead      int
}

// RespawnInfo 刷新信息
type RespawnInfo struct {
	ID        int `gorm:"primary_key"`
	MapID     int // 地图 ID
	MonsterID int // MonsterInfo.ID
	LocationX int
	LocationY int
	Count     int // 数量
	Spread    int // 范围
	Interval  int // 刷新时间（秒
}

// BaseStats 各职业的基础属性
type BaseStats struct {
	ID                  int     `gorm:"primary_key"` // 对应 cm.MirClass
	HPGain              float32 `gorm:"Column:hp_gain"`
	HPGainRate          float32 `gorm:"Column:hp_gain_rate"`
	MPGainRate          float32 `gorm:"Column:mp_gain_rate"`
	BagWeightGain       float32
	WearWeightGain      float32
	HandWeightGain      float32
	MinAC               int
	MaxAC               int
	MinMAC              int
	MaxMAC              int
	MinDC               int
	MaxDC               int
	MinMC               int
	MaxMC               int
	MinSC               int
	MaxSC               int
	StartAgility        int
	StartAccuracy       int
	StartCriticalRate   int
	StartCriticalDamage int
	CritialRateGain     float32
	CriticalDamageGain  float32
}

func (bs *BaseStats) ToServerBaseStats() *server.BaseStats {
	return &server.BaseStats{
		HPGain:              bs.HPGain,
		HPGainRate:          bs.HPGainRate,
		MPGainRate:          bs.MPGainRate,
		BagWeightGain:       bs.BagWeightGain,
		WearWeightGain:      bs.WearWeightGain,
		HandWeightGain:      bs.HandWeightGain,
		MinAC:               uint8(bs.MinAC),
		MaxAC:               uint8(bs.MaxAC),
		MinMAC:              uint8(bs.MinMAC),
		MaxMAC:              uint8(bs.MaxMAC),
		MinDC:               uint8(bs.MinDC),
		MaxDC:               uint8(bs.MaxDC),
		MinMC:               uint8(bs.MinMC),
		MaxMC:               uint8(bs.MaxMC),
		MinSC:               uint8(bs.MinSC),
		MaxSC:               uint8(bs.MaxSC),
		StartAgility:        uint8(bs.StartAgility),
		StartAccuracy:       uint8(bs.StartAccuracy),
		StartCriticalRate:   uint8(bs.StartCriticalRate),
		StartCriticalDamage: uint8(bs.StartCriticalDamage),
		CritialRateGain:     bs.CritialRateGain,
		CriticalDamageGain:  bs.CriticalDamageGain,
	}
}

// LevelMaxExperience 玩家等级和最大经验值对应关系
type LevelMaxExperience struct {
	ID            int `gorm:"primary_key"` // 对应等级
	MaxExperience int
}

// MagicInfo 魔法信息
type MagicInfo struct {
	ID              int `gorm:"primary_key"`
	Name            string
	Spell           int
	BaseCost        int
	LevelCost       int
	Icon            int
	Level1          int
	Level2          int
	Level3          int
	Need1           int
	Need2           int
	Need3           int
	DelayBase       int
	DelayReduction  int
	PowerBase       int
	PowerBonus      int
	MPowerBase      int
	MPowerBonus     int
	MagicRange      int
	MultiplierBase  float32
	MultiplierBonus float32
}
