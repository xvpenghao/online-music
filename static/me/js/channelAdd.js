

//创建渠道
function submitCreateChannel() {
    let channelObj = new Channel();
    let reqData = {'channelName':$('#channelName').val()};
    channelObj.CreateChannel(reqData);
}