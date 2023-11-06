package main

import (
	"flag"
)

const (
	todoFile = "todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	flag.Parse()

}
