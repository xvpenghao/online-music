

function adminLogin() {
    let form = $('#form-bLogin')[0];
    let formData = new FormData(form);
    let reqData = {
        userName:formData.get('userName'),
        password :formData.get('password'),
    };
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/bAdmin/login",
        type:"POST",
        data:JSON.stringify(reqData),
        dataType:"json",
        success:function (data,status) {
            //layui.layer.msg('提示：用户登录成功');
            window.location.href = "http://localhost:8080/admin/index"
        },
        error:function (err) {
            //错误提示
           // layui.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
}