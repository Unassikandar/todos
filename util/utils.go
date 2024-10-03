package util

import "database/sql"

func OpenDB() *sql.DB {
  database, err := sql.Open("sqlite3", "file:todos.db?_foreign_keys=on")
  if err != nil {
    panic("Could not open/create db")
  }
  return database
}
