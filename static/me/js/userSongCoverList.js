


$(function () {
    var songCoverObj = new SongCover();
    songCoverObj.loadUserSongListFunc();
});

function songCoveMouseover(ele) {
    $(ele).css("background-color","#E3E3E5");
}
function songCoveMouseout(ele) {
    $(ele).css("background-color","#fff");
}

//关闭弹窗
function closeUserCoverList() {
    console.log('closeUserCoverList');
    //父级窗口的变量
    let index = parent.ADD_SONG_TO_SONGCOVER_WINDOW_INDEX;
    console.log('parent.ADD_SONG_TO_SONGCOVER_WINDOW_INDEX',);
    parent.layer.close(index);
    //弹窗添加用户自定义歌单按钮
    parent.parent.songCoverObj.alertCreateSongCoverWindowFunc();
}

//添加歌曲到歌单中
function addSongToSongCover(ele) {
    var songCoverObj = new SongCover();
    let index = parent.ADD_SONG_TO_SONGCOVER_WINDOW_INDEX;
    songCoverObj.addSongToSongCoverFunc(ele,index);
}