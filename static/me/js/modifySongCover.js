
//编辑歌单,并关闭窗口
function modifySongCover(){

    let songCoverName = $('#modifySongCoverForm-songCoverName').val();
    let songCoverId = $('#modifySongCoverForm-songCoverId').val();

    let formData = {'songCoverId':songCoverId,'songCoverName':songCoverName};
    let index = parent.MODIFY_SONG_COVER_WINDOW_INDEX;
    var songCoverObj = new SongCover();
    songCoverObj.modifySongCoverFunc(formData,index)
}

//删除歌单-则会删除该歌单中的歌曲
function deleteSongCover(songCoverId){
    let songCover = new SongCover();
    let index = parent.MODIFY_SONG_COVER_WINDOW_INDEX;
    songCover.deleteSongCoverFunc(songCoverId,index);
}