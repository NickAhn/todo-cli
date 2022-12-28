/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"

	"github.com/NickAhn/todo/data"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new todo item",
	Long:  `Add creates a new todo item to the list.`,
	Run: func(cmd *cobra.Command, args []string) {
		items, _ := data.ReadItems(todo_list_path)
		for _, v := range args {
			item := data.Item{Text: v}
			item.SetPriority(priority)
			items = append(items, item)
			fmt.Println("\t\"" + item.ToString() + "\" has been added to the list")
		}

		// sort by Priority
		sort.Slice(items, func(i, j int) bool {
			return items[i].Priority < items[j].Priority
		})

		data.SaveItems(todo_list_path, items)
	},
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 4, "Priority:1,2,3")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
