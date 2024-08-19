package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/bedminer1/chapter1todo"
)

// Hardcode file name
const todoFileName = ".todo.json"

func main() {
	task := flag.String("task", "", "Task to be included in the to-do list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to bge completed")

	flag.Parse()

	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// switch statement based on flags
	switch {
	case *list:
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}