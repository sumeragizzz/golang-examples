package task

import (
	"os"
	"slices"
	"testing"
)

func TestList(t *testing.T) {
	store := createTempStoreFile(t)

	got, err := List(store)
	want := []Task{}

	if err != nil {
		t.Errorf("failed: %v", err)
	}

	if !slices.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestAdd(t *testing.T) {
	generateId = func() int64 {
		return 1234567890
	}

	store := createTempStoreFile(t)
	content := "test"

	got, err := Add(store, content)
	if err != nil {
		t.Errorf("failed: %v", err)
	}
	want := Task{Id: 1234567890, Content: "test", Status: "todo"}

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func createTempStoreFile(t *testing.T) string {
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer file.Close()

	file.WriteString("[]")

	return file.Name()
}
