var imgName = $('.slot');
var rewardName = $('.name');
var runing = true;
var trigger = true;
var num = 0;
var LotterynumberInit = 1; //设置单次抽奖人数
var Lotterynumber = LotterynumberInit; //设置单次抽奖人数
var speedInit = 50
var speed = speedInit //单位是毫秒
var t

var imgNames = [
    "candy",
    "pencil",
    "eraser",
    "tag",
    "coin",
    "notebook",
    "luckydraw1",
    "luckydraw3",
]


var imgs = imgNames.map(x => "img/" + x + ".jpg")


var rewards = [
    "糖果",
    "铅笔",
    "橡皮",
    "贴纸",
    "尚学币",
    "笔记本",
    "抽奖两次",
    "抽奖三次"
]

var imgCount = imgs.length - 1;//物品种类

//init func
// var evt = "onorientationchange" in window ? "orientationchange" : "resize";
// window.addEventListener(evt, resize, false);
// function resize() {

//     if (window.orientation == 0 || window.orientation == 180) { //横屏 oriental
//         //直接默认不让滑动
//         stopMove();
//     } else {
//         // 竖屏 vertical，微信有异常，手机或者电脑端网页都正常
//         if (!is_wexin()) {
//             stopMove();
//         }
//     }
// }

//禁止页面滑动for iphone
var preventFunc = function (e) {
    e.preventDefault();
};

function stopMove() {
    document.body.style.overflow = 'hidden';
    document.addEventListener("touchmove", preventFunc, {passive: false});//禁止页面滑动
}

function is_wexin() {
    var ua = navigator.userAgent.toLowerCase();
    if (ua.match(/MicroMessenger/i) == "micromessenger") {
        return true;
    } else {
        return false;
    }
}

function counter(){

    if(typeof(Storage)!=="undefined"){
        
        var num = localStorage.getItem("pageviews");
        var num2 = localStorage.getItem("num2");
        var num3 = localStorage.getItem("num3");

        if (num != null){
            localStorage.setItem("pageviews", Number(num) + 1);
        }else{
            localStorage.setItem("pageviews", 1);
        }

         num = localStorage.getItem("pageviews");

        document.getElementById("counter").innerHTML="Total：" + num+"  |Extra2："+num2+"  |Extra3："+num3;
    }else{
        document.getElementById("counter").innerHTML="Sorry, Your browser is unable to support web storage";
    }
}

//init
stopMove()
initNum()

function initNum(){
     if(typeof(Storage)!=="undefined"){

        var num2 = localStorage.getItem("num2");
        var num3 = localStorage.getItem("num3");

        if(num2 == null){
            localStorage.setItem("num2", 0);     
        }
        if(num3 == null){    
            localStorage.setItem("num3", 0);  
        }
      
     }
}

$("#counter").click(function(){
    clearLocalStorage()
})

function clearLocalStorage(){
    localStorage.setItem("pageviews", 0);
    localStorage.setItem("num2", 0); 
    localStorage.setItem("num3", 0); 
    document.getElementById("counter").innerHTML="Total：" + 0+"  |Extra2："+0+"  |Extra3："+0;
}


$(function () {
    imgName.css('background-image', 'url(' + imgs[0] + ')');
    rewardName.html(rewards[0]);
});

function clearList() {
    //清除中奖名单
    $('.luck-user-list').text("")
}

// 开始停止
function start() {
    //清除中奖名单
    $('.luck-user-list').text("")

    if (runing) {

        runing = false;
        $('#start').text('停止');
        startNum()

    } else {
        $('#start').text('自动抽取中...');
        print();
    }

}

// 开始抽奖
function startLuck() {
    runing = false;
    $('#btntxt').removeClass('start').addClass('stop');
    startNum()
}

// 循环参加名单
function startNum() {
    num = Math.floor(Math.random() * imgCount);
    imgName.css('background-image', 'url(' + imgs[num] + ')')
    rewardName.html(rewards[num]);

    t = setTimeout(startNum, speed);

}

function getRewardNumByProbability() {
    // 糖果一颗（15），铅笔一支（15），橡皮一个（15），贴纸一张（20），尚学币一个（5），笔记本一本（5），再抽奖一次（5）,再抽奖两次（1）
    //num=[0,80],rewardNum=[0,7]
    var rewardNum = 0
    num = Math.floor(Math.random() * 81);
    if (num >= 1 && num <= 15) {
        rewardNum = 0
    } else if (num >= 16 && num <= 30) {
        rewardNum = 1
    } else if (num >= 31 && num <= 45) {
        rewardNum = 2
    } else if (num >= 46 && num <= 65) {
        rewardNum = 3
    } else if (num >= 66 && num <= 70) {
        rewardNum = 4
    } else if (num >= 71 && num <= 75) {
        rewardNum = 5
    } else if (num >= 76 && num <= 80) {
        rewardNum = 6
    } else { //num == 0
        rewardNum = 7
    }

    console.log("rewardNum=", rewardNum)

    imgName.css('background-image', 'url(' + imgs[rewardNum] + ')')
    rewardName.html(rewards[rewardNum]);


    return rewardNum
}

// 停止跳动
function stop() {
    //统计使用次数
     counter()
    //imgCount = imgs.length-1;
    clearTimeout(t);
    t = 0;
}

function judgeBigRewards() {

    var extraNum = 1
    var text = $(".luck-user-list").html()
    if (text.lastIndexOf("luckydraw1") > -1) {
        extraNum = 2
        localStorage.setItem("num2", Number(localStorage.getItem("num2"))+1);
    } else if (text.lastIndexOf("luckydraw3") > -1) {
        extraNum = 3
        localStorage.setItem("num3", Number(localStorage.getItem("num3"))+1);
    }
    console.log("extraNum=", extraNum)
    return extraNum
}

// 打印中奖清单
function print() {
    if (trigger) {

        trigger = false;
        var i = 0;

        if (imgCount >= Lotterynumber) {
            stopTime = window.setInterval(function () {
                if (runing) {
                    runing = false;
                    $('#btntxt').removeClass('start').addClass('stop');
                    startNum();
                } else {
                    runing = true;
                    $('#btntxt').removeClass('stop').addClass('start');
                    stop();

                    i++;
                    Lotterynumber--;
                    if (Lotterynumber == 0) {
                        $('#start').text('自动抽取中...');
                    } else {
                        $('#start').text('自动抽取中(' + Lotterynumber + ')...');
                    }

                    //打印中奖物品清单
                    var rewardNum = getRewardNumByProbability()
                    $('.luck-user-list').prepend("<li><div class='portrait' style='background-image:url(" + imgs[rewardNum] + ")'></div><div class='luckuserName'>" + rewards[rewardNum] + "</div></li>")


                    if (i >= LotterynumberInit) {
                        console.log("抽奖结束");
                        window.clearInterval(stopTime);

                        LotterynumberInit = judgeBigRewards()
                        if (LotterynumberInit > 1) {
                            $('#start').text("开始x" + LotterynumberInit);
                        } else {
                            $('#start').text("开始");
                        }

                        Lotterynumber = LotterynumberInit;
                        trigger = true;

                    } else {
                        stop()
                        startLuck()
                    }

                }
            }, 1000);
        }
    }
}

