package game

const (
	MainKey           = "[@MAIN]"
	BuyKey            = "[@BUY]"
	SellKey           = "[@SELL]"
	BuySellKey        = "[@BUYSELL]"
	RepairKey         = "[@REPAIR]"
	SRepairKey        = "[@SREPAIR]"
	RefineKey         = "[@REFINE]"
	RefineCheckKey    = "[@REFINECHECK]"
	RefineCollectKey  = "[@REFINECOLLECT]"
	ReplaceWedRingKey = "[@REPLACEWEDDINGRING]"
	BuyBackKey        = "[@BUYBACK]"
	StorageKey        = "[@STORAGE]"
	ConsignKey        = "[@CONSIGN]"
	MarketKey         = "[@MARKET]"
	ConsignmentsKey   = "[@CONSIGNMENT]"
	CraftKey          = "[@CRAFT]"
	TradeKey          = "[TRADE]"
	RecipeKey         = "[RECIPE]"
	TypeKey           = "[TYPES]"
	QuestKey          = "[QUESTS]"
	GuildCreateKey    = "[@CREATEGUILD]"
	RequestWarKey     = "[@REQUESTWAR]"
	SendParcelKey     = "[@SENDPARCEL]"
	CollectParcelKey  = "[@COLLECTPARCEL]"
	AwakeningKey      = "[@AWAKENING]"
	DisassembleKey    = "[@DISASSEMBLE]"
	DowngradeKey      = "[@DOWNGRADE]"
	ResetKey          = "[@RESET]"
	PearlBuyKey       = "[@PEARLBUY]"
	BuyUsedKey        = "[@BUYUSED]"
)

type status int

const (
	SUCCESS status = 1
	FAILED  status = 2
	RUNNING status = 3
)
