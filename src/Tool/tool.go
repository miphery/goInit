package Tool

import (
	"fmt"
	"reflect"
	"time"
)

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func PrintInterfaceSlice(x []interface{}) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func PrintTimeCost(start time.Time) {
	end := time.Since(start)
	fmt.Println(end)
}

/*func Merge(s ...[]int) (slice []int) {
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
}*/



func Merge(s ...[]interface{}) (slice []interface{}) {
	switch len(s) {
	case 0:
		break
	case 1:
		slice = s[0]
		break
	default:
		s1 := s[0]
		s2 := Merge(s[1:]...) //...将数组元素打散
		slice = make([]interface{}, len(s1)+len(s2))
		copy(slice, s1)
		copy(slice[len(s1):], s2)
		break
	}

	return
}

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

func ToStr(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
