/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/NickAhn/todo/data"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a todo item",
	Long:  `del Deletes a todo item from the list at specified index.`,
	Run: func(cmd *cobra.Command, args []string) {
		items, _ := data.ReadItems(todo_list_path) // get saved todo-list
		sort.Strings(args)

		// deletedItems := make([]data.Item, len(args))
		fmt.Println("")
		for i := 0; i < len(args); i++ {
			index, _ := strconv.Atoi(args[i])
			delItem := items[index-i]
			// deletedItems = append(deletedItems, delItem)

			items = append(items[:index-i], items[index+1-i:]...)
			fmt.Println(data.Yellow, "\tItem \""+delItem.ToString()+"\" has been deleted from todo list", data.Reset)
		}
		fmt.Println("")

		sort.Slice(items, func(i, j int) bool {
			return items[i].Priority < items[j].Priority
		})

		data.PrintTODO(items)

		data.SaveItems(todo_list_path, items)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
