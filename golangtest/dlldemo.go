//package main
//
///*
//#include <stdlib.h>
//*/
//import "C"
//import (
//	"unsafe"
//)
//
//var out *C.char
//
////export Test
//func Test(cs *C.char, ptr *unsafe.Pointer) *C.char {
//	str := C.CString("这是一段测试文本：" + C.GoString(cs))
//	// 解引用后对齐进行赋值，把指针传出去
//	*ptr = unsafe.Pointer(str)
//	return str
//}
//
////export Free
//func Free(ptr unsafe.Pointer) {
//	// 释放内存
//	C.free(ptr) //C.CString 必须Free 否则会导致内存泄漏
//}
//////export requst
////func requst()  *C.char{
////	// 释放内存
////	req := url.NewRequest()
////	headers := url.NewHeaders()
////	headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
////	req.Headers = headers
////	//req.Cookies = url.ParseCookies(rawUrl,"_ga=GA1.1.630251354.1645893020; Hm_lvt_def79de877408c7bd826e49b694147bc=1647245863,1647936048,1648296630; Hm_lpvt_def79de877408c7bd826e49b694147bc=1648301329")
////	req.Ja3 = "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-34,0"
////	r, err := requests.Request("GET","https://ja3er.com/json", req)
////	//r, err := requests.Get("https://ja3er.com/json", req)
////	if err != nil {
////		return C.CString("err")
////		//fmt.Println(err)
////	}
////	//fmt.Println(r.Text)
////	return C.CString(r.Text)
////
////}
//
////set GOARCH=386
////set CGO_ENABLED=1
////go build -ldflags "-s -w" -buildmode=c-shared -o dlldemo.dll  dlldemo.go
//func main() {
//	// Need a main function to make CGO compile package as C shared library
//}
