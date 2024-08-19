# Task Tracker

Task Tracker is a simple command-line tool for managing tasks. You can add, update, list, and delete tasks, as well as mark them as in-progress or done.

## Features

- **Add Task:** Add a new task with a description.
- **Update Task:** Update the description of an existing task.
- **List Tasks:** List all tasks or filter tasks by their status (`todo`, `in-progress`, `done`).
- **Delete Task:** Delete a task by its ID.
- **Mark Task as In-Progress:** Mark a task as `in-progress`.
- **Mark Task as Done:** Mark a task as `done`.
## Usage

```bash
./task-tracker add "Your task description here"
```
```bash
./task-tracker list
```
```bash
./task-tracker list [status]
```

Replace `[status]` with `todo`, `in-progress`, or `done`.

```bash
./task-tracker update [task_id] "New task description"
```
```bash
./task-tracker delete [task_id]
```
```bash
./task-tracker mark-in-progress [task_id]
```
```bash
./task-tracker mark-done [task_id]
```
