/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/NickAhn/todo/data"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit todo list item",
	Long:  `Edit item in todo list by index`,
	Run: func(cmd *cobra.Command, args []string) {
		items, _ := data.ReadItems(todo_list_path)
		index, _ := strconv.Atoi(args[0])
		items[index].Text = args[1]
		data.SaveItems(todo_list_path, items)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}