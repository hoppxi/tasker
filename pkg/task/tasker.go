package task

type Task struct {
    Description string
}

func NewTask(description string) Task {
    return Task{Description: description}
}
