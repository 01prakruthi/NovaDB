package main

import (
	"fmt"

	"novadb1/internal/sstable"
)

func main() {

	value, found, err := sstable.SearchAll(
		"data",
		"4",
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Value:", value)
	fmt.Println("Found:", found)
}
