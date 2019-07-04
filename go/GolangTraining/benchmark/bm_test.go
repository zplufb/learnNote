/**
 * @Author: FB
 * @Description: 
 * @File:  bm_test.go
 * @Version: 1.0.0
 * @Date: 2019/7/4 16:34
 */

package main

import (
	"io/ioutil"
	"os"
	"bufio"
	"fmt"
	"testing"
	"strings"
)

var filename = "text.txt"
var context = "hello world!"

func init() {
	//initVar(1)
	//join string to long
	initVar(100)

}

func initVar(loop int) {
	b := strings.Builder{}
	for i := 0; i < loop; i++ {
		b.WriteString(context)
	}
	context = b.String()
	fmt.Println("len context=", len(context))
}

func BenchmarkFileW_IOutil(B *testing.B) {

	for i := 0; i < B.N; i++ {
		ioutil.WriteFile(filename, []byte(context), 0644)
	}

}

func BenchmarkFileW_OS(B *testing.B) {
	for i := 0; i < B.N; i++ {
		fp, _ := os.Create(filename)
		defer fp.Close()
		fp.WriteString(context)
		fp.Sync()

	}

}

func BenchmarkFileW_Bufio(B *testing.B) {

	for i := 0; i < B.N; i++ {
		//write
		fp, _ := os.Create(filename)
		defer fp.Close()
		writer := bufio.NewWriter(fp)
		writer.WriteString(context)
		writer.Flush()
	}

}

func BenchmarkFMT(B *testing.B) {
	for i := 0; i < B.N; i++ {
		fp, _ := os.Create(filename)
		defer fp.Close()
		fmt.Fprintf(fp, context)
	}

}

//read
func BenchmarkFileR_IOutil(B *testing.B) {

	for i := 0; i < B.N; i++ {
		ioutil.ReadFile(filename)
	}

}

func BenchmarkFileR_OS(B *testing.B) {
	for i := 0; i < B.N; i++ {
		fp2, _ := os.Open(filename)
		defer fp2.Close()
		buf := make([]byte, 1024)
		fp2.Read(buf)
	}

}

func BenchmarkFileR_Bufio(B *testing.B) {

	for i := 0; i < B.N; i++ {
		//read
		fp, _ := os.Open(filename)
		defer fp.Close()
		reader := bufio.NewReader(fp)
		buf := make([]byte, 1024)
		reader.Read(buf)
	}

}

//result as follow：

//initVar(100)
//len context= 1200
//goos: windows
//goarch: amd64
//BenchmarkFileW_IOutil-4   	   20000	     88600 ns/op
//BenchmarkFileW_OS-4       	     500	   3304000 ns/op
//BenchmarkFileW_Bufio-4    	   30000	     50566 ns/op
//BenchmarkFMT-4            	   30000	     43433 ns/op
//BenchmarkFileR_IOutil-4   	  100000	     23250 ns/op
//BenchmarkFileR_OS-4       	  100000	     20390 ns/op
//BenchmarkFileR_Bufio-4    	  100000	     26440 ns/op
//PASS
//
//
////initVar(1)
//len context= 12
//goos: windows
//goarch: amd64
//BenchmarkFileW_IOutil-4   	   30000	     46533 ns/op
//BenchmarkFileW_OS-4       	     500	   3002000 ns/op
//BenchmarkFileW_Bufio-4    	   30000	     60566 ns/op
//BenchmarkFMT-4            	   20000	     65300 ns/op
//BenchmarkFileR_IOutil-4   	  100000	     22360 ns/op
//BenchmarkFileR_OS-4       	  100000	     20230 ns/op
//BenchmarkFileR_Bufio-4    	  100000	     24600 ns/op
//PASS