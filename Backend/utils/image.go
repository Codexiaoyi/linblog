package utils

import (
	"math/rand"
	"time"
)

func GetRandomImageUrl() string {
	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	//生成10个0-99之间的随机数
	randNum := rand.Intn(100)
	return "https://picsum.photos/200/200/?random=" + string(rune(randNum))
}
