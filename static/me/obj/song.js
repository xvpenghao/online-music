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
