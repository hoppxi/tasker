package task

import "testing"

func TestNewTask(t *testing.T) {
    desc := "Test Task"
    task := NewTask(desc)
    if task.Description != desc {
        t.Errorf("Expected description %s, but got %s", desc, task.Description)
    }
}
