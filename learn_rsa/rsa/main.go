package main

import (
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"log"
	"strings"
)

const (
	PEM_BEGIN    = "-----BEGIN PRIVATE KEY-----\n"
	PEM_END      = "\n-----END PRIVATE KEY-----"
	PUBLIC_BEGIN = "-----BEGIN PUBLIC KEY-----\n"
	PUBLIC_END   = "\n-----END PUBLIC KEY-----"
)

// HmacMD5
//  @Description: HmacMD5方式计算签名
//  @param key 密钥，一般是appSecret
//  @param data 待签字符串
//  @return string 签名
//
func HmacMD5(key, data string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// HmacSHA256
//  @Description: HmacSHA256方式计算签名
//  @param key 密钥，一般是appSecret
//  @param data 待签字符串
//  @return string 签名
//
func HmacSHA256(key, data string) string {
	keys := []byte(key)
	h := hmac.New(sha256.New, keys)
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// SignSHA256withRSA
//  @Description: SignSHA256withRSA私钥加签 PKCS8格式
//  @param data 待签字符串
//  @param privateKeyBase64 私钥，PKCS8 Bas64格式
//  @return string 签名
//
func SignSHA256withRSA(data, privateKeyBase64 string) string {
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	privateKeyPem := FormatPrivateKey(privateKeyBase64)
	// 解析PEM格式的公钥
	block, _ := pem.Decode([]byte(privateKeyPem))
	if block == nil {
		panic(errors.New("private key error"))
	}

	private, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, private.(*rsa.PrivateKey), crypto.SHA256, hashed)
	if err != nil {
		panic(err)
	}
	// 签名做Bas64加密
	return base64.StdEncoding.EncodeToString(signature)
}

// VerySHA256withRSA
//  @Description: VerySHA256withRSA公钥验签
//  @param data 待签字符串
//  @param signature 传输的签名值
//  @param publicKeyBase64 公钥，Bas64格式
//  @return bool 比对计算签名值结果，一致返回true，错误不一致返回false
//
func VerySHA256withRSA(data, signature, publicKeyBase64 string) bool {
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	publicKeyPem := FormatPublicKey(publicKeyBase64)
	// 解析PEM格式的公钥
	block, _ := pem.Decode([]byte(publicKeyPem))
	if block == nil {
		panic(errors.New("public key error"))
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// 将base64的签名解码
	sign, _ := base64.StdEncoding.DecodeString(signature)

	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed, sign)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// FormatPrivateKey
//  @Description: FormatPrivateKey 格式化私钥
//  @param privateKey
//  @return string PEM文件格式
//
func FormatPrivateKey(privateKey string) string {

	if !strings.HasPrefix(privateKey, PEM_BEGIN) {

		privateKey = PEM_BEGIN + privateKey
	}
	if !strings.HasSuffix(privateKey, PEM_END) {

		privateKey = privateKey + PEM_END
	}
	return privateKey
}

// FormatPublicKey
//  @Description: FormatPublicKey 格式化公钥
//  @param publicKey
//  @return string PEM文件格式
//
func FormatPublicKey(publicKey string) string {

	if !strings.HasPrefix(publicKey, PUBLIC_BEGIN) {

		publicKey = PUBLIC_BEGIN + publicKey
	}
	if !strings.HasSuffix(publicKey, PUBLIC_END) {

		publicKey = publicKey + PUBLIC_END
	}
	return publicKey
}
