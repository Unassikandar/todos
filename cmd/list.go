package cmd

import (
	"strconv"
	"todos/util"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

type Todo struct {
  id int
  title string
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available todo lists",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
	  listTodos()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}


func listTodos() {
  database := util.OpenDB()
  rows, err := database.Query("SELECT id, title FROM todo")
  if err != nil {
    panic(err)
  }
  var id int
  var title string
  var list []Todo
  for rows.Next() {
    rows.Scan(&id, &title)
    list = append(list, Todo{id, title})
  }

  tableData := pterm.TableData{[]string{"index", "title", "tags"}}
  for i := range len(list) {
    item := list[i]
    tableData = append(tableData, []string{strconv.Itoa(item.id), item.title})
  }
  // pterm.DefaultArea.WithCenter().Start()
  pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
}
