package game

import "github.com/yenkeia/yams/game/proto/client"

type player struct {
	mapObject
	gameStage int
	accountID int
}

func (p *player) Turn(msg *client.Turn)                             {}
func (p *player) Walk(msg *client.Walk)                             {}
func (p *player) Run(msg *client.Run)                               {}
func (p *player) Chat(msg *client.Chat)                             {}
func (p *player) MoveItem(msg *client.MoveItem)                     {}
func (p *player) StoreItem(msg *client.StoreItem)                   {}
func (p *player) DepositRefineItem(msg *client.DepositRefineItem)   {}
func (p *player) RetrieveRefineItem(msg *client.RetrieveRefineItem) {}
func (p *player) RefineCancel(msg *client.RefineCancel)             {}
func (p *player) RefineItem(msg *client.RefineItem)                 {}
func (p *player) CheckRefine(msg *client.CheckRefine)               {}
func (p *player) ReplaceWedRing(msg *client.ReplaceWedRing)         {}
func (p *player) DepositTradeItem(msg *client.DepositTradeItem)     {}
func (p *player) RetrieveTradeItem(msg *client.RetrieveTradeItem)   {}
func (p *player) TakeBackItem(msg *client.TakeBackItem)             {}
func (p *player) MergeItem(msg *client.MergeItem)                   {}
func (p *player) EquipItem(msg *client.EquipItem)                   {}
func (p *player) RemoveItem(msg *client.RemoveItem)                 {}
func (p *player) RemoveSlotItem(msg *client.RemoveSlotItem)         {}
func (p *player) SplitItem(msg *client.SplitItem)                   {}
func (p *player) UseItem(msg *client.UseItem)                       {}
func (p *player) DropItem(msg *client.DropItem)                     {}
func (p *player) DropGold(msg *client.DropGold)                     {}
func (p *player) PickUp(msg *client.PickUp)                         {}
func (p *player) Inspect(msg *client.Inspect)                       {}
