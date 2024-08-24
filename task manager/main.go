package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task_manager/models"
	"task_manager/storage"
)

func main() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)
		if len(parts) < 1 {
			continue
		}

		switch parts[0] {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Usage: add <task_name>")
				continue
			}
			newTask := models.Task{
				ID:        len(tasks) + 1,
				Name:      strings.Join(parts[1:], " "),
				Completed: false,
			}
			tasks = append(tasks, newTask)
			if err := storage.SaveTasks(tasks); err != nil {
				fmt.Println("Error saving tasks:", err)
			}

		case "list":
			for _, task := range tasks {
				status := "incomplete"
				if task.Completed {
					status = "completed"
				}
				fmt.Printf("%d: %s [%s]\n", task.ID, task.Name, status)
			}

		case "complete":
			if len(parts) < 2 {
				fmt.Println("Usage: complete <task_id>")
				continue
			}
			id := parts[1]
			for i, task := range tasks {
				if fmt.Sprint(task.ID) == id {
					tasks[i].Completed = true
					if err := storage.SaveTasks(tasks); err != nil {
						fmt.Println("Error saving tasks:", err)
					}
					break
				}
			}

		case "delete":
			if len(parts) < 2 {
				fmt.Println("Usage: delete <task_id>")
				continue
			}
			id := parts[1]
			for i, task := range tasks {
				if fmt.Sprint(task.ID) == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					if err := storage.SaveTasks(tasks); err != nil {
						fmt.Println("Error saving tasks:", err)
					}
					break
				}
			}

		default:
			fmt.Println("Unknown command. Available commands: add, list, complete, delete.")
		}
	}
}
