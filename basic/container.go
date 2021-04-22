package main

import (
	"fmt"
)

func main() {
	jmlExportA:=1000.0
	jmlExportB:=100.0
	
	for i:=1; i<=18; i++ {
	if jmlExportA>jmlExportB {
		jmlExportA=jmlExportA-(50.0/100.0)
		jmlExportB=jmlExportB+(20.0/100.0)
		fmt.Println(jmlExportA, " ", jmlExportB)
	} else {
		fmt.Println("pajak pemerintah")
		
	}
}

}