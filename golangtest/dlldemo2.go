package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

var out *C.char

//export Test
func Test(cs *C.char, ptr *unsafe.Pointer) *C.char {
	str := C.CString("这是一段测试文本：" + C.GoString(cs))
	// 解引用后对齐进行赋值，把指针传出去
	*ptr = unsafe.Pointer(str)
	return str
}

//export Free
func Free(ptr unsafe.Pointer) {
	// 释放内存
	C.free(ptr) //C.CString 必须Free 否则会导致内存泄漏
}


//set GOARCH=386
//set CGO_ENABLED=1
//go build -ldflags "-s -w" -buildmode=c-shared -o dlldemo.dll  dlldemo.go
func main() {
	// Need a main function to make CGO compile package as C shared library
}
