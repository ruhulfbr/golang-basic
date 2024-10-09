package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Task struct representing a task
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// Task List - Slice of tasks
var tasks []Task

// Task ID Counter
var taskIDCounter int

// Function to add a task
func addTask(title string) {
	taskIDCounter++
	task := Task{
		ID:        taskIDCounter,
		Title:     title,
		Completed: false,
	}
	tasks = append(tasks, task)
	fmt.Println("Task added:", title)
}

// Function to list tasks
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available!")
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, task.Title, status)
	}
}

// Function to mark a task as done
func markTaskAsDone(taskID int) {
	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Completed = true
			fmt.Println("Task completed:", task.Title)
			return
		}
	}
	fmt.Println("Task ID not found!")
}

// Display the menu and options
func showMenu() {
	fmt.Println("\n=== Task App ===")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Mark Task as Done")
	fmt.Println("4. Exit")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()

		// Read user input for menu option
		fmt.Print("Choose an option: ")
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			// Add Task
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title := scanner.Text()
			addTask(title)

		case "2":
			// List Tasks
			listTasks()

		case "3":
			// Mark Task as Done
			fmt.Print("Enter task ID to mark as done: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid task ID!")
				continue
			}
			markTaskAsDone(id)

		case "4":
			// Exit
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
