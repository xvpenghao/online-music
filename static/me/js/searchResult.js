$(function () {

    let keyword = parent.$('#keyword').val();
    let channelId = parent.$('#select-music-platform').attr('name');
    let reqData  = {'keyWord':keyword,'channelId':channelId};
    console.log('reqData',reqData);

    let songObj  = new Song();
    songObj.querySongListByKeyWordFunc(reqData)

});