function BUser() {}

BUser.prototype.queryBUserList = function (bUser={},curPage=1) {
    console.log('queryBUserList');
    let bUserObj = this;

    let reqData = {...bUser,curPage:curPage};
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/admin/user/queryBUserList",
        type:"POST",
        data:JSON.stringify(reqData),
        dataType:"json",
        success:function (data,status) {
            if (data.list ==null || data.list.length===0){
                data.list= [];
            }

            let pageSize = (data.page.curPage-1)*data.page.limit;
            bUserObj.bUserListFunc(data.list,pageSize);
            //分页的列表的数据遍历
            bUserObj.renderObj = {
                elem: 'fenye' //注意，这里的 test1 是 ID，不用加 # 号
                ,count: data.page.count //数据总数，从服务端得到
                ,limit: data.page.limit  //每一页的大小
                ,curr:data.page.curPage
                ,groups: data.page.groups //连续出现的页码的个数
                ,layout: ['count', 'prev', 'page', 'next']
                ,jump: bUserObj.jumpFunc//分页切换时候触发
            };
            layui.use('laypage',function () {
                let laypage = layui.laypage;
                //执行一个laypage实例
                laypage.render(bUserObj.renderObj);
            });
        },
        error:function (err) {
            //错误提示
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};

//数据节点的生成
BUser.prototype.bUserListFunc = function(list,pageSize) {
    $('#tbody-buserList').empty();
    let trs = '';
    list.map((ele,index)=>{
        let tr = `
          <tr>
                <td data-field="number">
                    <div class="layui-table-cell laytable-cell-1-number">${index+1+pageSize}</div>
                </td>
                <td data-field="name">
                    <div class="layui-table-cell laytable-cell-1-name">${ele.userName}</div>
                </td>
                <td data-field="gender">
                    <div class="layui-table-cell laytable-cell-1-gender">${ele.gender}</div>
                </td>
                <td data-field="age">
                    <div class="layui-table-cell laytable-cell-1-age">${ele.age}</div>
                </td>
                <td data-field="mail">
                    <div class="layui-table-cell laytable-cell-1-mail">${ele.age}</div>
                </td>
                <td data-field="birthday">
                    <div class="layui-table-cell laytable-cell-1-birthday">${ele.birthday}</div>
                </td>
                <td data-field="handle" align="center" data-off="true">
    
                    <div class="layui-table-cell laytable-cell-1-handle">
                        &nbsp;&nbsp;
                        <a class="layui-btn layui-btn-xs" style="font-size: 14px;"
                           href="/admin/user/queryBUserByID/${ele.userId}" target="main">编辑</a>
                        &nbsp;&nbsp;
                        <a class="layui-btn layui-btn-danger layui-btn-xs"
                           href="#" style="font-size: 14px;"
                           onclick=""
                        >删除</a>
                    </div>
                </td>
            </tr>
        `
        trs += tr
    });
    $('#tbody-buserList').append(trs)

};

//切换页数是，会触发该事件
BUser.prototype.jumpFunc = function (obj,first) {
    let bUserObj = new BUser();
    //TODO 获取表单内容
    let form = $('#form-buser')[0];
    let formData = new FormData(form);
    let reqData ={
        userName:formData.get('userName'),
        email :formData.get('email'),
        age :parseInt(formData.get('age')),
        birthday :formData.get('birthday'),
        gender :formData.get('gender'),
    };

    if(!first) {//首次不会执行
        //改变也的大小
        console.log('jumpFunc-reqData',reqData);
        bUserObj.queryBUserList(reqData,obj.curr)
    }
};

//编辑用户
BUser.prototype.ModifyBUserFunc =function (reqData) {
    $.ajax({
        contentType:'application/json;charset=UTF-8',
        url:"http://localhost:8080/admin/user/modifyBUser",
        type:"PUT",
        data:JSON.stringify(reqData),
        dataType:"json",
        success:function (data,status) {
            parent.layer.msg('提示：修改用户信息成功');
            window.location.href = "http://localhost:8080/admin/user/bUserListUI"
        },
        error:function (err) {
            //错误提示
            parent.layer.msg('提示：'+err.responseJSON.resultMsg);
        }
    });
};

