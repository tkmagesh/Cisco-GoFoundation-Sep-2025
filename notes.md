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