# pythonify

A Go library that brings Python-like data structures and built-in functions to Go, with full generic type support.

This is a Go library extracted from [daedalus-mono](https://github.com/daedalus-mono/).

## Installation

```bash
go get github.com/dae-go/pythonify
```

## Features

### Data Structures

#### List
Python-like list with generic types:
```go
import "github.com/dae-go/pythonify/pkg"

// Create lists
list := py.NewList[int](1, 2, 3)
list.Append(4)
list.Extend(5, 6, 7)
list.Insert(0, 0)

// List operations
fmt.Println(list.Count(2))    // Count occurrences
fmt.Println(list.Index(3))    // Find index
list.Remove(2)                // Remove first occurrence
item := list.Pop(0)           // Remove and return item
list.Reverse()                // Reverse in place
list.Sort()                   // Sort (for ordered types)
```

#### Set
Generic set implementation:
```go
set := py.NewSet[string]("a", "b", "c")
set.Add("d")
set.Update(py.NewList("e", "f"))

// Set operations
other := py.NewSet[string]("c", "d", "e")
union := set.Union(other)
intersection := set.Intersection(other)
difference := set.Difference(other)

fmt.Println(set.Contains("a"))     // Check membership
fmt.Println(set.IsSubset(other))   // Set comparisons
```

#### Dict
Generic dictionary with comparable keys:
```go
dict := py.NewDict[string, int]()
dict["key1"] = 42
dict["key2"] = 84

value := dict.Get("key1")         // Get value
keys := dict.Keys()               // Get all keys
values := dict.Values()           // Get all values
items := dict.Items()             // Get key-value pairs

dict.Update(otherDict)            // Merge dictionaries
popped := dict.Pop("key1", 0)     // Remove and return with default
```

#### Tuple
Immutable sequence type:
```go
tuple := py.NewTuple[string]("a", "b", "c", "b")
fmt.Println(tuple.Count("b"))     // Count occurrences
fmt.Println(tuple.Index("c"))     // Find index
value, ok := tuple.Get(1)         // Safe index access
slice := tuple.ToSlice()          // Convert to slice
```

#### File
Python-like file operations with proper error handling:
```go
file := py.NewFile(osFile, "example.txt", "r")
data, err := file.Read(100)       // Read bytes
line, err := file.Readline()      // Read line
lines, err := file.Readlines()    // Read all lines
n, err := file.Write([]byte("hello"))  // Write bytes
file.Flush()                      // Flush buffers
file.Close()                      // Close file
```

### Built-in Functions

#### Type Conversion
```go
// Number base conversions
fmt.Println(py.Bin(42))          // "0b101010"
fmt.Println(py.Hex(255))         // "0xff"
fmt.Println(py.Oct(64))          // "0o100"

// Type conversions
py.Int("42", 10)                 // String to int with base
py.Float("3.14")                 // String to float
py.Bool(value)                   // Any value to bool (Python truthiness)
py.Chr(65)                       // Unicode code point to string
py.Ord("A")                      // String to Unicode code point
```

#### Sequence Operations
```go
// Using the generic functions
result := py.Filter(func(x int) bool { return x > 5 }, list)
doubled := py.Map(func(x int) int { return x * 2 }, list)
reversed := py.Reversed(list)
sorted := py.Sorted(list, false)

// Python-like zip
lists := py.Zip(list1, list2, list3)

// Aggregation
hasAny := py.Any(list, func(x int) bool { return x > 10 })
allPass := py.All(list, func(x int) bool { return x > 0 })
total := py.Sum(numbers)
```

#### Object Introspection
```go
// Object inspection
fmt.Println(py.Type(obj))         // Get type
attrs := py.Dir(obj)              // List attributes
hasAttr := py.HasAttr(obj, "field") // Check attribute
value := py.GetAttr(obj, "field", defaultVal) // Get attribute
vars := py.Vars(obj)              // Get struct fields as dict

// Object representation
fmt.Println(py.Repr(obj))         // Representation string
fmt.Println(py.Ascii(obj))        // ASCII representation
```

#### Iterators
Uses Go's `iter` package for proper iterator support:
```go
// Create iterators from various types
iterator := py.Iter(slice, nil)
iterator2 := py.Iter(channel, nil)
iterator3 := py.Iter(stringVal, nil)

// Use iterator
for {
    value := py.Next(iterator, nil) // Will panic on StopIteration
    // or with default:
    value := py.Next(iterator, "default")
}
```

### Low-level Functions (abf package)

For performance-critical code, use the `abf` package directly:
```go
import "github.com/dae-go/pythonify/pkg/abf"

// Direct slice operations (no type conversion overhead)
filtered := abf.Filter(predicate, slice)
mapped := abf.Map(transform, slice)
zipped := abf.Zip(slice1, slice2, slice3)
```

## Package Structure

- `pkg/` - Main package with generic types and Python-like functions
- `pkg/abf/` - Low-level functions operating directly on slices
- `pkg/list.go` - List, Alist, Olist implementations
- `pkg/set.go` - Set and FrozenSet implementations  
- `pkg/dict.go` - Generic Dict implementation
- `pkg/tuple.go` - Immutable Tuple implementation
- `pkg/file.go` - File operations with Python-like interface
- `pkg/utils.go` - High-level Python built-in function equivalents

## Type Hierarchy

```
Alist[T any]           - Base list for any type
├─ List[T comparable]  - List for comparable types (adds Index, Remove, Count)
└─ Olist[T ordered]    - List for ordered types (adds Sort)

Dict[K comparable, V any] - Generic dictionary

Set[T comparable]         - Mutable set
FrozenSet[T comparable]   - Immutable set

Tuple[T comparable]       - Immutable sequence
```

## Usage Patterns

### Type-safe Collections
```go
// Strongly typed collections
users := py.NewList[User]()
usersByID := py.NewDict[int, User]()
tags := py.NewSet[string]()
```

### Functional Programming
```go
// Chain operations
result := py.Filter(isValid, 
    py.Map(transform, 
        py.Sorted(data, false)))
```

### Python-like Iteration
```go
// Iterate like Python
for item := range py.Iter(collection, nil) {
    // process item
}
```

## Development

This library is automatically synced from the monorepo. Please make changes in the [main repository](https://github.com/icarus612/daedalus-monorepo/tree/main/libs/golang/pythonify).

## License

See the main repository for license information.