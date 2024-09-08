package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

type task struct {
	Name        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TaskList struct {
	Tasks []task
}

func (tl *TaskList) search(task string) (bool, int) {
	taskNames := []string{}
	for _, t := range tl.Tasks {
		taskNames = append(taskNames, t.Name)
	}
	sort.Strings(taskNames)
	i := sort.SearchStrings(taskNames, task)
	if i < len(tl.Tasks) && taskNames[i] == task {
		return true, i
	}

	return false, -1
}

func (tl *TaskList) Add(name string) error{
	if found, _ := tl.search(name); found {
		return fmt.Errorf("task Already Exists: %s", name)
	}

	newTask := task{
		Name: name,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}
	tl.Tasks = append(tl.Tasks, newTask)
	return nil
}

func (tl *TaskList) Remove(id int) error {
	id--
	if id < 0 || id >= len(tl.Tasks) {
		return fmt.Errorf("item %d does not exist", id+1)
	}
	tl.Tasks = append(tl.Tasks[:id], tl.Tasks[id+1:]...)
	return nil
}

func (tl *TaskList) List(out io.Writer) error {
	for i, t := range tl.Tasks {
		fmt.Fprintf(out, "%d: %s", i+1, t.Name)
		if (t.Done) {
			fmt.Fprintf(out, " - Done\n")
		} else {
			fmt.Fprintf(out, " - Not Done\n")
		}
	}
	return nil
}

func (tl *TaskList) Save(tasksFile string) error {
	js, err := json.Marshal(tl)
	if err != nil {
		return err
	}

	return os.WriteFile(tasksFile, js, 0644)
}

func (tl *TaskList) Load(tasksFile string) error {
	tf, err := os.ReadFile(tasksFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(tf) == 0 {
		return nil
	}

	return json.Unmarshal(tf, tl)
}