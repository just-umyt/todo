package main

import (
	"flag"

	todo "github.com/umyt-king/gitTutorial"
)

const (
	todoFile = "todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	flag.Parse()

	todo := todo.Todos{}
}
