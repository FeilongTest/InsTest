package utils

import (
	"encoding/base64"
	"github.com/twystd/tweetnacl-go/tweetnacl"
	"log"
	"testing"
)

func TestEncPassWord(t *testing.T) {
	pkey := "e126828bf03d60e43711e99254607132dc50a8351ed1cb113317fa6eb3f6fd2f"
	parsePkey(pkey)

	//时间戳
	time := "1656852053"
	log.Println([]byte(time))

}

func TestEncPassWord3(t *testing.T) {
	time := []byte("1656852053")
	password := []byte("789789")

	//生成随机的key 32位对应Aes-gcm-256 16位对应Aes-gcm-128
	aesKey := make([]byte, 32)
	//rand.Read(aesKey)

	log.Println(aesKey)
	iv := make([]byte, 12)
	encrypted, err := aesGmc(iv, aesKey, time, password)
	log.Println(base64.StdEncoding.EncodeToString(encrypted), err)

	ephKeyPair, err := tweetnacl.CryptoBoxKeyPair()
	if err != nil {
		panic("创建临时密钥对出错")
	}
	log.Println(ephKeyPair.PublicKey)

}

func TestEncPassWord4(t *testing.T) {

	a := make([]byte, 32)
	//b := make([]byte, 24)
	pkey := "e783d71cecf6ea55945d264c5caaa24807cf175ee43a277a6fe213f93144b853"
	pKeyByte := parsePkey(pkey)
	//创建nonce
	nonce := generateNonce(a, a)

	ret, err := tweetnacl.CryptoBox(a, nonce, pKeyByte, a)
	log.Println(ret, err)

}
