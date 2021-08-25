package order

import (
	"bytes"
	"encoding/json"
	"strings"
)

func Order(data interface{}) string {
	text, _ := jsonMarshal(data)
	temp := make(map[string]interface{})
	_ = json.Unmarshal(text, &temp)
	Filter(temp)
	result, _ := jsonMarshal(temp)
	return strings.TrimSuffix(string(result), "\n")
}

func Filter(data map[string]interface{}) {
	// 添加需要剔除的字段
	/*if _, ok := data["sign"]; ok {
		delete(data, "sign")
	}*/
	for _, value := range data {
		if elem, ok := value.(map[string]interface{}); ok {
			Filter(elem)
		}
	}
}

// json marshal escape &, <, and > to \u0026, \u003c, and \u003e
func jsonMarshal(data interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
