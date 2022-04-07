package utils

func MergeMap(mObj ...map[string]interface{}) map[string]interface{} {
	newObj := make(map[string]interface{})
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}

func MapSliceLarge(dict map[int64]int) int {
	result := 0
	for _, v := range dict {
		if v > result {
			result = v
		}
	}
	return result
}

func CountUint64ListWithoutZero(keyList []uint64) (dict map[uint64]int) {
	dict = make(map[uint64]int)
	for _, key := range keyList {
		if key == 0 {
			continue
		}
		dict[key]++
	}
	return dict
}