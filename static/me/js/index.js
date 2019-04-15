
/*选中的样式*/
const SELECT_STYLE_CLASS = "collect-select";

$(function(){ //页面加载完毕后再执行js
    addSongListClick();
    $('#createSongCover').click(function () {
        alertCreateSongCoverWindow();
    });

    //查询用户歌单信息
    queryUserSongCoverList();

});

//给每一个歌单[gd-select]都注册一个单击事件，
//单击事件触发时添加[.collect-select]样式，并清除其他歌单的[.collect-select]
/*为歌单添加onClick事件*/
function addSongListClick() {
    var $gdSelect = $(".gd-select"); //返回jq对象
    $.each($gdSelect, function (index, val) {
        //添加单击事件
        $(val).click(index, function () {
            $(this).addClass(SELECT_STYLE_CLASS);
            $.each($gdSelect, function (i, v) {
                if (i != index) {
                    $(v).removeClass(SELECT_STYLE_CLASS);
                }
            })

        })
    });
}

//弹出创建歌单窗口
function  alertCreateSongCoverWindow(){

    layer.config({title:"创建歌单"});
    layer.prompt(function(val, index){
        //得到输入的信息，并请求添加歌单请求，发送ajax
        $.ajax({
            contentType:'application/json;charset=UTF-8',
            url:"http://localhost:8080/v1/songCover/createSongCover",
            type:"POST",
            data:JSON.stringify({"songCoverName":val}),
            dataType:'json',
            success:function (data) {
                console.log('songCoverId',data.songCoverId);
                //如果用户的个歌单过程则截取
                let songCoverName  =  getSplitSongCoverName(val,7)
                let $div = `
                 <div class="collect-list gd-select" id="${data.songCoverId}">
                     <span>
                         <img src="/static/me/imgs/music.png">
                    </span>
                    <a href="/v1/song/songListUI" target="main">${songCoverName}</a>
                </div> `;
                $('#userSongCoverList').append($div);
            },
            error:function (err) {
                layer.msg('提示：'+err.responseJSON.resultMsg);
            }
        });
        layer.close(index);
    });
}

//查询用户歌单详情
function queryUserSongCoverList() {

    $.ajax({
        url:"http://localhost:8080/v1/songCover/queryUserSongCoverList",
        type:"GET",
        dataType:"json",
        success:function (data) {
            console.log(data);
            console.log('data.songCoverList',data.songCoverList);
            if (data.songCoverList != null){
                songCoverList(data.songCoverList,'#userSongCoverList','/v1/song/queryUserSongList');
            }
            if (data.collectList != null){
                songCoverList(data.collectList,'#userCollectSongCover','/v1/song/songListUI');
            }

        },
        error:function (err) {
            layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
}

//遍历歌单列表
function songCoverList(data,id,url) {
    // /v1/song/songListUI
    let userSongCoverList = "";
    data.map((ele,i)=>{
        let songCoverName = getSplitSongCoverName(ele.songCoverName,7);
        let $div = `
                 <div class="collect-list gd-select">
                     <span>
                         <img src="/static/me/imgs/music.png">
                    </span>
                    <a href="${url}/${ele.songCoverId}" id="${ele.songCoverId}" 
                                                  target="main">${songCoverName}</a>
               </div>`;
        userSongCoverList  += $div
    });
    $(id).append(userSongCoverList)
}
