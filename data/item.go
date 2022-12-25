package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	Text string
}

/*
SaveItems saves a list of Items as a json string in a specified path.
*/
func SaveItems(path string, items []Item) error {
	b, err := json.Marshal(items)
	check(err)
	err = os.WriteFile(path, b, 0644)
	check(err)

	fmt.Println(string(b))

	return nil
}

/*
check verifies if there was an error, throwing a panic if there is one.
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}
