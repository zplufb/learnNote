package main
import "C"
import "fmt"

//export myprint
func myprint(i C.int) {
	fmt.Printf("i = %v\n", uint32(i))
}

//export GoExportedFunc
func GoExportedFunc() {
	fmt.Println("I am a GoExportedFunc!")
}