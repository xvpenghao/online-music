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

//创建收藏歌单
function createCollectSongCover(songCover) {
    console.log('songCover',songCover);
    let songCoverName = getSplitSongCoverName(songCover.songCoverName,7);
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/songCover/createCollectSongCover",
        type:"POST",
        data:JSON.stringify(songCover),
        dataType:"json",
        success:function (data,status) {
            //layer.msg('提示：收藏成功');
            //添加收藏
            let $div = `
                 <div class="collect-list gd-select" id="${songCover.songCoverId}">
                    <span>
                         <img src="/static/me/imgs/music.png">
                    </span>
                    <a href="/v1/song/songListUI" target="main">${songCoverName}</a>
                </div>
            `;
            parent.$('#userCollectSongCover').append($div);
        },
        error:function (err) {
            //错误提示
            layer.msg('提示：'+err.responseJSON.msg);
        }
    });
}

//切割字符0-6，多于的加...
function getSplitSongCoverName(str,endIndex) {
    if (endIndex >= str.length){
        return str
    }
    if (str.length <7){
        return str
    }
    let result = str.slice(0,endIndex) + '...';
    return result

}
