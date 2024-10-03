package cmd

import (
	"fmt"
	"log"
	"strconv"
	"todos/util"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: ``,
  Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
    processGet(args[0])
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

var (
    reset = "\033[0m" 
    orange = "\033[38;5;14m"
    red = "\033[38;5;196m"
    green = "\033[38;5;120m"
    strike = "\033[9m"
)

func processGet(title string) {
  database := util.OpenDB()
  rows, err := database.Query("SELECT id, title FROM todo WHERE title = ?", title)
  if err != nil {
    panic(err)
  }

  defer database.Close()
  
  var todoId int
  var titleD string
  
  if rows.Next() {
    rows.Scan(&todoId, &titleD)
  } else {
    log.Fatalf("%sError:%s No todo found matching the title: %s%s\n", red, reset, green, title)
  }

  rows, err = database.Query("SELECT task, is_done FROM task WHERE todo_id = ?", todoId)
  if err != nil {
    panic(err)
  }

  var tasksList []string
  var render_str string
  var task string
  var isDone int
  
  for i := 0; rows.Next(); i++ {
    err := rows.Scan(&task, &isDone)
    if err != nil {
      panic(err)
    }
    tasksList = append(tasksList, task)
    if i != 0 {
      render_str += "\n"
    }
    if isDone == 0 {
      render_str += strconv.Itoa(i+1) + ". " + task
    } else if isDone == 1 {
      render_str += strike + strconv.Itoa(i+1) + ". " + task + reset 
    } else {
      panic("Corrupted task list")
    }
  }
  
  pterm.DefaultBox.WithTitle(orange + title + reset).Printf("%v", render_str)
  fmt.Println()
}
