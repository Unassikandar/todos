package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a todo list or a task",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
    //TODO Process Search
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
