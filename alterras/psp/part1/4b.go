package main

import (
	"fmt"
)

func main() {
	searchField := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
    searchNumber := 23

    fmt.Println("Running Program")
    fmt.Println("Searching list of numbers: ", searchField)
    fmt.Println("Searching for number: ", searchNumber)

    numFound := false
    result, _ := binarySearch(searchField, len(searchField), searchNumber, numFound)
    fmt.Println("Found! Your number is found in position: ", result)
   
}

func binarySearch(a []int, search int) (result int, searchCount int) {
    mid := len(a) / 2
    switch {
    case len(a) == 0:
        result = -1 
    case a[mid] > search:
        result, searchCount = binarySearch(a[:mid], search)
    case a[mid] < search:
        result, searchCount = binarySearch(a[mid+1:], search)
        if result >= 0 { 
            result += mid + 1
        }
    default: 
        result = mid 
    }
    searchCount++
    return
}