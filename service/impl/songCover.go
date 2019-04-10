package impl

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gocolly/colly"
	"online-music/common/constants"
	"online-music/models"
	"online-music/service/dbModel"
	"strings"
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
func (receiver *SongCoverService) QuerySongCoverList(req models.QuerySongCoverListReq) ([]dbModel.SongConver, error) {
	receiver.BeforeLog("QuerySongCoverList")
	var result []dbModel.SongConver
	var songCover dbModel.SongConver
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
			logs.Error("", "查询歌单列表错误：(%v)", e.Error())
			err = e
		}
	})
	if err != nil {
		logs.Error("", "查询歌单列表错误：(%v)", err.Error())
		return result, err
	}

	reqUrl := fmt.Sprintf(constants.SONG_PAGE_LIST_URL, constants.DEFAULT_PAGE_SIZE, req.CurPage)
	err = c.Visit(reqUrl)
	if err != nil {
		logs.Error("", "查询歌单列表-访问歌单链接错误：(%v)，歌单链接：(%v)", err.Error(), reqUrl)
		return result, err
	}

	return result, nil
}
