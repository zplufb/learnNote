/**
 * @Author: FB
 * @Description: 
 * @File:  solu.go
 * @Version: 1.0.0
 * @Date: 2019/7/11 10:43
 */

package main

import "fmt"


func main() {
	f := factorial2(4)

	for n := range f{
		fmt.Println("Total:", n)
	}

}

func factorial2(n int) chan int {
	out := make(chan int)

	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()

	return out
}