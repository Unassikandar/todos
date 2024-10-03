package cmd

import (
	"fmt"
	"strconv"
	"todos/util"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
  Long: `Marks a task as completed. 
Syntax: todo do <title> <task number>`,
  Args: cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
    pos, err := strconv.ParseInt(args[1], 10, 64)
    if err != nil {
      panic(err)
    }
    processDo(args[0], int(pos))
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}

func processDo(title string, pos int) {
  database := util.OpenDB()
  rows, err := database.Query(`
    SELECT id 
    FROM task 
    WHERE todo_id = (SELECT id FROM todo WHERE title = ?) 
    LIMIT 1 OFFSET ?`, title, pos-1)

  if err != nil {
    panic(err)
  }
  
  var taskId int

  if rows.Next() {
    rows.Scan(&taskId)
  }

  rows.Close()

  database.Ping()

  result, err := database.Prepare(`
    UPDATE task
    SET is_done = ?
    WhERE id = ?
    `)

  if err != nil {
    panic(err)
  }

  _, err = result.Exec(1, taskId)
  if err != nil {
    panic(err)
  }
  fmt.Println("Done")
}
