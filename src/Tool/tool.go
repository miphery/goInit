package Tool

import (
	"fmt"
	"time"
)

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func PrintTimeCost(start time.Time) {
	end := time.Since(start)
	fmt.Println(end)
}

func Merge(s ...[]int) (slice []int) {
	switch len(s) {
	case 0:
		break
	case 1:
		slice = s[0]
		break
	default:
		s1 := s[0]
		s2 := Merge(s[1:]...) //...将数组元素打散
		slice = make([]int, len(s1)+len(s2))
		copy(slice, s1)
		copy(slice[len(s1):], s2)
		break
	}
	return
}

type CommonFunc struct{}

var commonFunc CommonFunc

func (c *CommonFunc) Merge(s ...[]interface{}) (slice []interface{}) {
	switch len(s) {
	case 0:
		break
	case 1:
		slice = s[0]
		break
	default:
		s1 := s[0]
		s2 := commonFunc.Merge(s[1:]...) //...将数组元素打散
		slice = make([]interface{}, len(s1)+len(s2))
		copy(slice, s1)
		copy(slice[len(s1):], s2)
		break
	}

	return
}
