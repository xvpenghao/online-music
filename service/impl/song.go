package impl

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"online-music/common/constants"
	"online-music/common/utils"
	"online-music/models"
	"online-music/service/dbModel"
	"strings"
	"time"
)

type SongService struct {
	BaseService
}

/*
*@Title:查询歌曲详情
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/10
 */
func (receiver *SongService) QuerySongDetail(req models.QuerySongDetailReq) (dbModel.Song, error) {
	receiver.BeforeLog("QuerySongDetail")
	var result dbModel.Song
	var err error

	c := colly.NewCollector(
		colly.UserAgent(constants.USER_AGENT))

	var keywords string
	var imageUrl string
	c.OnHTML("meta[name='keywords'],meta[property='og:image']", func(e *colly.HTMLElement) {
		switch {
		case e.Attr("name") == "keywords":
			keywords = e.Attr("content")
		case e.Attr("property") == "og:image":
			imageUrl = e.Attr("content")
		}
	})

	c.OnError(func(response *colly.Response, e error) {
		if e != nil {
			err = e
		}
	})
	if err != nil {
		logs.Error("查询歌曲详情错误：(%v)", err.Error())
		return result, err
	}

	reqUrl := fmt.Sprintf(constants.SONG_URL, req.SongId)

	err = c.Visit(reqUrl)
	if err != nil {
		logs.Error("查询歌曲详情-根据歌曲id访问链接错误：(%v),访问链接:(%v)", err.Error(), reqUrl)
		return result, err
	}

	//设置专辑，歌手，歌曲封面链接，歌曲链接，歌词
	keywordss := strings.Split(keywords, "，")
	result.SongId = req.SongId
	result.SongName = keywordss[0]
	result.SongAlbum = keywordss[1]
	result.Singer = keywordss[2]
	result.SongCoverUrl = imageUrl
	//歌曲url
	result.SongPlayUrl = fmt.Sprintf(constants.SONG_PLAY_URL, req.SongId)

	//歌词url
	result.SongLyric, err = receiver.querySongLyricBySongId(req.SongId)
	if err != nil {
		logs.Error("查询歌曲详情-根据歌曲id获取歌词失败:(%v)", err.Error())
		return result, err
	}

	return result, nil
}

//根据歌曲id查询歌词
func (receiver *SongService) querySongLyricBySongId(songId string) (string, error) {

	type Lrc struct {
		Version int    `json:"version"`
		Lyric   string `json:"lyric"`
	}

	//定义临时结构体
	type Lyric struct {
		LrcStruct Lrc `json:"lrc"`
	}

	var lyric Lyric
	var err error
	c := colly.NewCollector(
		colly.UserAgent(constants.USER_AGENT))

	c.OnResponse(func(resp *colly.Response) {
		err = json.Unmarshal(resp.Body, &lyric)
	})

	if err != nil {
		logs.Error("查询歌曲详情错误：(%v)", err.Error())
		return lyric.LrcStruct.Lyric, err
	}

	reqUrl := fmt.Sprintf(constants.SONG_LYRIC_URL, songId)
	err = c.Visit(reqUrl)
	if err != nil {
		logs.Error("查询歌曲详情-访问歌曲链接错误：(%v),访问链接:(%v)", err.Error(), reqUrl)
		return lyric.LrcStruct.Lyric, err
	}

	return lyric.LrcStruct.Lyric, nil
}

/*
*@Title:查询歌曲详情
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/13
 */
func (receiver *SongService) QuerySongBaseInfo(req models.QuerySongDetailReq) (dbModel.Song, error) {
	receiver.BeforeLog("QuerySongBaseInfo")
	var result dbModel.Song
	var err error

	c := colly.NewCollector(
		colly.UserAgent(constants.USER_AGENT))

	var keywords string
	var imageUrl string
	c.OnHTML("meta[name='keywords'],meta[property='og:image']", func(e *colly.HTMLElement) {
		switch {
		case e.Attr("name") == "keywords":
			keywords = e.Attr("content")
		case e.Attr("property") == "og:image":
			imageUrl = e.Attr("content")
		}
	})

	c.OnError(func(response *colly.Response, e error) {
		if e != nil {
			err = e
		}
	})
	if err != nil {
		logs.Error("查询歌曲详情错误：(%v)", err.Error())
		return result, err
	}

	reqUrl := fmt.Sprintf(constants.SONG_URL, req.SongId)

	err = c.Visit(reqUrl)
	if err != nil {
		logs.Error("查询歌曲详情-根据歌曲id访问链接错误：(%v),访问链接:(%v)", err.Error(), reqUrl)
		return result, err
	}

	//设置专辑，歌手，歌曲封面链接，歌曲链接，歌词
	keywordss := strings.Split(keywords, "，")
	result.SongId = req.SongId
	result.SongName = keywordss[0]
	result.SongAlbum = keywordss[1]
	result.Singer = keywordss[2]
	result.SongCoverUrl = imageUrl
	//歌曲url
	result.SongPlayUrl = fmt.Sprintf(constants.SONG_PLAY_URL, req.SongId)

	return result, nil
}

/*
*@Title:根据歌单id查询歌曲列表
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/15 0015
*@Param
*@Return
 */
func (receiver *SongService) QueryUserSongList(req models.QueryUserSongListReq) ([]dbModel.UserSong, error) {
	receiver.BeforeLog("QueryUserSongList")

	db, err := receiver.GetConn()
	var result []dbModel.UserSong
	if err != nil {
		logs.Error("根据歌单id查询歌曲列表-数据库链接错误：(%v)", err.Error())
		return result, utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()

	sql := dbModel.QUERY_USER_SONG_LIST
	sqlParam := []interface{}{req.SongCoverId}

	err = db.Raw(sql, sqlParam...).Find(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Error("根据歌单id查询歌曲列表错误：(%v)", err.Error())
		return result, utils.NewDBErr("根据歌单id查询歌曲列表错误", err)
	}

	return result, nil
}

/*
*@Title:添加歌曲
*@Description: 根据歌曲id，将歌曲信息进行爬虫，然后添加歌曲，添加歌曲到选定的歌单中
*@User: 徐鹏豪
*@Date 2019/4/16 0016
*@Param
*@Return
 */
func (receiver *SongService) CreateSong(req models.CreateSongReq) error {
	receiver.BeforeLog("CreateSong")

	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("添加歌曲-数据库链接错误：(%v)", err.Error())
		return utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()

	//不可重复添加同一首歌曲到同一个歌单中
	var counts int
	err = db.Table("tb_song_cover_song").Where("song_id = ?", req.SongId).Where("song_cover_id = ?", req.SongCoverId).
		Where("user_id = ? ", receiver.BaseRequest.UserID).Count(&counts).Error
	if err != nil {
		logs.Error("添加歌曲-根据歌曲id，歌单id，用户id查询歌单歌曲数失败：(%v)", err.Error())
		return utils.NewDBErr("根据歌曲id，歌单id，用户id查询歌单歌曲数失败", err)
	}

	if counts > 0 {
		logs.Error("添加歌曲-重复添加同一个首歌曲到歌单中，歌单id(%v),歌曲id(%v),用户id:(%v)", req.SongCoverId, req.SongId,
			receiver.BaseRequest.UserID)
		return utils.NewSysErr("不能重复添加同一个首歌曲到歌单中")
	}

	//可能出现，同一个首个存放到其他的歌单中
	err = db.Table("tb_song").Where("song_id = ?", req.SongId).Count(&counts).Error
	if err != nil {
		logs.Error("添加歌曲-根据歌曲id查询歌曲数失败：(%v)", err.Error())
		return utils.NewDBErr("根据歌曲id，歌曲数失败", err)
	}

	nowTime := time.Now()
	//创建歌曲
	song := dbModel.SongTable{
		SongId:       req.SongInfo.SongId,
		SongName:     req.SongInfo.SongName,
		Singer:       req.SongInfo.Singer,
		SongAlbum:    req.SongInfo.SongAlbum,
		SongCoverUrl: req.SongInfo.SongCoverUrl,
		SongPlayUrl:  req.SongInfo.SongPlayUrl,
		SongLyric:    req.SongInfo.SongLyric,
		DelState:     constants.USER_NO_DEL_STATUS,
		CreatTime:    nowTime,
		CreateUser:   receiver.BaseRequest.UserName,
		CreateUserId: receiver.BaseRequest.UserID,
		UpdateTime:   nowTime,
		UpdateUser:   receiver.BaseRequest.UserName,
		UpdateUserId: receiver.BaseRequest.UserID,
	}
	//创建歌单歌曲信息
	songCoverSong := dbModel.SongCoverSongTable{
		SongCoverSongId: utils.GetUUID(),
		SongCoverId:     req.SongCoverId,
		UserId:          receiver.BaseRequest.UserID,
		SongId:          song.SongId,
		DelState:        constants.USER_NO_DEL_STATUS,
		CreatTime:       nowTime,
		CreateUser:      receiver.BaseRequest.UserName,
		CreateUserId:    receiver.BaseRequest.UserID,
		UpdateTime:      nowTime,
		UpdateUser:      receiver.BaseRequest.UserName,
		UpdateUserId:    receiver.BaseRequest.UserID,
	}

	//更新用户歌单图片为最新歌曲的封面链接
	updateSongCoverField := map[string]interface{}{
		"cover_url":      song.SongCoverUrl,
		"update_time":    nowTime,
		"update_user":    receiver.BaseRequest.UserName,
		"update_user_id": receiver.BaseRequest.UserID,
	}

	tx := db.Begin()
	//当歌曲不存在是，则添加
	if counts == 0 {
		err = tx.Table("tb_song").Create(&song).Error
		if err != nil {
			tx.Rollback()
			logs.Error("添加歌曲失败:(%v)", err.Error())
			return utils.NewDBErr("添加歌曲失败", err)
		}
	}
	err = tx.Table("tb_song_cover_song").Create(&songCoverSong).Error
	if err != nil {
		tx.Rollback()
		logs.Error("添加歌曲-添加歌单歌曲失败:(%v)", err.Error())
		return utils.NewDBErr("添加歌单歌曲失败", err)
	}

	err = tx.Table("tb_song_cover").Where("song_cover_id = ?", req.SongCoverId).
		Update(updateSongCoverField).Error
	if err != nil {
		tx.Rollback()
		logs.Error("添加歌曲-更新歌单信息失败:(%v)", err.Error())
		return utils.NewDBErr("更新歌单信息失败", err)
	}

	tx.Commit()
	return nil
}
