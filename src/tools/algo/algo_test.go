package algo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDoubleAverage(t *testing.T) {
	Convey("2倍均值算法", t, func() {
		count, amount := int64(10), int64(10000)
		remain := amount
		sum := int64(0)
		for i := int64(0); i < count; i++ {
			x := DoubleAverage(count-i, remain)
			remain -= x
			sum += x
		}
		So(sum, ShouldEqual, amount)
	})
}
