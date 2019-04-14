$(function(){

    detailHeadBtnAddMouseMove();

})
//为歌单详情的头部按钮-添加鼠标移动事件
function detailHeadBtnAddMouseMove(){
    //播放全部
    var $playAll = $(".playadd-button .play-list");
    //添加到当前播放列表
    var $addList = $(".playadd-button .add-list");
    //添加到收藏歌单中
    var $cloneBtn = $(".clone-button");
    //外部链接
    var $editBtn = $(".edit-button");
    var arr=[$playAll,$addList,$cloneBtn,$editBtn]
    arr.forEach(function(v,i,ar){
        mouserOverAndMouseOut(v)
    });
}

function mouserOverAndMouseOut($ele) {
    $ele.bind({
        mouseover:function(){
            $ele.css("background-color","#ccc");
        },
        mouseout:function(){
            $ele.css("background-color","#fff");
        }
    })
}

//为歌曲操作添加鼠标事件
function songHandleMouseOver(obj) {
    $curObj = $(obj)
    var $tools = $curObj.find(".tools2")
    $curObj.bind({
        mouseover:function(){
            $tools.removeClass("tools-hide")
        }
    })
}
//为歌曲操作添加鼠标事件
function songHandleMouseOut(obj) {
    $curObj = $(obj)
    var $tools = $curObj.find(".tools2")

    $curObj.bind({
        mouseout:function(){
            $tools.addClass("tools-hide")
        }
    })
}

//设置播放歌曲的url
function playSongByUrl(song) {
    console.log(song);
    //子页面获取父页面的数据
    let audio = parent.$('#audioTag').get(0);
    audio.src = song.playUrl;
    parent.$("#songCoverImg").attr({src:song.songCoverUrl});
    //设置图片的name为歌曲id
    parent.$("#songCoverImg").attr({name:song.songID});
    parent.$('.play-title').text(song.songName);

}