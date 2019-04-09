package request

import (
	"io"
	"net/http"
	"crypto/tls"
	"time"
	"bytes"
)

// GET 请求
/**
参数：
	url	请求的URL
	header	配置请求Header
	timeout	请求超时时间s
响应：
	code	响应码
	respIO	请求返回的IO，可使用ioutil.ReadAll(resp)读取
	respHeader	响应头
	err		错误信息
 */
func Get(url string, headers map[string]string,
	timeout time.Duration) (code int, respIO io.Reader, respHeader map[string][]string, err error) {
	// 新建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1, nil, nil, err
	}
	// 设置Header
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	// 忽略Https、设置超时
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
		Timeout: timeout * time.Second,
	}
	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, nil, nil, err
	}
	// 关闭连接
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	// 获取响应
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return resp.StatusCode, nil, nil, err
	}

	return resp.StatusCode, buf, resp.Header, nil
}

// POST 请求
/**
参数：
	url	请求的URL
	header	配置请求Header
	data	请求提交的数据
	timeout	请求超时时间s
响应：
	code	响应码
	respIO	请求返回的IO，可使用ioutil.ReadAll(resp)读取
	respHeader	响应头
	err		错误信息
 */
func Post(url string, headers map[string]string, data string,
	timeout time.Duration) (code int, respIO io.Reader, respHeader map[string][]string, err error) {
	// 新建请求
	reqBody := bytes.NewBuffer([]byte(data))
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return -1, nil, nil, err
	}
	// 设置Header
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	// 忽略Https、设置超时
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
		Timeout: timeout * time.Second,
	}
	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, nil, nil, err
	}
	// 关闭连接
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	// 获取响应
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return resp.StatusCode, nil, nil, err
	}
	return resp.StatusCode, buf, resp.Header, nil
}
