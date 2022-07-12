package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"testing"
)

func Sign(content, prvKey []byte) (sign string, err error) {
	block, _ := pem.Decode(prvKey)
	if block == nil {
		fmt.Println("pem.Decode err")
		return
	}
	var private interface{}
	private, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	privateKey := private.(*rsa.PrivateKey)
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(content))
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey,
		crypto.SHA1, hashed)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(signature)
	return
}

func RSAVerify(origdata, ciphertext string, publicKey []byte) (bool, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return false, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(origdata))
	digest := h.Sum(nil)
	body, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA1, digest, body)
	if err != nil {
		return false, err
	}
	return true, nil
}

func TestSign3(t *testing.T) {
	pubKey := "-----BEGIN PUBLIC KEY-----\n{公钥内容}\n-----END PUBLIC KEY-----"
	prvKey := "-----BEGIN RSA PRIVATE KEY-----\n{公钥内容}\n-----END RSA PRIVATE KEY-----"
	content := "模型训练需要花多长时间"

	sign, err := Sign([]byte(content), []byte(prvKey))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("sign签名结果：", sign)

	res, err := RSAVerify(content, sign, []byte(pubKey))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("验签结果：", res)

}

// ConvertPKCS PKCS相互转换
func ConvertPKCS(t *testing.T) {
	prvKey := "-----BEGIN RSA PRIVATE KEY-----\nMIICdwIBADANBgkqhkiG9w0........KwWPPXMiYR8=\n-----END RSA PRIVATE KEY-----"
	block, _ := pem.Decode([]byte(prvKey))
	if block == nil {
		fmt.Println("pem.Decode err")
		return
	}
	var private interface{}
	private, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	privateKey := private.(*rsa.PrivateKey)
	pkcs1 := x509.MarshalPKCS1PrivateKey(privateKey)
	t.Log("pkcs8 转 pkcs1：", base64.StdEncoding.EncodeToString(pkcs1))

	prvKey = fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\n%s\n-----END RSA PRIVATE KEY-----", pkcs1)
	pkcs8 := x509.MarshalPKCS1PrivateKey(privateKey)
	t.Log("pkcs1 转 pkcs8：", base64.StdEncoding.EncodeToString(pkcs8))
}
