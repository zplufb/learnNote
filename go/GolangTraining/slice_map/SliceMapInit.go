/**
 * @Author: FB
 * @Description: 
 * @File:  SliceMapInit.go
 * @Version: 1.0.0
 * @Date: 2019/7/5 14:39
 */

package main

import "fmt"

func main(){
	//InitSlice()
	InitMap()
}

func InitSlice(){
	//三个等价
	var a []int
	var a1 = []int{}
	var b = make([]int,0)
	var b2 = make([]int,0,0)

	//报错
	//a[0] = 1
	//b[0] = 2
	showVals(a,a1,b,b2)
	showVals("---------end 1---------")
	//三个等价
	var c = []int{0,0,0}
	var d = make([]int,3)
	var d2 = make([]int,3,3)

	var d3 = make([]int,3,5)

	c[0] = 1
	d[0] = 2
	showVals(c,d,d2)
	showVals("---------end 2---------")

	//报错，因为长度只有3
	//d[3] = 3
	showVals(d,len(d),cap(d))

	//报错，因为长度只有3
	//d3[3] = 3
	showVals(d3,len(d3),cap(d3))
	showVals("---------end 3---------")

	//append 以上两种情况都可以
	a = append(a,1,2,3)
	b = append(b,1,2,3)
	d = append(d,1,2,3)
	d3 = append(d3,1,2,3)
	showVals("a",a,len(a),cap(a))
	showVals("b",b,len(b),cap(b))
	showVals("d",d,len(d),cap(d))
	showVals("d3",d3,len(d3),cap(d3))

}

func InitMap(){
	var a map[string]string
	showVals(a)
	// add these lines:
	/*
		myGreeting["Tim"] = "Good morning."
		myGreeting["Jenny"] = "Bonjour."
	*/
	// and you will get this:
	// panic: assignment to entry in nil map

	//OK
	var b = make(map[string]string)
	//等价写法shorthand_make
	//b := make(map[string]string)

	//init
	b["Tim"] = "Good morning."
	b["Jenny"] = "Bonjour."
	showVals(b)

	//OK
	c := map[string]string{}
	//init
	c["Tim"] = "Good morning."
	c["Jenny"] = "Bonjour."
	showVals(c)

	//init
	c1 := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	c1["FB"] = "Hi"
	c1["Jenny"] = "Hello"
	c1["Jenny2"] = "Hello2"

	showVals(c1,len(c1))
	delete(c1,"Jenny2")
	showVals(c1)

	//key := "FB2"
	key := "FB"
	if val, exists :=  c1[key];exists{
		showVals(val)
		delete(c1,key)
		showVals(c1)
	}else{
		showVals("no found",key)
	}

	//遍历
	for key, val := range c1 {
		fmt.Println(key, " - ", val)
	}

}

func showVals(vals ...interface{}){
	for _,v :=range vals{
		fmt.Print(v,"\t")
	}
	fmt.Println()
}