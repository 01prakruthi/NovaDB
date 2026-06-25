package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	data := make(map[string]string)

	for i := 0; i < 100000; i++ {

		key := fmt.Sprintf("%d", i)
		value := fmt.Sprintf("User%d", i)

		data[key] = value
	}

	elapsed := time.Since(start)

	fmt.Println("Inserted 100000 records")
	fmt.Println("Time:", elapsed)
}
