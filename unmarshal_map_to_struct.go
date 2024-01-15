package confuse

import (
	"errors"
	"reflect"
	"strconv"
)

var (
	// ErrParentNotPointer is returned when the parent is not a pointer.
	ErrParentNotPointer = errors.New("parent must be a pointer")
	// ErrParentNotStruct is returned when the parent is not a struct.
	ErrParentNotStruct = errors.New("parent must be a struct")
	// ErrFieldCannotBeSet is returned when the field cannot be set.
	ErrFieldCannotBeSet = errors.New("field cannot be set")
)

func parseMapToStruct(parent any, m map[string]any) error {
	t := reflect.ValueOf(parent)

	// If the parent is not a pointer, we cannot set the field.
	if t.Kind() != reflect.Ptr {
		return ErrParentNotPointer
	}

	// Access the underlying value of the pointer.
	t = t.Elem()

	// If the parent is not a struct, we cannot set the field.
	if t.Kind() != reflect.Struct {
		return ErrParentNotStruct
	}

	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}

	// Iterate over the fields of the struct.
	// See if the field matches the key.
	for i := 0; i < t.NumField(); i++ {
		if index := findIndexOfField(t.Type().Field(i), keys); index >= 0 {
			field := t.Field(i)

			if !field.CanSet() {
				continue
			}

			key := keys[index]
			value := m[key]

			err := parseField(field, m, key, value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func parseField(field reflect.Value, m map[string]any, key string, value any) error {
	switch field.Kind() {
	case reflect.Ptr:
		return parseField(field.Elem(), m, key, value)
	case reflect.Struct:
		if mappedValue, ok := value.(map[string]any); ok {
			err := parseMapToStruct(field.Addr().Interface(), mappedValue)
			if err != nil {
				return err
			}
		} else {
			return ErrFieldCannotBeSet
		}
	case reflect.Slice:
		if mappedValue, ok := value.([]any); ok {
			if field.IsNil() {
				field.Set(reflect.MakeSlice(field.Type(), 0, len(mappedValue)))
			}

			for _, v := range mappedValue {
				parsedValue, err := getValueOf(field.Type().Elem(), v)
				if err != nil {
					return err
				}

				field.Set(reflect.Append(field, parsedValue))
			}
		} else {
			return ErrFieldCannotBeSet
		}
	case reflect.Map:
		if mappedValue, ok := value.(map[string]any); ok {
			if field.IsNil() {
				field.Set(reflect.MakeMap(field.Type()))
			}

			for k, v := range mappedValue {
				parsedKey, err := getValueOf(field.Type().Key(), k)
				if err != nil {
					return err
				}

				parsedValue, err := getValueOf(field.Type().Elem(), v)
				if err != nil {
					return err
				}

				field.SetMapIndex(parsedKey, parsedValue)
			}
		} else {
			return ErrFieldCannotBeSet
		}
	case reflect.Array:
		if mappedValue, ok := value.([]any); ok {
			if field.Len() != len(mappedValue) {
				return ErrFieldCannotBeSet
			}

			for i, v := range mappedValue {
				parsedValue, err := getValueOf(field.Index(i).Type(), v)
				if err != nil {
					return err
				}
				field.Index(i).Set(parsedValue)
			}
		} else {
			return ErrFieldCannotBeSet
		}
	default:
		parsedValue, err := getValueOf(field.Type(), value)
		if err != nil {
			return err
		}

		field.Set(parsedValue)
	}

	return nil
}

func getValueOf(t reflect.Type, value any) (reflect.Value, error) {
	if t.AssignableTo(reflect.TypeOf(value)) {
		return reflect.ValueOf(value), nil
	}

	// If we get here, we should check to see if the value can be parsed from a string.
	// This is mainly due to the ENV variables being strings.

	switch t.Kind() {
	case reflect.Ptr:
		return getValueOf(t.Elem(), value)
	case reflect.String:
		return reflect.ValueOf(value), nil
	case reflect.Bool:
		if v, err := strconv.ParseBool(value.(string)); err == nil {
			return reflect.ValueOf(v), nil
		}
	case reflect.Int:
		if v, err := strconv.ParseInt(value.(string), 10, 64); err == nil {
			return reflect.ValueOf(int(v)), nil
		}
	case reflect.Int8:
		if v, err := strconv.ParseInt(value.(string), 10, 8); err == nil {
			return reflect.ValueOf(int8(v)), nil
		}
	case reflect.Int16:
		if v, err := strconv.ParseInt(value.(string), 10, 16); err == nil {
			return reflect.ValueOf(int16(v)), nil
		}
	case reflect.Int32:
		if v, err := strconv.ParseInt(value.(string), 10, 32); err == nil {
			return reflect.ValueOf(int32(v)), nil
		}
	case reflect.Int64:
		if v, err := strconv.ParseInt(value.(string), 10, 64); err == nil {
			return reflect.ValueOf(v), nil
		}
	case reflect.Uint:
		if v, err := strconv.ParseUint(value.(string), 10, 64); err == nil {
			return reflect.ValueOf(uint(v)), nil
		}
	case reflect.Uint8:
		if v, err := strconv.ParseUint(value.(string), 10, 8); err == nil {
			return reflect.ValueOf(uint8(v)), nil
		}
	case reflect.Uint16:
		if v, err := strconv.ParseUint(value.(string), 10, 16); err == nil {
			return reflect.ValueOf(uint16(v)), nil
		}
	case reflect.Uint32:
		if v, err := strconv.ParseUint(value.(string), 10, 32); err == nil {
			return reflect.ValueOf(uint32(v)), nil
		}
	case reflect.Uint64:
		if v, err := strconv.ParseUint(value.(string), 10, 64); err == nil {
			return reflect.ValueOf(v), nil
		}
	case reflect.Float32:
		if v, err := strconv.ParseFloat(value.(string), 32); err == nil {
			return reflect.ValueOf(float32(v)), nil
		}
	case reflect.Float64:
		if v, err := strconv.ParseFloat(value.(string), 64); err == nil {
			return reflect.ValueOf(v), nil
		}
	}

	return reflect.Value{}, ErrFieldCannotBeSet
}
