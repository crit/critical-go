package form

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/fatih/structs"
)

type input struct {
	Name     string
	Value    string
	Active   bool
	Required bool
	Disabled bool
}

var FloatPrecision = 10

func Marshal(v interface{}) (out string, err error) {
	out += "<form>"

	if !structs.IsStruct(v) {
		return out, errors.New("can only marshal structs")
	}

	for _, field := range structs.Fields(v) {
		var el input
		el.Name = strings.ToLower(field.Name())
		el.Value = toString(field.Value())

		if field.Kind() == reflect.Bool {
			el.Active = field.Value().(bool)
		}

		var skip bool
		tags := field.Tag("form")

		if tags == "-" {
			continue
		}

		for i, tag := range strings.Split(field.Tag("form"), ",") {
			if i == 0 && tag != "" {
				el.Name = tag
				continue
			}

			switch tag {
			case "omitempty":
				if field.IsZero() {
					skip = true
					continue
				}
			case "required":
				el.Required = true
			case "disabled":
				el.Disabled = true
			}
		}

		if skip {
			continue
		}

		if strings.Contains(tags, "text") {
			out += toTextField(el)
			continue
		}

		if strings.Contains(tags, "hidden") {
			out += toHiddenField(el)
			continue
		}

		if strings.Contains(tags, "checkbox") {
			out += toCheckboxField(el)
		}

		switch field.Kind() {
		case reflect.Bool:
			out += toCheckboxField(el)
			continue
		case reflect.Map, reflect.Ptr, reflect.Struct, reflect.Slice:
			continue
		}

		out += toTextField(el)
	}

	out += "</form>"

	return out, nil
}

func toString(v interface{}) (out string) {
	switch v.(type) {
	case bool:
		out = strconv.FormatBool(v.(bool))

	case int:
		out = strconv.FormatInt(int64(v.(int)), 10)
	case int8:
		out = strconv.FormatInt(int64(v.(int8)), 10)
	case int16:
		out = strconv.FormatInt(int64(v.(int16)), 10)
	case int32:
		out = strconv.FormatInt(int64(v.(int32)), 10)
	case int64:
		out = strconv.FormatInt(v.(int64), 10)

	case uint:
		out = strconv.FormatUint(uint64(v.(uint)), 10)
	case uint8:
		out = strconv.FormatUint(uint64(v.(uint8)), 10)
	case uint16:
		out = strconv.FormatUint(uint64(v.(uint16)), 10)
	case uint32:
		out = strconv.FormatUint(uint64(v.(uint32)), 10)

	case float32:
		out = strconv.FormatFloat(float64(v.(float32)), 'f', FloatPrecision, 64)
	case float64:
		out = strconv.FormatFloat(v.(float64), 'f', FloatPrecision, 64)

	case string:
		out = v.(string)
	}

	return out
}
