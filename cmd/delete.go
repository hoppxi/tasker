package cmd

import (
    "github.com/spf13/cobra"
    "github.com/ermi111/tasker/pkg/storage"
    "fmt"
    "log"
)

var deleteCmd = &cobra.Command{
    Use:   "delete",
    Short: "Delete a task",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 1 {
            fmt.Println("Please provide a task description to delete.")
            return
        }
        description := args[0]
        err := storage.DeleteTask(description)
        if err != nil {
            log.Fatalf("Error deleting task: %v", err)
        }
        fmt.Println("Task deleted successfully!")
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}
