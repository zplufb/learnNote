//init global var
var g_arr = []

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
	var data=ev.dataTransfer.getData("Text");
	ev.target.appendChild(document.getElementById(data));
}

function newEvent(){

var list = $("#candicate_list")
itemArr_fliter = promptNewEvent()
list.append(itemArr_fliter)

}


function clearEventH(){

$("#level_high").html("")

}

function clearEventM(){

$("#level_mid").html("")

}

function clearEventL(){


$("#level_low").html("")

}

function clearAllEvent(){

$("#level_high").html("")
$("#level_mid").html("")
$("#level_low").html("")

}

function deleleEventCandicate(){

$("#candicate_list").html("")
}



function promptNewEvent(){

var tipName = "请输入事件名称(多个事件用空格隔开)"
	var eventName = prompt(tipName,"")
	eventName = eventName.trim()
	while(eventName === null || eventName === ""){
		eventName = prompt(tipName,"")
	}

	eventNameStr = eventName.replace(/\s+/g," ")
	eventNameArr = eventNameStr.split(" ")


respData = in_array_filter(g_arr,eventNameArr)
var eventNameArr_fliter = respData[0]
var flag = respData[1]

console.log("eventNameArr_fliter=",eventNameArr_fliter)
if(flag){
  alert("事件名称与之前重复，但已经进行过滤")
}

console.log("g_arr=",g_arr)

itemArr = []
for(var i=0;i<eventNameArr_fliter.length;i++){
	var ele = eventNameArr_fliter[i]
	item = '<div class="item" id="'+ele+'" draggable="true" ondragstart="drag(event)" width="auto" height="69">'+ele+'</div>'
	itemArr[i] = item
}

	return itemArr
}



function in_array_filter(arr,eleArr){

//先过滤自身重复元素
var eleArr_filter = []
for(var i=0;i<eleArr.length;i++){
  var ele = eleArr[i]
  if (eleArr_filter.indexOf(ele) == -1){
    eleArr_filter.push(ele)
  }
}

//过滤全局是否有重复元素
arr_filter = [].concat(eleArr_filter)

var eleArr_filter2 = []
var flag = false
len=arr_filter.length
for(var i=0;i<len;i++){
  var ele = eleArr[i]
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
