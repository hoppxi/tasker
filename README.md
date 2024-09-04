# Tasker

Hey there! Welcome to **Tasker**, a super simple CLI tool to help you manage your tasks. I made it just for fun.

## What It Does

- **Add**: Pop a new task into your list.
- **List**: See all your tasks.
- **Delete**: Get rid of a task you don't need.

## Setup

1. **Clone the Repo**

   ```bash
   git clone https://github.com/ermi111/tasker.git
   cd tasker
   ```

2. **Build It**

   ```bash
   go build -o tasker main.go
   ```

   Now you’ve got an executable called `tasker`.

## How to Use

### Add a Task

```bash
./tasker add "My awesome task"
```

### List All Tasks

```bash
./tasker list
```

### Delete a Task

```bash
./tasker delete "Task I don't need anymore"
```

## Notes

- Tasks are saved in `tasks.json` in the project folder.
- If you want to change the file name or path, you can do that in the code.

# License

Licensed under the [MIT license](https://opensource.org/license/mit). It's all yours and you can do whatever you want to do with it.