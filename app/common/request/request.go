package request

//LoginData 登录请求参数
type LoginData struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

//MessageData 发送私信所需参数
type MessageData struct {
	//私信内容

}
