package main

import (
	"fmt"

	"novadb1/internal/sstable"
)

func main() {

	value, found, err := sstable.Search(
		"data/sstable_1.txt",
		"2",
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Value:", value)
	fmt.Println("Found:", found)
}
