package main

/*

#include<stdio.h>
//extern void myprint(int i);
void myprint2(int i) {
	//fmt.Printf("i = %v\n", uint32(i))
	printf("\t%d",i);
}

void dofoo(void) {
    int i;
    for (i=0;i<10;i++) {
        myprint2(i);
    }
}

//extern void GoExportedFunc();
//
//void bar() {
//       printf("I am bar!\n");
//       GoExportedFunc();
//}
*/
import "C"
import "Demo/DemoCGO/Example"

func main() {
	C.dofoo()
	C.myprint2(222)
	Example.Example1()

	//C.bar()
}

