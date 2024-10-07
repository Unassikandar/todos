package cmd

import (
	"log"
	"strconv"
	"todos/util"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
  Short: "Add a task to an existing todo list.",
	Long: `Args: <todoId>, <task>`,
  Args: cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
    todoId, err := strconv.ParseInt(args[0], 10, 64)
    if err != nil {
      panic(err)
    }
    processAdd(int(todoId), args[1])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func processAdd(todoId int, content string) {
  database := util.OpenDB()
  statement, err := database.Prepare("INSERT INTO task (todo_id, task, is_done) values (?, ?, ?)")
  if err != nil {
    panic(err)
  }
  _, err = statement.Exec(todoId, content, 0)
  if err != nil {
    log.Fatal("\033[38;5;196mError:\033[0m could not add task. Is todo id correct?")
  }
}
