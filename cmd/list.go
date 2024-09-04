package cmd

import (
    "github.com/spf13/cobra"
    "github.com/ermi111/tasker/pkg/storage"
    "fmt"
    "log"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        tasks, err := storage.ListTasks()
        if err != nil {
            log.Fatalf("Error listing tasks: %v", err)
        }
        if len(tasks) == 0 {
            fmt.Println("No tasks found.")
            return
        }
        for _, t := range tasks {
            fmt.Printf("- %s\n", t.Description)
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
