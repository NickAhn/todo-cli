/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/NickAhn/todo/data"
	"github.com/spf13/cobra"
)

// undoCmd represents the undo command
var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo last command",
	Long:  `Undo last change in todo list.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("undo called")
		items, _ := data.ReadItems(todo_list_path)
		prevItems, _ := data.ReadItems(undo_list_path)

		data.SaveItems(todo_list_path, prevItems)
		data.SaveItems(undo_list_path, items)

		fmt.Println(data.Yellow, "\n\tUndo latest changes.", data.Reset+"\n")

		data.PrintTODO(prevItems)
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
