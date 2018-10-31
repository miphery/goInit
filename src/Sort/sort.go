package Sort

import (
	"../Tool"
)
/*
选择后，交换值
*/
func Select1(x []int) []int {
	len := len(x)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if x[i] > x[j] {
				tem := x[i]
				x[i] = x[j]
				x[j] = tem
			}
		}
	}
	return x
}

/*
选择后交换索引
*/
func Select2(x []int) []int {
	len := len(x)

	for i := 0; i < len-1; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			if x[j] < x[min] {
				min = j
			}
		}
		tmp := x[i]
		x[i] = x[min]
		x[min] = tmp
	}
	return x
}

/*
冒泡排序
*/
func Bubble(x []int) []int {
	len := len(x)
	for i := 0; i < len-1; i++ {
		canbreak := true
		for j := 0; j < len-1-i; j++ {
			if x[j] > x[j+1] {
				tmp := x[j]
				x[j+1] = x[j]
				x[j] = tmp
				canbreak = false
			}
		}
		if canbreak {
			break
		}
	}
	return x
}

/*
插入排序
*/

func Insert(x []int) []int {
	len := len(x)
	var preindex, current int
	for i := 1; i < len; i++ {
		preindex = i - 1
		current = x[i]
		for preindex >= 0 && x[preindex] > current {
			x[preindex+1] = current
			preindex--
		}
		x[preindex+1] = current
	}
	return x
}

/*
快速排序
*/
func Quick(x []interface{}) []interface{} {
	length := len(x)
	if length <= 1 {
		return x
	}
	current := x[0]
	var left, right []interface{}

	for i := 1; i < length; i++ {
		if Tool.ToStr(x[i]) < Tool.ToStr(current) {
			left = append(left, x[i])
		} else {
			right = append(right, x[i])
		}
	}
	left = append(left, current)
	//return Tool.Merge(Quick(left), Quick(right))
	left = Quick(left)
	right = Quick(right)
	return Tool.Merge( left, right)
	/*ll := len(Quick(left))
	var l []interface{}
	for li := 0;li < ll ; li++ {
		l = append(l, left[li])
	}
	rl := len(Quick(right))
	var r []interface{}
	for ri := 0;ri < rl ; ri++ {
		r = append(r, right[ri])
	}
	res := c.Merge( l, r)


	resl := len(res)
	var result []int
	for resi := 0;resi < resl ; resi++  {
		result = append(result, res[resi].(int))
	}
	return result*/
}
