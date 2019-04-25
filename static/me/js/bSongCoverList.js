$(function () {
    let bSongCoverObj = new SongCover();
    let reqData = {
        userName :'',
        songCoverName:'',
        type :0,
    };
    bSongCoverObj.queryPageSongCoverListFunc(reqData,1)
});

function searchSongCover() {
    let form = $('#form-queryPageSongCoverList')[0];
    let formData = new FormData(form);
    let reqData = {
        userName :formData.get('userName'),
        songCoverName:formData.get('songCoverName'),
        type :parseInt(formData.get('type')),
    };
    let bSongCoverObj = new SongCover();
    bSongCoverObj.queryPageSongCoverListFunc(reqData,1)
}

//删除歌单
function deleteBSongCover(songCoverId,userId) {
    let reqData = {
        songCoverId:songCoverId,
        userId:userId,
    };
    let bSongCoverObj = new SongCover();

    bSongCoverObj.deleteBSongCoverFunc(reqData)

}