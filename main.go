package main

import (
	"fmt"
	"io"
	"net/http"
)

// 基本的四种方法实现
// 步骤：构造请求体
// 设置client
// client.do
// 读取response

func main() {
	//get()
	//post()
	//put()
	//delete()

	//requestByHeaders()

	responseTest()
}

func get() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	// 每种请求都要关闭响应体
	defer func() { resp.Body.Close() }()

	// 读取body内容
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("读取到的内容 ", string(content))
}

func post() {
	resp, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		panic(err)
	}
	// 每种请求都要关闭响应体
	defer func() { resp.Body.Close() }()

	// 读取body内容
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("读取到的内容 ", string(content))
}

func put() {
	// 内置没有put请求，模拟get、post
	// 1. 构造请求，然后调用http.DefaultClient发送
	req, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(req)
	defer func() { response.Body.Close() }()

	if err != nil {
		panic(err)
	}
	// 读取body内容
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("读取到的内容 ", string(content))
}

func delete() {
	// 内置没有put请求，模拟get、post
	// 1. 构造请求，然后调用http.DefaultClient发送
	req, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(req)
	defer func() { response.Body.Close() }()

	if err != nil {
		panic(err)
	}
	// 读取body内容
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("读取到的内容 ", string(content))
}
