package cmd

import (
	. "github.com/nocd5/todo/internal/common"
	"os"
	"path"
)

var (
	todoList TodoList
	dbFile   = path.Join(os.Getenv("HOME"), ".todo-db.json")
	cfgFile  string
)
