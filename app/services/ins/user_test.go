package ins

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

var targetUrl = ShareDataUrl

func TestRun2(t *testing.T) {
	targetUrl = "http://ip123.in/ip.json"
	testHttp()
}

//http代理
func httpProxy(proxyUrl, user, pass string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("http proxy panic:%s", err)
		}
	}()
	urli := url.URL{}

	if !strings.Contains(proxyUrl, "http") {
		proxyUrl = fmt.Sprintf("http://%s", proxyUrl)
	}

	urlProxy, _ := urli.Parse(proxyUrl)
	//if user != "" && pass != "" {
	//	urlProxy.User = url.UserPassword(user, pass)
	//}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(urlProxy),
		},
	}

	rqt, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		panic(err)
		return
	}

	response, err := client.Do(rqt)
	if err != nil {
		panic(err)
		return
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	log.Printf("%s http success %s,%s", proxyUrl, response.Status, string(body))
	return
}

func testHttp() {
	httpProxy("", "", "")
}
