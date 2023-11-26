package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// 请求的查询参数
// 请求头
func printBody(resp *http.Response) {
	// 每种请求都要关闭响应体
	defer func() { resp.Body.Close() }()

	// 读取body内容
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("读取到的内容 ", string(content))
}

// get请求用
func requestByParams(request *http.Request, data map[string]string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)

	if err != nil {
		panic(err)
	}

	// 拼接请求参数
	params := make(url.Values)
	for key, val := range data {
		params.Add(key, val)
	}

	// 放到url上去
	request.URL.RawQuery = params.Encode()
	return request
}

func requestByHeaders() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)

	if err != nil {
		panic(err)
	}

	// 直接设置headers即可
	request.Header.Set("user-agent", "chrome")

	// params
	params := make(map[string]string)
	params["age"] = "22"
	params["name"] = "timotte"
	request = requestByParams(request, params)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	printBody(resp)
}
