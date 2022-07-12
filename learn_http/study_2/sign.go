package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// BuildWaitSignStr
// 构建待签字符串
//  @Description: 构建待签字符串
//  @param method 请求方式，大写
//  @param path 请求路径
//  @param appId 开发者应用id
//  @param timestamp 时间戳，单位秒
//  @param nonce 随机字符串
//  @param queryRaw 原始query参数
//  @param bodyRaw 原始body
//  @return string 待签字符串
//
func BuildWaitSignStr(method, path, appId, timestamp, nonce, queryRaw, bodyRaw string) string {
	arr := []string{method, path, appId, timestamp, nonce, queryRaw, bodyRaw}
	arr = filter(arr, func(s string) bool {
		if len(strings.TrimSpace(s)) <= 0 {
			return false
		}
		return true
	})
	return strings.Join(arr, "&")
}

// GetUnixTime
// 获取unix时间戳，单位秒
//  @Description:  获取unix时间戳，单位秒
//  @return string 字符串格式时间戳
//
func GetUnixTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// RandomString
// 随机字符串
//  @Description: 随机字符串
//  @param n
//  @return string
//
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//
//  filter 函数式编程-过滤器
//  @Description: 函数式编程-过滤器
//  @param arr 原始切片
//  @param function 比较方法
//  @return []string 返回切片
//
func filter(arr []string, function func(s string) bool) []string {
	var newArr []string
	for _, str := range arr {
		if function(str) == true {
			newArr = append(newArr, str)
		}
	}
	return newArr
}
