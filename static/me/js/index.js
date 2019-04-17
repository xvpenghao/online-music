
/*选中的样式*/
const SELECT_STYLE_CLASS = "collect-select";

var songCoverObj = null;

$(function(){ //页面加载完毕后再执行js
    songCoverObj = new SongCover();
    addSongListClick();
    $('#createSongCover').click(function () {
        songCoverObj.alertCreateSongCoverWindowFunc();
    });

    //查询用户歌单信息
    songCoverObj.queryUserSongCoverListFunc();

});

//给每一个歌单[gd-select]都注册一个单击事件，
//单击事件触发时添加[.collect-select]样式，并清除其他歌单的[.collect-select]
/*为歌单添加onClick事件*/
function addSongListClick() {
    var $gdSelect = $(".gd-select"); //返回jq对象
    $.each($gdSelect, function (index, val) {
        //添加单击事件
        $(val).click(index, function () {
            $(this).addClass(SELECT_STYLE_CLASS);
            $.each($gdSelect, function (i, v) {
                if (i != index) {
                    $(v).removeClass(SELECT_STYLE_CLASS);
                }
            })

        })
    });
}
