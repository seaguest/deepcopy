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
	if dstVal.Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}
	
	if dstVal.IsNil() {
		return fmt.Errorf("dst pointer cannot be nil")
	}

	srcVal := reflect.ValueOf(src)
	dstElem := dstVal.Elem()
	
	// Handle case where src is also a pointer
	if srcVal.Kind() == reflect.Ptr {
		if srcVal.IsNil() {
			// If src pointer is nil, set dst to zero value
			dstElem.Set(reflect.Zero(dstElem.Type()))
			return nil
		}
		
		// If dst is also expecting a pointer type, we need to handle pointer-to-pointer copying
		if dstElem.Kind() == reflect.Ptr {
			// Both src and dst are pointers - create target pointer and use copyRecursive
			elementType := dstElem.Type().Elem()
			newPtr := reflect.New(elementType)
			
			// Create temporary copy using the same logic as regular Copy
			cpy := reflect.New(srcVal.Type()).Elem()
			copyRecursive(srcVal, cpy)
			
			// Set the new pointer's element to point to our copied value
			if cpy.Kind() == reflect.Ptr {
				newPtr.Elem().Set(cpy.Elem())
			} else {
				newPtr.Elem().Set(cpy)
			}
			
			dstElem.Set(newPtr)
			return nil
		} else {
			// src is pointer, dst expects value - dereference src
			srcVal = srcVal.Elem()
		}
	} else {
		// src is not a pointer
		if dstElem.Kind() == reflect.Ptr {
			// src is value, dst expects pointer - create new pointer
			newPtr := reflect.New(srcVal.Type())
			copyRecursive(srcVal, newPtr.Elem())
			dstElem.Set(newPtr)
			return nil
		}
	}
	
	// Check type compatibility
	if !srcVal.Type().AssignableTo(dstElem.Type()) {
		return fmt.Errorf("cannot assign %v to %v", srcVal.Type(), dstElem.Type())
	}

	// Create a temporary copy
	cpy := reflect.New(srcVal.Type()).Elem()
	copyRecursive(srcVal, cpy)
	
	// Set the destination to the copy
	dstElem.Set(cpy)
	
	return nil
}

// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
func copyRecursive(original, cpy reflect.Value) {
	// check for implement deepcopy.Interface
	if original.CanInterface() {
		if copier, ok := original.Interface().(Interface); ok {
			// return if copier is nil
			if reflect.ValueOf(copier).IsNil() {
				return
			}
			cpy.Set(reflect.ValueOf(copier.DeepCopy()))
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
			copyKey := Copy(key.Interface())
			cpy.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}

	default:
		cpy.Set(original)
	}
}
