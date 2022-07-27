# InsTest 
#### 学习Instagram网页端逆向,仅限用作学习、交流，不可用作商业用途。
1. 使用golang实现登陆获取cookies、私信收发功能
2. 软件使用GIN提供REST服务，封装了登陆接口
> - http://127.0.0.1:8889/api/login
> - form-data
> - param1: username
> - param1: password
3. 使用paho.mqtt.golang对私信使用MQTT功能的实现，实现收、发
4. 实现http代理ip、mqtt协议代理ip
5. 登陆算法中的tweetnacl使用了cgo，需要安装对应环境