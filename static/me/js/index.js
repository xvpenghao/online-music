
/*选中的样式*/
const SELECT_STYLE_CLASS = "collect-select";

$(function(){ //页面加载完毕后再执行js
    addSongListClick();
    $('#createSongCover').click(function () {
        alertCreateSongCoverWindow();
    });
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
        let form  = document.createElement('form');
        form.action = 'http://localhost:8080/v1/songCover/createSongCover';
        form.method = 'post';
        let input = document.createElement('input');
        input.name = 'songCoverName';
        input.value = val;
        form.appendChild(input);
        $(document.body).append(form);
        form.submit();
      /*  //得到输入的信息，并请求添加歌单请求，发送ajax
        $.ajax({
            url:"http://localhost:8080/v1/songCover/createSongCover",
            type:"POST",
            data:{"songCoverName":val},
            success:function () {

            },
            error:function (err) {}
        });
        layer.close(index);*/
    });
}


