package cmd

import (
	"fmt"
	"todos/util"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: ``,
  // Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
    database := util.OpenDB()
    rows, err := database.Query("SELECT * from task")
    if err != nil {
      panic(err)
    }
    
    var (
      id int
      todoId int
      content string
    )
    for rows.Next() {
      rows.Scan(&id, &todoId, &content)
      fmt.Printf("ID: %d, TODO ID: %d, CONTENT: %s\n", id, todoId, content)
    }
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
