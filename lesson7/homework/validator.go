package homework

import (
	"github.com/pkg/errors"
	"reflect"
	"strconv"
	"strings"
)

var ErrNotStruct = errors.New("wrong argument given, should be a struct")
var ErrInvalidValidatorSyntax = errors.New("invalid validator syntax")
var ErrValidateForUnexportedFields = errors.New("validation for unexported field is not allowed")

type ValidationError struct {
	Err error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	// TODO: implement this
	allErrors := ""
	for _, err := range v {
		allErrors += err.Err.Error()
	}
	return allErrors
}

func checkExported(a reflect.StructField) bool {
	return a.PkgPath == ""
}

func contains(s []string, str any) bool {
	if len(s) == 1 && s[0] == "" { // пустой in
		return false
	}
	for _, v := range s {
		if val, ok := str.(string); ok {
			if v == val {
				return true
			}
		} else if val, ok := str.(int64); ok {
			if strconv.FormatInt(val, 10) == v {
				return true
			}
		}
	}
	return false
}

func Validate(v any) error {
	// TODO: implement this
	tp := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	if tp.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	var fail ValidationErrors = make([]ValidationError, 0)

	for i := 0; i < tp.NumField(); i++ {
		fil := tp.Field(i)

		switch fil.Type.Kind() {
		case reflect.String:
			s, ok := fil.Tag.Lookup("validate")
			if !ok {
				continue
			}

			if !checkExported(fil) {
				fail = append(fail, ValidationError{ErrValidateForUnexportedFields})
				continue
			}

			oper, ost, ok := strings.Cut(s, ":")
			if !ok {
				fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
				continue
			}

			switch oper {
			case "len":
				l, err := strconv.Atoi(ost)
				if err != nil {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
				if len(val.Field(i).String()) != l {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
			case "in":
				if !contains(strings.Split(ost, ","), val.Field(i).String()) {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
			case "min":
				l, err := strconv.Atoi(ost)
				if err != nil {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
				if len(val.Field(i).String()) < l {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
			case "max":
				l, err := strconv.Atoi(ost)
				if err != nil {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
				if len(val.Field(i).String()) > l {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
			}

		case reflect.Int:
			s, ok := fil.Tag.Lookup("validate")
			if !ok {
				continue
			}

			if !checkExported(fil) {
				fail = append(fail, ValidationError{ErrValidateForUnexportedFields})
				continue
			}

			oper, ost, ok := strings.Cut(s, ":")
			if !ok {
				fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
				continue
			}

			switch oper {
			case "in":
				if !contains(strings.Split(ost, ","), val.Field(i).Int()) {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
			case "min":
				l, err := strconv.ParseInt(ost, 10, 64)
				if err != nil {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
				if val.Field(i).Int() < l {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
			case "max":
				l, err := strconv.ParseInt(ost, 10, 64)
				if err != nil {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
				if val.Field(i).Int() > l {
					fail = append(fail, ValidationError{ErrInvalidValidatorSyntax})
					continue
				}
			}
		}
	}
	if len(fail) == 0 {
		return nil
	}
	return error(fail)
}
