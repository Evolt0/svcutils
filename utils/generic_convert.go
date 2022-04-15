package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

func ConvertStringToUInteger[T ~uint | ~uint8 | ~uint16 | ~uint64](opts ...string) ([]T, error) {
	list := make([]T, 0, len(opts))
	for _, value := range opts {
		uIntNum, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nil, err
		}
		list = append(list, T(uIntNum))
	}
	return list, nil
}
func ConvertStringToInteger[T ~int | ~int8 | ~int16 | ~int64](opts ...string) ([]T, error) {
	list := make([]T, 0, len(opts))
	for _, value := range opts {
		uIntNum, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nil, err
		}
		list = append(list, T(uIntNum))
	}
	return list, nil
}

// 在泛型中添加类型
func ConvertInterfaceToSlice[T ~string](obj interface{}) (list []T) {
	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		s := reflect.ValueOf(obj)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			list = append(list, ele.Interface().(T))
		}
	}
	return list
}

func SliceContainsAny[T ~string | ~int](s []T, sub T) error {
	for _, value := range s {
		if value == sub {
			return nil
		}
	}
	return fmt.Errorf("string slice not contain substring")
}
