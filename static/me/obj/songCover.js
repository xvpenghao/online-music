////////////////////////////////////////////歌单对象/////////////////////////////////////////////////////////////////
////////////////////////////////////////////歌单对象/////////////////////////////////////////////////////////////////
////////////////////////////////////////////歌单对象/////////////////////////////////////////////////////////////////
function SongCover() {}

//查询歌单列表
SongCover.prototype.queryUserSongCoverListFunc = function () {
    let songCoverObj = this;
    $.ajax({
        url:"http://localhost:8080/v1/songCover/queryUserSongCoverList",
        type:"GET",
        dataType:"json",
        success:function (data) {
            console.log(data);
            console.log('data.songCoverList',data.songCoverList);
            if (data.songCoverList != null){
                songCoverObj.songCoverListFunc(data.songCoverList,'#userSongCoverList','/v1/song/queryUserSongList');
            }
            if (data.collectList != null){
                songCoverObj.songCoverListFunc(data.collectList,'#userCollectSongCover','/v1/song/queryCollectSCoverSongList');
            }
            //加载歌单列表点单击事件
            songCoverObj.addSongCoverListSelect();
        },
        error:function (err) {
            layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};

//遍历歌单列表
SongCover.prototype.songCoverListFunc = function (data,id,url) {
    //将子元素变为空
    let utilsObj = new Utils();
    $(id).empty();
    let userSongCoverList = "";
    data.map((ele,i)=>{
        let songCoverName = utilsObj.getSplitSongCoverName(ele.songCoverName,7);
        let $div = `
                 <div class="collect-list gd-select" id="songCover-${ele.songCoverId}">
                     <span>
                         <img src="/static/me/imgs/music.png">
                    </span>
                    <a href="${url}/${ele.songCoverId}" id="${ele.songCoverId}" 
                                                  target="main">${songCoverName}</a>
               </div>`;
        userSongCoverList  += $div
    });
    $(id).append(userSongCoverList)
};

//弹出创建歌单窗口
SongCover.prototype.alertCreateSongCoverWindowFunc = function () {
    let utilsObj = new Utils();
    let songCoverUrl = '/v1/song/queryUserSongList';
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
                let songCoverName  =  utilsObj.getSplitSongCoverName(val,7)
                let $div = `
                 <div class="collect-list gd-select" id="${data.songCoverId}">
                     <span>
                         <img src="/static/me/imgs/music.png">
                    </span>
                    <a href="${songCoverUrl}/${data.songCoverId}" target="main">${songCoverName}</a>
                </div> `;
                $('#userSongCoverList').append($div);
            },
            error:function (err) {
                layer.msg('提示：'+err.responseJSON.resultMsg);
            }
        });
        layer.close(index);
    });
};

//创建收藏歌单
SongCover.prototype.createCollectSongCoverFunc = function (obj) {
    console.log('obj',obj);
    var utilsObj = new Utils();
    let songCoverName = utilsObj.getSplitSongCoverName(obj.songCoverName,7);
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/songCover/createCollectSongCover",
        type:"POST",
        data:JSON.stringify(obj),
        dataType:"json",
        success:function (data,status) {
            //layer.msg('提示：收藏成功');
            //添加收藏
            let $div = `
                 <div class="collect-list gd-select" id="${obj.songCoverId}">
                    <span>
                         <img src="/static/me/imgs/music.png">
                    </span>
                    <a href="/v1/song/songListUI" target="main">${songCoverName}</a>
                </div>
            `;
            parent.$('#userCollectSongCover').append($div);
        },
        error:function (err) {
            //错误提示
            layer.msg('提示：'+err.responseJSON.msg);
        }
    });
}

//添加歌曲到歌单窗口
SongCover.prototype.addSongToSongCoverWindowFunc =function (songId) {
    let index = layer.open({
                    type: 2,
                    title: '添加歌曲',
                    shadeClose: true,
                    shade: 0,
                    area: ['285px', '430px'],
                    content: `http://localhost:8080/v1/songCover/userSongCoverListUI/${songId}` //iframe的url
                });
    return index;
};

//加载用户歌单列表-用户在添加到歌曲到歌单的操作
SongCover.prototype.loadUserSongListFunc =function () {
    let songCoverObj = this;
    $.ajax({
        url: "http://localhost:8080/v1/songCover/queryUserSongCoverList",
        type: "GET",
        sync:false,
        dataType: "json",
        success: function (data) {
            if (data.songCoverList != null) {
                songCoverObj.userSongCoverFunc(data.songCoverList, '#dialog-playlist', '');
            }
        },
        error: function (err) {
            layer.msg('提示：' + err.responseJSON.resultMsg);
        }
    });
};

//用户歌单-用户在添加到歌曲到歌单的操作
SongCover.prototype.userSongCoverFunc = function (data,id,url) {
    let userSongCoverList = "";
    let utilsObj = new Utils();
    let songId = $('#songId').val();
    data.map((ele,i)=>{
        let songCoverName = utilsObj.getSplitSongCoverName(ele.songCoverName,7);
        let $div = `
                 <li class="playlist" songId="${songId}" songCoverId ="${ele.songCoverId}" onclick="addSongToSongCover(this)"
                                      onmouseover="songCoveMouseover(this)"
                                      onmouseout="songCoveMouseout(this)"
                     >
                    <img src="${ele.songCoverUrl}"/>
                    <h2>${songCoverName}</h2>
                </li>`;
        userSongCoverList  += $div
    });
    $(id).append(userSongCoverList);
};

//添加歌曲到歌单中
SongCover.prototype.addSongToSongCoverFunc = function (ele,index) {

    let songId = $(ele).attr('songId');
    //TODO 假象，后台在执行
    let songCoverId = $(ele).attr('songCoverId');
    let data = {'songId':songId,'songCoverId':songCoverId};
    //发送ajax请求
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/song/createSong",
        type:'POST',
        data:JSON.stringify(data),
        dataType:'json',
        success:function (msg) {
            parent.layer.msg('添加成功');
            parent.layer.close(index);
        },
        error:function (err) {
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
            parent.layer.close(index);
        }
    });
};

//编辑歌单的弹出框
SongCover.prototype.modifySongCoverLayerFunc = function (ele) {
    console.log(ele);
    let index = 0;
    let reqUrl = `songCoverId=${ele.songCoverId}&songCoverName=${ele.songCoverName}`;
    console.log('reqUrl',reqUrl);
    //将歌单的参数传递到后台，后台将歌单的信息在显示到前台
    index  = layer.open({
        type: 2,
        title: '编辑歌单',
        shadeClose: true,
        shade: 0,
        area: ['390px', '215px'],
        content: `http://localhost:8080/v1/songCover/modifySongCoverUI?${reqUrl}` //iframe的url
    });
    return index;
};

//编辑歌单
SongCover.prototype.modifySongCoverFunc = function (formData,index) {
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/songCover/modifySongCover",
        type:"POST",
        data:JSON.stringify(formData),
        dataType:"json",
        success:function (data) {
            parent.layer.close(index);
            parent.layer.msg('编辑成功');
            //重写加载页面
            parent.parent.window.location.reload();
        },
        error:function (err) {
            parent.layer.close(index);
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);

        }
    });
};

//删除用户的歌单
SongCover.prototype.deleteSongCoverFunc = function (songCoverId,index) {
    console.log('songCoverId',songCoverId);
    console.log('index',index);
    //删除歌单，删除组件
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/v1/songCover/deleteSongCover/"+songCoverId,
        type:"DELETE",
        dataType:"json",
        success:function (data) {
            parent.layer.close(index);
            parent.layer.msg('删除成功');
            //页面再次刷新
            parent.parent.window.location.reload();
        },
        error:function (err) {
            parent.layer.close(index);
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};

//给歌单列表添加事件
SongCover.prototype.addSongCoverListSelect = function () {
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
};

SongCover.prototype.renderObj ={};

//查询分页歌单列表
SongCover.prototype.queryPageSongCoverListFunc =function (bSongCover,curPage) {
    console.log('queryPageSongCoverListFunc');
    let songCoverObj = this;

    let reqData = {...bSongCover,curPage:curPage};
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/admin/songCover/queryPageSongCoverList",
        type:"POST",
        data:JSON.stringify(reqData),
        dataType:"json",
        success:function (data,status) {
            if (data.list ==null || data.list.length===0){
                data.list= [];
            }

            let pageSize = (data.page.curPage-1)*data.page.limit;
            songCoverObj.bSongCoverListFunc(data.list,pageSize);
            //分页的列表的数据遍历
            songCoverObj.renderObj = {
                elem: 'fenye' //注意，这里的 test1 是 ID，不用加 # 号
                ,count: data.page.count //数据总数，从服务端得到
                ,limit: data.page.limit  //每一页的大小
                ,curr:data.page.curPage
                ,groups: data.page.groups //连续出现的页码的个数
                ,layout: ['count', 'prev', 'page', 'next']
                ,jump: songCoverObj.jumpFunc//分页切换时候触发
            };
            layui.use('laypage',function () {
                let laypage = layui.laypage;
                //执行一个laypage实例
                laypage.render(songCoverObj.renderObj);
            });
        },
        error:function (err) {
            //错误提示
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};


//数据节点的生成
SongCover.prototype.bSongCoverListFunc = function(list,pageSize) {
    $('#tbody-queryPageSongCoverList').empty();
    let trs = '';
    list.map((ele,index)=>{
        let tr = `
            <tr>
                <td data-field="number">
                    <div class="layui-table-cell laytable-cell-1-number">${index+1+pageSize}</div>
                </td>
                <td data-field="name">
                    <div class="layui-table-cell laytable-cell-1-name">${ele.userName}</div>
                </td>
                <td data-field="songCoverName">
                    <div class="layui-table-cell laytable-cell-1-songCoverName">${ele.songCoverName}</div>
                </td>
                <td data-field="type">
                    <div class="layui-table-cell laytable-cell-1-type">${ele.type}</div>
                </td>
                <td data-field="handle" align="center" data-off="true">

                    <div class="layui-table-cell laytable-cell-1-handle">
                        &nbsp;&nbsp;
                        <a class="layui-btn layui-btn-xs" style="font-size: 14px;"
                           href="/admin/songCover/queryBSongCoverByID/${ele.songCoverId}/${ele.userId}" target="main">编辑</a>
                        &nbsp;&nbsp;
                        <a class="layui-btn layui-btn-danger layui-btn-xs"
                           href="#" style="font-size: 14px;"
                           onclick="deleteBSongCover(${ele.songCoverId},${ele.userId})"
                        >删除</a>
                    </div>
                </td>
            </tr>
        `;

        trs +=  tr;
    });

    $('#tbody-queryPageSongCoverList').append(trs);

};

//切换页数是，会触发该事件
SongCover.prototype.jumpFunc = function (obj,first) {
    let bSongCoverObj = new SongCover();

    let form = $('#form-queryPageSongCoverList')[0];
    let formData = new FormData(form);
    let reqData = {
        userName :formData.get('userName'),
        songCoverName:formData.get('songCoverName'),
        type :parseInt(formData.get('type')),
    };
    if(!first) {//首次不会执行
        //改变也的大小
        bSongCoverObj.queryPageSongCoverListFunc(reqData,obj.curr)
    }
};

SongCover.prototype.modifyBSongCoverFunc = function (reqData) {
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/admin/songCover/modifyBSongCover",
        type:"PUT",
        data:JSON.stringify(reqData),
        dataType:"json",
        success:function (data,status) {
            parent.layer.msg('提示：修改歌单信息成功');
            window.location.href = "http://localhost:8080/admin/songCover/bSongCoverListUI"
        },
        error:function (err) {
            //错误提示
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};
