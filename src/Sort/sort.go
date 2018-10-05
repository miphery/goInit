package Sort

import "../Tool"
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
		if canbreak  {
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
func Quick(x [] int) [] int{
	current := x[0]
	len := len(x);
	if(len <= 1) {
		return x;

	}
	var left, right [] int
	for i:=1; i < len;i++  {
		if(x[i] < current){
			left = append(left, x[i])
		}else{
			right = append(right, x[i])
		}
	}
	left = append(left, current)
	return Tool.Merge(Quick(left), Quick(right));
c := Tool.CommonFunc{}
c.Merge(Quick(left), Quick(right))
}