package util

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func GetMsgId() string {
	curNano := time.Now().UnixNano()
	r := rand.New(rand.NewSource(curNano))
	return fmt.Sprintf("%d%06v", curNano, r.Int31n(1000000))
}

// 判断文件或路径是否存在 存在true
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}