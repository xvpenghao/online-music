$(function () {
    
});


var MODIFY_SONG_COVER_WINDOW_INDEX =0;
//编辑歌单的弹出框
function modifySongCoverLayer(songCover) {

    var songCoverObj = new SongCover();
    MODIFY_SONG_COVER_WINDOW_INDEX = songCoverObj.modifySongCoverLayerFunc(songCover);
}

//删除歌曲
function deleteSong(obj){
    let songObj = new Song();
    parent.layer.open({
        title: '提示框',
        content: '你真的要删除',
        yes:function(index,layero){
            songObj.deleteSongFunc(obj,index);
        },
    });


}