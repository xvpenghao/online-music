package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["online-music/controllers:IndexController"] = append(beego.GlobalControllerRouter["online-music/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "ErrorUI",
			Router:           `/errorUI`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:IndexController"] = append(beego.GlobalControllerRouter["online-music/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "IndexUI",
			Router:           `/indexUI`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:LoginController"] = append(beego.GlobalControllerRouter["online-music/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "LoginIn",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:LoginController"] = append(beego.GlobalControllerRouter["online-music/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "LoginOut",
			Router:           `/loginOut`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:LoginController"] = append(beego.GlobalControllerRouter["online-music/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "LoginUI",
			Router:           `/loginUI`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongController"],
		beego.ControllerComments{
			Method:           "CreateSong",
			Router:           `/createSong`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongController"],
		beego.ControllerComments{
			Method:           "CreateSongPlayHistory",
			Router:           `/createSongPlayHistory`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongController"],
		beego.ControllerComments{
			Method:           "DeleteAllSongPlayHistory",
			Router:           `/deleteAllSongPlayHistory`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongController"],
		beego.ControllerComments{
			Method:           "DeleteSong",
			Router:           `/deleteSong`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongController"],
		beego.ControllerComments{
			Method:           "DeleteSongPlayHistory",
			Router:           `/deleteSongPlayHistory/:songId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongController"],
		beego.ControllerComments{
			Method:           "QueryCollectSCoverSongList",
			Router:           `/queryCollectSCoverSongList/:songCoverId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongController"],
		beego.ControllerComments{
			Method:           "QuerySongDetail",
			Router:           `/querySongDetail/:songID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongController"],
		beego.ControllerComments{
			Method:           "QuerySongListByKeyWord",
			Router:           `/querySongListByKeyWord`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongController"],
		beego.ControllerComments{
			Method:           "QuerySongListByKeyWordUI",
			Router:           `/querySongListByKeyWordUI`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongController"],
		beego.ControllerComments{
			Method:           "QuerySongPlayHistoryList",
			Router:           `/querySongPlayHistoryList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongController"],
		beego.ControllerComments{
			Method:           "QueryUserSongList",
			Router:           `/queryUserSongList/:songCoverId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongCoverController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongCoverController"],
		beego.ControllerComments{
			Method:           "CreateCollectSongCover",
			Router:           `/createCollectSongCover`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongCoverController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongCoverController"],
		beego.ControllerComments{
			Method:           "CreateSongCover",
			Router:           `/createSongCover`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongCoverController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongCoverController"],
		beego.ControllerComments{
			Method:           "DeleteSongCover",
			Router:           `/deleteSongCover/:songCoverId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongCoverController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongCoverController"],
		beego.ControllerComments{
			Method:           "ModifySongCover",
			Router:           `/modifySongCover`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongCoverController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongCoverController"],
		beego.ControllerComments{
			Method:           "ModifySongCoverUI",
			Router:           `/modifySongCoverUI`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongCoverController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongCoverController"],
		beego.ControllerComments{
			Method:           "QuerySongCoverList",
			Router:           `/querySongCoverList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongCoverController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongCoverController"],
		beego.ControllerComments{
			Method:           "QuerySongList",
			Router:           `/querySongList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongCoverController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongCoverController"],
		beego.ControllerComments{
			Method:           "QueryUserSongCoverList",
			Router:           `/queryUserSongCoverList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:SongCoverController"] = append(beego.GlobalControllerRouter["online-music/controllers:SongCoverController"],
		beego.ControllerComments{
			Method:           "UserSongCoverListUI",
			Router:           `/userSongCoverListUI/:songId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:UserController"] = append(beego.GlobalControllerRouter["online-music/controllers:UserController"],
		beego.ControllerComments{
			Method:           "CreateUser",
			Router:           `/createUser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:UserController"] = append(beego.GlobalControllerRouter["online-music/controllers:UserController"],
		beego.ControllerComments{
			Method:           "ModifyPwd",
			Router:           `/modifyPwd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:UserController"] = append(beego.GlobalControllerRouter["online-music/controllers:UserController"],
		beego.ControllerComments{
			Method:           "ModifyPwdUI",
			Router:           `/modifyPwdUI`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:UserController"] = append(beego.GlobalControllerRouter["online-music/controllers:UserController"],
		beego.ControllerComments{
			Method:           "ModifyUser",
			Router:           `/modifyUser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:UserController"] = append(beego.GlobalControllerRouter["online-music/controllers:UserController"],
		beego.ControllerComments{
			Method:           "RegisterUI",
			Router:           `/registerUI`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["online-music/controllers:UserController"] = append(beego.GlobalControllerRouter["online-music/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UserDetailUI",
			Router:           `/userDetailUI`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
