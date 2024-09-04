package cmd

import (
    "github.com/spf13/cobra"
    "github.com/ermi111/tasker/pkg/task"
    "github.com/ermi111/tasker/pkg/storage"
    "fmt"
    "log"
)

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new task",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 1 {
            fmt.Println("Please provide a task description.")
            return
        }
        description := args[0]
        t := task.NewTask(description)
        err := storage.AddTask(t)
        if err != nil {
            log.Fatalf("Error adding task: %v", err)
        }
        fmt.Println("Task added successfully!")
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
