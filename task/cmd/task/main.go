package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Invalid parameter\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
	case "add":
		addCommand := flag.NewFlagSet("add", flag.ExitOnError)
		status := addCommand.String("status", "todo", "status(ex: todo, doing, done)")
		addCommand.Usage = func() {
			fmt.Printf("Usage: task add [optoins] <content>\n")
			fmt.Printf("Options:\n")
			addCommand.PrintDefaults()
		}

		addCommand.Parse(os.Args[2:])
		if len(addCommand.Args()) < 1 {
			fmt.Println("タスク内容は必須です。")
			addCommand.Usage()
			os.Exit(1)
		}
		content := addCommand.Args()[0]
		fmt.Printf("add: content: %s, status: %s\n", content, *status)

	case "update":
	case "delete":
	case "done":
	default:
		fmt.Printf("コマンドが不正です。 %s\n", os.Args[1])
		os.Exit(1)
	}
}
