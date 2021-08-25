package utils

import (
	"reflect"
	"strconv"
)

func ConvertStringToUint64(opts ...string) ([]uint64, error) {
	list := make([]uint64, 0, len(opts))
	for _, value := range opts {
		intNum, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nil, err
		}
		list = append(list, intNum)
	}
	return list, nil
}

func ConvertInterfaceToStringSlice(obj interface{}) (list []string) {
	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		s := reflect.ValueOf(obj)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			list = append(list, ele.Interface().(string))
		}
	}
	return list
}