/*
Author      : lls
Time        : 2018/09/14
Description :
*/

package main

//void SayHello(const char* s);
import "C"

func main() {
	C.SayHello(C.CString("Hello cgo"))
}
