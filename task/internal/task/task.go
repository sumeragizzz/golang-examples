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

func Update(storePath string, id int64, content string) (Task, error) {
	if storePath == "" {
		return Task{}, errors.New("invalid parameter. store: blank")
	}
	if id == 0 {
		return Task{}, errors.New("invalid parameter. id: zero")
	}
	if content == "" {
		return Task{}, errors.New("invalid parameter. content: blank")
	}

	tasks, err := load(storePath)
	if err != nil {
		return Task{}, err
	}

	t, index, found := search(tasks, id)
	if !found {
		return Task{}, fmt.Errorf("not found id: %d", id)
	}

	t.Content = content
	tasks[index] = t

	if err := save(storePath, tasks); err != nil {
		return Task{}, err
	}

	return t, nil
}

func Delete(storePath string, id int64) (Task, error) {
	if storePath == "" {
		return Task{}, errors.New("invalid parameter. store: blank")
	}
	if id == 0 {
		return Task{}, errors.New("invalid parameter. id: zero")
	}

	tasks, err := load(storePath)
	if err != nil {
		return Task{}, err
	}

	t, index, found := search(tasks, id)
	if !found {
		return Task{}, fmt.Errorf("not found id: %d", id)
	}

	removedTasks := append(tasks[:index], tasks[index+1:]...)

	if err := save(storePath, removedTasks); err != nil {
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

func search(tasks []Task, id int64) (Task, int, bool) {
	for i, t := range tasks {
		if t.Id == id {
			return t, i, true
		}
	}
	return Task{}, -1, false
}
