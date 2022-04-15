package utils

func MergeMap[T ~string | ~int | ~uint64](mObj ...map[T]interface{}) map[T]interface{} {
	newObj := make(map[T]interface{})
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}

func MapSliceLarge[T ~uint | ~uint8 | ~uint16 | ~uint64 | ~int | ~int8 | ~int32 | ~int64](dict map[T]int) int {
	result := 0
	for _, v := range dict {
		if v > result {
			result = v
		}
	}
	return result
}

func CountListWithoutZero[T ~uint | ~uint8 | ~uint16 | ~uint64 | ~int | ~int8 | ~int32 | ~int64](keyList []T) (dict map[T]int) {
	dict = make(map[T]int)
	for _, key := range keyList {
		if key == 0 {
			continue
		}
		dict[key]++
	}
	return dict
}
