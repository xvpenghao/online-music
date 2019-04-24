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
            window.location.href = "http://localhost:8080/admin/channel/queryChannelList"
        },
        error:function (err) {
            //错误提示
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};