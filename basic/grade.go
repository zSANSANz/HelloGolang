package main

import (
    "fmt"
)

func main() {
    n := 0
    var grades []int
    fmt.Scan(&n)
    for i:=0; i<n; n++ {
        fmt.Scan(&grades[i])
    }
    
    for i:=1; i<=n; n++ {
        fmt.Println(grades[i])
    }
}
