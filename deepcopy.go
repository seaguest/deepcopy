// deepcopy makes deep copies of things. A standard copy will copy the
// pointers: deep copy copies the values pointed to.  Unexported field
// values are not copied.
//
// Copyright (c)2014-2016, Joel Scoble (github.com/mohae), all rights reserved.
// License: MIT, for more details check the included LICENSE file.
package deepcopy

import (
	"fmt"
	"reflect"
	"time"
)

// Interface for delegating copy process to type
type Interface interface {
	DeepCopy() interface{}
}

// Iface is an alias to Copy; this exists for backwards compatibility reasons.
func Iface(iface interface{}) interface{} {
	return Copy(iface)
}

// Copy creates a deep copy of whatever is passed to it and returns the copy
// in an interface{}.  The returned value will need to be asserted to the
// correct type.
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// Make the interface a reflect.Value
	original := reflect.ValueOf(src)

	// Make a copy of the same type as the original.
	cpy := reflect.New(original.Type()).Elem()

	// Recursively copy the original.
	copyRecursive(original, cpy)

	// Return the copy as an interface.
	return cpy.Interface()
}

// CopyTo copies the src value to dst. The dst must be a pointer to the target type.
// This allows copying to an existing variable without allocating a new one.
// Both src and dst can be pointers, and proper deep copying will be performed.
func CopyTo(src, dst interface{}) error {
	if src == nil {
		return nil
	}

	if dst == nil {
		return fmt.Errorf("dst cannot be nil")
	}

	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Ptr || dstVal.IsNil() {
		return fmt.Errorf("dst must be a non-nil pointer")
	}

	// Make a deep copy using the existing Copy function
	copied := Copy(src)
	if copied == nil {
		// Handle nil src by setting dst to zero value
		dstVal.Elem().Set(reflect.Zero(dstVal.Elem().Type()))
		return nil
	}

	copiedVal := reflect.ValueOf(copied)
	dstElem := dstVal.Elem()

	// Handle type mismatches with automatic conversion
	if copiedVal.Type() != dstElem.Type() {
		// Try to handle pointer/value mismatches
		if copiedVal.Kind() == reflect.Ptr && dstElem.Kind() != reflect.Ptr {
			// src is pointer, dst expects value - dereference
			if copiedVal.IsNil() {
				dstElem.Set(reflect.Zero(dstElem.Type()))
				return nil
			}
			copiedVal = copiedVal.Elem()
		} else if copiedVal.Kind() != reflect.Ptr && dstElem.Kind() == reflect.Ptr {
			// src is value, dst expects pointer - create pointer
			newPtr := reflect.New(copiedVal.Type())
			newPtr.Elem().Set(copiedVal)
			copiedVal = newPtr
		}

		// Final type check
		if !copiedVal.Type().AssignableTo(dstElem.Type()) {
			return fmt.Errorf("cannot assign %v to %v", copiedVal.Type(), dstElem.Type())
		}
	}

	dstElem.Set(copiedVal)
	return nil
}

// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
func copyRecursive(original, cpy reflect.Value) {
	// check for implement deepcopy.Interface
	if original.CanInterface() {
		if copier, ok := original.Interface().(Interface); ok {
			// return if copier is nil - safe nil check
			if copier == nil || (reflect.ValueOf(copier).Kind() == reflect.Ptr && reflect.ValueOf(copier).IsNil()) {
				return
			}

			deepCopyResult := copier.DeepCopy()
			if deepCopyResult == nil {
				// If DeepCopy returns nil, set the copy to the zero value of its type
				cpy.Set(reflect.Zero(cpy.Type()))
			} else {
				cpy.Set(reflect.ValueOf(deepCopyResult))
			}
			return
		}
	}

	// handle according to original's Kind
	switch original.Kind() {
	case reflect.Ptr:
		// Get the actual value being pointed to.
		originalValue := original.Elem()

		// if  it isn't valid, return.
		if !originalValue.IsValid() {
			return
		}
		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, cpy.Elem())

	case reflect.Interface:
		// If this is a nil, don't do anything
		if original.IsNil() {
			return
		}
		// Get the value for the interface, not the pointer.
		originalValue := original.Elem()

		// Get the value by calling Elem().
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		cpy.Set(copyValue)

	case reflect.Struct:
		t, ok := original.Interface().(time.Time)
		if ok {
			cpy.Set(reflect.ValueOf(t))
			return
		}
		// Go through each field of the struct and copy it.
		for i := 0; i < original.NumField(); i++ {
			// The Type's StructField for a given field is checked to see if StructField.PkgPath
			// is set to determine if the field is exported or not because CanSet() returns false
			// for settable fields.  I'm not sure why.  -mohae
			if original.Type().Field(i).PkgPath != "" {
				continue
			}
			copyRecursive(original.Field(i), cpy.Field(i))
		}

	case reflect.Slice:
		if original.IsNil() {
			return
		}
		// Make a new slice and copy each element.
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), cpy.Index(i))
		}

	case reflect.Map:
		if original.IsNil() {
			return
		}
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue)

			// Use copyRecursive for keys too for consistency
			copyKey := reflect.New(key.Type()).Elem()
			copyRecursive(key, copyKey)
			cpy.SetMapIndex(copyKey, copyValue)
		}

	case reflect.Array:
		// Handle arrays by copying each element
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), cpy.Index(i))
		}

	default:
		cpy.Set(original)
	}
}
