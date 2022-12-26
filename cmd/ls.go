/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/NickAhn/todo/data"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
/*
By default, ls displays todo list by creation order.
*/
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List todo list",
	Long: `The list command prints out the todo list.
	
	Usage: todo ls <flags>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		items, _ := data.ReadItems(todo_list_path)
		fmt.Println("TODO (" + fmt.Sprint(len(items)) + "):")
		for i, item := range items {
			fmt.Println("  ", fmt.Sprint(i)+".", item.Text, "(p"+fmt.Sprint(item.Priority)+")")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
