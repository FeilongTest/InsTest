package utils

import (
	"github.com/dchest/blake2b"
	"github.com/twystd/tweetnacl-go/tweetnacl"
)

func boxSeal(message, pKey []byte) []byte {
	var ret = make([]byte, 0)
	//create ephemeral keypair for sender 创建临时密钥对
	ephKeyPair, err := tweetnacl.CryptoBoxKeyPair()
	if err != nil {
		panic("创建临时密钥对出错")
	}
	//控制固定值变量
	//ephKeyPair.PublicKey = make([]byte, 32)
	//ephKeyPair.SecretKey = make([]byte, 32)

	ret = append(ret, ephKeyPair.PublicKey...)
	//创建nonce
	nonce := generateNonce(ephKeyPair.PublicKey, pKey)
	//创建box
	box, err := tweetnacl.CryptoBox(message, nonce, pKey, ephKeyPair.SecretKey)
	if err != nil {
		panic("创建盒子出错")
	}
	ret = append(ret[:len(ephKeyPair.PublicKey)], box...)
	return ret
}

func generateNonce(ephPubKey, pubKey []byte) []byte {
	blake := blake2b.NewMAC(24, nil)
	blake.Write(ephPubKey)
	blake.Write(pubKey)
	sum := blake.Sum(nil)
	return sum
}
