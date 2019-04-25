
//修改歌单
function modifyBSongCover(){
    let form = $('#form-modifyBSongCover')[0];
    let formData = new FormData(form);
    let reqData = {
        songCoverId:formData.get('songCoverId'),
        songCoverName:formData.get('songCoverName'),
        userId:formData.get("userId"),
    };
    let bsongCoverObj = new SongCover();
    bsongCoverObj.modifyBSongCoverFunc(reqData);
}