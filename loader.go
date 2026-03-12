package envloader

import (
	"os"
	"reflect"
)

func Load(cfg interface{}) error {

	v := reflect.ValueOf(cfg).Elem()

	return parseStruct(v, "")
}

func getEnv(key string) string {
	return os.Getenv(key)
}
