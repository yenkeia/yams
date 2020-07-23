package orm

// Account 账号信息
type Account struct {
	ID       int `gorm:"primary_key"`
	Username string
	Password string
}

// AccountCharacter 账号角色关联表
type AccountCharacter struct {
	ID          int `gorm:"primary_key"`
	AccountID   int
	CharacterID int
}

// Character 角色
type Character struct {
	ID               int `gorm:"primary_key"`
	Name             string
	Level            int
	Class            int
	Gender           int
	Hair             int
	CurrentMapID     int
	CurrentLocationX int
	CurrentLocationY int
	BindMapID        int
	BindLocationX    int
	BindLocationY    int
	Direction        int
	HP               int
	MP               int
	Experience       int
	AttackMode       int
	PetMode          int
	Gold             int
	AllowGroup       bool
}

// CharacterUserItem 角色物品关系
type CharacterUserItem struct {
	ID          int `gorm:"primary_key"`
	CharacterID int
	UserItemID  int
	Type        int // 	类型: Inventory / Equipment / QuestInventory
	Index       int //	所在类型格子的索引，比如在 Inventory 的第几个格子
}

// UserItem ..
type UserItem struct {
	ID          int `gorm:"primary_key"` // UniqueID     // uint64
	ItemID      int // int32
	CurrentDura int // uint16
	MaxDura     int // uint16
	Count       int // uint32
	AC          int // uint8
	MAC         int // uint8
	DC          int // uint8
	MC          int // uint8
	SC          int // uint8
	Accuracy    int // uint8
	Agility     int // uint8
	HP          int // uint8
	MP          int // uint8
	AttackSpeed int // int8
	Luck        int // int8
	SoulBoundID int // uint32
	/*
		Bools          int // uint8
		Strong         int // uint8
		MagicResist    int // uint8
		PoisonResist   int // uint8
		HealthRecovery int // uint8
		ManaRecovery   int // uint8
		PoisonRecovery int // uint8
		CriticalRate   int // uint8
		CriticalDamage int // uint8
		Freezing       int // uint8
		PoisonAttack   int // uint8
	*/
}
