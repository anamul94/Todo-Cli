package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github/anamul/todo-app"
	"io"
	"os"
	"strings"
)

const (
	todoFile = ".tood.json"
)

func main() {
	add := flag.Bool("add", false, "add a new tood")
	complete := flag.Int("complete", 0, "mark todo as completed")
	del := flag.Int("del", 0, "delete todo")
	list := flag.Bool("list", false, "print todo list")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			panic(err.Error())
		}
		todos.Add(task)
		store(todos)
	case *complete > 0:
		err := todos.MarkDone(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		store(todos)

	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		store(todos)

	case *list:
		todos.Print()

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}

}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return " ", err
	}

	if len(scanner.Text()) == 0 {
		return "", errors.New("emtpty todo is not allowed")
	}

	return scanner.Text(), nil
}

func store(todos *todo.Todos) {
	err := todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
