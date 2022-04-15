package utils

import "fmt"

func UniqueList[T ~string | ~int | ~uint64](keyList []T) (result []T) {
	dict := make(map[T]struct{})
	for _, key := range keyList {
		if _, ok := dict[key]; !ok {
			result = append(result, key)
			dict[key] = struct{}{}
		}
	}
	return result
}

func UniqueSliceCheck[T ~string | ~int | ~uint64](s []T) error {
	m := map[T]byte{} // 存放不重复主键
	for _, value := range s {
		l := len(m)
		m[value] = 0
		if len(m) == l { // 加入map后，map长度不变化，则元素重复
			return fmt.Errorf("slice has been repeat")
		}
	}
	return nil
}
