package util

import (
	"math/rand"
	"time"
)

// 随机分配字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiop")
	result := make([]byte, n)

	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
