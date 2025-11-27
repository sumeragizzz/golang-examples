package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func List(storePath string) ([]Task, error) {
	if storePath == "" {
		return nil, errors.New("invalid parameter. storePath: blank")
	}
	return load(storePath), nil
}

func Add(storePath string, content string) {
}

func load(storePath string) []Task {
	file, err := os.Open(storePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		panic(err)
	}

	return tasks
}

func save(storePath string, tasks []Task) {

}

type Task struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

func (t Task) Create(content string) {
	var task Task = Task{
		Id:      1,
		Content: content,
		Status:  "todo",
	}

	data, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}

	// TODO JSONファイルに保存
	fmt.Println(string(data))
}
