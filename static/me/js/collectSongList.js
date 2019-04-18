

//删除收藏歌单
function deleteCollectSongCover(songId) {
    let songCoverObj = new SongCover();

    parent.layer.open({
        title: '提示框',
        content: '你真的要删除',
        yes:function(index,layero){
            songCoverObj.deleteSongCoverFunc(songId,index)
        },
    });

}