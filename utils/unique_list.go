package utils

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

func UniqueList(keyList []interface{}) (result []interface{}) {
	dict := make(map[interface{}]struct{})
	for _, key := range keyList {
		if _, ok := dict[key]; !ok {
			result = append(result, key)
			dict[key] = struct{}{}
		}
	}
	return result
}
