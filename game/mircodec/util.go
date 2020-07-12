package mircodec

import "github.com/yenkeia/yams/game/proto/server"

func isNull(ui *server.UserItem) bool {
	if ui == nil || (ui.ID == 0 && ui.ItemID == 0) {
		return true
	}
	return false
}
