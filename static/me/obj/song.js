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
Song.prototype.createSongPlayHistoryFunc = function (song) {
    let songObj = this;
    let songData = JSON.stringify(song);
    songData = songData.replace(/"/g,"'");
    let liEle = `
            <li onmouseout="playHistoryMouseout(this)" onmouseover="playHistoryMouseover(this)"
                id="li-play-history-${song.songId}">
                <!--歌曲播放图标-->
                <div class="song-status">
                    <img style="display: none;" class="song-status-img" id="img-play-history-${song.songId}" src="/static/me/imgs/play_history2.png" title="播放"/>
                </div>
                <div class="song-title" onclick="selectPlayHistorySong(${songData})">
                    <span style="cursor: pointer;">${song.songName}</span>
                </div>
                <div class="song-singer">
                    <span style="cursor: pointer;">${song.singer}</span>
                </div>
                <div class="tools">
                    <a title="删除播放历史" class="detail-delete-button ng-hide" onclick="deleteSongPlayHistory(${song.songId})" >
                        <span class="li-del" style="cursor: pointer;">
                             <img  class="li-handle-img" src="/static/me/imgs/del.png">
                        </span>
                    </a>
                </div>
            </li>`;
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/song/createSongPlayHistory",
        type:"POST",
        data:JSON.stringify(song),
        dataType:"json",
        success:function (data) {
            //播放记录框中添加一条记录
            parent.$('.menu-list').empty();
            songObj.querySongPlayHistoryList2Func();
        },
        error:function (err) {
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};

//查询歌曲播放历史
Song.prototype.querySongPlayHistoryListFunc = function () {

    let liEles = '';
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/song/querySongPlayHistoryList",
        type:"GET",
        dataType:"json",
        success:function (data) {
            if (data.list == null || data.list.length==0){
                return
            }
            //播放记录框中添加一条记录
            data.list.map((ele,i)=>{
                let songData = JSON.stringify(ele);
                songData = songData.replace(/"/g,"'");
                let liEle = `
            <li onmouseout="playHistoryMouseout(this)" onmouseover="playHistoryMouseover(this)"
                id="li-play-history-${ele.songId}">
                <!--歌曲播放图标-->
                <div class="song-status">
                    <img style="display: none;" class="song-status-img" id="img-play-history-${ele.songId}" src="/static/me/imgs/play_history2.png" title="播放"/>
                </div>
                <div class="song-title" onclick="selectPlayHistorySong(${songData})">
                    <!-- <a href="#">当你老了</a>-->
                    <span style="cursor: pointer;">${ele.songName}</span>
                </div>
                <div class="song-singer">
                    <span style="cursor: pointer;">${ele.singer}</span>
                </div>
                <div class="tools">
                    <a title="删除播放历史" class="detail-delete-button ng-hide" onclick="deleteSongPlayHistory(${ele.songId})" >
                        <span class="li-del" style="cursor: pointer;">
                             <img  class="li-handle-img" src="/static/me/imgs/del.png">
                        </span>
                    </a>
                </div>
            </li>
    `;
                //liEles += liEle;
                $('.menu-list').append(liEle);
            });
            //设置歌曲总数
            let counts = $('.menu-list').children().length;
            $('#menu-title-total-song').text(`共${counts}首`);
        },
        error:function (err) {
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};


//查询歌曲播放历史
Song.prototype.querySongPlayHistoryList2Func = function () {
    let liEles = '';
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/song/querySongPlayHistoryList",
        type:"GET",
        dataType:"json",
        success:function (data) {
            if (data.list == null || data.list.length==0){
                return
            }
            //播放记录框中添加一条记录
            data.list.map((ele,i)=>{
                let songData = JSON.stringify(ele);
                songData = songData.replace(/"/g,"'");
                let liEle = `
            <li onmouseout="playHistoryMouseout(this)" onmouseover="playHistoryMouseover(this)"
                id="li-play-history-${ele.songId}">
                <!--歌曲播放图标-->
                <div class="song-status">
                    <img style="display: none;" class="song-status-img" id="img-play-history-${ele.songId}" src="/static/me/imgs/play_history2.png" title="播放"/>
                </div>
                <div class="song-title" onclick="selectPlayHistorySong(${songData})">
                    <!-- <a href="#">当你老了</a>-->
                    <span style="cursor: pointer;">${ele.songName}</span>
                </div>
                <div class="song-singer">
                    <span style="cursor: pointer;">${ele.singer}</span>
                </div>
                <div class="tools">
                    <a title="删除播放历史" class="detail-delete-button ng-hide" onclick="deleteSongPlayHistory(${ele.songId})" >
                        <span class="li-del" style="cursor: pointer;">
                             <img  class="li-handle-img" src="/static/me/imgs/del.png">
                        </span>
                    </a>
                </div>
            </li>
    `;
                //liEles += liEle;
                parent.$('.menu-list').append(liEle);
            });
            //设置歌曲总数
            let counts = parent.$('.menu-list').children().length;
            parent.$('#menu-title-total-song').text(`共${counts}首`);
        },
        error:function (err) {
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};


//清空播放历史
Song.prototype.clearAllHistorySongFunc = function () {
    console.log('clearAllHistorySongFunc');
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/song/deleteAllSongPlayHistory",
        type:"DELETE",
        dataType:"json",
        success:function (data) {
            $('.menu-list').empty();
            layer.msg('清空播放历史成功');
            //设置歌曲总数
            $('#menu-title-total-song').text(`共0首`);
        },
        error:function (err) {
            layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};

//删除播放历史歌曲
Song.prototype.deleteSongPlayHistory = function (songId) {
    console.log('deleteSongPlayHistory');
    let songObj = this;
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/song/deleteSongPlayHistory/"+songId,
        type:"DELETE",
        dataType:"json",
        success:function (data) {
            //删除节点
            let li = '#li-play-history-'+songId;
            $(li).remove();
            let counts = $('.menu-list').children().length;
            $('#menu-title-total-song').text(`共${counts}首`);
            layer.msg('删除成功');
        },
        error:function (err) {
            layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};
