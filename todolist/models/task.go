package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Task struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Progress int    `json:"progress"`
	User     string `json:"user"`
	Desc     string `json:"desc"`
	Status   string `json:"status"`
}

func loadTasks() ([]Task, error) {
	if file, err := ioutil.ReadFile("data/tasks.json"); err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		} else {
			return nil, err
		}
	} else {
		var tasks []Task
		if err := json.Unmarshal(file, &tasks); err == nil {
			return tasks, nil
		} else {
			return nil, err
		}
	}
}

func storeTasks(tasks []Task) error {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("data/tasks.json", bytes, 0644)
}

func GetTasks() []Task {
	tasks, err := loadTasks()
	if err == nil {
		return tasks
	}
	panic(err)
}

func GetTaskId() (int, error) {
	tasks, err := loadTasks()
	if err != nil {
		return -1, err
	}
	var id int
	for _, task := range tasks {
		if id < task.Id {
			id = task.Id
		}
	}
	return id + 1, nil
}

func CreateTask(name, user, desc string) {
	id, err := GetTaskId()
	if err != nil {
		panic(err)
	}
	task := Task{
		Id:       id,
		Name:     name,
		User:     user,
		Desc:     desc,
		Progress: 0,
		Status:   "new",
	}
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}
	tasks = append(tasks, task)
	storeTasks(tasks)
}

func GetTaskById(id int) (Task, error) {
	tasks, err := loadTasks()
	if err == nil {
		for _, task := range tasks {
			if id == task.Id {
				fmt.Println("====")
				fmt.Println(task)
				fmt.Println("====")
				return task, nil
			}
		}

	}
	return Task{}, errors.New("not found")
}

func ModifyTask(id int, name, desc string, progress int, user, status string) {
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}
	newTasks := make([]Task, len(tasks))
	for i, task := range tasks {
		if id == task.Id {
			task.Name = name
			task.Desc = desc
			task.Progress = progress
			task.User = user
			task.Status = status
		}
		newTasks[i] = task
	}
	storeTasks(newTasks)
}

func DeleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}
	newTask := make([]Task, 0)
	for _, task := range tasks {
		if id != task.Id {
			newTask = append(newTask, task)
		}
	}
	storeTasks(newTask)
}
