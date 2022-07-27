package utils

import (
	"crypto/aes"
	"crypto/cipher"
)

//参考自：https://blog.51cto.com/u_15069450/4256813
//参考：https://www.golangprograms.com/data-encryption-with-aes-gcm.html
func aesGmc(iv, key, additionalData, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //生成加解密用的block
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nil, iv, plaintext, additionalData) //加密,密文为:iv+密文+tag
	return ciphertext, nil
}
