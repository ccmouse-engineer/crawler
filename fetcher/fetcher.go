package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var limiter = time.Tick(10 * time.Millisecond)

// Fetch发送一个请求并获取响应内容
func Fetch(url string) ([]byte, error) {
	<-limiter
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d\n", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// determineEncoding通过检查确定HTML文档的编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	peek, err := r.Peek(1024)
	if err != nil {
		logrus.Printf("Fetcher error: %v\n", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(peek, "")
	return e
}

// get发送GET方式的HTTP请求
func get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("cookie", "sid=f9127967-5bff-46e4-a826-8a75821fc669; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1575434476; FSSBBIl1UgzbN7N443S=UBl0uRopClR39xc1smXI9Mo2MDlvznHl5iX0lBJAptp2U.2SPNrNX_2VFKhMnagI; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1575447468; FSSBBIl1UgzbN7N443T=4e6Pci82iTp78eR_gymot8NSrCaSuEuH.Rrt0s7_sRPulIPwYvpFVe01eJmqQ7S95JRMDgBR75v0rO4X68IFsBVkO7EEigIegzYO57r1I.pw7NRUdLeudeHVjZH6LTU1qoHxfo_Xazuv_bnml0wYFc.l.cF.RZWusJXdvnb2El9vSUy9vES2gdFERzS_gfYnPjk5qhYOH_E.xyKlLNrHU_qlz6LRZ7oTcWB9nLAE9vHRvhxaN9S62izVUTkzEhs6Gz5cxo.Nyh89sEznPirbTEvnrB5Pr6b9ea8IwIu1wi2MlpVjnaGgs2RYNm.ymu9T3nC9m2z.jSGKpYhYTXuQHi6OGAlaAOQJO2fAGi7cXCSeY1Q2HRYqA_ZK7Dp.6SKnek2G")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	client := &http.Client{}
	return client.Do(req)
}
