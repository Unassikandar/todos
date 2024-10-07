package cmd

import (
	"fmt"
	"todos/util"

	"github.com/spf13/cobra"
  "github.com/lithammer/fuzzysearch/fuzzy"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a todo list or a task",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
    tasks := getAllTasks()
    search := args[0]
    matchingTasks := fuzzySearch(search, tasks)
    if len(matchingTasks) == 0 {
      fmt.Println("\033[38;5;14mNo results found\033[0m")
      return
    }
    fmt.Println("\033[38;5;14mSearch results: \033[0m")
    for _, task := range matchingTasks {
      fmt.Printf(". %s: %s\n", task.Title, task.Task)
    }
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

type Task struct {
  Title string
  Task string
}

func getAllTasks() []Task {
  database := util.OpenDB()
  
  rows, err := database.Query(`
    SELECT todo.title, task.task
    FROM todo
    JOIN task
    ON todo.id = task.todo_id
    `)

  if err != nil {
    panic(err)
  }

  defer database.Close()

  var todoTitle string
  var task string
  var tasks []Task

  for rows.Next() {
    rows.Scan(&todoTitle, &task)
    task := Task{Title: todoTitle, Task: task}
    tasks = append(tasks, task)
    // fmt.Printf("%s: %s\n", todoTitle, task)
  }
  rows.Close()
  database.Ping()
  // fmt.Println("\033[38;5;14mFetch completed\n\n\033[0m")

  return tasks
}

func fuzzySearch(search string, tasks []Task) []Task {
  var matchingTasks []Task
  for _, task := range tasks {
    if fuzzy.Match(search, task.Task) {
      matchingTasks = append(matchingTasks, task)
    } else if fuzzy.Match(search, task.Title) {
      matchingTasks = append(matchingTasks, task)
    }
  }
  return matchingTasks
}

// implement a fuzzy search function that takes an array of strings and a search string as arguments
// and returns an array of strings that match the search string
// use Levenshtein distance to calculate the similarity between the search string and the strings in the array
// if the distance is less than 3, consider the string a match
// return an array of matching strings
// func fuzzySearch(search string, tasks []Task) []Task {
//   var matchingTasks []Task
//   for _, task := range tasks {
//     if levenshtein(search, task.Task) < 3 {
//       matchingTasks = append(matchingTasks, task)
//     } else if levenshtein(search, task.Title) < 3 {
//       matchingTasks = append(matchingTasks, task)
//     }
//   }
//   return matchingTasks
// }

