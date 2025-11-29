package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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

	case "show":

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
		updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
		store := updateCommand.String("store", "store.json", "store file path")
		updateCommand.Usage = func() {
			fmt.Printf("Usage: task update [optoins] <id> <content>\n")
			fmt.Printf("Options:\n")
			updateCommand.PrintDefaults()
		}
		updateCommand.Parse(os.Args[2:])
		if len(updateCommand.Args()) < 2 {
			fmt.Println("content, id is required")
			updateCommand.Usage()
			os.Exit(1)
		}
		id, err := strconv.ParseInt(updateCommand.Args()[0], 10, 64)
		if err != nil {
			fmt.Printf("invalid id. id: %s\n", updateCommand.Args()[0])
			os.Exit(1)
		}
		content := updateCommand.Args()[1]

		t, err := task.Update(*store, id, content)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%+v\n", t)

	case "delete":
		deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
		store := deleteCommand.String("store", "store.json", "store file path")
		deleteCommand.Usage = func() {
			fmt.Printf("Usage: task delete [optoins] <id>\n")
			fmt.Printf("Options:\n")
			deleteCommand.PrintDefaults()
		}
		deleteCommand.Parse(os.Args[2:])
		if len(deleteCommand.Args()) < 1 {
			fmt.Println("id is required")
			deleteCommand.Usage()
			os.Exit(1)
		}
		id, err := strconv.ParseInt(deleteCommand.Args()[0], 10, 64)
		if err != nil {
			fmt.Printf("invalid id. id: %s\n", deleteCommand.Args()[0])
			os.Exit(1)
		}

		t, err := task.Delete(*store, id)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%+v\n", t)

	case "done":
	default:
		fmt.Printf("invalid command. command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
