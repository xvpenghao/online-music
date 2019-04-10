package verify

import (
	"fmt"
	"online-music/common/utils"
	"online-music/models"
)

//查询歌单列表参数校验
func QuerySongCoverListReqVerify(req models.QuerySongCoverListReq) error {
	if !utils.CheckLegal(&utils.StrChecker{Value: req.ChannelId, MinLen: 1, MaxLen: 32}) {
		return fmt.Errorf("来源渠道id(%v)参数错误，取值(%v ~ %v)", req.ChannelId, 1, 32)
	}
	if !utils.CheckLegal(&utils.IntChecker{Value: req.CurPage, Min: 0, Max: utils.INT_MAX}) {
		return fmt.Errorf("当前页(%v)参数错误，取值(%v ~ %v)", req.CurPage, 0, utils.INT_MAX)
	}

	return nil
}
