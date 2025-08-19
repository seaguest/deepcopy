# deepcopy

[![GoDoc](https://godoc.org/github.com/mohae/deepcopy?status.svg)](https://godoc.org/github.com/mohae/deepcopy) [![Build Status](https://travis-ci.org/mohae/deepcopy.png)](https://travis-ci.org/mohae/deepcopy)

A Go library for making deep copies of data structures. Creates completely independent copies of complex data types including structs, slices, maps, and pointers.

## Installation

```bash
go get github.com/seaguest/deepcopy
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/seaguest/deepcopy"
)

func main() {
    original := map[string][]int{
        "numbers": {1, 2, 3},
        "more":    {4, 5, 6},
    }
    
    copy := deepcopy.Copy(original).(map[string][]int)
    
    // Modify the copy - original remains unchanged
    copy["numbers"][0] = 999
    
    fmt.Println("Original:", original["numbers"]) // [1 2 3]
    fmt.Println("Copy:", copy["numbers"])         // [999 2 3]
}
```

## Usage

### Basic Copy

```go
cpy := deepcopy.Copy(orig)
```

The returned value needs to be type-asserted to the correct type:

```go
originalSlice := []int{1, 2, 3}
copiedSlice := deepcopy.Copy(originalSlice).([]int)
```

### Copy to Existing Variable

Use `CopyTo` to copy into an existing variable without allocating a new one:

```go
var dst []int
src := []int{1, 2, 3}

err := deepcopy.CopyTo(src, &dst)
if err != nil {
    // Handle error
}
// dst now contains a deep copy of src
```

### Copying Complex Structures

```go
type Person struct {
    Name    string
    Age     int
    Friends []*Person
}

original := &Person{
    Name: "Alice",
    Age:  30,
    Friends: []*Person{
        {Name: "Bob", Age: 25},
        {Name: "Charlie", Age: 35},
    },
}

copy := deepcopy.Copy(original).(*Person)
```

### Custom Deep Copy Implementation

Implement the `Interface` to control how your types are copied:

```go
type MyStruct struct {
    Data string
}

func (m MyStruct) DeepCopy() interface{} {
    return MyStruct{Data: m.Data + " (copied)"}
}

original := MyStruct{Data: "test"}
copy := deepcopy.Copy(original).(MyStruct)
// copy.Data will be "test (copied)"
```

## Supported Types

deepcopy handles the following Go types:

- **Basic types**: strings, numbers, booleans
- **Pointers**: creates new pointers to deep copies of pointed-to values
- **Structs**: recursively copies all exported fields
- **Slices**: creates new slices with deep copies of elements
- **Maps**: creates new maps with deep copies of keys and values
- **Interfaces**: copies underlying concrete values
- **Time**: special handling for `time.Time` values
- **Custom types**: via the `Interface` implementation

## API Reference

### Functions

#### `Copy(src interface{}) interface{}`
Creates a deep copy of the source value and returns it as an `interface{}`.

#### `CopyTo(src, dst interface{}) error`
Copies the source value to the destination. The destination must be a pointer to the target type.

#### `Iface(iface interface{}) interface{}`
Alias to `Copy` for backwards compatibility.

### Interfaces

#### `Interface`
```go
type Interface interface {
    DeepCopy() interface{}
}
```
Implement this interface on your types to customize the copying behavior.

## Limitations

- **Unexported fields**: Values in unexported struct fields are not copied
- **Channels**: Not supported
- **Functions**: Not supported
- **Unsafe pointers**: Not supported

## Performance Notes

- Uses reflection, so it's slower than manual copying
- Memory allocations scale with the size and depth of the data structure
- Consider implementing the `Interface` for frequently copied types to optimize performance

## License

MIT - see LICENSE file for details
