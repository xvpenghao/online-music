const NG_HIDE = "ng-hide";
const DQXH = "单曲循环";
const XXBF = "顺序播放";
const SJBF = "随机播放";
let PLAY_MODEL = DQXH;
let lrcTimeArray = [];
let scrollLocation = 0;

const data =[
    {"song":"/static/me/temp/thh.mp3"},
];

$(function () {
    let audio = $('#audioTag').get(0);
    audio.src = data[0].song;
    audio.setAttribute("songIndex",0);

    audio.volume = 0.5;
    //歌曲的播放暂停
    playOrPause();
    //歌曲上下首切换
    upOrDownChange();
    //播放模式的切换
    changePlaymode();

    //获取音频的长度
    $('#audioTag').on("loadedmetadata",function () {
        $('.total').text(transTime(this.duration));
    });

    //监听音频播放，并更新进度条
    //timeupdate 事件是在播放位置改变时触发
    audio.addEventListener('timeupdate',updateProgress,false);

    /*设置音量*/
    setVolume();
    /*音量开关*/
    volumeOnOrOff();
    setInitVolume();
    //设置歌词播放时间
    setLyricTime(lyricData);
    //设置歌词详情和main的切换
    changeMainAndLyricDetail()

});


const PLAY_TITLE = "播放";
const PAUSE_TITLE = "暂停";
//歌曲播放暂停
function playOrPause() {
    let audio = $('#audioTag').get(0);

    $('#playPause').click(function(){
        //改变暂停/播放icon
        if(audio.paused){
            $(this).attr({src:"/static/me/imgs/bf_play.png",title:PAUSE_TITLE});
            audio.play();
            //选中第一个
            let $lrcP = $(".lrc-line");
            //当页面打开时，首先选中第一行
            $($lrcP[0]).removeClass("lyric");
            $($lrcP[0]).addClass("lrc-height-line");
        } else{
            $(this).attr({src:"/static/me/imgs/bf_pause.png",title:PLAY_TITLE});
            audio.pause();

        }
    })
}

//歌曲上下首切换
function upOrDownChange(){
    //根据不同的播放方式，来实现切换
    $('#upSong').click(clickUpPlay);

    $('#downSong').click(clickDownPlay);
}

function clickUpPlay() {
    let audio = $('#audioTag').get(0);
    //如果播放的方式等于随机
    if (PLAY_MODEL == SJBF){
        //到时候使用redis实现随机
        let curIndex = parseInt(data.length * Math.random());
        $('#audioTag').attr({"src":data[curIndex].song,"songIndex":curIndex});
        audio.play();
        return
    }

    if (PLAY_MODEL == XXBF ||PLAY_MODEL == DQXH){
        let curIndex = $('#audioTag').attr("songIndex");
        if (curIndex <=0){
            curIndex = data.length-1;
        }else{
            curIndex -=1;
        }
        $('#audioTag').attr({"src":data[curIndex].song,"songIndex":curIndex});
        audio.play();
    }
}

function clickDownPlay() {
    let audio = $('#audioTag').get(0);
    //如果播放的方式等于随机
    if (PLAY_MODEL == SJBF){
        //到时候使用redis实现随机
        let curIndex = parseInt(data.length * Math.random());
        $('#audioTag').attr({"src":data[curIndex].song,"songIndex":curIndex})
        audio.play();
    }

    if (PLAY_MODEL == XXBF || PLAY_MODEL == DQXH){
        let curIndex = $('#audioTag').attr("songIndex");
        curIndex = parseInt(curIndex)
        if (curIndex >= data.length-1){
            curIndex = 0;
        }else{
            curIndex = curIndex + 1;
        }
        $('#audioTag').attr({"src":data[curIndex].song,"songIndex":curIndex})
        audio.play();
    }
}


//更新进度条
function updateProgress() {
    let audio =document.getElementsByTagName('audio')[0]; //js获取的方式
    let value = Math.round((Math.floor(audio.currentTime) / Math.floor(audio.duration)) * 100, 0);

    $('.pgs-play').css('width', value + '%');
    $('.current').html(transTime(audio.currentTime));
    lrcMove(audio.currentTime)
}

//转换音频时长显示
function transTime(time) {
    var duration = parseInt(time);
    var minute = parseInt(duration/60);
    var sec = duration%60+'';
    var isM0 = ':';
    if(minute == 0){
        minute = '00';
    }else if(minute < 10 ){
        minute = '0'+minute;
    }
    if(sec.length == 1){
        sec = '0'+sec;
    }
    return minute+isM0+sec
}

/*播放方式的切换*/
function changePlaymode() {
    var $playModel = $(".bf-fs");
    $.each($playModel,function (index,ele) {
        //添加单击事件
        $(ele).click(index,function () {
            //添加隐藏class
            // $(this).addClass(NG_HIDE)
            //显示单击项的下一项
            $.each($playModel,function (i,ele2) {
                let showIndex = (index+1) %3;
                if (i== showIndex){
                    PLAY_MODEL = $(ele2).attr("title")
                    $(ele2).removeClass(NG_HIDE)
                }else{
                    $(ele2).addClass(NG_HIDE)
                }
            })
        })
    })
}

//设置初始化音量
function setInitVolume() {
    let elem = document.querySelector('input[type="range"]');
    //设置音量
    let audio = $('#audioTag').get(0);
    let newValue = elem.value;
    let max = elem.getAttribute("max");
    audio.volume = (newValue /max);
    let width = (91.3 / max * newValue) +"%";    //这里的91.3是用了整个滑块的宽度300减去拖动的那个圆形滑块的宽度30再加上圆形滑块的边框宽度4然后再除以300得来的，因为显示拖动距离的rang_width在绝对定位后在滑动过程中会遮挡住圆形滑块，导致圆形滑块无法被拖动，所以要适当的减少rang_width在滑动时的宽度，当然rang_width的宽度是根据你自己的实际需求来计算出来的，并不是一成不变的91.3%
    document.querySelector('.rang_width').style.width = width;
}

//设置音量
function setVolume(){
    let elem = document.querySelector('input[type="range"]');
    let rangeValue = function(){
        //设置音量
        let audio = $('#audioTag').get(0);
        let newValue = elem.value;
        if(newValue<=0){
            $(".on-off").attr({src:"/static/me/imgs/no_volume.png",title:OFF_VOLUME});
        }else{
            $(".on-off").attr({src:"/static/me/imgs/volume.png",title:ON_VOLUME});
        }

        let max = elem.getAttribute("max");
        audio.volume = (newValue /max);
        let width = (91.3 / max * newValue) +"%";    //这里的91.3是用了整个滑块的宽度300减去拖动的那个圆形滑块的宽度30再加上圆形滑块的边框宽度4然后再除以300得来的，因为显示拖动距离的rang_width在绝对定位后在滑动过程中会遮挡住圆形滑块，导致圆形滑块无法被拖动，所以要适当的减少rang_width在滑动时的宽度，当然rang_width的宽度是根据你自己的实际需求来计算出来的，并不是一成不变的91.3%
        document.querySelector('.rang_width').style.width = width;
    };
    elem.addEventListener("input", rangeValue);
}

const ON_VOLUME = "on";
const OFF_VOLUME = "off";
function volumeOnOrOff(){
    //给图片设置切换事件

    $(".on-off").click(function () {
        let value = $(".on-off").attr("title");
        if (value == ON_VOLUME){
            $(this).attr({src:"/static/me/imgs/no_volume.png",title:OFF_VOLUME});
            let audio = $('#audioTag').get(0);
            audio.volume = 0.0;
        }else{
            $(this).attr({src:"/static/me/imgs/volume.png",title:ON_VOLUME});
            let elem = document.querySelector('input[type="range"]');
            let audio = $('#audioTag').get(0);
            let newValue = elem.value;
            let max = elem.getAttribute("max");
            audio.volume = (newValue /max);
        }
    });
}

/* ***************************关于歌词的设置************************************************* */
let lyricData = [
    {
    "lyric":"[00:00.00] 作曲 : 金玟岐\n[00:01.00] 作词 : 金玟岐/梁振华\n[00:05.230]制作人/编曲：薛琳可\n[00:10.230]吉他演奏：薛峰\n[00:15.230]弦乐演奏：国际首席爱乐乐团\n[00:20.230]混音：赵靖@BIG.J Studio，Beijing\n[00:25.230]思美人兮 路长漫漫不可及\n[00:30.230]拭泪天涯无归期\n[00:36.120]千言万语 只如梗在心底\n[00:41.920]愁丝万缕一朝夕\n[00:47.840]思美人兮 悠悠浮云为我寄\n[00:53.390]奈何鸿飞不解意\n[00:58.550]归鸟乘风远飞\n[01:01.930]我折翼在原地\n[01:05.070]只见思念划天际\n[01:10.320]男儿志千古愁\n[01:13.360]溢于胸怀中\n[01:15.960]抛入一汪江水向东流\n[01:21.400]笑非笑 忧且忧\n[01:24.580]浊世谁人能懂\n[01:27.290]拂袖独行不回首\n[01:32.910]幽兰花异芬芳\n[01:36.110]含风影自香\n[01:38.800]问君可有识得她芳踪\n[01:44.160]趁年华未散尽\n[01:47.210]摘得芙蓉与共\n[01:49.990]莫让憾事绕心中\n[02:19.040]思美人兮 纵然九死无悔意\n[02:24.690]我心惜古人不及\n[02:30.370]江可竭山可移\n[02:33.210]唯志节不离弃\n[02:36.300]宁求上下无归期\n[02:41.610]男儿志千古愁\n[02:44.750]溢于胸怀中\n[02:47.110]抛入一汪江水向东流\n[02:52.770]笑非笑 忧且忧\n[02:55.880]浊世谁人能懂\n[02:58.710]拂袖独行不回首\n[03:04.360]幽兰花异芬芳\n[03:07.180]含风影自香\n[03:09.920]问君可有识得她芳踪\n[03:15.600]趁年华未散尽\n[03:18.620]摘得芙蓉与共\n[03:22.100]莫让憾事绕心中\n[03:27.420]男儿志千古愁\n[03:30.510]溢于胸怀中\n[03:32.860]抛入一汪江水向东流\n[03:38.460]笑非笑 忧且忧\n[03:41.410]浊世谁人能懂\n[03:44.650]拂袖独行不回首\n[03:49.950]幽兰花异芬芳\n[03:52.920]含风影自香\n[03:55.410]问君可有识得她芳踪\n[04:01.200]趁年华未散尽\n[04:04.320]摘得芙蓉与共\n[04:07.590]莫让憾事绕心中\n[04:13.210]思美人兮\n[04:16.160]漂泊南行无所依\n[04:19.360]我欲随风同归去\n",
    "code":200,
    }
];

//歌词段落的添加
function addLrc(lyrics) {


    let html = "<p></p>";
    let content = "";
    lyrics.map((v,i)=>{
        content = v.split(']')[1];
        if (content != undefined && content !=''){
            html += "<p class=\"lrc-line\" data-timeLine=\"" + lrcTimeArray[i] + "\">" + content + "</p>";
        }
    });
    $('.lyric').html(html);
}


//---------------------------------------------------------【歌词移动】
function lrcMove(currentTime){
    /*歌词高亮 #ff4444 红色*/

    let duration = parseInt(currentTime);
    let minute = parseInt(duration/60);
    let sec = duration%60;
    let totalSec = minute*60+sec;
    let curPlayLyric = 0;


    // let spaceHeiht = 18,
    let  lyricsHeight =8;
    let dataTimeLine = 0;
    //得到所有的歌词
    let scrolHeight = $('.lyric').scrollTop();
    // $('.lyric').scrollTop((scroolHeight+18)+19);
    let $lrcP = $(".lrc-line");
    //当页面打开时，首先选中第一行
    if (currentTime==0){
        return
    }
    //transTime
    $.each($lrcP,function (i,ele) {
        dataTimeLine =  $(ele).attr("data-timeLine");
        if(dataTimeLine > 0 && dataTimeLine == totalSec){
            //移出上一个class，选中下一个
            $($lrcP[i-1]).removeClass("lyric");
            if (i >=8){
                $('.lyric').scrollTop((scrolHeight)+lyricsHeight);
            }
            if (i-1>=0 ){
                $($lrcP[i-1]).removeClass("lrc-height-line");
            }
            $($lrcP[i]).addClass("lrc-height-line");
        }

    });

}

//设置歌词的时间，通过正则
function setLyricTime(lyricData){
    let lyrics = lyricData[0].lyric.split("\n");
    lyricData[0].lyric.replace(/\[(\d*):(\d*)([\.|\:]\d*)\]/g,function () {
        /*
        * arguments:
        * 0:"[00:01.00]"
        * 1:"00"
        * 2:"01"
        * 3:".00"
        * 4:
        * 5:all
        * */
        //console.log(arguments);
        let min = arguments[1] | 0, //分
            sec = arguments[2] | 0, //秒
            realMin = min * 60 + sec;//计算总秒数

        lrcTimeArray.push(realMin);
    });

    addLrc(lyrics);
}
//状态的切换
var mainAndLyricDetailstateChange = true;
//点击播放封面，完成事件的切换 显示歌词，不显示歌词详情
function changeMainAndLyricDetail(){
    $(".cover").click(function () {
        if (mainAndLyricDetailstateChange){
            //隐藏main，显示歌词详情
            $(".main").addClass("ng-hide");
            $(".songdetail-wrapper").removeClass("ng-hide");
            mainAndLyricDetailstateChange = false;
            //ajax加载数据
            let songID = $("#songCoverImg").attr('name');
            $.ajax({
                type:"GET",
                url:"http://localhost:8080/v1/song/querySongDetail/"+songID,
                dataType:"json",
                success:function (data) {
                    //设置歌曲详情页面属性
                    lyricData = [{'lyric':data.songLyric}];
                    $("#songName").text(data.songName);
                    $("#singer").text("歌手名称："+data.singer);
                    $("#album").text("专辑名称："+data.songAlbum);
                    $("#songDetailCoverImg").attr({src:data.songCoverUrl});
                    //先清空数组
                    lrcTimeArray = [];
                    setLyricTime(lyricData);
                },
                error:function (err) {
                    console.log(err);
                   alert("获取歌曲详情失败")
                }
            });
        }else{
            //隐藏歌词详情，显示main
            $(".songdetail-wrapper").addClass("ng-hide");
            $(".main").removeClass("ng-hide");
            mainAndLyricDetailstateChange = true
        }
    })
}

