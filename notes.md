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
- Higher Order functions (treating functions like data)
    - assign a function as a value to a variable
    - pass function as an argument to other function
    - return a function as a return value
- Deferred functions

## Collection
- Array
    - Fixed sized typed collection
- Slice
    - Varying sized typed collection
    - A pointer to an internal array
    - Functions
        - append()
        - len()
- Map
    - Varying sized typed collection of key/value pairs

## Error Handling
### Errors
- Values `returned` from a function
- Should implement `error` interface (best practice)
- Creation
    - errors.New()
    - fmt.Errorf()
    - Custom type implementing `error` interface

## Panic & Recovery
### Panic
- Represents the state of the application where the application execution is unable to proceed further
- When panic occurs, the application is shutdown **after all the scheduled deferred functions are executed**
- Use **panic()** to programmatically create and raise a panic

### Recovery
- Use **recover()** to recover from the panic
- Apt to be used in the `deferred` functions

## Modules & Packages
### Module
- Any code that need to be versioned and deployed together
- A module is just a folder with `go.mod` file
- **go.mod** file
    - the manifest for the module
    - contains the following
        - name of the module
            - Good to have the complete repo path in the module name
        - go version
        - references to external dependencies
- create a go.mod file
```shell
go mod init <module-name>
```
- run a module
```shell
go run .
```
- create a build of a module
```shell
go build .
```

### Package
- internal organization of code in a module
- just folders
- can have any number of .go files
- All the code in all the .go files of a package are considered to be part of the same package (ie., files loose their identity)
- Public entities name MUST start with upper case
- Private entities name MUST start with lower case