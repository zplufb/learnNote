/**
 * @Author: FB
 * @Description: 
 * @File:  StackDemo
 * @Version: 1.0.0
 * @Date: 2019/7/18 10:17
 */

package main

import "fmt"

type StackDemo struct {
	Data []interface{}
	Len  int
}

func ShowVals(eles ...interface{}) {
	for v := range eles {
		fmt.Println(v)
	}
	fmt.Println()
}

func ShowStackInfo(s *StackDemo) {
	fmt.Printf("info=%+v", *s)
	fmt.Println()
}

func main() {

	s := NewStack()
	ShowStackInfo(s)
	s.Pop()
	ShowStackInfo(s)

	s.Push(1)
	s.Push("a")
	s.Push(12)
	ShowStackInfo(s)
	s.Pop()
	ShowStackInfo(s)

}

func NewStack() (*StackDemo) {
	return &StackDemo{}
}

func (s *StackDemo) Pop() interface{} {
	var data interface{}
	if s.Len > 0 {
		data = s.Data[s.Len-1]
		s.Len = s.Len - 1
		s.Data = s.Data[:s.Len]
	}
	return data
}

func (s *StackDemo) Push(ele interface{}) {
	s.Data = append(s.Data, ele)
	s.Len = s.Len + 1
	return
}

func (s *StackDemo) Size(ele int) int {
	return s.Len
}
