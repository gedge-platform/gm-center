package api

import (
	"strconv"
)

func StringToInt(i string) int {
	v, _ := strconv.Atoi(i)
	return v
}
