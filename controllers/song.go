package controllers

type SongController struct {
	BaseController
}

//@Title SongBestCoverUI
//@Description 精选歌单UI
//@Failure exec error
//@router /songBestCoverUI [get]
func (receiver *SongController) SongBestCoverUI() error {
	receiver.BeforeStart("SongBestCoverUI")

	receiver.TplName = "song/songBestCover.html"
	return nil
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
