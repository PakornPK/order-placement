package util

import (
	"os"
	"strconv"
)

func GetEnvString(name, val string) string {
	e := os.Getenv(name)
	if e == "" {
		return val
	}
	return e
}

func GetEnvInt(name string, val int) int {
	e := os.Getenv(name)
	if e == "" {
		return val
	}
	v, err := strconv.Atoi(e)
	if err != nil {
		return val
	}
	return v
}
