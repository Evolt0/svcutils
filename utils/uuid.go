package utils

import (
	"github.com/go-courier/snowflakeid"
	"strconv"
)

var idGen, _ = snowflakeid.NewSnowflake(1)

// NewUUID id a Snowflake uuid generator
func NewUUID() (uint64, error) {
	id, err := idGen.ID()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewStringUUID() (string, error) {
	id, err := idGen.ID()
	if err != nil {
		return "", err
	}
	result := strconv.FormatUint(id, 10)
	return result, nil
}
