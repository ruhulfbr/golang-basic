package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID     int
	title  string
	status bool
}

var tasks []Task

var taskCounter int

func showOption() {
	fmt.Println()
	fmt.Println("Choose Option:")
	fmt.Println("1. Create Task")
	fmt.Println("2. Show Tasks")
	fmt.Println("3. Update Task")
	fmt.Println("4. Exit")
}

func createTask(title string) {
	taskCounter++
	task := Task{taskCounter, title, false}
	tasks = append(tasks, task)
	fmt.Println("Task created, ", title)
}

func showTasks() {

	fmt.Println()
	if len(tasks) == 0 {
		fmt.Println("No tasks found")

		return
	}

	fmt.Println("Task List : ")

	for _, task := range tasks {
		var status string = "Pending"
		if task.status {
			status = "Completed"
		}

		fmt.Println(task.ID, ". ", task.title, status)
	}
}

func markTaskComplete(ID int) {
	fmt.Println()

	for i, task := range tasks {
		if task.ID == ID {
			tasks[i].status = true
			fmt.Println("Task completed:", task.title)
			return
		}
	}

	fmt.Println("Invalid id, task not found with this ID")
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		showOption()

		fmt.Print("Choose an option: ")
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			fmt.Print("Enter Task title: ")
			scanner.Scan()
			title := scanner.Text()

			createTask(title)

		case "2":
			showTasks()

		case "3":
			fmt.Print("Enter Task Id: ")
			scanner.Scan()
			StrId := scanner.Text()

			taskId, err := strconv.Atoi(StrId)

			if err != nil {
				fmt.Println("Invalid Task Id")
				continue
			}

			markTaskComplete(taskId)

		case "4":
			fmt.Println("Goodbye!!")
			fmt.Println("Goodbye!!")
			return

		default:
			fmt.Println("Invalid option")
		}
	}

}
