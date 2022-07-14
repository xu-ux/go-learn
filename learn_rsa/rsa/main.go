package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

const (
	PEM_BEGIN         = "-----BEGIN PRIVATE KEY-----\n"
	PEM_END           = "\n-----END PRIVATE KEY-----"
	PUBLIC_BEGIN      = "-----BEGIN PUBLIC KEY-----\n"
	PUBLIC_END        = "\n-----END PUBLIC KEY-----"
	KEY_SIZE          = 2048
	MAX_ENCRYPT_BLOCK = KEY_SIZE/8 - 11
	MAX_DECRYPT_BLOCK = KEY_SIZE / 8
)

//
// RsaEncrypt
//  @Description: 公钥加密
//  @param data 原始明文
//  @param publicKeyBase64 公钥（base64格式）
//  @return string 密文（base64格式）
//  @return error 错误
//
func RsaEncrypt(data, publicKeyBase64 string) (string, error) {
	//解密pem格式的公钥
	publicKeyPem := FormatPublicKey(publicKeyBase64)
	block, _ := pem.Decode([]byte(publicKeyPem))
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	originData := []byte(data)
	tmp := make([]byte, 0)
	if len(originData) < MAX_ENCRYPT_BLOCK {
		ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, originData[:len(originData)])
		if err != nil {
			return "", err
		}
		tmp = append(tmp, ciphertext...)
		// 分段加密
	} else {
		for i := 0; i < len(originData); i += MAX_ENCRYPT_BLOCK {
			ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, originData[i:i+MAX_ENCRYPT_BLOCK])
			if err != nil {
				return "", err
			}
			tmp = append(tmp, ciphertext...)
		}

	}
	encryptData := base64.StdEncoding.EncodeToString(tmp)
	return encryptData, nil
}

//
// RsaDecrypt
//  @Description: 私钥解密
//  @param encryptData 密文（base64格式）
//  @param privateKeyBase64 私钥（base64格式）
//  @return string 原始明文
//  @return error 错误
//
func RsaDecrypt(encryptData, privateKeyBase64 string) (string, error) {
	ciphertext, _ := base64.StdEncoding.DecodeString(encryptData)
	//获取私钥
	privateKeyPem := FormatPrivateKey(privateKeyBase64)
	block, _ := pem.Decode([]byte(privateKeyPem))
	if block == nil {
		return "", errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	private, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	tmp := make([]byte, 0)
	// 分段解密
	for i := 0; i < len(ciphertext); i += MAX_DECRYPT_BLOCK {
		data, err := rsa.DecryptPKCS1v15(rand.Reader, private.(*rsa.PrivateKey), ciphertext[i:(i+MAX_DECRYPT_BLOCK)])
		if err != nil {
			return "", err
		}
		tmp = append(tmp, data...)
	}
	return string(tmp), nil
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

func main() {
	privateKey := "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCfCEjKlE8mxAeKnG61B2rYmtaR35EIjUgd/EpZrKzfCDofVi3KAwQ4KI0HbzeUmHVPBNqs/lqnSdCPC2OuSprwlFj+8EsnfHe1NeXXABCvXiD4WTifpzDktjknNUwiM+iXVe421upWFzD7bmsPmLDVovvZKUi4jozF18AbTmxULM0HEuM9CFCIp+1j64esDwN+tuZXQnWJvcc4SZL6VmEgHCBaIDoWASUd1eflJDjWkv7Q2Uw1rT+9AtpDiy2jYUv7zYtGGQj6yuHXAc/UGCo//B4hWY85tiHJidwshCs2zMCFnVy8KRPSecc7JGJiQIiz6tSS2LaO82w533KAgZAjAgMBAAECggEABFnUq/4swDHCxw8KlFdUnAJ4dls4e5Rp8bASVKu9uwhdlrfj3tAAUI0Ddr2bNqyJIRVP6kW3MzZ4x0EXhBA0dvqICMmINNdZ6xJDbINq8XFQ05qVSwDm/Irju2fg4lqrNWC7eLKejKZrx6U4tp+FzwJ7g3B2td3oig0iC10054SbEChpKWDq/XAgrIrjLZ/bgxDFVVUWVzXAFL4akQPd2astIFDi/jRfauyTWgIsfX/f2morHyuSuG4jK4WRD/m4PNDp2/Rm7yz+dNSpP0MkAlMLuSh9wx9kfuqcfnfcGiiWeJhQoY5Yduw6rGArx9kWl3nCGhmJeuX1k9stED2SkQKBgQDvHC6Nbt1tE+O+gVk6o8GVQ2ekUfkjLVsrSTQqHx5Maf25pz1VB28vMWWEfVFClQWbBIJp7yZbZN9Mw5gsrjMrrOh8++gdHjmRC7ey4b3y2tXIpsGzlVLwabPmsXAsmHCLPYE8Mrh8jxfqy/WttUiBIT13Ndmn2RcObtt6tzHIawKBgQCqRA9oig0jXp55UbcI2sBY03hSI8fehqJjm/X0VTtiTlDUtb1xZOX3vMCXAfLfdVQV8XVclmISLOXJg9eBzvJqML9EIr9JTF9lYQvZ2irrFeaBbNseFrrpp+xQ/WInvGNB0euxeXO6ojrww98oCYHjzE8xH141Di1lYvdFrvMlKQKBgQDsx1mKENkQZPvH8MrteLAAIWmGnO47WXTInosbkwkr3mG08NmZU+1ULHQ9COPpLS0J3yNNx9aR9ofxulb9F9vwSh9HdSTbgMy8x3+3kjfJP88oDYoPTbV+AQ53Sgqs/p+kItnRRODP59tlVWgKBlSwGryFSjwpLJ7aWgjZsoOH1QKBgC4as9lo2Fnlex/6wodBRKhIyuHjEnHtHve9+YGpuqTJ9BVFCQE1gxfsInJBctSTXqt6cH8bsX6ebbJ9YtOhh/69KG14wzdD2OkIuD7LVqfFjF8rbMHfAcnXUKQ1mGiOGIpwH1Q1QOMenrsnLrwWpvdaEW+JwOa46g30GGTkFK7RAoGAY9kVgkNtOM4Hc5qQMFy9Lp3PL5yZ+DI4+EB3kRv1TmHx7+IGZq88abHG9BvxAbpXKPlHVV6xlqbM4jj72KCm5xaJQ32m2z6rbTXHfaSSLYRVJESPgO9bfNeNgbwOZcLsae4YFgtgIHDKgcnzVUcQAr7DBjPRMhd+Q9h8vXBQ1po="
	publicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnwhIypRPJsQHipxutQdq2JrWkd+RCI1IHfxKWays3wg6H1YtygMEOCiNB283lJh1TwTarP5ap0nQjwtjrkqa8JRY/vBLJ3x3tTXl1wAQr14g+Fk4n6cw5LY5JzVMIjPol1XuNtbqVhcw+25rD5iw1aL72SlIuI6MxdfAG05sVCzNBxLjPQhQiKftY+uHrA8DfrbmV0J1ib3HOEmS+lZhIBwgWiA6FgElHdXn5SQ41pL+0NlMNa0/vQLaQ4sto2FL+82LRhkI+srh1wHP1BgqP/weIVmPObYhyYncLIQrNszAhZ1cvCkT0nnHOyRiYkCIs+rUkti2jvNsOd9ygIGQIwIDAQAB"
	originData := `待签值：12345678
待加密值：12345678
AppSecret：20220711testappsecret
`

	encryptData, e := RsaEncrypt(originData, publicKey)
	if e != nil {
		panic(e)
	}
	println(encryptData)
	data, e := RsaDecrypt(encryptData, privateKey)
	if e != nil {
		panic(e)
	}
	println(data)

	testEncryptData := "L3FEz6AKyu1F8NHUcDNuuPPKv4dAMg+SNNRAMzjgGAuqVc4s/q3IRcqG7+70L0tvTbL9giNARyuimgmAKZuuuAFPnU9/URkjelc+jQ4Wu+bLskq9pAvCz8aw8G+VGyJZjL6neWbVs1NpzoyrZ3VJlcQ2t6fh9DC5Cj5nZNCVwZG/k0UDMjKoU2CGQgolhVqw342EUlBs6w/8ppt9emYxoeOVBrukcxPOk+HCGp70QK8gEWLAmyVZAoVDeBKysuztIn7PAenNihEtYhwWyksuiqe6jg7X8m5+FhnDLcQrMIx92dBUHTWD9J2XZTnjUFELauC+3EvHu6/UnB49r0wLiA=="
	testData, e := RsaDecrypt(testEncryptData, privateKey)
	if e != nil {
		panic(e)
	}
	println(testData)
}
