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
		_, store := createCommand("list", "[optoins]", os.Args[2:])

		tasks, err := task.List(*store)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, t := range tasks {
			fmt.Printf("%+v\n", t)
		}

	case "add":
		command, store := createCommand("add", "[optoins] <content>", os.Args[2:])
		if len(command.Args()) < 1 {
			fmt.Fprintln(os.Stderr, "content is required")
			command.Usage()
			os.Exit(1)
		}
		content := command.Args()[0]

		t, err := task.Add(*store, content)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("%+v\n", t)

	case "update":
		command, store := createCommand("update", "[optoins] <id> <content>", os.Args[2:])
		if len(command.Args()) < 2 {
			fmt.Fprintln(os.Stderr, "content, id is required")
			command.Usage()
			os.Exit(1)
		}
		id, err := strconv.ParseInt(command.Args()[0], 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		content := command.Args()[1]

		t, err := task.Update(*store, id, content)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("%+v\n", t)

	case "delete":
		command, store := createCommand("delete", "[optoins] <id>", os.Args[2:])
		if len(command.Args()) < 1 {
			fmt.Fprintln(os.Stderr, "id is required")
			command.Usage()
			os.Exit(1)
		}
		id, err := strconv.ParseInt(command.Args()[0], 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		t, err := task.Delete(*store, id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("%+v\n", t)

	case "doing":
		command, store := createCommand("doing", "[optoins] <id>", os.Args[2:])
		if len(command.Args()) < 1 {
			fmt.Println("id is required")
			command.Usage()
			os.Exit(1)
		}
		id, err := strconv.ParseInt(command.Args()[0], 10, 64)
		if err != nil {
			fmt.Printf("invalid id. id: %s\n", command.Args()[0])
			os.Exit(1)
		}

		t, err := task.Doing(*store, id)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%+v\n", t)

	case "done":
		command, store := createCommand("done", "[optoins] <id>", os.Args[2:])
		if len(command.Args()) < 1 {
			fmt.Println("id is required")
			command.Usage()
			os.Exit(1)
		}
		id, err := strconv.ParseInt(command.Args()[0], 10, 64)
		if err != nil {
			fmt.Printf("invalid id. id: %s\n", command.Args()[0])
			os.Exit(1)
		}

		t, err := task.Done(*store, id)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%+v\n", t)

	default:
		fmt.Printf("invalid command. command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func createCommand(name string, usage string, args []string) (*flag.FlagSet, *string) {
	command := flag.NewFlagSet(name, flag.ExitOnError)
	store := command.String("store", "store.json", "store file path")
	command.Usage = func() {
		fmt.Printf("Usage: task %s %s\n", name, usage)
		fmt.Printf("Options:\n")
		command.PrintDefaults()
	}
	command.Parse(args)
	return command, store
}
