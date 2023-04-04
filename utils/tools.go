package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

func GetNowFormatTodayTime() string {

	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()),
		now.Day())

	return dateStr
}

func HashString(s string) string {

	// 计算SHA256哈希值
	hash := sha256.Sum256([]byte(s))

	// 对哈希值进行base64编码
	encoded := base64.StdEncoding.EncodeToString(hash[:])

	fmt.Println("Message:", s)
	fmt.Println("Hash:", encoded)
	return encoded
}
