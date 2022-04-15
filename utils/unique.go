package utils

import "fmt"

func UniqueUint64List(keyList []uint64) (result []uint64) {
	dict := make(map[uint64]struct{})
	for _, key := range keyList {
		if _, ok := dict[key]; !ok {
			result = append(result, key)
			dict[key] = struct{}{}
		}
	}
	return result
}

func UniqueAnyList(keyList []interface{}) (result []interface{}) {
	dict := make(map[interface{}]struct{})
	for _, key := range keyList {
		if _, ok := dict[key]; !ok {
			result = append(result, key)
			dict[key] = struct{}{}
		}
	}
	return result
}

func UniqueStingSliceCheck(s []string) error {
	m := map[string]byte{} // 存放不重复主键
	for _, value := range s {
		l := len(m)
		m[value] = 0
		if len(m) == l { // 加入map后，map长度不变化，则元素重复
			return fmt.Errorf("slice has been repeat")
		}
	}
	return nil
}
