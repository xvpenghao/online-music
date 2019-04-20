
/*选中的样式*/
const SELECT_STYLE_CLASS = "collect-select";

var songCoverObj = null;

$(function(){ //页面加载完毕后再执行js
    songCoverObj = new SongCover();
    //addSongListClick();
    $('#createSongCover').click(function () {
        songCoverObj.alertCreateSongCoverWindowFunc();
    });

    //查询用户歌单信息
    songCoverObj.queryUserSongCoverListFunc();
    //加载歌曲播放历史
    let songObj = new Song();
    songObj.querySongPlayHistoryListFunc();

});

//给每一个歌单[gd-select]都注册一个单击事件，
//单击事件触发时添加[.collect-select]样式，并清除其他歌单的[.collect-select]
/*为歌单添加onClick事件*/
function addSongListClick() {
    var $gdSelect = $(".gd-select"); //返回jq对象
    console.log('$gdSelect.length',$gdSelect.length);
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


function songPlayHistoryLayer() {
    $('#song-play-history').toggle();
}

function playHistoryMouseover(ele) {
    //显示删除按钮，改变背景颜色
    let $curObj = $(ele);
    let $tools = $curObj.find(".tools");

    $curObj.css("background-color","#E5E5E5");
    $curObj.bind({
        mouseover:function(){
            //去掉隐藏
            $tools.find(".detail-delete-button").removeClass("ng-hide");
        }
    })
}

function playHistoryMouseout(ele) {
    //不显示删除按钮。删除背景颜色
    let $curObj = $(ele);
    let $tools = $curObj.find(".tools");

    $curObj.css("background-color","#fff");
    $curObj.bind({
        mouseout:function(){
            //去掉隐藏
            $tools.find(".detail-delete-button").addClass("ng-hide");
        }
    })

}

function selectPlayHistorySong(song) {

    console.log('selectPlayHistorySong-song',song);

    let liEle = '#li-play-history-'+song.songId;
    let imgEle ='#img-play-history-'+song.songId;

    //清除全部的
    $('.menu-list li').each(function () {
        $(this).removeClass('playing-song');
    });
    $('.song-status-img').each(function () {
        $(this).css({'display':'none'});
    });

    //显示播放图片
    $(imgEle).css({'display':'inline'});
    //给选择的li添加样式
    $(liEle).addClass('playing-song');

    //播放歌曲
    playSongHistory(song)
}

function playSongHistory(song) {
    let audio = $('#audioTag').get(0);
    audio.src = song.playUrl;
    //切换播放图片
    $('#playPause').attr({src:"/static/me/imgs/bf_play.png",title:PAUSE_TITLE});
    audio.play();

    //选中第一个
    let $lrcP = $(".lrc-line");
    //当页面打开时，首先选中第一行
    $($lrcP[0]).removeClass("lyric");
    $($lrcP[0]).addClass("lrc-height-line");

    $("#songCoverImg").attr({src:song.songCoverUrl});
    //设置图片的name为歌曲id
    $("#songCoverImg").attr({name:song.songId});
    $('.play-title').text(song.songName);

    //设置当前播放的歌曲的id
    $('#play-history-sasc').attr({'name':song.songId});
}

function deleteSongPlayHistory(songId) {
    let songObj = new Song();
    songObj.deleteSongPlayHistory(songId);
}

//情况歌曲播放历史
function clearAllHistorySong() {
    let songObj = new Song();
    songObj.clearAllHistorySongFunc();

}

//历史歌曲添加到歌单
function historySongAddToSC() {

    let songCoverObj = new SongCover();
    let songId = $('#play-history-sasc').attr('name');
    songCoverObj.addSongToSongCoverWindowFunc(songId)
}

let MUSIC_PLATFORM = 1;
function changeSelectMusicPlatform(ele) {

    MUSIC_PLATFORM = ele.options[ele.options.selectedIndex].value;
    $(ele).attr({'name':MUSIC_PLATFORM});
}

let reqData = {};

function Search() {
    //切换界面
    $('#my-iframe').attr({'src':'http://localhost:8080/v1/song/querySongListByKeyWordUI'});
    //获取搜索的key 和选择的平台
}