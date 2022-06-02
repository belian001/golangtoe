package main

import (
	"fmt"
	"github.com/CUCyber/ja3transport"
	tls "github.com/refraction-networking/utls"
	"io/ioutil"
	"net/http"
)

type Browser struct {
	Ja3       string
	UserAgent string
}

func req(browser Browser) {
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	tr, _ := ja3transport.NewTransportWithConfig(browser.Ja3, &config)

	client := &http.Client{
		Transport: tr,
	}
	req, _ := http.NewRequest("GET", "https://ja3er.com/json", nil)
	req.Header.Set("User-Agent", browser.UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}

func main() {
	//ja3List := []Browser{
	//	{Ja3: "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53-10,0-23-65281-10-11-35-16-5-13-18-51-45-43-27,29-23-24,0",UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Safari/605.1.15"},
	//	{Ja3: "771,4865-4866-4867-49196-49195-52393-49200-49199-52392-49188-49187-49162-49161-49192-49191-49172-49171-157-156-61-60-53-47-49160-49170-10,0-23-65281-10-11-16-5-13-18-51-45-43-21,29-23-24-25,0", UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 Edg/92.0.902.67"},
	//	{Ja3: "771,4865-4866-4867-49196-49195-49188-49187-49162-49161-52393-49200-49199-49192-49191-49172-49171-52392-157-156-61-60-53-47-49160-49170-10,65281-0-23-13-5-18-16-11-51-45-43-10-21,29-23-24-25,0", UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 13_1_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.1 Mobile/15E148 Safari/604.1"},
	//}
	//for _, ja3 := range ja3List{
	//	req(ja3)
	//}
	ja3List := Browser{Ja3: "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-17513-39578-21,29-23-24,0", UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 13_1_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.1 Mobile/15E148 Safari/604.1"}
	req(ja3List)
}
