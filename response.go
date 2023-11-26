package main

import (
	"fmt"
	"io"
	"net/http"
)

func responseBody(resp *http.Response) {
	//// 利用io.Read直接读出来
	//content, _ := io.ReadAll(resp.Body)
	//fmt.Println("body内容：", string(content))

	var content2 string
	// 1024字节的缓冲区
	buf := make([]byte, 1024)
	// 分块读
	for {
		n, err := resp.Body.Read(buf)

		if err == io.EOF {
			break
		}

		content2 += string(buf[:n])
	}
	fmt.Println("分块读：", content2)
}

func status(resp *http.Response) {
	fmt.Printf("状态码：%d，状态信息: %s", resp.StatusCode, resp.Status)
}

func header(resp *http.Response) {
	fmt.Println("查看某个头 ", resp.Header.Get("content-type"))
}

// 获得body的编码请求
func encoding() {

}

func responseTest() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer func() { resp.Body.Close() }()

	responseBody(resp)
	status(resp)
	header(resp)
}
