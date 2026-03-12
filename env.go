package envloader

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

func setValue(field reflect.Value, value string) error {

	switch field.Kind() {

	case reflect.String:
		field.SetString(value)

	case reflect.Int:
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		field.SetInt(int64(i))

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(b)

	case reflect.Slice:

		if field.Type().Elem().Kind() == reflect.String {
			items := strings.Split(value, ",")
			field.Set(reflect.ValueOf(items))
		}

	case reflect.Int64:

		if field.Type().String() == "time.Duration" {

			d, err := time.ParseDuration(value)
			if err != nil {
				return err
			}

			field.SetInt(int64(d))
		}

	}

	return nil
}
