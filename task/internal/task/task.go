package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

func (t Task) Create(content string) Task {
	return Task{
		Id:      time.Now().UnixNano(),
		Content: content,
		Status:  "todo",
	}
}

func List(storePath string) ([]Task, error) {
	if storePath == "" {
		return nil, errors.New("invalid parameter. store: blank")
	}

	return load(storePath)
}

func Add(storePath string, content string) (Task, error) {
	if storePath == "" {
		return Task{}, errors.New("invalid parameter. store: blank")
	}
	if content == "" {
		return Task{}, errors.New("invalid parameter. content: blank")
	}

	t := Task{}.Create(content)

	tasks, err := load(storePath)
	_ = err

	tasks = append(tasks, t)

	if err := save(storePath, tasks); err != nil {
		return Task{}, err
	}

	return t, nil
}

func load(storePath string) ([]Task, error) {
	file, err := os.Open(storePath)
	if err != nil {
		return []Task{}, fmt.Errorf("file open error. store: %s: %w", storePath, err)
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return []Task{}, fmt.Errorf("json decode error. store: %s: %w", storePath, err)
	}

	return tasks, nil
}

func save(storePath string, tasks []Task) error {
	file, err := os.Create(storePath)
	if err != nil {
		return fmt.Errorf("file open error. store: %s: %w", storePath, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("json encode error. tasks: %+v: %w", tasks, err)
	}

	return nil
}
