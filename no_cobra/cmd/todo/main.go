package main

import (
	"fmt"
	"os"
	"strings"

	todo "github.com/bedminer1/chapter1todo"
)

// Hardcode file name
const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	// Use Get method to read from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on num of arguments
	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}

	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}