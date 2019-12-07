package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var agents = []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60",
	"Opera/8.0 (Windows NT 5.1; U; en)",
	"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.1) Gecko/20061208 Firefox/2.0.0 Opera 9.50",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 9.50",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0",
	"Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.57.2 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.133 Safari/534.16",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.11 TaoBrowser/2.0 Safari/536.11",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.71 Safari/537.1 LBBROWSER",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; LBBROWSER)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E; LBBROWSER)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; QQBrowser/7.0.3698.400)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)",
	"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.84 Safari/535.11 SE 2.X MetaSr 1.0",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SV1; QQDownload 732; .NET4.0C; .NET4.0E; SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Maxthon/4.4.3.4000 Chrome/30.0.1599.101 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 UBrowser/4.0.3214.0 Safari/537.36",
}

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

// getNewUserAgent用于获取一个新的agent
func getNewUserAgent() string {
	rand.Seed(time.Now().UnixNano())
	return agents[rand.Intn(10)]
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
	req.Header.Set("user-agent", getNewUserAgent())
	client := &http.Client{}
	return client.Do(req)
}
