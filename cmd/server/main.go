package main

import (
	"fmt"

	"novadb1/internal/mvcc"
)

func main() {

	store := mvcc.New()

	store.Put("2", "Bob", 100)
	store.Put("2", "Robert", 200)
	store.Put("2", "Bobby", 300)

	tx := mvcc.NewTransaction(250)

	value, found := tx.Read(
		store,
		"2",
	)

	fmt.Println("Snapshot Value:", value)
	fmt.Println("Found:", found)
}
