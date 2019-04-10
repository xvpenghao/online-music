package controllers

type SongController struct {
	BaseController
}

//@Title SongDetailUI
//@Description 歌曲列表UI
//@Failure exec error
//@router /songListUI [get]
func (receiver *SongController) SongListUI() error {
	receiver.BeforeStart("SongListUI")

	receiver.TplName = "song/songList.html"
	return nil
}
