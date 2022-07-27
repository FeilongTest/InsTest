package ins

import (
	"github.com/levigross/grequests"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

func GetIP() *url.URL {
	ipUrl := "http://list.rola.info:8088/user_get_ip_list?token=lWzOrLFISJHm8PvL1650006153915&qty=1&country=&state=&city=&time=5&format=txt&protocol=http&filter=1"
Retry:
	resp, err := grequests.Get(ipUrl,
		&grequests.RequestOptions{})
	if err != nil {
		log.Println("获取代理ip出错", err)
		log.Println(resp.String())
		os.Exit(0)
		return nil
	} else {
		urli := url.URL{}
		urlproxy, err := urli.Parse("http://" + strings.TrimSpace(resp.String()))
		if err != nil {
			log.Printf("ip转换Url出错")
			log.Println("请检查是否添加代理白名单")
			log.Println(resp.String())
			if strings.Index(resp.String(), "秒后再试") != -1 {
				time.Sleep(2 * time.Second)
				goto Retry
			}
			os.Exit(0)
			return nil
		}
		return urlproxy
	}
}

func (c *Client) GetMyIP() {
	resp, err := c.Session.Get("http://ip123.in/ip.json",
		&grequests.RequestOptions{
			DialTimeout:    10 * time.Second,
			RequestTimeout: 10 * time.Second,
		})
	if err != nil {
		log.Println("err", err)
	} else {
		log.Println(resp.String())
	}
}

func (c *Client) GetIns() {
	resp, err := c.Session.Get("https://www.instagram.com",
		&grequests.RequestOptions{
			DialTimeout:    10 * time.Second,
			RequestTimeout: 10 * time.Second,
		})
	if err != nil {
		log.Println("err", err)
	} else {
		log.Println(resp.String())
	}
}
