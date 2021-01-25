package util

import (
	"testing"
)

func TestDecide(t *testing.T) {
	// 类型判断
	//r1 := isType("123", reflect.String)
	//r2 := isType(123, reflect.Int)
	//fmt.Println(r1, r2)

	// 字符串
	//var str1  = []string{"Jack", "Eric", "Alec"}
	//s1 := IsContainStr("Eric", str1)
	//fmt.Println(s1)

	// 整型
	//var int1 = []int{1,2,3}
	//s2 := IsContainInt(1, int1)
	//fmt.Println(s2)

	// 性能对比
	//i1 := make([]int, 1000000)
	//i2 := make([]int, 1000000)
	//for i:=0;i<100001;i++ {
	//	i1 = append(i1, i)
	//}
	//st := time.Now()
	//r11 := IsContainCapacity(100000, i1)
	//et := time.Since(st)
	//fmt.Println(et, r11)
	//
	//st = time.Now()
	//r11 = IsContainInt(100000, i1)
	//et = time.Since(st)
	//fmt.Println(et, r11)
	//
	//st = time.Now()
	//r12 := IsContainIntMap(i1)
	//et = time.Since(st)
	//fmt.Println(et, r12(100000))
	// 二分查找
	//var i3 = []int{3, 6, 20, 24, 59, 123}
	//fmt.Println(sort.SearchInts([]int{1, 2, 6, 8, 9, 11}, 6)) // 2
	//fmt.Println(sort.SearchInts([]int{1, 2, 6, 8, 9, 11}, 7)) // 3

	// map key
	//var i4 = []int{3, 6, 20, 24, 59, 123}

}
