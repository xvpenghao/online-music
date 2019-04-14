package impl

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gocolly/colly"
	"online-music/common/constants"
	"online-music/common/utils"
	"online-music/models"
	"online-music/service/dbModel"
	"strconv"
	"strings"
	"time"
)

type SongCoverService struct {
	BaseService
}

/*
*@Title:查询歌单列表
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/10
 */
func (receiver *SongCoverService) QuerySongCoverList(req models.QuerySongCoverListReq) ([]dbModel.SongCover, error) {
	receiver.BeforeLog("QuerySongCoverList")
	var result []dbModel.SongCover
	var songCover dbModel.SongCover
	var err error
	c := colly.NewCollector(
		colly.UserAgent(constants.USER_AGENT))

	//歌单列表的爬取
	c.OnHTML("ul[class='m-cvrlst f-cb']", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, element *colly.HTMLElement) {
			img := element.ChildAttr("div[class='u-cover u-cover-1'] img", "src")
			desTitle := element.ChildAttr("p[class='dec'] a", "title")
			playHref := element.ChildAttr("p[class='dec'] a", "href")
			songCover.CoverImgUrl = img
			songCover.Description = desTitle
			///playlist?id=2708450548，切割
			songCover.SongCoverId = strings.Split(playHref, "=")[1]
			result = append(result, songCover)
		})
	})

	c.OnError(func(response *colly.Response, e error) {
		if e != nil {
			err = e
		}
	})
	if err != nil {
		logs.Error("查询歌单列表错我：(%v)", err.Error())
		return result, err
	}

	reqUrl := fmt.Sprintf(constants.SONG_PAGE_LIST_URL, constants.DEFAULT_PAGE_SIZE, strconv.Itoa(req.CurPage))
	err = c.Visit(reqUrl)
	if err != nil {
		logs.Error("查询歌单列表-访问歌单链接错误：(%v)，歌单链接：(%v)", err.Error(), reqUrl)
		return result, err
	}

	return result, nil
}

/*
*@Title:根据歌单id查询歌曲列表
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/10
*@Param
*@Return
 */
func (receiver *SongCoverService) QuerySongList(req models.QuerySongListReq) ([]dbModel.Song, error) {
	receiver.BeforeLog("QuerySongList")
	var result []dbModel.Song
	var song dbModel.Song
	var err error

	c := colly.NewCollector(
		colly.UserAgent(constants.USER_AGENT))

	//歌曲列表
	c.OnHTML("ul[class='f-hide']", func(e *colly.HTMLElement) {
		var songId string
		e.ForEach("li", func(i int, ele *colly.HTMLElement) {
			if len(result) == 10 {
				return
			}
			songId = ele.ChildAttr("a", "href")
			///song?id=1350336759 只要id
			song.SongId = strings.Split(songId, "=")[1]
			song.SongName = ele.Text
			result = append(result, song)
		})
	})

	c.OnError(func(response *colly.Response, e error) {
		if e != nil {
			err = e
		}
	})
	if err != nil {
		logs.Error("根据歌单id查询歌曲列表错误：(%v)", err.Error())
		return result, err
	}

	//到时候根据前台传递的歌单id来得到该得到的歌曲列表
	reqUrl := fmt.Sprintf(constants.SONG_COVER_DETAIL_URL, req.SongCoverId)
	err = c.Visit(reqUrl)
	if err != nil {
		logs.Error("根据歌单id查询歌曲列表-访问链接错误:(%v),链接:(%v)", err.Error(), reqUrl)
		return result, err
	}

	return result, nil

}

/*
*@Title:创建歌单
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/14 0014
*@Param
*@Return
 */
func (receiver *SongCoverService) CreateSongCover(req models.CreateSongCoverReq) error {
	receiver.BeforeLog("CreateSongCover")

	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("创建歌单-数据库链接错误：(%v)", err.Error())
		return utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()
	//该用户的创建的自定义歌单名不能重复
	sql := dbModel.QUERY_USER_COVER_COUNT_BY_SONG_COVER_NAME
	sqlParam := []interface{}{receiver.BaseRequest.UserID, req.SongCoverName}
	var counts int
	err = db.Raw(sql, sqlParam...).Count(&counts).Error
	if err != nil {
		logs.Error("创建歌单-根据歌单名查询用户歌单失败：(%v)", err.Error())
		return utils.NewDBErr("根据歌单名查询用户歌单失败", err)
	}

	if counts > 0 {
		logs.Error("创建歌单-根据歌单名查询用户歌单，歌单名重复，歌单名称：(%v)", req.SongCoverName)
		return utils.NewSysErr("根据歌单名查询用户歌单,歌单名重复")
	}

	nowTime := time.Now()
	songCover := dbModel.SongCoverInfo{
		ID:            utils.GetUUID(),
		Type:          constants.SONG_COVER_TYPE_CUSTOMER,
		SongCoverName: req.SongCoverName,
		DelState:      constants.USER_NO_DEL_STATUS,
		CreatTime:     nowTime,
		CreateUser:    receiver.BaseRequest.UserName,
		CreateUserId:  receiver.BaseRequest.UserID,
		UpdateTime:    nowTime,
		UpdateUser:    receiver.BaseRequest.UserName,
		UpdateUserId:  receiver.BaseRequest.UserID,
	}
	userSongCover := dbModel.UserSongCover{
		ID:           utils.GetUUID(),
		UserId:       receiver.BaseRequest.UserID,
		SongCoverId:  songCover.ID,
		DelState:     constants.USER_NO_DEL_STATUS,
		CreatTime:    nowTime,
		CreateUser:   receiver.BaseRequest.UserName,
		CreateUserId: receiver.BaseRequest.UserID,
		UpdateTime:   nowTime,
		UpdateUser:   receiver.BaseRequest.UserName,
		UpdateUserId: receiver.BaseRequest.UserID,
	}

	tx := db.Begin()
	err = tx.Table("tb_song_cover").Create(&songCover).Error
	if err != nil {
		tx.Rollback()
		logs.Error("创建歌单错误：(%v)", err.Error())
		return utils.NewDBErr("创建歌单错误", err)
	}

	err = tx.Table("tb_user_song_cover").Create(&userSongCover).Error
	if err != nil {
		tx.Rollback()
		logs.Error("创建歌单-创建用户歌单失败：(%v)", err.Error())
		return utils.NewDBErr("创建用户歌单失败", err)
	}
	tx.Commit()
	return nil
}
