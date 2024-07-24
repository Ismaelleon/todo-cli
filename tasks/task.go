package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Task struct {
	Name string
	Done bool
}

var Tasks []Task

func Add (args []string) {
	// If command has no argument show error
	if len(args) != 1 {
		errorString := fmt.Sprintf("todo-cli: Wrong format for argument in 'add' command\n\nExample:\ntodo-cli add \"Buy dog food\"")

		fmt.Println(errors.New(errorString))
	}

	// Get task name
	taskName := args[0]
	
	// Add new task to tasks slice
	Tasks = append(Tasks, Task{Name: taskName, Done: false})

	// Save tasks
	Save()

	// Show message
	fmt.Printf("Your task '%s' was added!\n", taskName)

	// List tasks
	List()
}

func Set (args []string) {
	// If command has no argument show error
	if len(args) != 1 {
		errorString := fmt.Sprintf("todo-cli: Wrong format for argument in 'set' command\n\nExample:\ntodo-cli set 0")

		fmt.Println(errors.New(errorString))
	}

	// Get the task by index
	i, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("todo-cli: Argument for 'set' command is not a task index (a number)")
	}

	// Change the state of the task
	if Tasks[i].Done {
		Tasks[i].Done = false
	} else {
		Tasks[i].Done = true 
	}

	// Save tasks
	Save()

	// List tasks
	List()
}

func Delete(args []string) {
	// If command has no argument show error
	if len(args) != 1 {
		errorString := fmt.Sprintf("todo-cli: Wrong format for argument in 'delete' command\n\nExample:\ntodo-cli delete 0")

		fmt.Println(errors.New(errorString))
	}

	// Get the task by index
	i, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("todo-cli: Argument for 'set' command is not a task index (a number)")
	}

	// Delete the task
	Tasks = append(Tasks[:i], Tasks[i+1:]...)

	// Save tasks
	Save()

	// Show message
	fmt.Printf("Task at index %d deleted!\n", i)

	// List tasks
	List()
}

func List() {
	// Loop through all tasks
	for i := 0; i < len(Tasks); i++ {
		if Tasks[i].Done {
			fmt.Printf("%d \u2192 %s \u2713\n", i, Tasks[i].Name)
		} else {
			fmt.Printf("%d \u2192 %s\n", i, Tasks[i].Name)
		}
	}
}

func Load() {
	// Get the home directory
	homeDir, err := os.UserHomeDir()

	// Check for error
	if err != nil {
		log.Fatal(err)
	}

	// Read tasks file (not throwing error if file not found)
	rawData, _ := os.ReadFile(fmt.Sprintf("%s/tasks.json", homeDir))

	// Convert json to slice and update the tasks slice
	json.Unmarshal([]byte(rawData), &Tasks)
}

func Save() {
	// Convert tasks slice to json
	json, _ := json.Marshal(Tasks)

	// Convert json to byte array
	jsonString := string(json)
	fileContent := []byte(jsonString)

	// Get the home directory
	homeDir, err := os.UserHomeDir()

	// Check for error
	if err != nil {
		log.Fatal(err)
	}

	// Write the file
	err = os.WriteFile(fmt.Sprintf("%s/tasks.json", homeDir), fileContent, 0644)

	// Check for error
	if err != nil {
		fmt.Println(errors.New("Error saving tasks file"))
	}
}
