```go
package main

import (

    "fmt"
)

func main() {
    var m map[string]int
    if m == nil {
        fmt.Println("Map is nil") // Handle the nil case appropriately
    } else {
        fmt.Println(m["a"])
    }
}
```