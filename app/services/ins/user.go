package ins

import (
	"InsTest/app/common/request"
	"InsTest/app/model"
	"InsTest/global"
	"InsTest/utils"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/levigross/grequests"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (c *Client) getKey() (data model.ShareDataResp, err error) {
	resp, err := c.Session.Get(ShareDataUrl, &grequests.RequestOptions{
		Headers:   c.Header,
		UserAgent: c.UserAgent,
	})
	if err != nil {
		err = fmt.Errorf("获取密钥出错%v", err)
	} else {
		err = jsoniter.Unmarshal(resp.Bytes(), &data)
		if err == nil {
			c.Header["cookie"] = c.makeCookies(resp.Header["Set-Cookie"])
			global.App.Log.Info("登录前的Cookies为" + c.Header["cookie"])
		}
	}
	return
}

func Login(params request.LoginData) (response model.LoginResponse, err error) {
	//首先获取csrfToken

	c := Client{
		Header: map[string]string{
			"origin":           "https://www.instagram.com",
			"referer":          "https://www.instagram.com/",
			"x-requested-with": "XMLHttpRequest",
		},
	}
	ipProxy := GetIP()
	//urli := url.URL{}
	//ipProxy, _ := urli.Parse("http://" + strings.TrimSpace("149.28.144.253:7408"))

	netTransport := &http.Transport{
		Proxy: http.ProxyURL(ipProxy),
	}
	httpClient := &http.Client{
		Timeout:   time.Second * 60,
		Transport: netTransport,
	}
	c.Session = grequests.NewSession(&grequests.RequestOptions{
		HTTPClient: httpClient,
	})
	c.GetMyIP()

	keyData, err := c.getKey()
	if err == nil {
		data := model.LoginData{
			KeyId:     keyData.Encryption.KeyId,
			PublicKey: keyData.Encryption.PublicKey,
			Time:      strconv.Itoa(int(time.Now().Unix())),
		}
		//开始登陆
		pwd, _ := utils.Encrypt(data.KeyId, data.PublicKey, []byte(params.Password), []byte(data.Time))

		var header1 = make(map[string]string)
		header1 = utils.CopyMap(header1, c.Header)
		header1["content-type"] = "application/x-www-form-urlencoded"
		header1["x-csrftoken"] = keyData.Config.CsrfToken
		resp, err := c.Session.Post(LoginUrl, &grequests.RequestOptions{
			Data: map[string]string{
				"enc_password":         pwd,
				"username":             params.Username,
				"queryParams":          "{}",
				"optIntoOneTap":        "false",
				"stopDeletionNonce":    "",
				"trustedDeviceRecords": "{}",
			},
			Headers:   header1,
			UserAgent: UserAgent,
		})
		if err != nil {
			global.App.Log.Error("登录出错", zap.String("login error", err.Error()))
		} else {
			//生成登录中的Cookies
			c.Header["cookie"] = c.makeCookies(resp.Header["Set-Cookie"])
			c.getFinalCookie(header1)
			global.App.Log.Info("登录成功" + resp.String())
			global.App.Log.Info("最终的Cookies为" + c.Header["cookie"])

			var result model.LoginResponse
			err = jsoniter.Unmarshal(resp.Bytes(), &result)
			if err != nil {
				global.App.Log.Error("登录成功但是json解析出错了", zap.String("login json parse error", err.Error()))
			} else {
				if result.Authenticated == false && result.User == true {
					result.Info = "用户存在但鉴权失败"
				} else if result.User == false {
					result.Info = "用户不存在"
				} else if result.Authenticated == true && result.User == true {
					result.Info = "登录成功"
					result.Cookies = c.Header["cookie"]
				} else {
					log.Println("未知情况")
				}
			}

			return result, err
		}
	}
	return response, err
}

func (c *Client) getFinalCookie(header1 map[string]string) {
	delete(header1, "origin")
	delete(header1, "content-type")
	delete(header1, "x-csrftoken")
	resp, err := c.Session.Get("https://www.instagram.com/accounts/onetap/?next=%2F", &grequests.RequestOptions{
		Headers:   header1,
		UserAgent: UserAgent,
	})
	if err != nil {
		log.Println("出错")
	} else {
		//生成登录后的Cookies
		c.Header["cookie"] = c.makeCookies(resp.Header["Set-Cookie"])
	}
}

func (c *Client) makeCookies(responseCk []string) string {
	var cookies = ""
	if c.Header["cookie"] != "" {
		cookies = c.Header["cookie"]
	}
	for _, ck := range responseCk {
		if strings.Contains(ck, "csrftoken=") && c.Header["cookie"] != "" {
			continue
		}
		if strings.Contains(ck, "ds_user_id") && strings.Contains(cookies, "ds_user_id") {
			continue
		}
		if strings.Contains(ck, "rur=") && len(ck) < 100 {
			continue
		}
		index := strings.Index(ck, "; ")
		cookies += ck[:index+2]
	}
	return cookies
}
