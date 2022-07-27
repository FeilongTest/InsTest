package model

type LoginData struct {
	KeyId     string
	PublicKey string
	Time      string
}

type ShareDataResp struct {
	Config struct {
		CsrfToken string `json:"csrf_token"`
	} `json:"config"`
	Encryption struct {
		KeyId     string `json:"key_id"`
		PublicKey string `json:"public_key"`
		Version   string `json:"version"`
	} `json:"encryption"`
}

type LoginResponse struct {
	Info          string `json:"info,omitempty"`
	User          bool   `json:"user,omitempty"`
	UserId        string `json:"userId,omitempty"`
	Authenticated bool   `json:"authenticated,omitempty"`
	OneTapPrompt  bool   `json:"oneTapPrompt,omitempty"`
	Status        string `json:"status,omitempty"`
	Cookies       string `json:"cookies,omitempty"`
}
