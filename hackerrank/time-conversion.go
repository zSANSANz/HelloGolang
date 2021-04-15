package main

import(
    "fmt"
)

func main() {
    var jam, menit, detik int
    var suffix string
    fmt.Scanf("%d:%d:%d%s", &jam, &menit, &detik, &suffix)
    if suffix == "PM" {
        if jam != 12 {
			jam = jam + 12
		}
    }

	if suffix == "AM" && jam == 12 {
		jam = 0
	}
    fmt.Printf("%02d:%02d:%02d", jam, menit, detik)
}