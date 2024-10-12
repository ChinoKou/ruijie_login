package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func (c *Config) connTest() int {
	client := &http.Client{
		Timeout: time.Second * 5, // 设置超时时间为5秒
	}
	resp, err := client.Get("http://www.baidu.com")
	if err != nil {
		return ERR_TIMEOUT
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态码
	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
		return ERR_TIMEOUT
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ERR_TIMEOUT
	}

	// 检查响应内容
	if strings.HasPrefix(string(body), "<script>top.self.location.href=") {
		return NEED_LOGIN
	}

	return ONLINE
}

func (c *Config) connTestLoop(doLoginReqChan chan bool, loginResultChan chan int) {
	for {
		testResult := c.connTest()
		if testResult != ONLINE {
			for {
				log.Println("offline! re-connecting")
				doLoginReqChan <- true
				loginResult := <-loginResultChan
				if loginResult == ERR_TIMEOUT || loginResult == ERR_LOGIN_FAILED {
					time.Sleep(time.Duration(c.RetryInterval) * time.Second)
					break
				}
				if loginResult == LOGIN_SUCCESS {
					break
				}
			}
		} else {
			log.Println("connectivity check success")
			time.Sleep(time.Second * time.Duration(c.TTLInterval))
		}
	}
}
