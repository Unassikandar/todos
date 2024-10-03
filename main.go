package main

import (
	"todos/cmd"
  "todos/util"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
  database := util.OpenDB()
  statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS todo (id INTEGER PRIMARY KEY, title TEXT UNIQUE)")
  if err != nil {
    panic(err)
  }
  statement.Exec()

  statement, err = database.Prepare(
    `CREATE TABLE IF NOT EXISTS task (
      id INTEGER PRIMARY KEY, 
      todo_id INTEGER, 
      task TEXT, 
      is_done INTEGER,
    FOREIGN KEY (todo_id) REFERENCES todo(id)
    )`,
  )
  if err != nil {
    panic(err)
  }
  _, err = statement.Exec()
  if err != nil {
    panic(err)
  }

	cmd.Execute()
}
