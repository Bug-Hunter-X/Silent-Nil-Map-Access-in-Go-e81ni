```go
func main() {
    var m map[string]int
    fmt.Println(m["a"]) // This will not panic, it returns the zero value of int (0)
}
```