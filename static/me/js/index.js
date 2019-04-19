
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

function selectPlayHistorySong(id) {

    let liEle = '#li-play-history-'+id;
    let imgEle ='#img-play-history-'+id;

    console.log('id',id);

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
}