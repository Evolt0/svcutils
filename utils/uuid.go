package utils

import (
	"github.com/go-courier/snowflakeid"
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
