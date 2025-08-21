package deepcopy

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

// just basic is this working stuff
func TestSimple(t *testing.T) {
	Strings := []string{"a", "b", "c"}
	cpyS := Copy(Strings).([]string)
	if (*reflect.SliceHeader)(unsafe.Pointer(&Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyS)).Data {
		t.Error("[]string: expected SliceHeader data pointers to point to different locations, they didn't")
		goto CopyBools
	}
	if len(cpyS) != len(Strings) {
		t.Errorf("[]string: len was %d; want %d", len(cpyS), len(Strings))
		goto CopyBools
	}
	for i, v := range Strings {
		if v != cpyS[i] {
			t.Errorf("[]string: got %v at index %d of the copy; want %v", cpyS[i], i, v)
		}
	}

CopyBools:
	Bools := []bool{true, true, false, false}
	cpyB := Copy(Bools).([]bool)
	if (*reflect.SliceHeader)(unsafe.Pointer(&Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyB)).Data {
		t.Error("[]bool: expected SliceHeader data pointers to point to different locations, they didn't")
		goto CopyBytes
	}
	if len(cpyB) != len(Bools) {
		t.Errorf("[]bool: len was %d; want %d", len(cpyB), len(Bools))
		goto CopyBytes
	}
	for i, v := range Bools {
		if v != cpyB[i] {
			t.Errorf("[]bool: got %v at index %d of the copy; want %v", cpyB[i], i, v)
		}
	}

CopyBytes:
	Bytes := []byte("hello")
	cpyBt := Copy(Bytes).([]byte)
	if (*reflect.SliceHeader)(unsafe.Pointer(&Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyBt)).Data {
		t.Error("[]byte: expected SliceHeader data pointers to point to different locations, they didn't")
		goto CopyInts
	}
	if len(cpyBt) != len(Bytes) {
		t.Errorf("[]byte: len was %d; want %d", len(cpyBt), len(Bytes))
		goto CopyInts
	}
	for i, v := range Bytes {
		if v != cpyBt[i] {
			t.Errorf("[]byte: got %v at index %d of the copy; want %v", cpyBt[i], i, v)
		}
	}

CopyInts:
	Ints := []int{42}
	cpyI := Copy(Ints).([]int)
	if (*reflect.SliceHeader)(unsafe.Pointer(&Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyI)).Data {
		t.Error("[]int: expected SliceHeader data pointers to point to different locations, they didn't")
		goto CopyUints
	}
	if len(cpyI) != len(Ints) {
		t.Errorf("[]int: len was %d; want %d", len(cpyI), len(Ints))
		goto CopyUints
	}
	for i, v := range Ints {
		if v != cpyI[i] {
			t.Errorf("[]int: got %v at index %d of the copy; want %v", cpyI[i], i, v)
		}
	}

CopyUints:
	Uints := []uint{1, 2, 3, 4, 5}
	cpyU := Copy(Uints).([]uint)
	if (*reflect.SliceHeader)(unsafe.Pointer(&Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyU)).Data {
		t.Error("[]: expected SliceHeader data pointers to point to different locations, they didn't")
		goto CopyFloat32s
	}
	if len(cpyU) != len(Uints) {
		t.Errorf("[]uint: len was %d; want %d", len(cpyU), len(Uints))
		goto CopyFloat32s
	}
	for i, v := range Uints {
		if v != cpyU[i] {
			t.Errorf("[]uint: got %v at index %d of the copy; want %v", cpyU[i], i, v)
		}
	}

CopyFloat32s:
	Float32s := []float32{3.14}
	cpyF := Copy(Float32s).([]float32)
	if (*reflect.SliceHeader)(unsafe.Pointer(&Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyF)).Data {
		t.Error("[]float32: expected SliceHeader data pointers to point to different locations, they didn't")
		goto CopyInterfaces
	}
	if len(cpyF) != len(Float32s) {
		t.Errorf("[]float32: len was %d; want %d", len(cpyF), len(Float32s))
		goto CopyInterfaces
	}
	for i, v := range Float32s {
		if v != cpyF[i] {
			t.Errorf("[]float32: got %v at index %d of the copy; want %v", cpyF[i], i, v)
		}
	}

CopyInterfaces:
	Interfaces := []interface{}{"a", 42, true, 4.32}
	cpyIf := Copy(Interfaces).([]interface{})
	if (*reflect.SliceHeader)(unsafe.Pointer(&Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyIf)).Data {
		t.Error("[]interfaces: expected SliceHeader data pointers to point to different locations, they didn't")
		return
	}
	if len(cpyIf) != len(Interfaces) {
		t.Errorf("[]interface{}: len was %d; want %d", len(cpyIf), len(Interfaces))
		return
	}
	for i, v := range Interfaces {
		if v != cpyIf[i] {
			t.Errorf("[]interface{}: got %v at index %d of the copy; want %v", cpyIf[i], i, v)
		}
	}
}

type Basics struct {
	String      string
	Strings     []string
	StringArr   [4]string
	Bool        bool
	Bools       []bool
	Byte        byte
	Bytes       []byte
	Int         int
	Ints        []int
	Int8        int8
	Int8s       []int8
	Int16       int16
	Int16s      []int16
	Int32       int32
	Int32s      []int32
	Int64       int64
	Int64s      []int64
	Uint        uint
	Uints       []uint
	Uint8       uint8
	Uint8s      []uint8
	Uint16      uint16
	Uint16s     []uint16
	Uint32      uint32
	Uint32s     []uint32
	Uint64      uint64
	Uint64s     []uint64
	Float32     float32
	Float32s    []float32
	Float64     float64
	Float64s    []float64
	Complex64   complex64
	Complex64s  []complex64
	Complex128  complex128
	Complex128s []complex128
	Interface   interface{}
	Interfaces  []interface{}
}

// These tests test that all supported basic types are copied correctly.  This
// is done by copying a struct with fields of most of the basic types as []T.
func TestMostTypes(t *testing.T) {
	test := Basics{
		String:      "kimchi",
		Strings:     []string{"uni", "ika"},
		StringArr:   [4]string{"malort", "barenjager", "fernet", "salmiakki"},
		Bool:        true,
		Bools:       []bool{true, false, true},
		Byte:        'z',
		Bytes:       []byte("abc"),
		Int:         42,
		Ints:        []int{0, 1, 3, 4},
		Int8:        8,
		Int8s:       []int8{8, 9, 10},
		Int16:       16,
		Int16s:      []int16{16, 17, 18, 19},
		Int32:       32,
		Int32s:      []int32{32, 33},
		Int64:       64,
		Int64s:      []int64{64},
		Uint:        420,
		Uints:       []uint{11, 12, 13},
		Uint8:       81,
		Uint8s:      []uint8{81, 82},
		Uint16:      160,
		Uint16s:     []uint16{160, 161, 162, 163, 164},
		Uint32:      320,
		Uint32s:     []uint32{320, 321},
		Uint64:      640,
		Uint64s:     []uint64{6400, 6401, 6402, 6403},
		Float32:     32.32,
		Float32s:    []float32{32.32, 33},
		Float64:     64.1,
		Float64s:    []float64{64, 65, 66},
		Complex64:   complex64(-64 + 12i),
		Complex64s:  []complex64{complex64(-65 + 11i), complex64(66 + 10i)},
		Complex128:  complex128(-128 + 12i),
		Complex128s: []complex128{complex128(-128 + 11i), complex128(129 + 10i)},
		Interfaces:  []interface{}{42, true, "pan-galactic"},
	}

	cpy := Copy(test).(Basics)

	// see if they point to the same location
	if fmt.Sprintf("%p", &cpy) == fmt.Sprintf("%p", &test) {
		t.Error("address of copy was the same as original; they should be different")
		return
	}

	// Go through each field and check to see it got copied properly
	if cpy.String != test.String {
		t.Errorf("String: got %v; want %v", cpy.String, test.String)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Strings)).Data {
		t.Error("Strings: address of copy was the same as original; they should be different")
		goto StringArr
	}

	if len(cpy.Strings) != len(test.Strings) {
		t.Errorf("Strings: len was %d; want %d", len(cpy.Strings), len(test.Strings))
		goto StringArr
	}
	for i, v := range test.Strings {
		if v != cpy.Strings[i] {
			t.Errorf("Strings: got %v at index %d of the copy; want %v", cpy.Strings[i], i, v)
		}
	}

StringArr:
	if unsafe.Pointer(&test.StringArr) == unsafe.Pointer(&cpy.StringArr) {
		t.Error("StringArr: address of copy was the same as original; they should be different")
		goto Bools
	}
	for i, v := range test.StringArr {
		if v != cpy.StringArr[i] {
			t.Errorf("StringArr: got %v at index %d of the copy; want %v", cpy.StringArr[i], i, v)
		}
	}

Bools:
	if cpy.Bool != test.Bool {
		t.Errorf("Bool: got %v; want %v", cpy.Bool, test.Bool)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Bools)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Bools)).Data {
		t.Error("Bools: address of copy was the same as original; they should be different")
		goto Bytes
	}
	if len(cpy.Bools) != len(test.Bools) {
		t.Errorf("Bools: len was %d; want %d", len(cpy.Bools), len(test.Bools))
		goto Bytes
	}
	for i, v := range test.Bools {
		if v != cpy.Bools[i] {
			t.Errorf("Bools: got %v at index %d of the copy; want %v", cpy.Bools[i], i, v)
		}
	}

Bytes:
	if cpy.Byte != test.Byte {
		t.Errorf("Byte: got %v; want %v", cpy.Byte, test.Byte)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Bytes)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Bytes)).Data {
		t.Error("Bytes: address of copy was the same as original; they should be different")
		goto Ints
	}
	if len(cpy.Bytes) != len(test.Bytes) {
		t.Errorf("Bytes: len was %d; want %d", len(cpy.Bytes), len(test.Bytes))
		goto Ints
	}
	for i, v := range test.Bytes {
		if v != cpy.Bytes[i] {
			t.Errorf("Bytes: got %v at index %d of the copy; want %v", cpy.Bytes[i], i, v)
		}
	}

Ints:
	if cpy.Int != test.Int {
		t.Errorf("Int: got %v; want %v", cpy.Int, test.Int)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Ints)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Ints)).Data {
		t.Error("Ints: address of copy was the same as original; they should be different")
		goto Int8s
	}
	if len(cpy.Ints) != len(test.Ints) {
		t.Errorf("Ints: len was %d; want %d", len(cpy.Ints), len(test.Ints))
		goto Int8s
	}
	for i, v := range test.Ints {
		if v != cpy.Ints[i] {
			t.Errorf("Ints: got %v at index %d of the copy; want %v", cpy.Ints[i], i, v)
		}
	}

Int8s:
	if cpy.Int8 != test.Int8 {
		t.Errorf("Int8: got %v; want %v", cpy.Int8, test.Int8)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Int8s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Int8s)).Data {
		t.Error("Int8s: address of copy was the same as original; they should be different")
		goto Int16s
	}
	if len(cpy.Int8s) != len(test.Int8s) {
		t.Errorf("Int8s: len was %d; want %d", len(cpy.Int8s), len(test.Int8s))
		goto Int16s
	}
	for i, v := range test.Int8s {
		if v != cpy.Int8s[i] {
			t.Errorf("Int8s: got %v at index %d of the copy; want %v", cpy.Int8s[i], i, v)
		}
	}

Int16s:
	if cpy.Int16 != test.Int16 {
		t.Errorf("Int16: got %v; want %v", cpy.Int16, test.Int16)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Int16s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Int16s)).Data {
		t.Error("Int16s: address of copy was the same as original; they should be different")
		goto Int32s
	}
	if len(cpy.Int16s) != len(test.Int16s) {
		t.Errorf("Int16s: len was %d; want %d", len(cpy.Int16s), len(test.Int16s))
		goto Int32s
	}
	for i, v := range test.Int16s {
		if v != cpy.Int16s[i] {
			t.Errorf("Int16s: got %v at index %d of the copy; want %v", cpy.Int16s[i], i, v)
		}
	}

Int32s:
	if cpy.Int32 != test.Int32 {
		t.Errorf("Int32: got %v; want %v", cpy.Int32, test.Int32)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Int32s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Int32s)).Data {
		t.Error("Int32s: address of copy was the same as original; they should be different")
		goto Int64s
	}
	if len(cpy.Int32s) != len(test.Int32s) {
		t.Errorf("Int32s: len was %d; want %d", len(cpy.Int32s), len(test.Int32s))
		goto Int64s
	}
	for i, v := range test.Int32s {
		if v != cpy.Int32s[i] {
			t.Errorf("Int32s: got %v at index %d of the copy; want %v", cpy.Int32s[i], i, v)
		}
	}

Int64s:
	if cpy.Int64 != test.Int64 {
		t.Errorf("Int64: got %v; want %v", cpy.Int64, test.Int64)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Int64s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Int64s)).Data {
		t.Error("Int64s: address of copy was the same as original; they should be different")
		goto Uints
	}
	if len(cpy.Int64s) != len(test.Int64s) {
		t.Errorf("Int64s: len was %d; want %d", len(cpy.Int64s), len(test.Int64s))
		goto Uints
	}
	for i, v := range test.Int64s {
		if v != cpy.Int64s[i] {
			t.Errorf("Int64s: got %v at index %d of the copy; want %v", cpy.Int64s[i], i, v)
		}
	}

Uints:
	if cpy.Uint != test.Uint {
		t.Errorf("Uint: got %v; want %v", cpy.Uint, test.Uint)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Uints)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Uints)).Data {
		t.Error("Uints: address of copy was the same as original; they should be different")
		goto Uint8s
	}
	if len(cpy.Uints) != len(test.Uints) {
		t.Errorf("Uints: len was %d; want %d", len(cpy.Uints), len(test.Uints))
		goto Uint8s
	}
	for i, v := range test.Uints {
		if v != cpy.Uints[i] {
			t.Errorf("Uints: got %v at index %d of the copy; want %v", cpy.Uints[i], i, v)
		}
	}

Uint8s:
	if cpy.Uint8 != test.Uint8 {
		t.Errorf("Uint8: got %v; want %v", cpy.Uint8, test.Uint8)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Uint8s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Uint8s)).Data {
		t.Error("Uint8s: address of copy was the same as original; they should be different")
		goto Uint16s
	}
	if len(cpy.Uint8s) != len(test.Uint8s) {
		t.Errorf("Uint8s: len was %d; want %d", len(cpy.Uint8s), len(test.Uint8s))
		goto Uint16s
	}
	for i, v := range test.Uint8s {
		if v != cpy.Uint8s[i] {
			t.Errorf("Uint8s: got %v at index %d of the copy; want %v", cpy.Uint8s[i], i, v)
		}
	}

Uint16s:
	if cpy.Uint16 != test.Uint16 {
		t.Errorf("Uint16: got %v; want %v", cpy.Uint16, test.Uint16)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Uint16s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Uint16s)).Data {
		t.Error("Uint16s: address of copy was the same as original; they should be different")
		goto Uint32s
	}
	if len(cpy.Uint16s) != len(test.Uint16s) {
		t.Errorf("Uint16s: len was %d; want %d", len(cpy.Uint16s), len(test.Uint16s))
		goto Uint32s
	}
	for i, v := range test.Uint16s {
		if v != cpy.Uint16s[i] {
			t.Errorf("Uint16s: got %v at index %d of the copy; want %v", cpy.Uint16s[i], i, v)
		}
	}

Uint32s:
	if cpy.Uint32 != test.Uint32 {
		t.Errorf("Uint32: got %v; want %v", cpy.Uint32, test.Uint32)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Uint32s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Uint32s)).Data {
		t.Error("Uint32s: address of copy was the same as original; they should be different")
		goto Uint64s
	}
	if len(cpy.Uint32s) != len(test.Uint32s) {
		t.Errorf("Uint32s: len was %d; want %d", len(cpy.Uint32s), len(test.Uint32s))
		goto Uint64s
	}
	for i, v := range test.Uint32s {
		if v != cpy.Uint32s[i] {
			t.Errorf("Uint32s: got %v at index %d of the copy; want %v", cpy.Uint32s[i], i, v)
		}
	}

Uint64s:
	if cpy.Uint64 != test.Uint64 {
		t.Errorf("Uint64: got %v; want %v", cpy.Uint64, test.Uint64)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Uint64s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Uint64s)).Data {
		t.Error("Uint64s: address of copy was the same as original; they should be different")
		goto Float32s
	}
	if len(cpy.Uint64s) != len(test.Uint64s) {
		t.Errorf("Uint64s: len was %d; want %d", len(cpy.Uint64s), len(test.Uint64s))
		goto Float32s
	}
	for i, v := range test.Uint64s {
		if v != cpy.Uint64s[i] {
			t.Errorf("Uint64s: got %v at index %d of the copy; want %v", cpy.Uint64s[i], i, v)
		}
	}

Float32s:
	if cpy.Float32 != test.Float32 {
		t.Errorf("Float32: got %v; want %v", cpy.Float32, test.Float32)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Float32s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Float32s)).Data {
		t.Error("Float32s: address of copy was the same as original; they should be different")
		goto Float64s
	}
	if len(cpy.Float32s) != len(test.Float32s) {
		t.Errorf("Float32s: len was %d; want %d", len(cpy.Float32s), len(test.Float32s))
		goto Float64s
	}
	for i, v := range test.Float32s {
		if v != cpy.Float32s[i] {
			t.Errorf("Float32s: got %v at index %d of the copy; want %v", cpy.Float32s[i], i, v)
		}
	}

Float64s:
	if cpy.Float64 != test.Float64 {
		t.Errorf("Float64: got %v; want %v", cpy.Float64, test.Float64)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Float64s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Float64s)).Data {
		t.Error("Float64s: address of copy was the same as original; they should be different")
		goto Complex64s
	}
	if len(cpy.Float64s) != len(test.Float64s) {
		t.Errorf("Float64s: len was %d; want %d", len(cpy.Float64s), len(test.Float64s))
		goto Complex64s
	}
	for i, v := range test.Float64s {
		if v != cpy.Float64s[i] {
			t.Errorf("Float64s: got %v at index %d of the copy; want %v", cpy.Float64s[i], i, v)
		}
	}

Complex64s:
	if cpy.Complex64 != test.Complex64 {
		t.Errorf("Complex64: got %v; want %v", cpy.Complex64, test.Complex64)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Complex64s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Complex64s)).Data {
		t.Error("Complex64s: address of copy was the same as original; they should be different")
		goto Complex128s
	}
	if len(cpy.Complex64s) != len(test.Complex64s) {
		t.Errorf("Complex64s: len was %d; want %d", len(cpy.Complex64s), len(test.Complex64s))
		goto Complex128s
	}
	for i, v := range test.Complex64s {
		if v != cpy.Complex64s[i] {
			t.Errorf("Complex64s: got %v at index %d of the copy; want %v", cpy.Complex64s[i], i, v)
		}
	}

Complex128s:
	if cpy.Complex128 != test.Complex128 {
		t.Errorf("Complex128s: got %v; want %v", cpy.Complex128s, test.Complex128s)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Complex128s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Complex128s)).Data {
		t.Error("Complex128s: address of copy was the same as original; they should be different")
		goto Interfaces
	}
	if len(cpy.Complex128s) != len(test.Complex128s) {
		t.Errorf("Complex128s: len was %d; want %d", len(cpy.Complex128s), len(test.Complex128s))
		goto Interfaces
	}
	for i, v := range test.Complex128s {
		if v != cpy.Complex128s[i] {
			t.Errorf("Complex128s: got %v at index %d of the copy; want %v", cpy.Complex128s[i], i, v)
		}
	}

Interfaces:
	if cpy.Interface != test.Interface {
		t.Errorf("Interface: got %v; want %v", cpy.Interface, test.Interface)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Interfaces)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Interfaces)).Data {
		t.Error("Interfaces: address of copy was the same as original; they should be different")
		return
	}
	if len(cpy.Interfaces) != len(test.Interfaces) {
		t.Errorf("Interfaces: len was %d; want %d", len(cpy.Interfaces), len(test.Interfaces))
		return
	}
	for i, v := range test.Interfaces {
		if v != cpy.Interfaces[i] {
			t.Errorf("Interfaces: got %v at index %d of the copy; want %v", cpy.Interfaces[i], i, v)
		}
	}
}

// not meant to be exhaustive
func TestComplexSlices(t *testing.T) {
	orig3Int := [][][]int{{[]int{13}, []int{11, 22}}, {[]int{66, 88, 99}}}
	cpyI := Copy(orig3Int).([][][]int)
	if (*reflect.SliceHeader)(unsafe.Pointer(&orig3Int)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyI)).Data {
		t.Error("[][][]int: address of copy was the same as original; they should be different")
		return
	}
	if len(orig3Int) != len(cpyI) {
		t.Errorf("[][][]int: len of copy was %d; want %d", len(cpyI), len(orig3Int))
		goto sliceMap
	}
	for i, v := range orig3Int {
		if len(v) != len(cpyI[i]) {
			t.Errorf("[][][]int: len of element %d was %d; want %d", i, len(cpyI[i]), len(v))
			continue
		}
		for j, vv := range v {
			if len(vv) != len(cpyI[i][j]) {
				t.Errorf("[][][]int: len of element %d:%d was %d, want %d", i, j, len(cpyI[i][j]), len(vv))
				continue
			}
			for k, vvv := range vv {
				if vvv != cpyI[i][j][k] {
					t.Errorf("[][][]int: element %d:%d:%d was %d, want %d", i, j, k, cpyI[i][j][k], vvv)
				}
			}
		}

	}

sliceMap:
	slMap := []map[int]string{{0: "a", 1: "b"}, {11: "l", 12: "m"}}
	cpyM := Copy(slMap).([]map[int]string)
	if (*reflect.SliceHeader)(unsafe.Pointer(&slMap)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyM)).Data {
		t.Error("[]map[int]string: address of copy was the same as original; they should be different")
	}
	if len(slMap) != len(cpyM) {
		t.Errorf("[]map[int]string: len of copy was %d; want %d", len(cpyM), len(slMap))
		goto done
	}
	for i, v := range slMap {
		if len(v) != len(cpyM[i]) {
			t.Errorf("[]map[int]string: len of element %d was %d; want %d", i, len(cpyM[i]), len(v))
			continue
		}
		for k, vv := range v {
			val, ok := cpyM[i][k]
			if !ok {
				t.Errorf("[]map[int]string: element %d was expected to have a value at key %d, it didn't", i, k)
				continue
			}
			if val != vv {
				t.Errorf("[]map[int]string: element %d, key %d: got %s, want %s", i, k, val, vv)
			}
		}
	}
done:
}

type A struct {
	Int    int
	String string
	UintSl []uint
	NilSl  []string
	Map    map[string]int
	MapB   map[string]*B
	SliceB []B
	B
	T time.Time
}

type B struct {
	Vals []string
}

var AStruct = A{
	Int:    42,
	String: "Konichiwa",
	UintSl: []uint{0, 1, 2, 3},
	Map:    map[string]int{"a": 1, "b": 2},
	MapB: map[string]*B{
		"hi":  {Vals: []string{"hello", "bonjour"}},
		"bye": {Vals: []string{"good-bye", "au revoir"}},
	},
	SliceB: []B{
		{Vals: []string{"Ciao", "Aloha"}},
	},
	B: B{Vals: []string{"42"}},
	T: time.Now(),
}

func TestStructA(t *testing.T) {
	cpy := Copy(AStruct).(A)
	if &cpy == &AStruct {
		t.Error("expected copy to have a different address than the original; it was the same")
		return
	}
	if cpy.Int != AStruct.Int {
		t.Errorf("A.Int: got %v, want %v", cpy.Int, AStruct.Int)
	}
	if cpy.String != AStruct.String {
		t.Errorf("A.String: got %v; want %v", cpy.String, AStruct.String)
	}
	if (*reflect.SliceHeader)(unsafe.Pointer(&cpy.UintSl)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&AStruct.UintSl)).Data {
		t.Error("A.Uintsl: expected the copies address to be different; it wasn't")
		goto NilSl
	}
	if len(cpy.UintSl) != len(AStruct.UintSl) {
		t.Errorf("A.UintSl: got len of %d, want %d", len(cpy.UintSl), len(AStruct.UintSl))
		goto NilSl
	}
	for i, v := range AStruct.UintSl {
		if cpy.UintSl[i] != v {
			t.Errorf("A.UintSl %d: got %d, want %d", i, cpy.UintSl[i], v)
		}
	}

NilSl:
	if cpy.NilSl != nil {
		t.Error("A.NilSl: expected slice to be nil, it wasn't")
	}

	if *(*uintptr)(unsafe.Pointer(&cpy.Map)) == *(*uintptr)(unsafe.Pointer(&AStruct.Map)) {
		t.Error("A.Map: expected the copy's address to be different; it wasn't")
		goto AMapB
	}
	if len(cpy.Map) != len(AStruct.Map) {
		t.Errorf("A.Map: got len of %d, want %d", len(cpy.Map), len(AStruct.Map))
		goto AMapB
	}
	for k, v := range AStruct.Map {
		val, ok := cpy.Map[k]
		if !ok {
			t.Errorf("A.Map: expected the key %s to exist in the copy, it didn't", k)
			continue
		}
		if val != v {
			t.Errorf("A.Map[%s]: got %d, want %d", k, val, v)
		}
	}

AMapB:
	if *(*uintptr)(unsafe.Pointer(&cpy.MapB)) == *(*uintptr)(unsafe.Pointer(&AStruct.MapB)) {
		t.Error("A.MapB: expected the copy's address to be different; it wasn't")
		goto ASliceB
	}
	if len(cpy.MapB) != len(AStruct.MapB) {
		t.Errorf("A.MapB: got len of %d, want %d", len(cpy.MapB), len(AStruct.MapB))
		goto ASliceB
	}
	for k, v := range AStruct.MapB {
		val, ok := cpy.MapB[k]
		if !ok {
			t.Errorf("A.MapB: expected the key %s to exist in the copy, it didn't", k)
			continue
		}
		if unsafe.Pointer(val) == unsafe.Pointer(v) {
			t.Errorf("A.MapB[%s]: expected the addresses of the values to be different; they weren't", k)
			continue
		}
		// the slice headers should point to different data
		if (*reflect.SliceHeader)(unsafe.Pointer(&v.Vals)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&val.Vals)).Data {
			t.Errorf("%s: expected B's SliceHeaders to point to different Data locations; they did not.", k)
			continue
		}
		for i, vv := range v.Vals {
			if vv != val.Vals[i] {
				t.Errorf("A.MapB[%s].Vals[%d]: got %s want %s", k, i, vv, val.Vals[i])
			}
		}
	}

ASliceB:
	if (*reflect.SliceHeader)(unsafe.Pointer(&AStruct.SliceB)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.SliceB)).Data {
		t.Error("A.SliceB: expected the copy's address to be different; it wasn't")
		goto B
	}

	if len(AStruct.SliceB) != len(cpy.SliceB) {
		t.Errorf("A.SliceB: got length of %d; want %d", len(cpy.SliceB), len(AStruct.SliceB))
		goto B
	}

	for i := range AStruct.SliceB {
		if unsafe.Pointer(&AStruct.SliceB[i]) == unsafe.Pointer(&cpy.SliceB[i]) {
			t.Errorf("A.SliceB[%d]: expected them to have different addresses, they didn't", i)
			continue
		}
		if (*reflect.SliceHeader)(unsafe.Pointer(&AStruct.SliceB[i].Vals)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.SliceB[i].Vals)).Data {
			t.Errorf("A.SliceB[%d]: expected B.Vals SliceHeader.Data to point to different locations; they did not", i)
			continue
		}
		if len(AStruct.SliceB[i].Vals) != len(cpy.SliceB[i].Vals) {
			t.Errorf("A.SliceB[%d]: expected B's vals to have the same length, they didn't", i)
			continue
		}
		for j, val := range AStruct.SliceB[i].Vals {
			if val != cpy.SliceB[i].Vals[j] {
				t.Errorf("A.SliceB[%d].Vals[%d]: got %v; want %v", i, j, cpy.SliceB[i].Vals[j], val)
			}
		}
	}
B:
	if unsafe.Pointer(&AStruct.B) == unsafe.Pointer(&cpy.B) {
		t.Error("A.B: expected them to have different addresses, they didn't")
		goto T
	}
	if (*reflect.SliceHeader)(unsafe.Pointer(&AStruct.B.Vals)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.B.Vals)).Data {
		t.Error("A.B.Vals: expected the SliceHeaders.Data to point to different locations; they didn't")
		goto T
	}
	if len(AStruct.B.Vals) != len(cpy.B.Vals) {
		t.Error("A.B.Vals: expected their lengths to be the same, they weren't")
		goto T
	}
	for i, v := range AStruct.B.Vals {
		if v != cpy.B.Vals[i] {
			t.Errorf("A.B.Vals[%d]: got %s want %s", i, cpy.B.Vals[i], v)
		}
	}
T:
	if fmt.Sprintf("%p", &AStruct.T) == fmt.Sprintf("%p", &cpy.T) {
		t.Error("A.T: expected them to have different addresses, they didn't")
		return
	}
	if AStruct.T != cpy.T {
		t.Errorf("A.T: got %v, want %v", cpy.T, AStruct.T)
	}
}

type Unexported struct {
	A  string
	B  int
	aa string
	bb int
	cc []int
	dd map[string]string
}

func TestUnexportedFields(t *testing.T) {
	u := &Unexported{
		A:  "A",
		B:  42,
		aa: "aa",
		bb: 42,
		cc: []int{1, 2, 3},
		dd: map[string]string{"hello": "bonjour"},
	}
	cpy := Copy(u).(*Unexported)
	if cpy == u {
		t.Error("expected addresses to be different, they weren't")
		return
	}
	if u.A != cpy.A {
		t.Errorf("Unexported.A: got %s want %s", cpy.A, u.A)
	}
	if u.B != cpy.B {
		t.Errorf("Unexported.A: got %d want %d", cpy.B, u.B)
	}
	if cpy.aa != "" {
		t.Errorf("Unexported.aa: unexported field should not be set, it was set to %s", cpy.aa)
	}
	if cpy.bb != 0 {
		t.Errorf("Unexported.bb: unexported field should not be set, it was set to %d", cpy.bb)
	}
	if cpy.cc != nil {
		t.Errorf("Unexported.cc: unexported field should not be set, it was set to %#v", cpy.cc)
	}
	if cpy.dd != nil {
		t.Errorf("Unexported.dd: unexported field should not be set, it was set to %#v", cpy.dd)
	}
}

// Note: this test will fail until https://github.com/golang/go/issues/15716 is
// fixed and the version it is part of gets released.
type T struct {
	time.Time
}

func TestTimeCopy(t *testing.T) {
	tests := []struct {
		Y    int
		M    time.Month
		D    int
		h    int
		m    int
		s    int
		nsec int
		TZ   string
	}{
		{2016, time.July, 4, 23, 11, 33, 3000, "America/New_York"},
		{2015, time.October, 31, 9, 44, 23, 45935, "UTC"},
		{2014, time.May, 5, 22, 01, 50, 219300, "Europe/Prague"},
	}

	for i, test := range tests {
		l, err := time.LoadLocation(test.TZ)
		if err != nil {
			t.Errorf("%d: unexpected error: %s", i, err)
			continue
		}
		var x T
		x.Time = time.Date(test.Y, test.M, test.D, test.h, test.m, test.s, test.nsec, l)
		c := Copy(x).(T)
		if fmt.Sprintf("%p", &c) == fmt.Sprintf("%p", &x) {
			t.Errorf("%d: expected the copy to have a different address than the original value; they were the same: %p %p", i, &c, &x)
			continue
		}
		if x.UnixNano() != c.UnixNano() {
			t.Errorf("%d: nanotime: got %v; want %v", i, c.UnixNano(), x.UnixNano())
			continue
		}
		if x.Location() != c.Location() {
			t.Errorf("%d: location: got %q; want %q", i, c.Location(), x.Location())
		}
	}
}

func TestPointerToStruct(t *testing.T) {
	type Foo struct {
		Bar int
	}

	f := &Foo{Bar: 42}
	cpy := Copy(f)
	if f == cpy {
		t.Errorf("expected copy to point to a different location: orig: %p; copy: %p", f, cpy)
	}
	if !reflect.DeepEqual(f, cpy) {
		t.Errorf("expected the copy to be equal to the original (except for memory location); it wasn't: got %#v; want %#v", f, cpy)
	}
}

func TestIssue9(t *testing.T) {
	// simple pointer copy
	x := 42
	testA := map[string]*int{
		"a": nil,
		"b": &x,
	}
	copyA := Copy(testA).(map[string]*int)
	if unsafe.Pointer(&testA) == unsafe.Pointer(&copyA) {
		t.Fatalf("expected the map pointers to be different: testA: %v\tcopyA: %v", unsafe.Pointer(&testA), unsafe.Pointer(&copyA))
	}
	if !reflect.DeepEqual(testA, copyA) {
		t.Errorf("got %#v; want %#v", copyA, testA)
	}
	if testA["b"] == copyA["b"] {
		t.Errorf("entries for 'b' pointed to the same address: %v; expected them to point to different addresses", testA["b"])
	}

	// map copy
	type Foo struct {
		Alpha string
	}

	type Bar struct {
		Beta  string
		Gamma int
		Delta *Foo
	}

	type Biz struct {
		Epsilon map[int]*Bar
	}

	testB := Biz{
		Epsilon: map[int]*Bar{
			0: {},
			1: {
				Beta:  "don't panic",
				Gamma: 42,
				Delta: nil,
			},
			2: {
				Beta:  "sudo make me a sandwich.",
				Gamma: 11,
				Delta: &Foo{
					Alpha: "okay.",
				},
			},
		},
	}

	copyB := Copy(testB).(Biz)
	if !reflect.DeepEqual(testB, copyB) {
		t.Errorf("got %#v; want %#v", copyB, testB)
		return
	}

	// check that the maps point to different locations
	if unsafe.Pointer(&testB.Epsilon) == unsafe.Pointer(&copyB.Epsilon) {
		t.Fatalf("expected the map pointers to be different; they weren't: testB: %v\tcopyB: %v", unsafe.Pointer(&testB.Epsilon), unsafe.Pointer(&copyB.Epsilon))
	}

	for k, v := range testB.Epsilon {
		if v == nil && copyB.Epsilon[k] == nil {
			continue
		}
		if v == nil && copyB.Epsilon[k] != nil {
			t.Errorf("%d: expected copy of a nil entry to be nil; it wasn't: %#v", k, copyB.Epsilon[k])
			continue
		}
		if v == copyB.Epsilon[k] {
			t.Errorf("entries for '%d' pointed to the same address: %v; expected them to point to different addresses", k, v)
			continue
		}
		if v.Beta != copyB.Epsilon[k].Beta {
			t.Errorf("%d.Beta: got %q; want %q", k, copyB.Epsilon[k].Beta, v.Beta)
		}
		if v.Gamma != copyB.Epsilon[k].Gamma {
			t.Errorf("%d.Gamma: got %d; want %d", k, copyB.Epsilon[k].Gamma, v.Gamma)
		}
		if v.Delta == nil && copyB.Epsilon[k].Delta == nil {
			continue
		}
		if v.Delta == nil && copyB.Epsilon[k].Delta != nil {
			t.Errorf("%d.Delta: got %#v; want nil", k, copyB.Epsilon[k].Delta)
		}
		if v.Delta == copyB.Epsilon[k].Delta {
			t.Errorf("%d.Delta: expected the pointers to be different, they were the same: %v", k, v.Delta)
			continue
		}
		if v.Delta.Alpha != copyB.Epsilon[k].Delta.Alpha {
			t.Errorf("%d.Delta.Foo: got %q; want %q", k, v.Delta.Alpha, copyB.Epsilon[k].Delta.Alpha)
		}
	}

	// test that map keys are deep copied
	testC := map[*Foo][]string{
		{Alpha: "Henry Dorsett Case"}: []string{
			"Cutter",
		},
		{Alpha: "Molly Millions"}: []string{
			"Rose Kolodny",
			"Cat Mother",
			"Steppin' Razor",
		},
	}

	copyC := Copy(testC).(map[*Foo][]string)
	if unsafe.Pointer(&testC) == unsafe.Pointer(&copyC) {
		t.Fatalf("expected the map pointers to be different; they weren't: testB: %v\tcopyB: %v", unsafe.Pointer(&testB.Epsilon), unsafe.Pointer(&copyB.Epsilon))
	}

	// make sure the lengths are the same
	if len(testC) != len(copyC) {
		t.Fatalf("got len %d; want %d", len(copyC), len(testC))
	}

	// check that everything was deep copied: since the key is a pointer, we check to
	// see if the pointers are different but the values being pointed to are the same.
	for k, v := range testC {
		for kk, vv := range copyC {
			if *kk == *k {
				if kk == k {
					t.Errorf("key pointers should be different: orig: %p; copy: %p", k, kk)
				}
				// check that the slices are the same but different
				if !reflect.DeepEqual(v, vv) {
					t.Errorf("expected slice contents to be the same; they weren't: orig: %v; copy: %v", v, vv)
				}

				if (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&vv)).Data {
					t.Errorf("expected the SliceHeaders.Data to point to different locations; they didn't: %v", (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data)
				}
				break
			}
		}
	}

	type Bizz struct {
		*Foo
	}

	testD := map[Bizz]string{
		{&Foo{"Neuromancer"}}: "Rio",
		{&Foo{"Wintermute"}}:  "Berne",
	}
	copyD := Copy(testD).(map[Bizz]string)
	if len(copyD) != len(testD) {
		t.Fatalf("copy had %d elements; expected %d", len(copyD), len(testD))
	}

	for k, v := range testD {
		var found bool
		for kk, vv := range copyD {
			if reflect.DeepEqual(k, kk) {
				found = true
				// check that Foo points to different locations
				if unsafe.Pointer(k.Foo) == unsafe.Pointer(kk.Foo) {
					t.Errorf("Expected Foo to point to different locations; they didn't: orig: %p; copy %p", k.Foo, kk.Foo)
					break
				}
				if *k.Foo != *kk.Foo {
					t.Errorf("Expected copy of the key's Foo field to have the same value as the original, it wasn't: orig: %#v; copy: %#v", k.Foo, kk.Foo)
				}
				if v != vv {
					t.Errorf("Expected the values to be the same; the weren't: got %v; want %v", vv, v)
				}
			}
		}
		if !found {
			t.Errorf("expected key %v to exist in the copy; it didn't", k)
		}
	}
}

type I struct {
	A string
}

func (i *I) DeepCopy() interface{} {
	return &I{A: "custom copy"}
}

type NestI struct {
	I *I
}

func TestInterface(t *testing.T) {
	i := &I{A: "A"}
	copied := Copy(i).(*I)
	if copied.A != "custom copy" {
		t.Errorf("expected value %v, but it's %v", "custom copy", copied.A)
	}
	// check for nesting values
	ni := &NestI{I: &I{A: "A"}}
	copiedNest := Copy(ni).(*NestI)
	if copiedNest.I.A != "custom copy" {
		t.Errorf("expected value %v, but it's %v", "custom copy", copiedNest.I.A)
	}
}

// TestPrimitiveTypes tests deep copying of primitive types
func TestPrimitiveTypes(t *testing.T) {
	tests := []struct {
		name     string
		original interface{}
		modifier func(interface{}) interface{}
	}{
		{
			name:     "int",
			original: 42,
			modifier: func(v interface{}) interface{} { return v.(int) + 1 },
		},
		{
			name:     "int8",
			original: int8(8),
			modifier: func(v interface{}) interface{} { return v.(int8) + 1 },
		},
		{
			name:     "int16", 
			original: int16(16),
			modifier: func(v interface{}) interface{} { return v.(int16) + 1 },
		},
		{
			name:     "int32",
			original: int32(32),
			modifier: func(v interface{}) interface{} { return v.(int32) + 1 },
		},
		{
			name:     "int64",
			original: int64(64),
			modifier: func(v interface{}) interface{} { return v.(int64) + 1 },
		},
		{
			name:     "uint",
			original: uint(42),
			modifier: func(v interface{}) interface{} { return v.(uint) + 1 },
		},
		{
			name:     "uint8",
			original: uint8(8),
			modifier: func(v interface{}) interface{} { return v.(uint8) + 1 },
		},
		{
			name:     "uint16",
			original: uint16(16),
			modifier: func(v interface{}) interface{} { return v.(uint16) + 1 },
		},
		{
			name:     "uint32",
			original: uint32(32),
			modifier: func(v interface{}) interface{} { return v.(uint32) + 1 },
		},
		{
			name:     "uint64",
			original: uint64(64),
			modifier: func(v interface{}) interface{} { return v.(uint64) + 1 },
		},
		{
			name:     "float32",
			original: float32(3.14),
			modifier: func(v interface{}) interface{} { return v.(float32) + 1.0 },
		},
		{
			name:     "float64",
			original: 3.14159,
			modifier: func(v interface{}) interface{} { return v.(float64) + 1.0 },
		},
		{
			name:     "complex64",
			original: complex64(1 + 2i),
			modifier: func(v interface{}) interface{} { return v.(complex64) + complex64(1+1i) },
		},
		{
			name:     "complex128",
			original: complex128(2 + 3i),
			modifier: func(v interface{}) interface{} { return v.(complex128) + complex128(1+1i) },
		},
		{
			name:     "string",
			original: "hello",
			modifier: func(v interface{}) interface{} { return v.(string) + " world" },
		},
		{
			name:     "bool true",
			original: true,
			modifier: func(v interface{}) interface{} { return !v.(bool) },
		},
		{
			name:     "bool false",
			original: false,
			modifier: func(v interface{}) interface{} { return !v.(bool) },
		},
		{
			name:     "byte",
			original: byte('A'),
			modifier: func(v interface{}) interface{} { return byte(v.(byte) + 1) },
		},
		{
			name:     "rune",
			original: rune('世'),
			modifier: func(v interface{}) interface{} { return rune(v.(rune) + 1) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			original1 := test.original
			original2 := test.original
			
			// Deep copy one of them
			copied := Copy(original1)
			
			// Modify the original
			original1 = test.modifier(original1)
			
			// Assert that the copy remains unchanged
			if !reflect.DeepEqual(copied, original2) {
				t.Errorf("%s: copy was modified when original changed. Got %v, want %v", test.name, copied, original2)
			}
		})
	}
}

// TestCollections tests deep copying of various collections
func TestCollections(t *testing.T) {
	t.Run("slice of ints", func(t *testing.T) {
		original1 := []int{1, 2, 3, 4, 5}
		original2 := []int{1, 2, 3, 4, 5}
		
		// Deep copy
		copied := Copy(original1).([]int)
		
		// Modify original
		original1[0] = 999
		original1 = append(original1, 6)
		
		// Assert copy unchanged
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("slice copy was modified. Got %v, want %v", copied, original2)
		}
		if len(copied) != len(original2) {
			t.Errorf("slice length changed. Got %d, want %d", len(copied), len(original2))
		}
	})

	t.Run("slice of strings", func(t *testing.T) {
		original1 := []string{"hello", "world", "test"}
		original2 := []string{"hello", "world", "test"}
		
		copied := Copy(original1).([]string)
		
		original1[0] = "modified"
		original1 = append(original1, "new")
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("string slice copy was modified. Got %v, want %v", copied, original2)
		}
	})

	t.Run("slice of slices", func(t *testing.T) {
		original1 := [][]int{{1, 2}, {3, 4}, {5, 6}}
		original2 := [][]int{{1, 2}, {3, 4}, {5, 6}}
		
		copied := Copy(original1).([][]int)
		
		original1[0][0] = 999
		original1[0] = append(original1[0], 999)
		original1 = append(original1, []int{7, 8})
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("nested slice copy was modified. Got %v, want %v", copied, original2)
		}
	})

	t.Run("array", func(t *testing.T) {
		original1 := [5]int{1, 2, 3, 4, 5}
		original2 := [5]int{1, 2, 3, 4, 5}
		
		copied := Copy(original1).([5]int)
		
		original1[0] = 999
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("array copy was modified. Got %v, want %v", copied, original2)
		}
	})

	t.Run("map string to int", func(t *testing.T) {
		original1 := map[string]int{"one": 1, "two": 2, "three": 3}
		original2 := map[string]int{"one": 1, "two": 2, "three": 3}
		
		copied := Copy(original1).(map[string]int)
		
		original1["one"] = 999
		original1["four"] = 4
		delete(original1, "two")
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("map copy was modified. Got %v, want %v", copied, original2)
		}
	})

	t.Run("map with slice values", func(t *testing.T) {
		original1 := map[string][]int{
			"evens": {2, 4, 6},
			"odds":  {1, 3, 5},
		}
		original2 := map[string][]int{
			"evens": {2, 4, 6},
			"odds":  {1, 3, 5},
		}
		
		copied := Copy(original1).(map[string][]int)
		
		original1["evens"][0] = 999
		original1["evens"] = append(original1["evens"], 8)
		original1["new"] = []int{10, 11}
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("map with slice values copy was modified. Got %v, want %v", copied, original2)
		}
	})

	t.Run("nested maps", func(t *testing.T) {
		original1 := map[string]map[string]int{
			"group1": {"a": 1, "b": 2},
			"group2": {"c": 3, "d": 4},
		}
		original2 := map[string]map[string]int{
			"group1": {"a": 1, "b": 2},
			"group2": {"c": 3, "d": 4},
		}
		
		copied := Copy(original1).(map[string]map[string]int)
		
		original1["group1"]["a"] = 999
		original1["group1"]["new"] = 100
		original1["group3"] = map[string]int{"e": 5}
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("nested map copy was modified. Got %v, want %v", copied, original2)
		}
	})
}

// TestCustomStructs tests deep copying of custom structs
func TestCustomStructs(t *testing.T) {
	type Person struct {
		Name   string
		Age    int
		Active bool
	}

	t.Run("simple struct", func(t *testing.T) {
		original1 := Person{Name: "Alice", Age: 30, Active: true}
		original2 := Person{Name: "Alice", Age: 30, Active: true}
		
		copied := Copy(original1).(Person)
		
		original1.Name = "Bob"
		original1.Age = 25
		original1.Active = false
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("struct copy was modified. Got %v, want %v", copied, original2)
		}
	})

	type Employee struct {
		Person
		ID       int
		Skills   []string
		Projects map[string]bool
	}

	t.Run("struct with embedded struct and collections", func(t *testing.T) {
		original1 := Employee{
			Person:   Person{Name: "Charlie", Age: 35, Active: true},
			ID:       123,
			Skills:   []string{"Go", "Python", "JavaScript"},
			Projects: map[string]bool{"proj1": true, "proj2": false},
		}
		original2 := Employee{
			Person:   Person{Name: "Charlie", Age: 35, Active: true},
			ID:       123,
			Skills:   []string{"Go", "Python", "JavaScript"},
			Projects: map[string]bool{"proj1": true, "proj2": false},
		}
		
		copied := Copy(original1).(Employee)
		
		// Modify embedded struct
		original1.Person.Name = "Dave"
		original1.Person.Age = 40
		// Modify slice
		original1.Skills[0] = "Rust"
		original1.Skills = append(original1.Skills, "C++")
		// Modify map
		original1.Projects["proj1"] = false
		original1.Projects["proj3"] = true
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("complex struct copy was modified. Got %v, want %v", copied, original2)
		}
	})

	type Node struct {
		Value int
		Next  *Node
	}

	t.Run("struct with pointer to same type", func(t *testing.T) {
		node2 := &Node{Value: 2, Next: nil}
		original1 := &Node{Value: 1, Next: node2}
		
		// Create identical structure for comparison
		node2Copy := &Node{Value: 2, Next: nil}
		original2 := &Node{Value: 1, Next: node2Copy}
		
		copied := Copy(original1).(*Node)
		
		// Modify original
		original1.Value = 999
		original1.Next.Value = 888
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("linked structure copy was modified. Got %v, want %v", copied, original2)
		}
		
		// Verify that pointers are different
		if copied == original1 {
			t.Error("copied root node points to same address as original")
		}
		if copied.Next == original1.Next {
			t.Error("copied next node points to same address as original")
		}
	})
}

// TestNestedStructures tests deep copying of complex nested structures
func TestNestedStructures(t *testing.T) {
	type Address struct {
		Street  string
		City    string
		Country string
	}
	
	type Contact struct {
		Email   string
		Phone   string
		Address *Address
	}
	
	type Company struct {
		Name      string
		Employees []Contact
		Locations map[string]*Address
	}

	t.Run("deeply nested structures", func(t *testing.T) {
		addr1 := &Address{Street: "123 Main St", City: "New York", Country: "USA"}
		addr2 := &Address{Street: "456 Oak Ave", City: "Los Angeles", Country: "USA"}
		
		original1 := Company{
			Name: "TechCorp",
			Employees: []Contact{
				{Email: "john@techcorp.com", Phone: "555-1234", Address: addr1},
				{Email: "jane@techcorp.com", Phone: "555-5678", Address: addr2},
			},
			Locations: map[string]*Address{
				"HQ": addr1,
				"West": addr2,
			},
		}
		
		// Create identical structure for comparison
		addr1Copy := &Address{Street: "123 Main St", City: "New York", Country: "USA"}
		addr2Copy := &Address{Street: "456 Oak Ave", City: "Los Angeles", Country: "USA"}
		
		original2 := Company{
			Name: "TechCorp",
			Employees: []Contact{
				{Email: "john@techcorp.com", Phone: "555-1234", Address: addr1Copy},
				{Email: "jane@techcorp.com", Phone: "555-5678", Address: addr2Copy},
			},
			Locations: map[string]*Address{
				"HQ": addr1Copy,
				"West": addr2Copy,
			},
		}
		
		copied := Copy(original1).(Company)
		
		// Modify nested structures
		original1.Name = "NewCorp"
		original1.Employees[0].Email = "changed@email.com"
		original1.Employees[0].Address.Street = "999 Changed St"
		original1.Employees = append(original1.Employees, Contact{Email: "new@email.com", Phone: "999-0000", Address: nil})
		original1.Locations["East"] = &Address{Street: "789 East St", City: "Miami", Country: "USA"}
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("deeply nested structure copy was modified. Got %v, want %v", copied, original2)
		}
		
		// Verify address pointers are different
		if copied.Employees[0].Address == original1.Employees[0].Address {
			t.Error("nested address pointer should be different")
		}
		if copied.Locations["HQ"] == original1.Locations["HQ"] {
			t.Error("map value pointer should be different")
		}
	})

	type TreeNode struct {
		Value    int
		Children []*TreeNode
		Metadata map[string]interface{}
	}

	t.Run("tree structure with interface values", func(t *testing.T) {
		child1 := &TreeNode{
			Value: 2,
			Metadata: map[string]interface{}{
				"type": "leaf",
				"data": []int{1, 2, 3},
			},
		}
		child2 := &TreeNode{
			Value: 3,
			Metadata: map[string]interface{}{
				"type": "leaf",
				"data": "text data",
			},
		}
		
		original1 := &TreeNode{
			Value:    1,
			Children: []*TreeNode{child1, child2},
			Metadata: map[string]interface{}{
				"type":  "root",
				"count": 2,
			},
		}
		
		// Create identical structure
		child1Copy := &TreeNode{
			Value: 2,
			Metadata: map[string]interface{}{
				"type": "leaf",
				"data": []int{1, 2, 3},
			},
		}
		child2Copy := &TreeNode{
			Value: 3,
			Metadata: map[string]interface{}{
				"type": "leaf",
				"data": "text data",
			},
		}
		
		original2 := &TreeNode{
			Value:    1,
			Children: []*TreeNode{child1Copy, child2Copy},
			Metadata: map[string]interface{}{
				"type":  "root",
				"count": 2,
			},
		}
		
		copied := Copy(original1).(*TreeNode)
		
		// Modify tree structure
		original1.Value = 999
		original1.Children[0].Value = 888
		original1.Children[0].Metadata["type"] = "modified"
		original1.Metadata["count"] = 999
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("tree structure copy was modified. Got %v, want %v", copied, original2)
		}
		
		// Verify pointers are different
		if copied == original1 {
			t.Error("root nodes should have different addresses")
		}
		if copied.Children[0] == original1.Children[0] {
			t.Error("child nodes should have different addresses")
		}
	})
}

// Test types for interfaces
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// TestPointersAndInterfaces tests deep copying of pointers and interface types
func TestPointersAndInterfaces(t *testing.T) {
	t.Run("pointer to primitive", func(t *testing.T) {
		value1 := 42
		value2 := 42
		original1 := &value1
		original2 := &value2
		
		copied := Copy(original1).(*int)
		
		*original1 = 999
		
		if *copied != *original2 {
			t.Errorf("pointer copy was modified. Got %v, want %v", *copied, *original2)
		}
		
		// Verify different addresses
		if copied == original1 {
			t.Error("copied pointer should have different address")
		}
	})

	t.Run("slice of pointers", func(t *testing.T) {
		a1, b1, c1 := 1, 2, 3
		a2, b2, c2 := 1, 2, 3
		original1 := []*int{&a1, &b1, &c1}
		original2 := []*int{&a2, &b2, &c2}
		
		copied := Copy(original1).([]*int)
		
		*original1[0] = 999
		original1[1] = nil
		
		if *copied[0] != *original2[0] {
			t.Errorf("pointer in slice was modified. Got %v, want %v", *copied[0], *original2[0])
		}
		if copied[1] == nil {
			t.Error("pointer in copied slice should not be nil")
		}
		
		// Verify different pointer addresses
		if copied[0] == original1[0] {
			t.Error("pointers in slice should have different addresses")
		}
	})

	t.Run("interface slice", func(t *testing.T) {
		original1 := []Shape{
			Rectangle{Width: 10, Height: 5},
			Circle{Radius: 3},
		}
		original2 := []Shape{
			Rectangle{Width: 10, Height: 5},
			Circle{Radius: 3},
		}
		
		copied := Copy(original1).([]Shape)
		
		// Modify original (this creates new interface values)
		original1[0] = Rectangle{Width: 20, Height: 10}
		original1 = append(original1, Circle{Radius: 5})
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("interface slice copy was modified. Got %v, want %v", copied, original2)
		}
		
		// Test that copied interfaces work correctly
		if copied[0].Area() != 50.0 {
			t.Errorf("copied interface method failed. Got %v, want 50.0", copied[0].Area())
		}
	})

	t.Run("map with interface values", func(t *testing.T) {
		original1 := map[string]interface{}{
			"string": "hello",
			"int":    42,
			"slice":  []int{1, 2, 3},
			"struct": Rectangle{Width: 5, Height: 4},
		}
		original2 := map[string]interface{}{
			"string": "hello",
			"int":    42,
			"slice":  []int{1, 2, 3},
			"struct": Rectangle{Width: 5, Height: 4},
		}
		
		copied := Copy(original1).(map[string]interface{})
		
		// Modify original
		original1["string"] = "modified"
		original1["int"] = 999
		original1["slice"] = []int{9, 8, 7}
		original1["new"] = "new value"
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("interface map copy was modified. Got %v, want %v", copied, original2)
		}
		
		// Verify that slice in interface was deep copied
		originalSlice := original1["slice"].([]int)
		copiedSlice := copied["slice"].([]int)
		
		if reflect.ValueOf(originalSlice).Pointer() == reflect.ValueOf(copiedSlice).Pointer() {
			t.Error("slice in interface should have different memory addresses")
		}
	})

	t.Run("pointer to interface", func(t *testing.T) {
		var shape Shape = Rectangle{Width: 8, Height: 6}
		original1 := &shape
		
		var shapeCopy Shape = Rectangle{Width: 8, Height: 6}
		original2 := &shapeCopy
		
		copied := Copy(original1).(*Shape)
		
		// Modify original
		*original1 = Circle{Radius: 10}
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("pointer to interface copy was modified. Got %v, want %v", *copied, *original2)
		}
		
		// Verify different addresses
		if copied == original1 {
			t.Error("copied pointer should have different address")
		}
		
		// Test interface method on copy
		if (*copied).Area() != 48.0 {
			t.Errorf("copied interface method failed. Got %v, want 48.0", (*copied).Area())
		}
	})
}

// TestEdgeCases tests edge cases like nil values and empty collections
func TestEdgeCases(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		var original1 *int
		var original2 *int
		
		copied := Copy(original1).(*int)
		
		if copied != original2 {
			t.Errorf("nil pointer copy should be nil. Got %v, want %v", copied, original2)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var original1 []int
		var original2 []int
		
		copied := Copy(original1).([]int)
		
		if copied != nil {
			t.Errorf("nil slice copy should be nil. Got %v, want %v", copied, original2)
		}
	})

	t.Run("nil map", func(t *testing.T) {
		var original1 map[string]int
		var original2 map[string]int
		
		copied := Copy(original1).(map[string]int)
		
		if copied != nil {
			t.Errorf("nil map copy should be nil. Got %v, want %v", copied, original2)
		}
	})

	t.Run("nil interface", func(t *testing.T) {
		var original1 interface{}
		var original2 interface{}
		
		copied := Copy(original1)
		
		if copied != original2 {
			t.Errorf("nil interface copy should be nil. Got %v, want %v", copied, original2)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		original1 := []int{}
		original2 := []int{}
		
		copied := Copy(original1).([]int)
		
		original1 = append(original1, 1)
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("empty slice copy was modified. Got %v, want %v", copied, original2)
		}
		if len(copied) != 0 {
			t.Errorf("empty slice should remain empty. Got len %d", len(copied))
		}
	})

	t.Run("empty map", func(t *testing.T) {
		original1 := make(map[string]int)
		original2 := make(map[string]int)
		
		copied := Copy(original1).(map[string]int)
		
		original1["key"] = 123
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("empty map copy was modified. Got %v, want %v", copied, original2)
		}
		if len(copied) != 0 {
			t.Errorf("empty map should remain empty. Got len %d", len(copied))
		}
	})

	t.Run("slice with nil elements", func(t *testing.T) {
		original1 := []*int{nil, nil}
		original2 := []*int{nil, nil}
		
		copied := Copy(original1).([]*int)
		
		value := 42
		original1[0] = &value
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("slice with nil elements copy was modified. Got %v, want %v", copied, original2)
		}
		if copied[0] != nil || copied[1] != nil {
			t.Error("copied slice should maintain nil elements")
		}
	})

	t.Run("map with nil values", func(t *testing.T) {
		original1 := map[string]*int{
			"nil1": nil,
			"nil2": nil,
		}
		original2 := map[string]*int{
			"nil1": nil,
			"nil2": nil,
		}
		
		copied := Copy(original1).(map[string]*int)
		
		value := 42
		original1["nil1"] = &value
		original1["new"] = &value
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("map with nil values copy was modified. Got %v, want %v", copied, original2)
		}
		if copied["nil1"] != nil || copied["nil2"] != nil {
			t.Error("copied map should maintain nil values")
		}
	})

	type StructWithNils struct {
		Ptr    *int
		Slice  []string
		Map    map[string]int
		Iface  interface{}
	}

	t.Run("struct with nil fields", func(t *testing.T) {
		original1 := StructWithNils{
			Ptr:   nil,
			Slice: nil,
			Map:   nil,
			Iface: nil,
		}
		original2 := StructWithNils{
			Ptr:   nil,
			Slice: nil,
			Map:   nil,
			Iface: nil,
		}
		
		copied := Copy(original1).(StructWithNils)
		
		// Modify original
		value := 123
		original1.Ptr = &value
		original1.Slice = []string{"test"}
		original1.Map = map[string]int{"key": 1}
		original1.Iface = "interface"
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("struct with nil fields copy was modified. Got %v, want %v", copied, original2)
		}
		
		// Verify all fields are nil
		if copied.Ptr != nil || copied.Slice != nil || copied.Map != nil || copied.Iface != nil {
			t.Error("copied struct should maintain nil fields")
		}
	})

	t.Run("zero values", func(t *testing.T) {
		type ZeroStruct struct {
			Int     int
			String  string
			Bool    bool
			Float   float64
			Complex complex128
		}
		
		original1 := ZeroStruct{}
		original2 := ZeroStruct{}
		
		copied := Copy(original1).(ZeroStruct)
		
		// Modify original
		original1.Int = 42
		original1.String = "modified"
		original1.Bool = true
		original1.Float = 3.14
		original1.Complex = 1 + 2i
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("zero value struct copy was modified. Got %v, want %v", copied, original2)
		}
	})
}

// Custom types for testing
type Event struct {
	Name      string
	Timestamp time.Time
	Duration  time.Duration
}

type CustomCopyable struct {
	Value string
	Count int
}

func (c *CustomCopyable) DeepCopy() interface{} {
	return &CustomCopyable{
		Value: c.Value + "_copied",
		Count: c.Count + 1,
	}
}

type GenericCopyable struct {
	Data string
}

func (g *GenericCopyable) DeepCopy() *GenericCopyable {
	return &GenericCopyable{Data: g.Data + "_generic"}
}

// TestSpecialTypes tests time.Time and custom DeepCopy interface
func TestSpecialTypes(t *testing.T) {
	t.Run("time.Time", func(t *testing.T) {
		original1 := time.Date(2023, time.December, 25, 10, 30, 0, 123456789, time.UTC)
		original2 := time.Date(2023, time.December, 25, 10, 30, 0, 123456789, time.UTC)
		
		copied := Copy(original1).(time.Time)
		
		// Modify original (this doesn't actually modify since time.Time is immutable,
		// but we test the pattern anyway)
		original1 = original1.Add(time.Hour)
		
		if !copied.Equal(original2) {
			t.Errorf("time.Time copy was modified. Got %v, want %v", copied, original2)
		}
		
		// Verify exact equality including nanoseconds and location
		if copied != original2 {
			t.Errorf("time.Time should be exactly equal. Got %v, want %v", copied, original2)
		}
	})

	t.Run("struct with time.Time", func(t *testing.T) {
		original1 := Event{
			Name:      "Meeting",
			Timestamp: time.Date(2023, time.June, 15, 14, 30, 0, 0, time.UTC),
			Duration:  2 * time.Hour,
		}
		original2 := Event{
			Name:      "Meeting",
			Timestamp: time.Date(2023, time.June, 15, 14, 30, 0, 0, time.UTC),
			Duration:  2 * time.Hour,
		}
		
		copied := Copy(original1).(Event)
		
		// Modify original
		original1.Name = "Modified Meeting"
		original1.Timestamp = time.Now()
		original1.Duration = 1 * time.Hour
		
		if !reflect.DeepEqual(copied, original2) {
			t.Errorf("struct with time.Time copy was modified. Got %v, want %v", copied, original2)
		}
	})

	t.Run("custom DeepCopy interface", func(t *testing.T) {
		original1 := &CustomCopyable{Value: "test", Count: 5}
		
		copied := Copy(original1).(*CustomCopyable)
		
		// The custom DeepCopy should have been called
		if copied.Value != "test_copied" {
			t.Errorf("custom DeepCopy not called. Got %s, want test_copied", copied.Value)
		}
		if copied.Count != 6 {
			t.Errorf("custom DeepCopy not called. Got %d, want 6", copied.Count)
		}
		
		// Verify different addresses
		if copied == original1 {
			t.Error("custom copy should have different address")
		}
		
		// Modify original to ensure copy is independent
		original1.Value = "modified"
		original1.Count = 999
		
		if copied.Value == "modified" || copied.Count == 999 {
			t.Error("custom copy was affected by original modification")
		}
	})

	t.Run("generic Copier interface", func(t *testing.T) {
		original1 := &GenericCopyable{Data: "generic_test"}
		
		// Note: The current implementation only checks for the Interface type,
		// not the generic Copier[T] interface, so this test verifies the current behavior
		copied := Copy(original1).(*GenericCopyable)
		
		// Since it doesn't implement Interface, it should do regular deep copy
		if copied.Data != "generic_test" {
			t.Errorf("expected regular copy, got %s", copied.Data)
		}
		
		// Verify different addresses
		if copied == original1 {
			t.Error("copy should have different address")
		}
	})

	t.Run("nil custom interface", func(t *testing.T) {
		var original1 *CustomCopyable
		
		copied := Copy(original1).(*CustomCopyable)
		
		if copied != nil {
			t.Errorf("nil custom interface copy should be nil. Got %v", copied)
		}
	})
}

// TestCopyTo tests the CopyTo function that copies from src to dst
func TestCopyTo(t *testing.T) {
	t.Run("copy primitive to existing variable", func(t *testing.T) {
		src := 42
		var dst int
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed: %v", err)
		}
		
		if dst != 42 {
			t.Errorf("CopyTo failed. Got %v, want 42", dst)
		}
		
		// Modify src to ensure independence
		src = 999
		if dst != 42 {
			t.Errorf("dst was affected by src change. Got %v, want 42", dst)
		}
	})

	t.Run("copy slice to existing variable", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5}
		var dst []int
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed: %v", err)
		}
		
		if !reflect.DeepEqual(dst, []int{1, 2, 3, 4, 5}) {
			t.Errorf("CopyTo failed. Got %v, want [1 2 3 4 5]", dst)
		}
		
		// Modify src to ensure independence
		src[0] = 999
		src = append(src, 6)
		
		if !reflect.DeepEqual(dst, []int{1, 2, 3, 4, 5}) {
			t.Errorf("dst was affected by src change. Got %v", dst)
		}
	})

	t.Run("copy map to existing variable", func(t *testing.T) {
		src := map[string]int{"a": 1, "b": 2}
		var dst map[string]int
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed: %v", err)
		}
		
		expected := map[string]int{"a": 1, "b": 2}
		if !reflect.DeepEqual(dst, expected) {
			t.Errorf("CopyTo failed. Got %v, want %v", dst, expected)
		}
		
		// Modify src to ensure independence
		src["a"] = 999
		src["c"] = 3
		
		if !reflect.DeepEqual(dst, expected) {
			t.Errorf("dst was affected by src change. Got %v", dst)
		}
	})

	t.Run("copy struct to existing variable", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		
		src := Person{Name: "Alice", Age: 30}
		var dst Person
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed: %v", err)
		}
		
		expected := Person{Name: "Alice", Age: 30}
		if !reflect.DeepEqual(dst, expected) {
			t.Errorf("CopyTo failed. Got %v, want %v", dst, expected)
		}
		
		// Modify src to ensure independence
		src.Name = "Bob"
		src.Age = 25
		
		if !reflect.DeepEqual(dst, expected) {
			t.Errorf("dst was affected by src change. Got %v", dst)
		}
	})

	t.Run("copy complex nested structure", func(t *testing.T) {
		type Address struct {
			Street string
			City   string
		}
		
		type Person struct {
			Name      string
			Addresses []Address
			Metadata  map[string]interface{}
		}
		
		src := Person{
			Name: "Charlie",
			Addresses: []Address{
				{Street: "123 Main St", City: "NYC"},
				{Street: "456 Oak Ave", City: "LA"},
			},
			Metadata: map[string]interface{}{
				"age":    35,
				"active": true,
				"tags":   []string{"employee", "manager"},
			},
		}
		
		var dst Person
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed: %v", err)
		}
		
		expected := Person{
			Name: "Charlie",
			Addresses: []Address{
				{Street: "123 Main St", City: "NYC"},
				{Street: "456 Oak Ave", City: "LA"},
			},
			Metadata: map[string]interface{}{
				"age":    35,
				"active": true,
				"tags":   []string{"employee", "manager"},
			},
		}
		
		if !reflect.DeepEqual(dst, expected) {
			t.Errorf("CopyTo failed. Got %v, want %v", dst, expected)
		}
		
		// Modify src to ensure independence
		src.Name = "David"
		src.Addresses[0].Street = "999 Changed St"
		src.Metadata["age"] = 40
		
		if !reflect.DeepEqual(dst, expected) {
			t.Errorf("dst was affected by src change. Got %v", dst)
		}
	})

	// Error cases
	t.Run("error cases", func(t *testing.T) {
		src := 42
		var dst int
		
		// Test nil dst
		err := CopyTo(src, nil)
		if err == nil {
			t.Error("expected error for nil dst")
		}
		
		// Test non-pointer dst
		err = CopyTo(src, dst)
		if err == nil {
			t.Error("expected error for non-pointer dst")
		}
		
		// Test nil pointer dst
		var nilPtr *int
		err = CopyTo(src, nilPtr)
		if err == nil {
			t.Error("expected error for nil pointer dst")
		}
		
		// Test type mismatch
		var stringDst string
		err = CopyTo(src, &stringDst)
		if err == nil {
			t.Error("expected error for type mismatch")
		}
	})

	t.Run("copy with custom DeepCopy interface", func(t *testing.T) {
		src := &CustomCopyable{Value: "test", Count: 5}
		var dst *CustomCopyable
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed: %v", err)
		}
		
		// The custom DeepCopy should have been called
		if dst.Value != "test_copied" {
			t.Errorf("custom DeepCopy not called. Got %s, want test_copied", dst.Value)
		}
		if dst.Count != 6 {
			t.Errorf("custom DeepCopy not called. Got %d, want 6", dst.Count)
		}
	})

	t.Run("copy nil src", func(t *testing.T) {
		var dst int
		
		err := CopyTo(nil, &dst)
		if err != nil {
			t.Errorf("CopyTo should handle nil src without error: %v", err)
		}
	})

	t.Run("pointer to pointer copying", func(t *testing.T) {
		value := 42
		src := &value
		var dst *int
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed for pointer to pointer: %v", err)
		}
		
		if dst == nil {
			t.Error("dst should not be nil after copying")
		}
		
		if *dst != 42 {
			t.Errorf("CopyTo failed. Got %v, want 42", *dst)
		}
		
		// Ensure they are different pointers
		if src == dst {
			t.Error("src and dst should be different pointers")
		}
		
		// Modify original to ensure independence
		*src = 999
		if *dst != 42 {
			t.Errorf("dst should remain unchanged when src is modified. Got %v, want 42", *dst)
		}
	})

	t.Run("pointer to value copying", func(t *testing.T) {
		value := 42
		src := &value
		var dst int
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed for pointer to value: %v", err)
		}
		
		if dst != 42 {
			t.Errorf("CopyTo failed. Got %v, want 42", dst)
		}
		
		// Modify original to ensure independence
		*src = 999
		if dst != 42 {
			t.Errorf("dst should remain unchanged when src is modified. Got %v, want 42", dst)
		}
	})

	t.Run("value to pointer copying", func(t *testing.T) {
		src := 42
		var dst *int
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed for value to pointer: %v", err)
		}
		
		if dst == nil {
			t.Error("dst should not be nil after copying")
		}
		
		if *dst != 42 {
			t.Errorf("CopyTo failed. Got %v, want 42", *dst)
		}
		
		// Modify original to ensure independence
		src = 999
		if *dst != 42 {
			t.Errorf("dst should remain unchanged when src is modified. Got %v, want 42", *dst)
		}
	})

	t.Run("nil pointer source", func(t *testing.T) {
		var src *int
		var dst *int
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed for nil pointer src: %v", err)
		}
		
		if dst != nil {
			t.Error("dst should be nil when src pointer is nil")
		}
	})

	t.Run("nil pointer to value", func(t *testing.T) {
		var src *int
		var dst int
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed for nil pointer to value: %v", err)
		}
		
		if dst != 0 {
			t.Errorf("dst should be zero value when src pointer is nil. Got %v, want 0", dst)
		}
	})

	t.Run("complex struct with pointers", func(t *testing.T) {
		type Node struct {
			Value int
			Next  *Node
		}
		
		node3 := &Node{Value: 3, Next: nil}
		node2 := &Node{Value: 2, Next: node3}
		src := &Node{Value: 1, Next: node2}
		
		var dst *Node
		
		err := CopyTo(src, &dst)
		if err != nil {
			t.Errorf("CopyTo failed for complex struct with pointers: %v", err)
		}
		
		if dst == nil {
			t.Error("dst should not be nil")
		}
		
		if dst.Value != 1 {
			t.Errorf("dst.Value = %v, want 1", dst.Value)
		}
		
		if dst.Next == nil || dst.Next.Value != 2 {
			t.Error("dst.Next not properly copied")
		}
		
		if dst.Next.Next == nil || dst.Next.Next.Value != 3 {
			t.Error("dst.Next.Next not properly copied")
		}
		
		// Ensure independence
		if src == dst || src.Next == dst.Next || src.Next.Next == dst.Next.Next {
			t.Error("copied structure shares pointers with original")
		}
		
		// Test modification independence
		src.Value = 999
		src.Next.Value = 888
		src.Next.Next.Value = 777
		
		if dst.Value != 1 || dst.Next.Value != 2 || dst.Next.Next.Value != 3 {
			t.Error("dst should remain unchanged when src is modified")
		}
	})
}

func TestChannelTypes(t *testing.T) {
	t.Run("unbuffered channel", func(t *testing.T) {
		original := make(chan int)
		copied := Copy(original).(chan int)
		
		if copied == nil {
			t.Error("copied channel should not be nil")
		}
		
		// Note: Current implementation copies channels by value (same reference)
		// This is the documented behavior for reference types
		if copied != original {
			t.Error("channel copy should reference the same channel (current implementation)")
		}
		
		// Test that channels share the same underlying channel
		go func() {
			original <- 42
			close(original)
		}()
		
		// Give some time for the goroutine to send
		val := <-copied
		if val != 42 {
			t.Errorf("copied channel should receive value from original: got %d, want 42", val)
		}
	})
	
	t.Run("buffered channel", func(t *testing.T) {
		original := make(chan string, 2)
		original <- "test1"
		original <- "test2"
		
		copied := Copy(original).(chan string)
		
		if copied == nil {
			t.Error("copied channel should not be nil")
		}
		
		// Current implementation: channels are copied by reference
		if copied != original {
			t.Error("channel copy should reference the same channel (current implementation)")
		}
		
		// Both should see the same data
		if len(original) != 2 {
			t.Errorf("original channel should have 2 items, got %d", len(original))
		}
		
		if len(copied) != 2 {
			t.Errorf("copied channel should also have 2 items (shared), got %d", len(copied))
		}
		
		// Reading from copied affects original (shared channel)
		val := <-copied
		if val != "test1" && val != "test2" {
			t.Errorf("unexpected value from copied channel: %s", val)
		}
		
		if len(original) != 1 {
			t.Error("original channel should now have 1 item after reading from copy")
		}
	})
	
	t.Run("nil channel", func(t *testing.T) {
		var original chan int
		copied := Copy(original).(chan int)
		
		if copied != nil {
			t.Error("copied nil channel should remain nil")
		}
	})
	
	t.Run("struct with channels", func(t *testing.T) {
		type ChannelStruct struct {
			Ch1 chan int
			Ch2 chan string
			Data int
		}
		
		original := ChannelStruct{
			Ch1: make(chan int, 1),
			Ch2: make(chan string),
			Data: 42,
		}
		original.Ch1 <- 99
		
		copied := Copy(original).(ChannelStruct)
		
		// Current implementation: channels are copied by reference
		if copied.Ch1 != original.Ch1 {
			t.Error("copied channel should reference same channel (current implementation)")
		}
		
		if copied.Ch2 != original.Ch2 {
			t.Error("copied channel should reference same channel (current implementation)")
		}
		
		if copied.Data != 42 {
			t.Errorf("copied data should be 42, got %d", copied.Data)
		}
		
		// Verify channels are shared (same underlying channel)
		if len(copied.Ch1) != 1 {
			t.Error("copied channel should share data with original (1 item)")
		}
		
		if len(original.Ch1) != 1 {
			t.Error("original channel should still have 1 item")
		}
		
		// Reading from copy affects original
		val := <-copied.Ch1
		if val != 99 {
			t.Errorf("expected 99 from channel, got %d", val)
		}
		
		if len(original.Ch1) != 0 {
			t.Error("original channel should now be empty after reading from copy")
		}
	})
}

func TestFunctionTypes(t *testing.T) {
	t.Run("function field", func(t *testing.T) {
		type FuncStruct struct {
			Fn   func(int) int
			Data int
		}
		
		original := FuncStruct{
			Fn: func(x int) int { return x * 2 },
			Data: 10,
		}
		
		copied := Copy(original).(FuncStruct)
		
		if copied.Data != 10 {
			t.Errorf("copied data should be 10, got %d", copied.Data)
		}
		
		// Functions should be copied by value (same reference)
		if copied.Fn == nil {
			t.Error("copied function should not be nil")
		}
		
		// Test that the function works
		if copied.Fn != nil && copied.Fn(5) != 10 {
			t.Errorf("copied function(5) = %d, want 10", copied.Fn(5))
		}
		
		// Modify original data to ensure independence
		original.Data = 999
		if copied.Data != 10 {
			t.Error("copied struct should be independent of original")
		}
	})
	
	t.Run("nil function", func(t *testing.T) {
		type FuncStruct struct {
			Fn func(int) int
		}
		
		original := FuncStruct{Fn: nil}
		copied := Copy(original).(FuncStruct)
		
		if copied.Fn != nil {
			t.Error("copied nil function should remain nil")
		}
	})
	
	t.Run("slice of functions", func(t *testing.T) {
		add := func(a, b int) int { return a + b }
		mul := func(a, b int) int { return a * b }
		
		original := []func(int, int) int{add, mul, nil}
		copied := Copy(original).([]func(int, int) int)
		
		if len(copied) != 3 {
			t.Errorf("copied slice should have 3 elements, got %d", len(copied))
		}
		
		if copied[0] == nil || copied[0](2, 3) != 5 {
			t.Error("first function not copied correctly")
		}
		
		if copied[1] == nil || copied[1](2, 3) != 6 {
			t.Error("second function not copied correctly")
		}
		
		if copied[2] != nil {
			t.Error("nil function should remain nil")
		}
		
		// Ensure slice independence
		original[0] = nil
		if copied[0] == nil {
			t.Error("copied slice should be independent")
		}
	})
}

func TestUnsafePointerTypes(t *testing.T) {
	t.Run("uintptr field", func(t *testing.T) {
		type UintptrStruct struct {
			Ptr  uintptr
			Data int
		}
		
		x := 42
		original := UintptrStruct{
			Ptr:  uintptr(unsafe.Pointer(&x)),
			Data: 100,
		}
		
		copied := Copy(original).(UintptrStruct)
		
		if copied.Ptr != original.Ptr {
			t.Error("uintptr should be copied by value")
		}
		
		if copied.Data != 100 {
			t.Errorf("data should be 100, got %d", copied.Data)
		}
		
		// Modify original to ensure independence
		original.Data = 999
		if copied.Data != 100 {
			t.Error("copied struct should be independent")
		}
	})
	
	t.Run("unsafe.Pointer field", func(t *testing.T) {
		type UnsafeStruct struct {
			Ptr  unsafe.Pointer
			Data string
		}
		
		x := 42
		original := UnsafeStruct{
			Ptr:  unsafe.Pointer(&x),
			Data: "test",
		}
		
		copied := Copy(original).(UnsafeStruct)
		
		// unsafe.Pointer should be copied by value
		if copied.Ptr != original.Ptr {
			t.Error("unsafe.Pointer should be copied by value")
		}
		
		if copied.Data != "test" {
			t.Errorf("data should be 'test', got %s", copied.Data)
		}
		
		// Both pointers should point to the same memory
		if *(*int)(copied.Ptr) != 42 {
			t.Error("unsafe pointer should still point to original value")
		}
	})
}

func TestMultiLevelPointers(t *testing.T) {
	t.Run("triple pointer", func(t *testing.T) {
		value := 42
		ptr1 := &value
		ptr2 := &ptr1
		original := &ptr2
		
		copied := Copy(original).(***int)
		
		// Verify all levels are properly copied
		if copied == original {
			t.Error("top level pointer should have different address")
		}
		
		if *copied == *original {
			t.Error("second level pointer should have different address")
		}
		
		if **copied == **original {
			t.Error("third level pointer should have different address")
		}
		
		// But the final value should be equal
		if ***copied != ***original {
			t.Errorf("final value should be equal: got %d, want %d", ***copied, ***original)
		}
		
		// Test independence by modifying original
		***original = 999
		if ***copied != 42 {
			t.Error("copied value should remain unchanged")
		}
	})
	
	t.Run("mixed pointer levels in struct", func(t *testing.T) {
		type ComplexStruct struct {
			Single   *int
			Double   **string
			Triple   ***bool
			Regular  int
		}
		
		val1 := 10
		val2 := "hello"
		val3 := true
		ptr2 := &val2
		ptr3 := &val3
		ptr3_2 := &ptr3
		
		original := ComplexStruct{
			Single:  &val1,
			Double:  &ptr2,
			Triple:  &ptr3_2,
			Regular: 100,
		}
		
		copied := Copy(original).(ComplexStruct)
		
		// Verify independence at all levels
		if copied.Single == original.Single {
			t.Error("single pointer should have different address")
		}
		
		if copied.Double == original.Double {
			t.Error("double pointer level 1 should have different address")
		}
		
		if *copied.Double == *original.Double {
			t.Error("double pointer level 2 should have different address")
		}
		
		if copied.Triple == original.Triple {
			t.Error("triple pointer level 1 should have different address")
		}
		
		// Values should be equal
		if *copied.Single != 10 {
			t.Error("single pointer value should be 10")
		}
		
		if **copied.Double != "hello" {
			t.Error("double pointer value should be 'hello'")
		}
		
		if ***copied.Triple != true {
			t.Error("triple pointer value should be true")
		}
		
		if copied.Regular != 100 {
			t.Error("regular value should be 100")
		}
		
		// Test modifications
		*original.Single = 999
		**original.Double = "modified"
		***original.Triple = false
		original.Regular = 888
		
		if *copied.Single != 10 || **copied.Double != "hello" || ***copied.Triple != true || copied.Regular != 100 {
			t.Error("copied values should remain unchanged after original modification")
		}
	})
}

func TestArrayTypes(t *testing.T) {
	t.Run("fixed size arrays", func(t *testing.T) {
		original := [5]int{1, 2, 3, 4, 5}
		copied := Copy(original).([5]int)
		
		// Arrays are value types - check contents are same initially
		for i := range original {
			if copied[i] != original[i] {
				t.Errorf("element %d: got %d, want %d", i, copied[i], original[i])
			}
		}
		
		// Test independence - arrays should be copied by value
		original[0] = 999
		if copied[0] != 1 {
			t.Error("copied array should be independent")
		}
	})
	
	t.Run("array of pointers", func(t *testing.T) {
		val1, val2, val3 := 10, 20, 30
		original := [3]*int{&val1, &val2, &val3}
		copied := Copy(original).([3]*int)
		
		// Arrays now properly deep copy their elements (fixed implementation)
		for i := range original {
			if copied[i] == original[i] {
				t.Errorf("element %d pointer should have different address after fix", i)
			}
			if *copied[i] != *original[i] {
				t.Errorf("element %d value: got %d, want %d", i, *copied[i], *original[i])
			}
		}
		
		// Test independence - array elements should now be deep copied
		*original[0] = 999
		if *copied[0] != 10 {
			t.Error("copied array elements should be independent after fix")
		}
	})
	
	t.Run("multidimensional array", func(t *testing.T) {
		original := [2][3]string{
			{"a", "b", "c"},
			{"d", "e", "f"},
		}
		copied := Copy(original).([2][3]string)
		
		for i := range original {
			for j := range original[i] {
				if copied[i][j] != original[i][j] {
					t.Errorf("element [%d][%d]: got %s, want %s", i, j, copied[i][j], original[i][j])
				}
			}
		}
		
		// Test independence
		original[0][0] = "modified"
		if copied[0][0] != "a" {
			t.Error("copied multidimensional array should be independent")
		}
	})
}

// Types for cyclic reference testing
type CyclicNodeA struct {
	Value int
	B     *CyclicNodeB
}

type CyclicNodeB struct {
	Value int
	A     *CyclicNodeA
}

func TestCyclicReferences(t *testing.T) {
	t.Run("self-referencing struct", func(t *testing.T) {
		t.Skip("Current implementation doesn't handle circular references - causes stack overflow")
		
		type CyclicNode struct {
			Value int
			Self  *CyclicNode
		}
		
		original := &CyclicNode{Value: 42}
		original.Self = original
		
		// This would cause stack overflow in current implementation
		// copied := Copy(original).(*CyclicNode)
	})
	
	t.Run("mutual reference structs", func(t *testing.T) {
		t.Skip("Current implementation doesn't handle circular references - causes stack overflow")
		
		a := &CyclicNodeA{Value: 1}
		b := &CyclicNodeB{Value: 2}
		a.B = b
		b.A = a
		
		// This would cause stack overflow in current implementation
		// copiedA := Copy(a).(*CyclicNodeA)
	})
	
	t.Run("non-circular complex structures", func(t *testing.T) {
		// Test complex structures without circular references
		a := &CyclicNodeA{Value: 1}
		b := &CyclicNodeB{Value: 2}
		
		a.B = b
		// Don't create circular reference: b.A = a
		
		// This should work fine
		copiedA := Copy(a).(*CyclicNodeA)
		
		if copiedA == a {
			t.Error("copied NodeA should have different address")
		}
		
		if copiedA.Value != 1 {
			t.Error("NodeA value should be copied")
		}
		
		if copiedA.B == nil {
			t.Error("NodeB should be copied")
		}
		
		if copiedA.B == b {
			t.Error("copied NodeB should have different address")
		}
		
		if copiedA.B.Value != 2 {
			t.Error("NodeB value should be copied")
		}
		
		if copiedA.B.A != nil {
			t.Error("NodeB.A should be nil (no circular reference)")
		}
	})
}

func TestLargeDataStructures(t *testing.T) {
	t.Run("large slice", func(t *testing.T) {
		const size = 100000
		original := make([]int, size)
		for i := range original {
			original[i] = i
		}
		
		copied := Copy(original).([]int)
		
		if len(copied) != size {
			t.Errorf("copied slice length: got %d, want %d", len(copied), size)
		}
		
		// Verify independence with spot checks
		original[0] = -1
		original[size/2] = -1
		original[size-1] = -1
		
		if copied[0] != 0 || copied[size/2] != size/2 || copied[size-1] != size-1 {
			t.Error("large slice copy should be independent")
		}
	})
	
	t.Run("large map", func(t *testing.T) {
		const size = 10000
		original := make(map[int]string, size)
		for i := 0; i < size; i++ {
			original[i] = fmt.Sprintf("value_%d", i)
		}
		
		copied := Copy(original).(map[int]string)
		
		if len(copied) != size {
			t.Errorf("copied map length: got %d, want %d", len(copied), size)
		}
		
		// Spot check values
		if copied[0] != "value_0" || copied[size/2] != fmt.Sprintf("value_%d", size/2) {
			t.Error("large map values not copied correctly")
		}
		
		// Test independence
		original[0] = "modified"
		if copied[0] != "value_0" {
			t.Error("large map copy should be independent")
		}
	})
}

// Types for interface testing
type TestReader interface {
	Read() string
}

type TestStringReader struct {
	Data string // Make it exported so it gets copied
}

func (s *TestStringReader) Read() string {
	return s.Data
}

type TestContainer struct {
	Reader TestReader
	Name   string
}

type TestStringer interface {
	String() string
}

type TestMyString string
func (m TestMyString) String() string { return string(m) }

type TestMyInt int
func (m TestMyInt) String() string { return fmt.Sprintf("%d", int(m)) }

func TestComplexInterface(t *testing.T) {
	t.Run("struct with embedded interface", func(t *testing.T) {
		original := TestContainer{
			Reader: &TestStringReader{Data: "test data"},
			Name:   "container1",
		}
		
		copied := Copy(original).(TestContainer)
		
		if copied.Name != "container1" {
			t.Error("container name should be copied")
		}
		
		if copied.Reader == nil {
			t.Error("interface should not be nil")
		}
		
		if copied.Reader == nil {
			t.Fatal("copied reader is nil")
		}
		
		readValue := copied.Reader.Read()
		if readValue != "test data" {
			t.Errorf("interface method should work on copied value: got %q, want %q", readValue, "test data")
		}
		
		// Test independence (modify original)
		if sr, ok := original.Reader.(*TestStringReader); ok {
			sr.Data = "modified"
		}
		original.Name = "modified_container"
		
		if copied.Name != "container1" || copied.Reader.Read() != "test data" {
			t.Error("copied container should be independent")
		}
	})
	
	t.Run("slice of interfaces", func(t *testing.T) {
		original := []TestStringer{
			TestMyString("hello"),
			TestMyInt(42),
			nil,
		}
		
		copied := Copy(original).([]TestStringer)
		
		if len(copied) != 3 {
			t.Errorf("copied slice length: got %d, want 3", len(copied))
		}
		
		if copied[0].String() != "hello" {
			t.Error("first interface not copied correctly")
		}
		
		if copied[1].String() != "42" {
			t.Error("second interface not copied correctly")
		}
		
		if copied[2] != nil {
			t.Error("nil interface should remain nil")
		}
		
		// Ensure slice independence
		original[0] = TestMyString("modified")
		if copied[0].String() != "hello" {
			t.Error("interface slice should be independent")
		}
	})
}

// Additional types for comprehensive DeepCopy interface testing
type NilReturningCopyable struct {
	Value string
}

func (n *NilReturningCopyable) DeepCopy() interface{} {
	// This implementation returns nil
	return nil
}

type PanicCopyable struct {
	Value string
}

func (p *PanicCopyable) DeepCopy() interface{} {
	panic("DeepCopy panic test")
}

type SelfReferencingCopyable struct {
	Value string
	Self  *SelfReferencingCopyable
}

func (s *SelfReferencingCopyable) DeepCopy() interface{} {
	return &SelfReferencingCopyable{
		Value: s.Value + "_copied",
		Self:  s, // Return reference to original (testing shallow behavior)
	}
}

type SliceCopyable struct {
	Items []string
}

func (s *SliceCopyable) DeepCopy() interface{} {
	// Custom slice copying logic
	newItems := make([]string, len(s.Items))
	for i, item := range s.Items {
		newItems[i] = item + "_copied"
	}
	return &SliceCopyable{Items: newItems}
}

type PointerFieldCopyable struct {
	Value *string
	Count int
}

func (p *PointerFieldCopyable) DeepCopy() interface{} {
	var newValue *string
	if p.Value != nil {
		val := *p.Value + "_copied"
		newValue = &val
	}
	return &PointerFieldCopyable{
		Value: newValue,
		Count: p.Count * 2,
	}
}

type NestedCopyable struct {
	Inner *CustomCopyable
	Name  string
}

func (n *NestedCopyable) DeepCopy() interface{} {
	var innerCopy *CustomCopyable
	if n.Inner != nil {
		innerCopy = n.Inner.DeepCopy().(*CustomCopyable)
	}
	return &NestedCopyable{
		Inner: innerCopy,
		Name:  n.Name + "_nested_copied",
	}
}

func TestDeepCopyInterface(t *testing.T) {
	t.Run("nil returning DeepCopy", func(t *testing.T) {
		original := &NilReturningCopyable{Value: "test"}
		copied := Copy(original).(*NilReturningCopyable)
		
		// The DeepCopy returns nil, so copied should be nil
		if copied != nil {
			t.Errorf("expected nil from DeepCopy that returns nil, got %v", copied)
		}
	})
	
	t.Run("panic in DeepCopy", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic from DeepCopy, but didn't panic")
			} else if r != "DeepCopy panic test" {
				t.Errorf("unexpected panic message: %v", r)
			}
		}()
		
		original := &PanicCopyable{Value: "test"}
		Copy(original)
	})
	
	t.Run("nil pointer with DeepCopy interface", func(t *testing.T) {
		var original *CustomCopyable
		copied := Copy(original).(*CustomCopyable)
		
		if copied != nil {
			t.Error("copying nil pointer should result in nil")
		}
	})
	
	t.Run("self-referencing DeepCopy", func(t *testing.T) {
		original := &SelfReferencingCopyable{Value: "test"}
		original.Self = original
		
		copied := Copy(original).(*SelfReferencingCopyable)
		
		if copied == nil {
			t.Error("copied should not be nil")
		}
		
		if copied.Value != "test_copied" {
			t.Errorf("copied value should be 'test_copied', got %s", copied.Value)
		}
		
		// The DeepCopy implementation returns reference to original
		if copied.Self != original {
			t.Error("DeepCopy should return reference to original as implemented")
		}
	})
	
	t.Run("slice manipulation in DeepCopy", func(t *testing.T) {
		original := &SliceCopyable{Items: []string{"a", "b", "c"}}
		copied := Copy(original).(*SliceCopyable)
		
		if copied == nil {
			t.Error("copied should not be nil")
		}
		
		if len(copied.Items) != 3 {
			t.Errorf("copied should have 3 items, got %d", len(copied.Items))
		}
		
		expected := []string{"a_copied", "b_copied", "c_copied"}
		for i, item := range copied.Items {
			if item != expected[i] {
				t.Errorf("item %d should be %s, got %s", i, expected[i], item)
			}
		}
		
		// Modify original to test independence
		original.Items[0] = "modified"
		if copied.Items[0] != "a_copied" {
			t.Error("copied slice should be independent of original")
		}
	})
	
	t.Run("pointer field handling in DeepCopy", func(t *testing.T) {
		value := "original"
		original := &PointerFieldCopyable{Value: &value, Count: 5}
		copied := Copy(original).(*PointerFieldCopyable)
		
		if copied == nil {
			t.Error("copied should not be nil")
		}
		
		if copied.Value == nil {
			t.Error("copied value pointer should not be nil")
		}
		
		if *copied.Value != "original_copied" {
			t.Errorf("copied value should be 'original_copied', got %s", *copied.Value)
		}
		
		if copied.Count != 10 {
			t.Errorf("copied count should be 10, got %d", copied.Count)
		}
		
		// Test pointer independence
		if copied.Value == original.Value {
			t.Error("copied pointer should have different address")
		}
		
		*original.Value = "modified"
		if *copied.Value != "original_copied" {
			t.Error("copied value should be independent")
		}
	})
	
	t.Run("nil pointer field in DeepCopy", func(t *testing.T) {
		original := &PointerFieldCopyable{Value: nil, Count: 5}
		copied := Copy(original).(*PointerFieldCopyable)
		
		if copied == nil {
			t.Error("copied should not be nil")
		}
		
		if copied.Value != nil {
			t.Error("copied value should remain nil")
		}
		
		if copied.Count != 10 {
			t.Errorf("copied count should be 10, got %d", copied.Count)
		}
	})
	
	t.Run("nested DeepCopy implementations", func(t *testing.T) {
		inner := &CustomCopyable{Value: "inner", Count: 3}
		original := &NestedCopyable{Inner: inner, Name: "outer"}
		
		copied := Copy(original).(*NestedCopyable)
		
		if copied == nil {
			t.Error("copied should not be nil")
		}
		
		if copied.Name != "outer_nested_copied" {
			t.Errorf("copied name should be 'outer_nested_copied', got %s", copied.Name)
		}
		
		if copied.Inner == nil {
			t.Error("copied inner should not be nil")
		}
		
		if copied.Inner.Value != "inner_copied" {
			t.Errorf("copied inner value should be 'inner_copied', got %s", copied.Inner.Value)
		}
		
		if copied.Inner.Count != 4 {
			t.Errorf("copied inner count should be 4, got %d", copied.Inner.Count)
		}
		
		// Test independence
		if copied.Inner == original.Inner {
			t.Error("copied inner should have different address")
		}
		
		original.Inner.Value = "modified"
		if copied.Inner.Value != "inner_copied" {
			t.Error("copied inner should be independent")
		}
	})
	
	t.Run("nil nested DeepCopy", func(t *testing.T) {
		original := &NestedCopyable{Inner: nil, Name: "outer"}
		copied := Copy(original).(*NestedCopyable)
		
		if copied == nil {
			t.Error("copied should not be nil")
		}
		
		if copied.Name != "outer_nested_copied" {
			t.Errorf("copied name should be 'outer_nested_copied', got %s", copied.Name)
		}
		
		if copied.Inner != nil {
			t.Error("copied inner should remain nil")
		}
	})
	
	t.Run("slice of DeepCopy implementations", func(t *testing.T) {
		original := []*CustomCopyable{
			{Value: "first", Count: 1},
			{Value: "second", Count: 2},
			nil,
			{Value: "fourth", Count: 4},
		}
		
		copied := Copy(original).([]*CustomCopyable)
		
		if len(copied) != 4 {
			t.Errorf("copied should have 4 elements, got %d", len(copied))
		}
		
		// Check first element
		if copied[0] == nil {
			t.Error("first element should not be nil")
		}
		if copied[0].Value != "first_copied" {
			t.Errorf("first element value should be 'first_copied', got %s", copied[0].Value)
		}
		if copied[0].Count != 2 {
			t.Errorf("first element count should be 2, got %d", copied[0].Count)
		}
		
		// Check nil element
		if copied[2] != nil {
			t.Error("third element should remain nil")
		}
		
		// Test independence
		if copied[0] == original[0] {
			t.Error("copied elements should have different addresses")
		}
		
		original[0].Value = "modified"
		if copied[0].Value != "first_copied" {
			t.Error("copied elements should be independent")
		}
	})
	
	t.Run("map with DeepCopy implementations", func(t *testing.T) {
		original := map[string]*CustomCopyable{
			"key1": {Value: "value1", Count: 1},
			"key2": {Value: "value2", Count: 2},
			"key3": nil,
		}
		
		copied := Copy(original).(map[string]*CustomCopyable)
		
		if len(copied) != 3 {
			t.Errorf("copied map should have 3 elements, got %d", len(copied))
		}
		
		// Check regular elements
		if copied["key1"] == nil {
			t.Error("key1 should not be nil")
		}
		if copied["key1"].Value != "value1_copied" {
			t.Errorf("key1 value should be 'value1_copied', got %s", copied["key1"].Value)
		}
		if copied["key1"].Count != 2 {
			t.Errorf("key1 count should be 2, got %d", copied["key1"].Count)
		}
		
		// Check nil element
		if copied["key3"] != nil {
			t.Error("key3 should remain nil")
		}
		
		// Test independence
		if copied["key1"] == original["key1"] {
			t.Error("copied map elements should have different addresses")
		}
		
		original["key1"].Value = "modified"
		if copied["key1"].Value != "value1_copied" {
			t.Error("copied map elements should be independent")
		}
	})
}
