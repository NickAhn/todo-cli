/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
		temp, _ := data.ReadItems(todo_list_path)
		index, _ := strconv.Atoi(args[0])
		items := append(temp[:index], temp[index+1:]...)

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
