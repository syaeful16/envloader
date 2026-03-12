package envloader

import (
	"fmt"
	"reflect"
)

func parseStruct(v reflect.Value, parentPrefix string) error {

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {

		field := v.Field(i)
		fieldType := t.Field(i)

		if field.Kind() == reflect.Struct && fieldType.Anonymous == false {
			err := parseStruct(field, parentPrefix)
			if err != nil {
				return err
			}
			continue
		}

		envKey := fieldType.Tag.Get("env")
		defaultValue := fieldType.Tag.Get("default")
		required := fieldType.Tag.Get("required")
		prefix := fieldType.Tag.Get("prefix")

		if prefix != "" {
			envKey = prefix + "_" + envKey
		}

		value := getEnv(envKey)

		if value == "" {
			value = defaultValue
		}

		if value == "" && required == "true" {
			return fmt.Errorf("required env missing: %s", envKey)
		}

		err := setValue(field, value)
		if err != nil {
			return err
		}
	}

	return nil
}
