package util

import "reflect"

// 变量的判断类型（data 数据、t 类型）
// 支持判断类型 reflect.Kind
func isType(data interface{}, t reflect.Kind) bool {
	val := reflect.ValueOf(data)
	// 判断数据是否符合传入的类型
	if val.Kind() == t {
		return true
	}
	return false
}

// 基于反射来支持所有元素和数组类型（性能损耗）
func IsContainCapacity(value interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i <= s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

/*
	判断一个元素是否在一个切片中
	性能：快
*/

// 字符串
func IsContainStr(value string, array []string) bool {
	for _, eachItem := range array {
		if value == eachItem {
			return true
		}
	}
	return false
}

// 整数
func IsContainInt(value int, array []int) bool {
	for _, eachItem := range array {
		if value == eachItem {
			return true
		}
	}
	return false
}

/*
	基于map实现判断数组内包含指定元素
	性能：中
*/

// 整型
func IsContainIntMap(array []int) func(int) bool {
	set := make(map[int]struct{})

	for _, e := range array {
		set[e] = struct{}{}
	}

	return func(needle int) bool {
		_, ok := set[needle]
		return ok
	}
}

// 字符串
func IsContainStrMap(array []string) func(string) bool {
	set := make(map[string]struct{})

	for _, e := range array {
		set[e] = struct{}{}
	}

	return func(needle string) bool {
		_, ok := set[needle]
		return ok
	}
}
