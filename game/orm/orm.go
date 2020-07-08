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
