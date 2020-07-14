package orm

import "github.com/yenkeia/yams/game/cm"

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
	NoTeleport      int
	NoReconnect     int
	NoRandom        int
	NoEscape        int
	NoRecall        int
	NoDrug          int
	NoPosition      int
	NoFight         int
	NoThrowItem     int
	NoDropPlayer    int
	NoDropMonster   int
	NoNames         int
	NoMount         int
	NeedBridle      int
	Fight           int
	Fire            int
	Lightning       int
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
	SpellRecovery  int // uint8
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
}

// NPCInfo ...
type NPCInfo struct {
	ID        int `gorm:"primary_key"`
	MapID     int
	Filename  string `gorm:"Column:file_name"`
	Name      string
	LocationX int `gorm:"Column:location_x"`
	LocationY int `gorm:"Column:location_y"`
}
