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

var generateId = func() int64 {
	return time.Now().UnixNano()
}

func (t Task) Create(content string) Task {
	return Task{
		Id:      generateId(),
		Content: content,
		Status:  "todo",
	}
}

func List(store string) ([]Task, error) {
	if store == "" {
		return nil, errors.New("invalid parameter. store: blank")
	}

	return load(store)
}

func Add(store string, content string) (Task, error) {
	if store == "" {
		return Task{}, errors.New("invalid parameter. store: blank")
	}
	if content == "" {
		return Task{}, errors.New("invalid parameter. content: blank")
	}

	tasks, _ := load(store)

	t := Task{}.Create(content)
	tasks = append(tasks, t)

	if err := save(store, tasks); err != nil {
		return Task{}, err
	}

	return t, nil
}

func Update(store string, id int64, content string) (Task, error) {
	if store == "" {
		return Task{}, errors.New("invalid parameter. store: blank")
	}
	if id == 0 {
		return Task{}, errors.New("invalid parameter. id: zero")
	}
	if content == "" {
		return Task{}, errors.New("invalid parameter. content: blank")
	}

	return update(store, id, content, "")
}

func Delete(store string, id int64) (Task, error) {
	if store == "" {
		return Task{}, errors.New("invalid parameter. store: blank")
	}
	if id == 0 {
		return Task{}, errors.New("invalid parameter. id: zero")
	}

	tasks, err := load(store)
	if err != nil {
		return Task{}, err
	}

	t, index, found := search(tasks, id)
	if !found {
		return Task{}, fmt.Errorf("not found id: %d", id)
	}

	removedTasks := append(tasks[:index], tasks[index+1:]...)

	if err := save(store, removedTasks); err != nil {
		return Task{}, err
	}

	return t, nil
}

func Doing(store string, id int64) (Task, error) {
	if store == "" {
		return Task{}, errors.New("invalid parameter. store: blank")
	}
	if id == 0 {
		return Task{}, errors.New("invalid parameter. id: zero")
	}

	return update(store, id, "", "doing")
}

func Done(store string, id int64) (Task, error) {
	if store == "" {
		return Task{}, errors.New("invalid parameter. store: blank")
	}
	if id == 0 {
		return Task{}, errors.New("invalid parameter. id: zero")
	}

	return update(store, id, "", "done")
}

func load(store string) ([]Task, error) {
	file, err := os.Open(store)
	if err != nil {
		return []Task{}, fmt.Errorf("file open error. store: %s: %w", store, err)
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return []Task{}, fmt.Errorf("json decode error. store: %s: %w", store, err)
	}

	return tasks, nil
}

func save(store string, tasks []Task) error {
	file, err := os.Create(store)
	if err != nil {
		return fmt.Errorf("file open error. store: %s: %w", store, err)
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

func update(store string, id int64, content string, status string) (Task, error) {
	tasks, err := load(store)
	if err != nil {
		return Task{}, err
	}

	t, index, found := search(tasks, id)
	if !found {
		return Task{}, fmt.Errorf("not found id: %d", id)
	}

	if content != "" {
		t.Content = content
	}
	if status != "" {
		t.Status = status
	}
	tasks[index] = t

	if err := save(store, tasks); err != nil {
		return Task{}, err
	}

	return t, nil
}
