package ins

import (
	"InsTest/app/model"
	"InsTest/utils"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/levigross/grequests"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

//此处手动填写cookies
const cookies = ``
const xCsrfToken = ``
const xIgAppId = ``

// Ms 私信
type Ms struct {
	MqttClient mqtt.Client
}

func Listen() {

	sessionId := rand.Uint64()
	id, _ := uuid.NewRandom()
	clientId := id.String()

	h := http.Header{}
	h.Add("Cookie", cookies)
	h.Add("User-Agent", UserAgent)
	h.Add("Accept", "*/*")
	h.Add("Referer", "https://www.instagram.com")
	h.Add("Host", "edge-chat.instagram.com")
	h.Add("Origin", "https://www.instagram.com")

	wsOpts := mqtt.WebsocketOptions{}
	connOpts := mqtt.
		NewClientOptions().
		SetWebsocketOptions(&wsOpts).
		AddBroker("wss://edge-chat.instagram.com/chat").
		SetClientID("mqttwsclient").
		SetUsername(fmt.Sprintf(`{"u": "%s","s": %d,"cp": 1,"ecp": 0,"chat_on": true,"fg": true,"d": "%s","ct": "cookie_auth","mqtt_sid": "","aid": 936619743392459,"st": [],"pm": [],"dc": "","no_auto_fg": true,"asi":{"Accept-Language":"zh-cn"},"a":"%s"}`,
			"00000000000", sessionId, clientId, UserAgent)). //此处a对应登陆获取到的userid
		SetProtocolVersion(3).
		SetHTTPHeaders(h).
		SetAutoReconnect(true).
		SetResumeSubs(true).
		SetCleanSession(true).
		SetKeepAlive(30 * time.Second)
	connOpts.OnConnect = connectHandler
	connOpts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln(fmt.Sprintf("cannot mqtt connect: %s\n", token.Error()))
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(30 * time.Second)
	}()
	wg.Wait()
	client.Disconnect(250)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("接收到来自: %s 的消息: %s\n", msg.Topic(), msg.Payload())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("已连接,开始处理连接")
	for _, topic := range []string{
		"/ig_message_sync",
		"/ig_send_message_response",
		"/pubsub",
	} {
		if token := client.Subscribe(topic, byte(0), messagePubHandler); token.Wait() && token.Error() != nil {
			log.Println(fmt.Sprintf("cannot subscribe on mqtt topic %s: %s\n", topic, token.Error()))
		}
	}
	reSub(client, "/ig_sub_iris_response")
	publishSeqId(client) //获取并发送当前Seq_id 私信序号 第6条
	pubUnSub(client)
	reSub(client, "/pubsub")
	pubSub(client) //第10条

	//获取实时消息的4条区别在于
	id, _ := uuid.NewRandom()
	clientId1 := id.String()
	id, _ = uuid.NewRandom()
	clientId2 := id.String()

	realtimeUnSub(client, clientId1, clientId2)
	reSub(client, "/ig_realtime_sub")
	realtimeSub(client, clientId1, clientId2)
	for _, topic := range []string{
		"/ig_message_sync",
		"/ig_send_message_response",
		"/pubsub",
		"/ig_sub_iris_response",
		"/ig_realtime_sub",
	} {
		if token := client.Subscribe(topic, byte(1), messagePubHandler); token.Wait() && token.Error() != nil {
			log.Println(fmt.Sprintf("cannot subscribe on mqtt topic %s: %s\n", topic, token.Error()))
		}
	}
	log.Println("开始发送私信")
	sendText(client)
}

func reSub(client mqtt.Client, topic string) {
	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		log.Println(fmt.Sprintf("cannot subscribe on mqtt topic %s: %s\n", topic, token.Error()))
	}
	if token := client.Subscribe(topic, byte(1), messagePubHandler); token.Wait() && token.Error() != nil {
		log.Println(fmt.Sprintf("cannot subscribe on mqtt topic %s: %s\n", topic, token.Error()))
	}
}

func publishSeqId(client mqtt.Client) {
	seqId, err := fetchSeqId()
	if err != nil {
		panic("获取SeqId出错" + err.Error())
	}

	token := client.Publish("/ig_sub_iris", byte(1), false, fmt.Sprintf(`{"seq_id":%s,"snapshot_at_ms":%d,"snapshot_app_version":"web","subscription_type":"message"}`, seqId, time.Now().UnixNano()/1e6))

	token.Wait()
	if token.Error() != nil {
		log.Println(fmt.Sprintf("cannot publish messenger_sync_create_queue: %s", token.Error()))
	}
}

func pubUnSub(client mqtt.Client) {
	token := client.Publish("/pubsub", 1, false, fmt.Sprintf(`{"unsub":["ig/u/v1/54082540125"]}`))
	token.Wait()
	if token.Error() != nil {
		log.Println(fmt.Sprintf("cannot publish messenger_sync_create_queue: %s", token.Error()))
	}
}

func pubSub(client mqtt.Client) {
	token := client.Publish("/pubsub", 1, false, fmt.Sprintf(`{"sub":["ig/u/v1/54082540125"]}`))
	token.Wait()
	if token.Error() != nil {
		log.Println(fmt.Sprintf("cannot publish messenger_sync_create_queue: %s", token.Error()))
	}
}

func realtimeUnSub(client mqtt.Client, id1, id2 string) {
	token := client.Publish("/ig_realtime_sub", 1, false, fmt.Sprintf(`{"unsub":["1/graphqlsubscriptions/17935802131473058/{\"input_data\":{\"client_subscription_id\":\"%s\"}}","1/graphqlsubscriptions/17846944882223835/{\"input_data\":{\"client_subscription_id\":\"%s\"}}"]}`, id1, id2))
	token.Wait()
	if token.Error() != nil {
		log.Println(fmt.Sprintf("cannot publish messenger_sync_create_queue: %s", token.Error()))
	}
}

func realtimeSub(client mqtt.Client, id1, id2 string) {
	token := client.Publish("/ig_realtime_sub", 1, false, fmt.Sprintf(`{"sub":["1/graphqlsubscriptions/17935802131473058/{\"input_data\":{\"client_subscription_id\":\"%s\"}}","1/graphqlsubscriptions/17846944882223835/{\"input_data\":{\"client_subscription_id\":\"%s\"}}"]}`, id1, id2))
	token.Wait()
	if token.Error() != nil {
		log.Println(fmt.Sprintf("cannot publish messenger_sync_create_queue: %s", token.Error()))
	}
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("连接已断开: %v\n", err)
}

//同步seqId
func fetchSeqId() (string, error) {
	resp, err := grequests.Get("https://i.instagram.com/api/v1/direct_v2/inbox/?persistentBadging=true&folder=&limit=10&thread_message_limit=10", &grequests.RequestOptions{
		Headers: map[string]string{
			"cookie":      cookies,
			"x-csrftoken": xCsrfToken,
			"x-ig-app-id": xIgAppId,
		},
		UserAgent: UserAgent,
	})
	if err != nil {
		log.Println("访问出错")
		return "", err
	} else {
		var response model.InboxResponse
		err = jsoniter.Unmarshal(resp.Bytes(), &response)
		if err != nil {
			log.Println("fetchSeqId json解析出错", err)
		} else {
			for i, thread := range response.Inbox.Threads {
				log.Println("thread_id为", i, thread.ThreadId)
			}
		}
		return jsoniter.Get(resp.Bytes(), "seq_id").ToString(), nil
	}
}

func generateClientId() int64 {
	var a = time.Now().UnixNano() / 1e6
	var b = utils.GenerateRangeNum(0, 4294967295)
	b2 := "0000000000000000000000" + strconv.FormatInt(int64(b), 2)
	b2 = b2[len(b2)-22:] //slice(-22);
	a2 := strconv.FormatInt(a, 2) + b2
	p, _ := strconv.ParseInt(a2[len(a2)-63:], 2, 64) // 2进制转10进制
	return p
}

//私信发送文本
func sendText(client mqtt.Client) {
	//发送好友的threadId
	threadId := "34028236684xxxxxxxxxxxxxxxxxx3064073114"
	id := "680C5652-XXXX-XXXX-XXXX-F28E99B1D98B"
	clientId := generateClientId()
	//开始发送激活窗口
	token := client.Publish("/ig_send_message", 1, false, fmt.Sprintf(`{"client_context":"%d","device_id":"%s","action":"indicate_activity","thread_id":"%s","activity_status":0}`, clientId, id, threadId))
	token.Wait()
	if token.Error() != nil {
		log.Println(fmt.Sprintf("cannot publish messenger_sync_create_queue: %s", token.Error()))
	}
	//发送内容
	time.Sleep(500 * time.Millisecond)
	token = client.Publish("/ig_send_message", 1, false, fmt.Sprintf(`{"client_context":"%d","device_id":"%s","action":"send_item","item_type":"text","mutation_token":"%d","text":"测试内容","thread_id":"%s"}`, clientId, id, generateClientId(), threadId))
	token.Wait()
	if token.Error() != nil {
		log.Println(fmt.Sprintf("cannot publish messenger_sync_create_queue: %s", token.Error()))
	}
	//发送完成
	token = client.Publish("/ig_send_message", 1, false, fmt.Sprintf(`{"client_context":"%d","device_id":"%s","action":"indicate_activity","thread_id":"%s","activity_status":1}`, clientId, id, threadId))
	token.Wait()
	if token.Error() != nil {
		log.Println(fmt.Sprintf("cannot publish messenger_sync_create_queue: %s", token.Error()))
	}
}

//私信发送图片
func sendImg(client mqtt.Client) {

}
