package utils

import (
	"math/rand"
	"time"
)

//GenerateRangeNum 生成范围随机数
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}
