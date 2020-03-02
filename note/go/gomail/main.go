/**
 * @Author: FB
 * @Description: 
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 14:20
 */

package main

import (
	_ "gomail/ServerInit"
	"net/http"
	"fmt"
)

func main() {
	// 开启web服务器，并监听58899端口
	fmt.Println("Gomail is Running On IP 192.168.3.123 Port 58899")
	err := http.ListenAndServe("0.0.0.0:58899", nil)
	//err := http.ListenAndServeTLS("0.0.0.0:58888", "cert.crt", "key.crt", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
