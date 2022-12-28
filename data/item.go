package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	Text     string
	Priority int
}

/* Set todo Item's Priority (Default is 2). */
func (i *Item) SetPriority(priority int) {
	switch priority {
	case 1:
		i.Priority = 1
	case 2:
		i.Priority = 2
	case 3:
		i.Priority = 3
	default:
		i.Priority = 4
	}
}

func (i *Item) ToString() string {
	return i.Text + " p(" + fmt.Sprint(i.Priority) + ")"
}

/*
SaveItems saves a list of Items as a json string in a specified path.
*/
func SaveItems(path string, items []Item) error {
	b, err := json.Marshal(items)
	check(err)
	err = os.WriteFile(path, b, 0644)
	check(err)

	return nil
}

func ReadItems(path string) ([]Item, error) {
	b, err := os.ReadFile(path)
	check(err)

	items := []Item{}
	if err := json.Unmarshal(b, &items); err != nil {
		fmt.Println("Error in item.go:ReadItems() - ", err)
		return []Item{}, err
	}

	return items, nil
}

/*
check verifies if there was an error, throwing a panic if there is one.
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}
