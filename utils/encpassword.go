package utils

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
)

func Encrypt(key string, pkey string, password []byte, time []byte) (result string, err error) {
	var pLength = 100 + len(password)
	if 64 != len(pkey) {
		err = fmt.Errorf("%s", "Public key is not valid hex string")
	}

	//parsePublicKey
	pkeyByte := parsePkey(pkey)

	//log.Println(pLength, pkeyByte)

	//转换字符类型key为unit8
	keyInt, _ := strconv.Atoi(key)

	var y = make([]byte, pLength)
	f := 0
	y[f] = 1
	f += 1
	y[f] = byte(keyInt)
	f += 1

	//Generate Key
	//生成随机的key 32位对应Aes-gcm-256 16位对应Aes-gcm-128
	aesKey := make([]byte, 32)
	rand.Read(aesKey)
	//iv = 12位空[]byte
	iv := make([]byte, 12)

	//additionalData = time的[]byte
	additionalData := time

	cipherText, err := aesGmc(iv, aesKey, additionalData, password)
	if err != nil {
		panic("aes加密出错" + err.Error())
	}

	sealed := boxSeal(aesKey, pkeyByte)
	y[f] = byte(255 & len(sealed))
	y[f+1] = byte(len(sealed) >> 8 & 255)
	f += 2
	y = append(y[:f], sealed...)
	f += 32
	f += 48
	if len(sealed) != 32+48 {
	}

	var s = make([]byte, len(cipherText))
	copy(s, cipherText)
	c := s[len(s)-16:]
	h := s[:len(s)-16]
	y = append(y[:f], c...)
	f += 16
	y = append(y[:f], h...)

	appVersion := 10
	result = fmt.Sprintf("%s:%d:%s:%s", "#PWD_INSTAGRAM_BROWSER", appVersion, string(time), base64.StdEncoding.EncodeToString(y))
	return
}

/*
* parsePkey 字符串转16进制字节数组
* Parse publicKey
 */
func parsePkey(pKey string) []byte {
	t, _ := hex.DecodeString(pKey)
	return t
}
