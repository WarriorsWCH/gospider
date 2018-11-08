package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	// "io"
	"bufio"
	// "regexp"
	"log"
)


func Fetch(url string) ([]byte, error) {

	// resp, err := http.Get(url)
	// 直接用http.Get(url)进行获取信息会报错：Error: status code 403
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	// 查看自己浏览器中的User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 如果状态码不是OK 打印错误
	if resp.StatusCode != http.StatusOK{
		// fmt.Println("Error: status code:",resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	// 猜html的charset的encoding是什么
	e := detemineEncoding(bodyReader)

	utf8Reader := transform.NewReader(
		resp.Body, e.NewDecoder())

	return  ioutil.ReadAll(utf8Reader)
}

func detemineEncoding(r *bufio.Reader) encoding.Encoding {
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 字节数据
	bytes, err :=  r.Peek(1024)
	if err != nil {
		// panic(err)
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}















