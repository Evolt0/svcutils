package utils

import (
	"encoding/json"
	"fmt"
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

func ConvertInterfaceToMap(src interface{}) (map[string]interface{}, error) {
	marshal, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}
	srcMap := make(map[string]interface{})
	err = json.Unmarshal(marshal, &srcMap)
	if err != nil {
		return nil, err
	}
	return srcMap, nil
}

func StringSliceContainsAny(s []string, sub string) error {
	for _, value := range s {
		if value == sub {
			return nil
		}
	}
	return fmt.Errorf("string slice not contain substring")
}

func UniqueSliceCheck(s []string) error {
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

func ConvertToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch v := value.(type) {
	case float64:
		key = strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		key = strconv.FormatFloat(float64(v), 'f', -1, 64)
	case int:
		key = strconv.Itoa(v)
	case uint:
		key = strconv.Itoa(int(v))
	case int8:
		key = strconv.Itoa(int(v))
	case uint8:
		key = strconv.Itoa(int(v))
	case int16:
		key = strconv.Itoa(int(v))
	case uint16:
		key = strconv.Itoa(int(v))
	case int32:
		key = strconv.Itoa(int(v))
	case uint32:
		key = strconv.Itoa(int(v))
	case int64:
		key = strconv.FormatInt(v, 10)
	case uint64:
		key = strconv.FormatUint(v, 10)
	case string:
		key = v
	case []byte:
		key = string(v)
	default:
		newValue, _ := json.Marshal(v)
		key = string(newValue)
	}

	return key
}
