package cmd

import (
	"fmt"
	"todos/util"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "for testing purposes, ignore.",
	Long: ``,
  // Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
    database := util.OpenDB()
    rows, err := database.Query("SELECT id, todo_id, task, is_done FROM task WHERE todo_id = 2 LIMIT 1 OFFSET 2")
    if err != nil {
      panic(err)
    }
    
    var (
      id int
      todoId int
      content string
      isDone int
    )
    for rows.Next() {
      rows.Scan(&id, &todoId, &content, &isDone)
      fmt.Printf("ID: %d, TODO ID: %d, CONTENT: %s, IS_DONE: %d\n", id, todoId, content, isDone)
    }
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
