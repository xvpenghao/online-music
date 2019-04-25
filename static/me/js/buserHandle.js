
//修改用户信息
function modifyBUser() {
    let form = $('#form-modifyBUser')[0];
    let formData = new FormData(form);

    let reqData ={
        userId:formData.get('userId'),
        userName:formData.get('userName'),
        email :formData.get('email'),
        age :parseInt(formData.get('age')),
        birthday :formData.get('birthday'),
        gender :formData.get('gender'),
    };
    let bUserObj = new BUser();
    bUserObj.ModifyBUserFunc(reqData)
}