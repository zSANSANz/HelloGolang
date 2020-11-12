package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var sourceString []string

	// mencari file atau folder dari folder source
	errSource := filepath.Walk("./source",
		func(path string, info os.FileInfo, errSource error) error {
			if errSource != nil {
				return errSource
			}
			sourceString = append(sourceString, path)
			return nil
		})
	if errSource != nil {
		log.Println(errSource)
	}

	var targetString []string

	// mencari file atau folder dari folder target
	errTarget := filepath.Walk("./target",
		func(path string, info os.FileInfo, errTarget error) error {
			if errTarget != nil {
				return errTarget
			}
			targetString = append(targetString, path)
			return nil
		})
	if errTarget != nil {
		log.Println(errTarget)
	}

	fmt.Println(" ")
	fmt.Println(" ")

	for i := 0; i < len(sourceString); i++ {
		destroySource := strings.Replace(sourceString[i], "source", "", -1)
		destroyTarget := strings.Replace(targetString[i], "target", "", -1)

		if destroySource == destroyTarget {
			// kalau ketemu file yang sama di beri tulisan modified
			fmt.Println(destroySource, "MODIFIED")
		} else if strings.Contains(destroySource, "txt") {
			// kalau ada di source tapi tidak ada di target berikan keterangan new
			fmt.Println(destroySource, "NEW")
		} else if strings.Contains(destroyTarget, "txt") {
			// Jika file tidak ada di source tapi ada di target berikan keterangan deleted
			fmt.Println(destroyTarget, "DELETED")
		}

		fmt.Println()
	}

}
