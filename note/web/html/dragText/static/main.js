//init global let
$().ready(()=>{


function init(){
if(typeof(Storage)=="undefined"){
	console.log("the browser don't support Storage")	
}else{
	//initData
	initData()
}
}

function initData(){
	str_g_arr = localStorage.getItem("str_g_arr");
	if( str_g_arr == null  ){
		g_arr = []
	}else{
		g_arr = str_g_arr.split(",")

	var htmlHigh = localStorage.getItem("htmlHigh");
	var htmlMid = localStorage.getItem("htmlMid");
	var htmlLow = localStorage.getItem("htmlLow");
	var htmlCan = localStorage.getItem("htmlCan");
	var htmlRec = localStorage.getItem("htmlRec");

		 $("#level_high").html(htmlHigh)
		 $("#level_mid").html(htmlMid)
		 $("#level_low").html(htmlLow)
		 $("#candicate_list").html(htmlCan)
		 $("#level_recycle").html(htmlRec)
	}
}

//初始化
init()


})


let g_arr = []

function newEvent(){

let list = $("#candicate_list")
itemArr_fliter = promptNewEvent()
list.append(itemArr_fliter)
refresh()
}


function clearEventH(){

$("#level_high").html("")
refresh()

}

function clearEventM(){

$("#level_mid").html("")
refresh()
}

function clearEventL(){


$("#level_low").html("")
refresh()
}

function clearEventR(){


$("#level_recycle").html("")
refresh()
}

function clearAllEvent(){

$("#level_high").html("")
$("#level_mid").html("")
$("#level_low").html("")
refresh()
}

function deleleEventCandicate(){

$("#candicate_list").html("")
refresh()
}



function promptNewEvent(){

let tipName = "请输入事件名称(多个事件用空格隔开)"
	let eventName = prompt(tipName,"")
	eventName = eventName.trim()
	while(eventName === null || eventName === ""){
		eventName = prompt(tipName,"")
	}

	eventNameStr = eventName.replace(/\s+/g," ")
	eventNameArr = eventNameStr.split(" ")


respData = in_array_filter(g_arr,eventNameArr)
let eventNameArr_fliter = respData[0]
let flag = respData[1]

console.log("eventNameArr_fliter=",eventNameArr_fliter)
if(flag){
  alert("事件名称与之前重复，但已经进行过滤")
}


localStorage.setItem("str_g_arr",g_arr);
console.log("g_arr=",g_arr)

itemArr = []
for(let i=0;i<eventNameArr_fliter.length;i++){
	let ele = eventNameArr_fliter[i]
	item = '<div class="item" id="'+ele+'" draggable="true" ondragstart="drag(event)" width="auto" height="69">'+ele+'</div>'
	itemArr[i] = item
}

	return itemArr
}




function allowDrop(ev)
{
	ev.preventDefault();
}

function drag(ev)
{
	ev.dataTransfer.setData("Text",ev.target.id);
}

function drop(ev)
{
	ev.preventDefault();
	let data=ev.dataTransfer.getData("Text");
	ev.target.appendChild(document.getElementById(data));
	refresh()
}

function in_array_filter(arr,eleArr){

//先过滤自身重复元素
let eleArr_filter = []
for(let i=0;i<eleArr.length;i++){
  let ele = eleArr[i]
  if (eleArr_filter.indexOf(ele) == -1){
    eleArr_filter.push(ele)
  }
}

//过滤全局是否有重复元素
arr_filter = [].concat(eleArr_filter)

let eleArr_filter2 = []
let flag = false
len=arr_filter.length
for(let i=0;i<len;i++){
  let ele = eleArr[i]
  if (arr.indexOf(ele) == -1){
    eleArr_filter2.push(ele)
    arr.push(ele)
  }else{
    flag = true
  }
}

console.log("arr",arr)


return [eleArr_filter2,flag]

}

function refresh(){
	let htmlHigh = $("#level_high").html()
	let htmlMid = $("#level_mid").html()
	let htmlLow = $("#level_low").html()
	let htmlCan = $("#candicate_list").html()
	let htmlRec = $("#level_recycle").html()


	localStorage.setItem("htmlHigh",htmlHigh);
	localStorage.setItem("htmlMid",htmlMid);
	localStorage.setItem("htmlLow",htmlLow);
	localStorage.setItem("htmlCan",htmlCan);
	localStorage.setItem("htmlRec",htmlRec);

}