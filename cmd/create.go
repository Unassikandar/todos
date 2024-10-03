package cmd

import (
	"fmt"
	"todos/util"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new TODO list",
	Long: ``,
  Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
    processCreate(args[0])
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func processCreate(title string) {
  database := util.OpenDB()
  statement, err := database.Prepare("INSERT INTO todo (title) values (?)")
  if err != nil {
    panic(err)
  }
  statement.Exec(title)
  fmt.Printf("Todo created, title: %s \n", title)
}
