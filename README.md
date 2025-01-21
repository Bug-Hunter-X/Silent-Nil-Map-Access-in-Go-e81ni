# Silent Nil Map Access in Go

This repository demonstrates a potential pitfall in Go when working with maps: accessing elements in a nil map. Unlike some other languages that might throw an exception, Go silently returns the zero value of the map's value type.

## The Bug

The `bug.go` file contains a simple program that showcases this behavior.  Accessing a key in an uninitialized (nil) map does not cause a runtime panic.  Instead, it returns 0 (the zero value for `int`)

## The Solution

The recommended solution is to always check for `nil` before accessing elements in a map.  The `bugSolution.go` file provides a corrected version that explicitly checks if the map is nil before accessing keys.  This prevents unexpected results and makes the code more robust.

## How to reproduce:

1. Clone this repository.
2. Navigate to the repository's directory.
3. Run `go run bug.go` to see the unexpected output.
4. Run `go run bugSolution.go` to see the corrected behaviour.
