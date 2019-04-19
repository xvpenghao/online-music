function Song() {}

Song.prototype.deleteSongFunc = function (obj,index) {
    console.log('deleteSongFunc');
    let delLiID = '#deleteSong-'+obj.songId;
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/song/deleteSong",
        type:"DELETE",
        data:JSON.stringify(obj),
        dataType:"json",
        success:function (data) {
            parent.layer.close(index);
            parent.layer.msg('删除成功');
            $(delLiID).remove();
        },
        error:function (err) {
            parent.layer.close(index);
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};

//创建歌曲播放历史
Song.prototype.createSongPlayHistory = function (song) {
    console.log('song',song);
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/song/createSongPlayHistory",
        type:"POST",
        data:JSON.stringify(song),
        dataType:"json",
        success:function (data) {
            //播放记录框中添加一条记录
        },
        error:function (err) {
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};
