$(function () {
    //加载页面，请求数据
    let bUserObj = new BUser();
    let reqData ={
        userName:'',
        email :'',
        age :0,
        birthday :'',
        gender :'',
    };
    bUserObj.queryBUserList(reqData,1);
});

function searchUserForm() {
    let form = $('#form-buser')[0];
    let formData = new FormData(form);
    let reqData ={
        userName:formData.get('userName'),
        email :formData.get('email'),
        age :parseInt(formData.get('age')),
        birthday :formData.get('birthday'),
        gender :formData.get('gender'),
    };

    let bUserObj = new BUser();
    console.log('searchUserForm',reqData);
    bUserObj.queryBUserList(reqData,1);
}