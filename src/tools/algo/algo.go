package algo

import (
	"math/rand"
	"time"
)

var min = int64(1)

func DoubleAverage(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	//算出最大可用金额
	max := amount - min*count
	//计算平均值
	av := max/count + min
	//计算2倍均值
	av = 2 * av
	//随机一个数
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(av) + min
	return x
}
