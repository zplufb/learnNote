package Example
/*
typedef int (*intFunc) ();

int
bridge_int_func(intFunc f)
{
     return f();
}

int fortytwo()
{
     return 42;
}*/
import "C"
import "fmt"

func Example1() {
	f := C.intFunc(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))
}
