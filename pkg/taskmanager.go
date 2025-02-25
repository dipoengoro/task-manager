package pkg

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	"encoding/json"
)

var tasks []Task
var FilteredTasks []Task

func AddTask(task string, description string, deadline string) {
	if task == "" {
		fmt.Println("Error: Task tidak boleh kosong")
		return
	}

	newTask:= Task{
		ID: uuid.New(),
		Task: task,
		Description: description,
		Deadline: deadline,
		Completed: false,
	}

	tasks = append(tasks, newTask)

	SaveTasks()

	fmt.Println("Task berhasil ditambahkan!")
}

func ListTasks(filterCompleted *bool)  {
	if filterCompleted == nil {
		FilteredTasks = tasks
		fmt.Println("pake nil")
	}

	for _, task:= range tasks {
		if task.Completed == *filterCompleted {
			FilteredTasks = append(FilteredTasks, task)
		}
	}
}

func CompletedTask(id string)  {
	// TODO("implement Completed")
}

func DeleteTask(id string)  {
	// TODO("Implement Delete")
}

func SaveTasks()  {
	data, err:= json.Marshal(tasks)
	CheckError(err)

	err = os.WriteFile("tasks.json", data, 0644)
	CheckError(err)
}

func LoadTasks() {
	data, err:= os.ReadFile("tasks.json")
	if err!= nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error:", err)
		}
		return
	}

	err = json.Unmarshal(data, &tasks)
	CheckError(err)
}