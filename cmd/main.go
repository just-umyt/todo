package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/just-umyt/cli/pkg"
)

const (
	todoFile = ".todo.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	ls := flag.Bool("ls", false, "show a list")
	complete := flag.Int("comp", 0, "mark a todo as complete")
	del := flag.Int("del", 0, "delete a todo")
	flags := flag.Bool("flags", false, "show all flags")

	flag.Parse()

	todos := &pkg.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task := getInput(flag.Args()...)
		todos.Add(task)
		todos.Store(todoFile)
	case *complete > 0:
		todos.Complete(*complete)
		todos.Store(todoFile)
	case *del > 0:
		todos.Delete(*del)
		todos.Store(todoFile)
	case *ls:
		todos.Print()
	case *flags:
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Println("Flag: ", f.Name, " | Description:  ", f.Usage,)
		})
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}
}

func getInput(args ...string) string {
	if len(args) <= 0 {
		fmt.Fprintln(os.Stderr, errors.New("empty string"))
		os.Exit(1)
	}

	return strings.Join(args, " ")

}
