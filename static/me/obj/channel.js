function Channel() {

}

//创建渠道
Channel.prototype.CreateChannel = function (reqData) {

    console.log('reqData',reqData);
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/admin/channel/createChannel",
        type:"POST",
        data:JSON.stringify(reqData),
        dataType:"json",
        success:function (data,status) {
            parent.layer.msg('提示：添加平台成功');
            window.location.href = "http://localhost:8080/admin/channel/queryChannelListUI"
        },
        error:function (err) {
            //错误提示
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};

Channel.prototype.renderObj = {};

//查询渠道信息
Channel.prototype.queryChannelListFunc = function (channelName='',curPage =1) {
    console.log('queryChannelList');
    let channelObj = this;
    let reqData = {channelName:channelName,curPage:curPage};
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/admin/channel/queryChannelList",
        type:"POST",
        data:JSON.stringify(reqData),
        dataType:"json",
        success:function (data,status) {
            if (data.list ==null || data.list.length===0){
                data.list= [];
            }

            let pageSize = (data.page.curPage-1)*data.page.limit;
            channelObj.channelListsFunc(data.list,pageSize);
            //分页的列表的数据遍历
            channelObj.renderObj = {
                elem: 'fenye' //注意，这里的 test1 是 ID，不用加 # 号
                ,count: data.page.count //数据总数，从服务端得到
                ,limit: data.page.limit  //每一页的大小
                ,curr:data.page.curPage
                ,groups: data.page.groups //连续出现的页码的个数
                ,layout: ['count', 'prev', 'page', 'next']
                ,jump: channelObj.jumpFunc//分页切换时候触发
            };
            layui.use('laypage',function () {
                let laypage = layui.laypage;
                //执行一个laypage实例
                laypage.render(channelObj.renderObj);
            });
        },
        error:function (err) {
            //错误提示
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};

Channel.prototype.channelListsFunc = function(list,pageSize){
    $('#tbody-channelList').empty();
    let lis = "";
    list.map((ele,index)=>{
        let li = `
             <tr>
                <td data-field="number">
                    <div class="layui-table-cell laytable-cell-1-number">${index+1+pageSize}</div>
                </td>
                <td data-field="channelName">
                    <div class="layui-table-cell laytable-cell-1-channelName">${ele.channelName}</div>
                </td>
                <td data-field="createName">
                    <div class="layui-table-cell laytable-cell-1-createName">${ele.createUser}</div>
                </td>
                <td data-field="updateTime">
                    <div class="layui-table-cell laytable-cell-1-updateTime">${ele.updateTime}</div>
                </td>
                <td data-field="handle" align="center" data-off="true">

                    <div class="layui-table-cell laytable-cell-1-handle">
                        &nbsp;&nbsp;
                        <a class="layui-btn layui-btn-xs" style="font-size: 14px;"
                           href="/admin/channel/queryChannelDetail/${ele.channelId}" target="main">编辑</a>
                        &nbsp;&nbsp;
                        <a class="layui-btn layui-btn-danger layui-btn-xs"
                           href="#" style="font-size: 14px;"
                           onclick="deleteChannel(this,${ele.channelId})"
                        >删除</a>
                    </div>
                </td>
            </tr>
        `

        lis += li;
    });
    $('#tbody-channelList').append(lis);
};

Channel.prototype.jumpFunc = function (obj,first) {
    let channelObj = new Channel();
    let channelName = $('#channelName').val();
    if(!first) {//首次不会执行
        //改变也的大小
        console.log('obj.curr',obj.curr);
        channelObj.queryChannelListFunc(channelName,obj.curr)
    }
};

Channel.prototype.deleteChannelFunc =function (obj,channelId) {

    console.log('deleteChannelFunc',channelId);
    //TODO 成功后删除该节点，后台更新数据
};

Channel.prototype.modifyChannelFunc =function (formData) {
    let reqData = {channelName:formData.get('channelName'),
        channelId:formData.get('channelId'),
    };
    console.log('reqData',reqData);
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/admin/channel/modifyChannel",
        type:"PUT",
        data:JSON.stringify(reqData),
        dataType:"json",
        success:function (data,status) {
            parent.layer.msg('提示：修改渠道成功');
            window.location.href = "http://localhost:8080/admin/channel/queryChannelListUI"
        },
        error:function (err) {
            //错误提示
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};