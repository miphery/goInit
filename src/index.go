package main

import (
	S "./Sort"
	"./Tool"
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	numbers := []int{33, 22, 77, 44, 11, 99, 66, 55, 88}
	Tool.PrintSlice(numbers)
/*	//选择排序1
	t1 := time.Now()
	fmt.Print("选择排序：")
	numbers = S.Select1(numbers)
	Tool.PrintSlice(numbers)
	Tool.PrintTimeCost(t1)

	//选择排序2
	t1 = time.Now()
	fmt.Print("选择排序：")
	numbers = S.Select2(numbers)
	Tool.PrintSlice(numbers)
	Tool.PrintTimeCost(t1)

	// 冒泡排序
	t1 = time.Now()
	fmt.Print("冒泡排序：")
	numbers = S.Bubble(numbers)
	Tool.PrintSlice(numbers)
	Tool.PrintTimeCost(t1)

	//插入排序
	t1 = time.Now()
	fmt.Print("插入排序：")
	numbers = S.Insert(numbers)
	Tool.PrintSlice(numbers)
	Tool.PrintTimeCost(t1)*/

	//快速排序
	t1 := time.Now()
	fmt.Print("快速排序：")
	var numbersInterface []interface{}
	for i := 0; i < len(numbers); i++ {
		numbersInterface = append(numbersInterface, numbers[i] )
	}
	fmt.Printf("type =%v left:=%v\n",reflect.TypeOf(numbersInterface), numbersInterface)
	numbersInterface = S.Quick(numbersInterface)
	Tool.PrintInterfaceSlice(numbersInterface)
	Tool.PrintTimeCost(t1)
}
