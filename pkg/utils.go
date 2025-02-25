package pkg

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err!= nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func PrintTask(tasks []Task) {
	var status string

	if len(tasks) == 0 {
		fmt.Println("Tidak ada task.")
		return
	}

	for i, task:= range tasks {
		if !task.Completed {
			status = "Not Completed"
		} else {
			status = "Completed"
		}
		fmt.Printf("%d. %s - [%s]\n", i+1, task.Task, status)
	}
}