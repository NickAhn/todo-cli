/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/NickAhn/todo/data"
	"github.com/spf13/cobra"
)

// editpCmd represents the editp command
var editpCmd = &cobra.Command{
	Use:   "editp",
	Short: "Edit priority of todo Item",
	Long:  `Edit the priority of todo item.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get item list
		items, err := data.ReadItems(todo_list_path)
		if err != nil {
			fmt.Println(err)
		}

		// get todo item index
		index, err := strconv.Atoi(args[0])

		// Get user input for new priority
		fmt.Print("\t- Enter new priority for item \"" + items[index].Text +
			" (p" + fmt.Sprint(items[index].Priority) + ")\": ")
		var newPriority int
		fmt.Scanln(&newPriority)

		items[index].SetPriority(newPriority)

		fmt.Println("\tPriority has been set to p" + fmt.Sprint(items[index].Priority))

		// Save Changes
		data.SaveItems(todo_list_path, items)

	},
}

func init() {
	rootCmd.AddCommand(editpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
