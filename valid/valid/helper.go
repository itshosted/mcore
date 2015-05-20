package valid

import (
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
	"reflect"
	"strconv"
)

func FnGetInt(i interface{}) (int64, error) {
	min, err := strconv.ParseInt(i.(string), 0, 0)
	if err != nil {
		return 0, err
	}

	return min, nil
}

func FnGetStr(i interface{}) (string, error) {
	return i.(string), nil
}

func FnGetStrSlice(i interface{}) ([]string, error) {
	return i.([]string), nil
}

/* Reflect struct */
func Validate(t interface{}) bool {
	/* Loop through each field in given struct */
	s := reflect.Indirect(reflect.ValueOf(t))
	for num := 0; num < s.NumField(); num++ {
		name := s.Type().Field(num).Name
		if name == "_" || !s.Field(num).CanInterface() {
			continue
		}

		/* Exported field, specific field rule */
		value := s.Field(num).Interface()

		/* Get validation rule */
		rule := s.Type().Field(num).Tag.Get("validate")
		if len(rule) == 0 {
			continue
		}

		/* Create parser for this rule and pass the context to it */
		l := new(Valdsl)
		l.Debug = true
		err, valid := l.Parse(t, rule, value)
		if err != nil {
			/* Deverror in rule */
			panic(err)
		}

		if !valid {
			return false
		}

		/* Is this a slice? */
		if s.Type().Field(num).Type.Kind() == reflect.Slice {
			/* Loop through slice and validate each item */
			s := reflect.ValueOf(s.Field(num).Interface())
			for i := 0; i < s.Len(); i++ {
				ret := Validate(s.Index(i).Interface())
				fmt.Println("...", ret)
				if !ret {
					return false
				}
			}
		} else if s.Type().Field(num).Type.Kind() == reflect.Struct {
			/* If it's a struct, only validate the struct */
			ret := Validate(s.Field(num).Interface())
			fmt.Println("...", ret)
			if !ret {
				return false
			}
		}
	}

	return true
}

func ParseForm(input interface{}, r *http.Request) error {
	r.ParseForm()
	decoder := schema.NewDecoder()
	err := decoder.Decode(input, r.PostForm)
	return err
}
