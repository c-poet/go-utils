// Package types provides reflection helpers.
package types

import "reflect"

// IsStruct reports whether value is a struct or a pointer to a struct.
func IsStruct(value any) bool {
	if value == nil {
		return false
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Kind() == reflect.Struct
}
