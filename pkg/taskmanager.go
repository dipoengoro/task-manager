package pkg

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	"encoding/json"
)

var tasks []Task

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

func ListTasks(completed bool)  {
	status := "Not Completed"
	if len(tasks) == 0 {
		fmt.Println("Tidak ada task.")
		return
	}

	for i, task:= range tasks {
		if task.Completed { status = "Completed" }
		if task.Completed == completed {
			fmt.Printf("%d. %s - [%s]\n", i+1, task.Task, status)
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