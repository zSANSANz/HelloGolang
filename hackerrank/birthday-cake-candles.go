package main

import (
    "fmt"
)

func main() {
    n, t := 0, 0
    fmt.Scan(&n)
    a := []int{}
    
    for i:=1; i<=n; i++ {
        fmt.Scan(&t)
        a = append(a, t)
        
    }
    
    temp := 0
    for _, values := range a {
        if temp < values {
            temp = values
        }
    }
    
    sum := 0
    
    for _, values := range a {
        if temp == values {
            sum += 1
        }
    }
    
    fmt.Println(sum)
}
