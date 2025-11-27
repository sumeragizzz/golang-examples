package main

import (
	"flag"
	"fmt"
	"os"
	"task/internal/task"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("invalid parameter. args: %v\n", os.Args)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listCommand := flag.NewFlagSet("list", flag.ExitOnError)
		store := listCommand.String("store", "store.json", "store file path")
		listCommand.Usage = func() {
			fmt.Printf("Usage: task list [optoins]\n")
			fmt.Printf("Options:\n")
			listCommand.PrintDefaults()
		}
		listCommand.Parse(os.Args[2:])

		tasks, err := task.List(*store)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		for _, t := range tasks {
			fmt.Printf("%+v\n", t)
		}

	case "add":
		addCommand := flag.NewFlagSet("add", flag.ExitOnError)
		store := addCommand.String("store", "store.json", "store file path")
		addCommand.Usage = func() {
			fmt.Printf("Usage: task add [optoins] <content>\n")
			fmt.Printf("Options:\n")
			addCommand.PrintDefaults()
		}
		addCommand.Parse(os.Args[2:])
		if len(addCommand.Args()) < 1 {
			fmt.Println("content is required")
			addCommand.Usage()
			os.Exit(1)
		}
		content := addCommand.Args()[0]

		t, err := task.Add(*store, content)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%+v\n", t)

	case "update":
	case "delete":
	case "done":
	default:
		fmt.Printf("invalid command. command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
