package storage

import (
    "github.com/ermi111/tasker/pkg/task"
    "os"
    "testing"
)

func TestAddAndListTasks(t *testing.T) {
    // Clear tasks file before test
    os.Remove("tasks.json")

    testTask := task.NewTask("Test Task")
    err := AddTask(testTask)
    if err != nil {
        t.Fatalf("Error adding task: %v", err)
    }

    tasks, err := ListTasks()
    if err != nil {
        t.Fatalf("Error listing tasks: %v", err)
    }

    if len(tasks) != 1 || tasks[0].Description != "Test Task" {
        t.Errorf("Expected 1 task with description 'Test Task', but got %v", tasks)
    }
}

func TestDeleteTask(t *testing.T) {
    // Clear tasks file before test
    os.Remove("tasks.json")

    testTask := task.NewTask("Test Task")
    AddTask(testTask)

    err := DeleteTask("Test Task")
    if err != nil {
        t.Fatalf("Error deleting task: %v", err)
    }

    tasks, err := ListTasks()
    if err != nil {
        t.Fatalf("Error listing tasks: %v", err)
    }

    if len(tasks) != 0 {
        t.Errorf("Expected no tasks, but got %v", tasks)
    }
}
