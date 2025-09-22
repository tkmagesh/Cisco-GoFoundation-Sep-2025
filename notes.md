# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
| What | When |
| ----- | ----- |
| Commence | 9:00 AM |
| Tea Break | 10:30 AM (20 mins) |
| Lunch Break | 12:30 PM (1 hr) |
| Tea Break | 3:00 PM (20 mins) |
| Wind up | 4:30 PM |

## Methodology
- No powerpoint
- Only code & discussion

## Software Prerequisities
- Any Editor (Visual Studio Code)
- Go Tools (https://go.dev/dl)

### Go Tools installation verification
```shell
go version
```

## Repository
- https://github.com/tkmagesh/cisco-gofoundation-sep-2025

## Why Go?
1. Simplicity
    - ONLY 25 keywords
        - package, import, func, var, const, if, else, switch, case, default, for, range, return, go, select, break, continue, goto, type
    - No access modifiers (public/private/protected)
    - No reference types (everything is a value) 
    - Use pointers for handling references
        - But pointer arithmatic is not supported in the `language`
    - No classes (only structs)
    - No inheritance (only composition)
    - No exceptions (only errors)
    - No try-catch-finally construct
    - No implicity type conversion
2. Performance
    - compiled to native code
        - tooling support for cross compilation
    - on par with c++
3. Concurrency
    - Cheaper & lighter alternative to OS Threads called `goroutines`
    - OS Threads = ~2MB, Goroutines = ~2KB
    - Support for concurrency is **built in the language**
        - `go` keyword
        - `chan` datatype
        - `<-` operator
        - `range` construct
        - `select-case` construct
    - API support
        - `sync` package
        - `sync/atomic` package

## Compilation & Execution
### Compilation
```shell
go build <filename.go>
```

```shell
go build -o <binary-name> <filename.go>
```
### Compile & Execute
```shell
go run <filename.go>
```
### Cross Compilation
#### To get the list of go environment variables and their values
```shell
go env
```

```shell
go env <var_1> <var_2> ....
# ex
go env gopath gobin
```

#### To change the environment variable values
```shell
go env -w <var_1>=<value_1> <var_2>=<value_2> ...
```

#### Environment variables for cross-compilation
- GOOS
- GOARCH

#### To get the list of suppported platforms (for cross-compilation)
```shell
go tool dist list
```

#### Cross compilation
```shell
GOOS=<target_os> GOARCH=<target_processor_arch> go build <filename.go>
# ex
GOOS=windows GOARCH=amd64 go build 01-hello-world.go
```

##### In windows (powershell)
```powershell
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o myapp-windows-amd64.exe
```

## Standard Library documentation
- https://pkg.go.dev/std

## Data Types
- string
- bool
- integers family
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integers family
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating points family
    - float32
    - float64
- complex family
    - complex64 ( real[float32] + imaginary[float32] )
    - complex128 ( real[float64] + imaginary[float64] )
- aliases
    - byte (alias for unsigned int)
    - rune (alias for unicode code point)

### Zero values
| Data Type | Zero value |
------------ | ------------- |
|int family     | 0 |
|uint family    | 0 |
|complex family | (0+0i) |
|string         | "" (empty string) |
|bool           | false |
|byte           | 0 |
|interface      | nil |
|pointer        | nil |
|function       | nil |
|struct         | struct instance |

### Function Scope
- CANNOT have unused variables
- Can use ":="

### Package Scope
- Can have unused variables
- Cannot use ":="

## Programming constructs
- if else
- switch case
- for

## Functions
- Named results
- Variadic functions
- Anonymous functions
    - functions defined in another functions
    - cannot have a name
    - they have to be used
- Higher Order functions
- Deferred functions