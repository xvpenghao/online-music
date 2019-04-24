$(function () {
    let channelObj = new Channel();
    channelObj.queryChannelListFunc('',1);
});

function searchChannelForm() {
    //获取搜索内容
    let channelName = $('#channelName').val();
    let channelObj = new Channel();
    channelObj.queryChannelListFunc(channelName,1)
} 

function deleteChannel(obj,channelId) {
    let channelObj = new Channel();
    channelObj.deleteChannelFunc(obj,channelId)
}