package storage

import (
    "github.com/ermi111/tasker/pkg/task"
    "os"
    "encoding/json"
    "errors"
)

var tasksFile = "tasks.json"

func AddTask(t task.Task) error {
    tasks, err := loadTasks()
    if err != nil {
        return err
    }
    tasks = append(tasks, t)
    return saveTasks(tasks)
}

func ListTasks() ([]task.Task, error) {
    return loadTasks()
}

func DeleteTask(description string) error {
    tasks, err := loadTasks()
    if err != nil {
        return err
    }
    var updatedTasks []task.Task
    found := false
    for _, t := range tasks {
        if t.Description != description {
            updatedTasks = append(updatedTasks, t)
        } else {
            found = true
        }
    }
    if !found {
        return errors.New("task not found")
    }
    return saveTasks(updatedTasks)
}

func loadTasks() ([]task.Task, error) {
    var tasks []task.Task
    file, err := os.Open(tasksFile)
    if err != nil {
        if os.IsNotExist(err) {
            return tasks, nil
        }
        return nil, err
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&tasks)
    if err != nil {
        return nil, err
    }
    return tasks, nil
}

func saveTasks(tasks []task.Task) error {
    file, err := os.Create(tasksFile)
    if err != nil {
        return err
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    return encoder.Encode(tasks)
}
