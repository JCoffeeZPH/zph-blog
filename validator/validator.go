package validator

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("in", In)
		v.RegisterValidation("not_empty_max", NotEmptyMaxLen)
		v.RegisterValidation("not_empty_min", NotEmptyMinLen)
		v.RegisterValidation("empty_or_max", EmptyOrMax)
	}
}

func In(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	if v, ok := field.Interface().(string); ok {
		if len(v) == 0 {
			return true
		}
		params := strings.Split(param, " ")
		for _, p := range params {
			if strings.TrimSpace(p) == v {
				return true
			}
		}
		return false
	}
	if v, ok := field.Interface().(int32); ok {

		params := strings.Split(param, " ")
		for _, p := range params {
			if val, e := strconv.ParseInt(p, 10, 64); e == nil && int32(val) == v {
				return true
			}
		}

		return false
	}

	if v, ok := field.Interface().(uint32); ok {

		params := strings.Split(param, " ")
		for _, p := range params {
			if val, e := strconv.ParseUint(p, 10, 64); e == nil && uint32(val) == v {
				return true
			}
		}

		return false
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func NotEmptyMaxLen(fl validator.FieldLevel) bool {
	param := fl.Param()
	field := fl.Field()
	fieldKind := field.Kind()
	switch fieldKind {

	case reflect.String:
		if len(field.String()) == 0 {
			return false
		}
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			panic(err)
		}

		return int64(utf8.RuneCountInString(field.String())) <= p
	}
	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func NotEmptyMinLen(fl validator.FieldLevel) bool {
	param := fl.Param()
	field := fl.Field()
	fieldKind := field.Kind()
	switch fieldKind {

	case reflect.String:
		if len(field.String()) == 0 {
			return false
		}
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			panic(err)
		}

		return int64(utf8.RuneCountInString(field.String())) >= p
	}
	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func EmptyOrMax(fl validator.FieldLevel) bool {
	param := fl.Param()
	field := fl.Field()
	fieldKind := field.Kind()
	switch fieldKind {

	case reflect.String:
		if len(field.String()) == 0 {
			return true
		}
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			panic(err)
		}

		return int64(utf8.RuneCountInString(field.String())) <= p
	}
	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
