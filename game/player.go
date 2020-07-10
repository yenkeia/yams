package game

import "github.com/yenkeia/yams/game/proto/client"

type player struct {
	gameStage int
	accountID int
	objectID  int
}

func (p *player) id() int {
	return p.objectID
}

func (p *player) turn(msg *client.Turn)                             {}
func (p *player) walk(msg *client.Walk)                             {}
func (p *player) run(msg *client.Run)                               {}
func (p *player) chat(msg *client.Chat)                             {}
func (p *player) moveItem(msg *client.MoveItem)                     {}
func (p *player) storeItem(msg *client.StoreItem)                   {}
func (p *player) depositRefineItem(msg *client.DepositRefineItem)   {}
func (p *player) retrieveRefineItem(msg *client.RetrieveRefineItem) {}
func (p *player) refineCancel(msg *client.RefineCancel)             {}
func (p *player) refineItem(msg *client.RefineItem)                 {}
func (p *player) checkRefine(msg *client.CheckRefine)               {}
func (p *player) replaceWedRing(msg *client.ReplaceWedRing)         {}
func (p *player) depositTradeItem(msg *client.DepositTradeItem)     {}
func (p *player) retrieveTradeItem(msg *client.RetrieveTradeItem)   {}
func (p *player) takeBackItem(msg *client.TakeBackItem)             {}
func (p *player) mergeItem(msg *client.MergeItem)                   {}
func (p *player) equipItem(msg *client.EquipItem)                   {}
func (p *player) removeItem(msg *client.RemoveItem)                 {}
func (p *player) removeSlotItem(msg *client.RemoveSlotItem)         {}
func (p *player) splitItem(msg *client.SplitItem)                   {}
func (p *player) useItem(msg *client.UseItem)                       {}
func (p *player) dropItem(msg *client.DropItem)                     {}
func (p *player) dropGold(msg *client.DropGold)                     {}
func (p *player) pickUp(msg *client.PickUp)                         {}
func (p *player) inspect(msg *client.Inspect)                       {}
