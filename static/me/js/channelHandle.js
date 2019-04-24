

//创建渠道
function submitCreateChannel() {
    let channelObj = new Channel();
    let reqData = {'channelName':$('#channelName').val()};
    channelObj.CreateChannel(reqData);
}

function modifyChannel() {
    let form=document.querySelector("#form-modify-channel");
    let formData = new FormData(form);

    let channelObj = new Channel();
    channelObj.modifyChannelFunc(formData)

}