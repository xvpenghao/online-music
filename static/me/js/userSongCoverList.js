$(function () {

    //加载用户歌单列表
    loadUserSongList();
});

function loadUserSongList() {
    $.ajax({
        url: "http://localhost:8080/v1/songCover/queryUserSongCoverList",
        type: "GET",
        sync:false,
        dataType: "json",
        success: function (data) {
            if (data.songCoverList != null) {
                userSongCover(data.songCoverList, '#dialog-playlist', '');
            }
        },
        error: function (err) {
            layer.msg('提示：' + err.responseJSON.resultMsg);
        }
    });
}

//用户歌单
function userSongCover(data,id,url) {
    let userSongCoverList = "";
    let songId = $('#songId').val();
    data.map((ele,i)=>{
        let songCoverName = getSplitSongCoverName(ele.songCoverName,7);
        let $div = `
                 <li class="playlist" onclick="addSongToSongCover({
                                     'songId':${songId},
                                     'songCoverId':${ele.songCoverId},
                                     })"
                                      onmouseover="songCoveMouseover(this)"
                                      onmouseout="songCoveMouseout(this)"
                     >
                    <img src="${ele.songCoverUrl}"/>
                    <h2>${songCoverName}</h2>
                </li>`;
        userSongCoverList  += $div
    });
    $(id).append(userSongCoverList);
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
    parent.parent.alertCreateSongCoverWindow();
}