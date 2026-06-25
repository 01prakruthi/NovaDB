package main

import (
	"fmt"

	"novadb1/internal/bloom"
)

func main() {

	bf := bloom.New(100)

	bf.Add("2")
	bf.Add("4")
	bf.Add("5")

	fmt.Println("Contains 2:", bf.MightContain("2"))
	fmt.Println("Contains 4:", bf.MightContain("4"))
	fmt.Println("Contains 99:", bf.MightContain("99"))
}
