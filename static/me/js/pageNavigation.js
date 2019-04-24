$(function() {

	//ajax请求数据
	//当前页，和查询条件，请求数据
	let reqData = {'curPage':1,'changeName':'QQ云音乐'};

	channels = [
		{'channelId':1,'channelName':'张三',
			'createName':'王五','updateTime':'2019-4-24'},
		{'channelId':1,'channelName':'张三',
			'createName':'王五','updateTime':'2019-4-24'},
		{'channelId':1,'channelName':'张三',
			'createName':'王五','updateTime':'2019-4-24'},
		{'channelId':1,'channelName':'张三',
			'createName':'王五','updateTime':'2019-4-24'},
	];

	pageData ={count:20,groups:5,limit:3};
	//TODO 响应数据
	let respData = {
		list:channels,pageData:pageData};

	//TODO 数据的渲染

	renderObj = {
		elem: 'fenye' //注意，这里的 test1 是 ID，不用加 # 号
		,count: respData.pageData.count //数据总数，从服务端得到
		,limit: respData.pageData.limit  //每一页的大小
		,curr:1
		,groups: respData.pageData.groups //连续出现的页码的个数
		//,curr:location.hash.replace('#fenye','')//获取起始页
		//,hash:'fenye'//自定义hash值
		,layout: ['count', 'prev', 'page', 'next']
		,jump: jumpFunc//分页切换时候触发
	};

	layui.use('laypage',renderFunc );
});
let renderObj = null;

function renderFunc() {
	var laypage = layui.laypage;
	//执行一个laypage实例
	laypage.render(renderObj);
}

function jumpFunc(obj,first) {
	console.log('你好');
	if(!first) {//首次不会执行
		//改变也的大小
		console.log('renderObj.curr',renderObj.curr)
		//doGoPage( );
	}
}

//用户表单的搜索
function searchChannelForm() {

	//TODO 响应数据
	pageData ={count:8};
	let respData = {pageData:pageData};

	//TODO 数据的渲染
	renderObj = {
		...renderObj,
		count: respData.pageData.count //数据总数，从服务端得到
		,curr:1
	};

	layui.use('laypage',renderFunc);
}

//当你切换分页是触发
function doGoPage(pCur){
	/* var pCurId = document.getElementById("pCurId");
     pCurId.setAttribute("value",pCur);
     document.forms[0];
     document.forms[0].action=url;
     document.forms[0].submit();*/
}



