package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"bufio"
	"fmt"
	"github.com/wangluozhe/requests"
	"github.com/wangluozhe/requests/url"
	"strconv"
	"strings"
	"unsafe"
	//"fmt"
	//"github.com/wangluozhe/requests"
	//"github.com/wangluozhe/requests/url"
)

var out *C.char

//export Http
func Http(method *C.char, rawurl *C.char, body *C.char, cook *C.char, header *C.char, ja3 *C.char, proxy *C.char, ptr *unsafe.Pointer) *C.char {

	D := Requst(C.GoString(method), C.GoString(rawurl), C.GoString(body), C.GoString(cook), C.GoString(header), C.GoString(ja3), C.GoString(proxy))
	result := C.CString(D)
	//req := url.NewRequest()
	//headers := url.NewHeaders()
	//headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	////str := C.CString("这是一段测试文本：" + C.GoString(method))
	//r, err := requests.Get("https://ja3er.com/json", req)
	//// 解引用后对齐进行赋值，把指针传出去
	//if err != nil {
	//	//fmt.Println(err)
	//	result := C.CString("err")
	//	//return C.CString("err")
	//	*ptr = unsafe.Pointer(result)
	//	return result
	//}
	//result :=  C.CString(r.Text)
	*ptr = unsafe.Pointer(result)
	////C.free( unsafe.Pointer(str))
	return result
}

//export Testja3
func Testja3(proxy *C.char, ptr *unsafe.Pointer) *C.char {

	req := url.NewRequest()
	headers := url.NewHeaders()
	headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	////str := C.CString("这是一段测试文本：" + C.GoString(method))
	req.Headers = headers
	r, err := requests.Get("https://ja3er.com/json", req)
	//// 解引用后对齐进行赋值，把指针传出去

	if err != nil {
		//fmt.Println(err)
		result := C.CString("err")
		//return C.CString("err")
		*ptr = unsafe.Pointer(result)
		return result
	}
	result := C.CString(r.Text)
	*ptr = unsafe.Pointer(result)
	////C.free( unsafe.Pointer(str))
	return result
}

//export Test
func Test(cs *C.char, ptr *unsafe.Pointer) *C.char {
	//str := C.CString("这是一段测试文本：" + C.GoString(cs))
	// 解引用后对齐进行赋值，把指针传出去
	s := strings.Split(C.GoString(cs), `----`)
	O := strconv.Itoa(len(s))
	str := C.CString(O)
	*ptr = unsafe.Pointer(str)
	return str
}

//export Free
func Free(ptr unsafe.Pointer) {
	// 释放内存
	C.free(ptr) //C.CString 必须Free 否则会导致内存泄漏
}

func Requst(method string, rawurl string, body string, cook string, header string, ja3 string, proxy string) string {
	// 释放内存
	req := url.NewRequest()
	headers := url.NewHeaders()
	headerlen := SplitLines(header)
	if header != "" {
		for i := 0; i < len(headerlen); i++ {
			c := strings.Split(headerlen[i], `: `)
			headers.Set(c[0], strings.Replace(c[1], " ", "", -1))
		}
	} else {
		headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	}

	if proxy != "" {
		req.Proxies = proxy
	}
	if ja3 != "" {
		req.Ja3 = ja3
	}
	if cook != "" {
		req.Cookies = url.ParseCookies(rawurl, cook)
	}
	if body != "" {
		data := url.NewData()
		data.Set("body", body)
		req.Data = data
	}

	req.AllowRedirects = false

	req.Headers = headers
	//req.Cookies = url.ParseCookies(rawUrl,"_ga=GA1.1.630251354.1645893020; Hm_lvt_def79de877408c7bd826e49b694147bc=1647245863,1647936048,1648296630; Hm_lpvt_def79de877408c7bd826e49b694147bc=1648301329")
	//req.Ja3 = "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-34,0"
	r, err := requests.Request(method, rawurl, req)
	//r, err := requests.Get("https://ja3er.com/json", req)
	if err != nil {
		fmt.Println(err)

		//return C.CString("err")
		//*ptr = unsafe.Pointer(result)
		return "err"
	}
	//result := C.CString(r.Text)
	//fmt.Println(unsafe.Pointer(result))
	//V :=C.CString(P)
	//fmt.Println((*string)(unsafe.Pointer(result)))
	//Free(unsafe.Pointer(result))
	//fmt.Println(V)
	//b := []byte(P )
	//fmt.Println(b)
	//K :=C.GoString(V)
	//fmt.Println(K)
	return r.Text
}
func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

//set GOARCH=386
//set CGO_ENABLED=1
//go build -ldflags "-s -w" -buildmode=c-shared -o dlldemo.dll  dlldemo.go
func main() {
	//c :=Demoja3("demo ")
	//fmt.Println(string(c))
	header := `User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36
Accept-Encoding: gzip, deflate
Accept: application/json, text/javascript, */*; q=0.01
Connection: keep-alive
Accept-Language: zh-CN,zh;q=0.9
Content-Type: application/x-www-form-urlencoded; charset=UTF-8
Origin: https://www.bilibili.com
Referer: https://www.bilibili.com/video/BV1b34y1Y7ae?spm_id_from=333.1073.sub_channel.dynamic_video.click
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: same-site
sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"
sec-ch-ua-mobile: ?0
sec-ch-ua-platform: "Windows"
`

	//Host: api.bilibili.com
	headerlen := SplitLines(header)
	fmt.Println(headerlen)
	fmt.Println(len(headerlen))
	if header != "" {
		for i := 0; i < len(headerlen); i++ {
			c := strings.Split(headerlen[i], `: `)
			fmt.Println(c[0] + strings.Replace(c[1], " ", "", -1))
		}
	}

	//127.0.0.1:12230"

	//771,39578-4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,31354-0-23-65281-10-11-35-16-5-13-18-51-45-43-27-17513-51914-21,31354-29-23-24,0

	//771,14906-4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,2570-0-23-65281-10-11-35-16-5-13-18-51-45-43-27-17513-6682-21,27242-29-23-24,0
	pro := "http://114.239.198.13:4245"
	pro = "http://111.77.127.167:18442"

	ja3 := "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0"
	//C := Requst("POST","https://api.bilibili.com/x/click-interface/web/heartbeat","aid=725914505&cid=581243339&bvid=BV1PS4y1a7CK&mid=0&csrf=&played_time=0&real_played_time=0&realtime=0&start_ts=1650609750&type=3&dt=2&play_type=1&from_spmid=333.6.0.0&spmid=333.788.0.0&auto_continued_play=0&refer_url=&bsource=","buvid3=1F621DFF-3B78-0D05-C160-F6E960CC271952218infoc; b_nut=1650264253; buvid4=A4E42930-0058-2588-055F-904C2684EFD652218-022041814-4tYW16MtCosARcVuctiLBg==; innersign=0; i-wanna-go-back=-1; b_ut=7; b_lsid=5ABFBDA10_1803B6A3BC8; _uuid=EC5631033-B1021-49BA-1EC7-9AEAB9FA244A59535infoc",header,ja3,pro)

	//C := Requst("POST","https://api.bilibili.com/x/click-interface/click/web/h5https://api.bilibili.com/x/click-interface/click/web/h5","","",header,ja3,pro)

	C := Requst("GET", "https://ja3er.com/json", "", "", header, ja3, pro)
	//C := Requst("GET","https://www.bilibili.com","","",header,ja3,pro)
	fmt.Println(C)
	// Need a main function to make CGO compile package as C shared library
}
