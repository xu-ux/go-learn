package main

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
)

//
// AesEncryptECB
//  @Description: AES/ECB/PKCS5Padding方式，AES加密方法
//  @param content 明文
//  @param secret 密钥（需要做散列计算）
//  @return encryptData 密文（base64字符串）
//  @return err 错误
//
func AesEncryptECB(content string, secret string) (encryptData string, err error) {
	originData := []byte(content)
	// 密钥 16 24 32长度
	cipher, e := aes.NewCipher(Md5([]byte(secret)))
	if e != nil {
		return "", e
	}
	plain := PKCS5Padding(originData, aes.BlockSize)
	encrypt := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(originData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypt[bs:be], plain[bs:be])
	}
	// 密文做base64
	data := base64.StdEncoding.EncodeToString(encrypt)
	return data, nil
}

//
// AesDecryptECB
//  @Description:  AES/ECB/PKCS5Padding方式，AES解密方法
//  @param content 密文（base64字符串）
//  @param secret 密钥（需要做散列计算）
//  @return decryptData 明文
//  @return err 错误
//
func AesDecryptECB(content string, secret string) (decryptData string, err error) {
	// 密文base64解码
	encryptData, _ := base64.StdEncoding.DecodeString(content)
	cipher, e := aes.NewCipher(Md5([]byte(secret)))
	if e != nil {
		return "", e
	}
	decrypt := make([]byte, len(encryptData))
	// 分块解密
	for bs, be := 0, cipher.BlockSize(); bs < len(encryptData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypt[bs:be], encryptData[bs:be])
	}
	decrypt = PKCS5UnPadding(decrypt)
	return string(decrypt), nil
}

//
// PKCS5Padding
//  @Description: 对不足的块进行填充
//  @param ciphertext 内容
//  @param blockSize 块大小
//  @return []byte 填充后内容
//
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	// 需要填充的长度
	padding := blockSize - len(ciphertext)%blockSize
	// 填充的长度 = 填充的byte内容
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

//
// PKCS5UnPadding
//  @Description: 去除填充的字节
//  @param originData 原始内容
//  @return []byte 去除填充后的内容
//
func PKCS5UnPadding(originData []byte) []byte {
	length := len(originData)
	// 去掉最后一个字节 unPadding 次
	unPadding := int(originData[length-1])
	return originData[:(length - unPadding)]
}

// Md5 md5计算摘要
func Md5(data []byte) []byte {
	md5 := md5.New()
	md5.Write(data)
	md5Data := md5.Sum([]byte(nil))
	return md5Data
}

func main() {
	var secret string = "20220711testappsecret"
	encrypt, _ := AesEncryptECB("12345678", secret)
	println("加密密文：", encrypt)
	decrypt, _ := AesDecryptECB(encrypt, secret)
	println("解密明文：", decrypt)
}

func generateKey(key []byte) (genKey []byte) {
	return Md5(key)
}

func searchByteSliceIndex(bSrc []byte) int {
	var b byte = bSrc[len(bSrc)-1:][0]
	for i := 0; i < len(bSrc); i++ {
		if bSrc[i] == b {
			return i
		}
	}
	return len(bSrc)
}
