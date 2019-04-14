package impl

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gocolly/colly"
	"online-music/common/constants"
	"online-music/models"
	"online-music/service/dbModel"
	"strings"
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
