package tests

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	p, _ := rsa.GenerateKey(rand.Reader, 888)

	derStream := x509.MarshalPKCS1PrivateKey(p)

	block := &pem.Block{
		Type:  "SAM PRIVATE KEY",
		Bytes: derStream,
	}
	b := pem.EncodeToMemory(block)
	fmt.Println(string(b))

	der, err := x509.MarshalPKIXPublicKey(&p.PublicKey)
	if err != nil {
		fmt.Println(err)
	}
	block = &pem.Block{
		Type:  "SAM PUBLIC KEY",
		Bytes: der,
	}

	pu := pem.EncodeToMemory(block)
	fmt.Println(string(pu))
}

func TestEncrypt(t *testing.T) {

}

func TestDecrypt(t *testing.T) {

}

//
//// 加密
//func RsaEncrypt(origData []byte) ([]byte, error) {
//	block, _ := pem.Decode(publicKey)
//	if block == nil {
//		return nil, errors.New("public key error")
//	}
//	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
//	if err != nil {
//		return nil, err
//	}
//	pub := pubInterface.(*rsa.PublicKey)
//	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
//}
//
//// 解密
//func RsaDecrypt(ciphertext []byte) ([]byte, error) {
//	block, _ := pem.Decode(privateKey)
//	if block == nil {
//		return nil, errors.New("private key error!")
//	}
//	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
//	if err != nil {
//		return nil, err
//	}
//	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
//
//}