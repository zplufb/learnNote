package main

import (
	"os"
	"fmt"
)

var fileName = "E:/ZTestData/tvmao_sjdzd.txt"

func init() {
	//删除指定文件
	err := os.Remove(fileName)
	if err != nil {
		fmt.Errorf("err=%v", err)
	}

}
func main() {
	StartCollect()
}
