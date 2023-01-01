package data

import (
	"encoding/json"
	"fmt"
	"os"
)

var Reset = "\033[0m"
var Red = "\033[31m" // p1
// var Green = "\033[32m"
var Yellow = "\033[33m"

// var Blue = "\033[34m"
var Purple = "\033[35m" //p2
var Cyan = "\033[36m"   //p3
// var Gray = "\033[37m"
var White = "\033[97m" //p4

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
	str := i.Text + " p(" + fmt.Sprint(i.Priority) + ")"
	// Add color by Priority
	return str
}

func (i *Item) colorPriority() string {
	switch i.Priority {
	case 1:
		return Red + "p(" + fmt.Sprint(i.Priority) + ")" + Reset
	case 2:
		return Purple + "p(" + fmt.Sprint(i.Priority) + ")" + Reset
		// str = Purple + str
	case 3:
		return Cyan + "p(" + fmt.Sprint(i.Priority) + ")" + Reset
	default:
		return "p(" + fmt.Sprint(i.Priority) + ")" + Reset
	}
}

/*
SaveItems saves a list of Items as a json string in a specified path.
*/
func SaveItems(path string, items []Item) error {
	b, err := json.Marshal(items)
	Check(err)
	err = os.WriteFile(path, b, 0644)
	Check(err)

	return nil
}

func ReadItems(path string) ([]Item, error) {
	b, err := os.ReadFile(path)
	Check(err)

	items := []Item{}
	if err = json.Unmarshal(b, &items); err != nil {
		return []Item{}, nil
	}

	return items, nil
}

func PrintTODO(items []Item) {
	fmt.Println("TODO (" + fmt.Sprint(len(items)) + "):")
	for i, item := range items {
		fmt.Println("   ", fmt.Sprint(i)+".", item.Text, item.colorPriority())
	}

}

/*
Error handling function
*/
func Check(e error) {
	if e != nil {
		fmt.Println(Red, e, Reset)
		os.Exit(1)
	}
}
