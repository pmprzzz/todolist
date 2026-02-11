package functions

import (
	"CLI-TodoList/task_struct"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func CreateIfNotEx() {
	if _, err := os.Stat("task.json"); os.IsNotExist(err) {
		os.WriteFile("task.json", []byte("[]"), 0777)
	}
}

func Clean() {

	var err error
	var file *os.File

	if file, err = os.OpenFile("task.json", os.O_TRUNC, 0777); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

}

func Constructor(len int, task string) task_struct.Task {

	var err error
	var file *os.File

	if file, err = os.Open("task.json"); err != nil {
		log.Fatal(err)
	}

	var tasks []task_struct.Task

	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	file.Close()

	var temp_max, id_int int = 0, 0

	for _, task := range tasks {
		if id_int, err = strconv.Atoi(task.Id); err != nil {
			log.Fatal(err)
		}
		if id_int > temp_max {
			temp_max = id_int
		}
	}

	task_res := task_struct.Task{
		Id:          strconv.Itoa(temp_max + 1),
		Description: task,
		Status:      "todo",
		CreatedAt:   time.Now().Format("02-01-2006 15.04.05"),
		UpdatedAt:   time.Now().Format("02-01-2006 15.04.05"),
	}
	return task_res
}

func TaskAdder(task string) {

	var err error
	var file *os.File

	if file, err = os.Open("task.json"); err != nil {
		log.Fatal(err)
	}

	var tasks []task_struct.Task

	if err = json.NewDecoder(file).Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	file.Close()

	new_task := Constructor(len(tasks), task)

	tasks = append(tasks, new_task)

	if file, err = os.OpenFile("task.json",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err = json.NewEncoder(file).Encode(tasks); err != nil {
		log.Fatal(err)
	}

}

func TaskEraser(id string) {

	var err error
	var file *os.File

	if file, err = os.Open("task.json"); err != nil {
		log.Fatal(err)
	}

	var task_list []task_struct.Task

	if err = json.NewDecoder(file).Decode(&task_list); err != nil {
		log.Fatal(err)
	}

	file.Close()

	var flag bool

	for i := 0; i < len(task_list); i++ {

		if task_list[i].GetId() == id {
			flag = true
			if i+1 == len(task_list) {
				continue
			} else {
				temp := task_list[i+1]
				task_list[i+1] = task_list[i]
				task_list[i] = temp
			}
		}
	}

	if flag == false {
		log.Fatal("Incorrect id")
	}

	if file, err = os.OpenFile("task.json",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777); err != nil {
		log.Fatal(err)

	}

	if err = json.NewEncoder(file).Encode(task_list[:len(task_list)-1]); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

}

func ShowTasksDone() {
	var err error
	var file *os.File

	if file, err = os.Open("task.json"); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var task_list []task_struct.Task

	if err = json.NewDecoder(file).Decode(&task_list); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done tasks list: ")

	for _, value := range task_list {
		if value.GetStatus() == "done" {
			fmt.Println("ID:", value.Id, " Task:", value.GetDescription())
		}
	}
}

func ShowTasksInProgress() {

	var err error
	var file *os.File

	if file, err = os.Open("task.json"); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var task_list []task_struct.Task

	if err = json.NewDecoder(file).Decode(&task_list); err != nil {
		log.Fatal(err)
	}

	fmt.Println("In-progress tasks list: ")

	for _, value := range task_list {
		if value.GetStatus() == "in-progress" {
			fmt.Println("ID:", value.Id, " Task:", value.GetDescription())
		}
	}
}

func ShowTasksTodo() {

	var err error
	var file *os.File

	if file, err = os.Open("task.json"); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var task_list []task_struct.Task

	if err = json.NewDecoder(file).Decode(&task_list); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done tasks list: ")

	for _, value := range task_list {
		if value.GetStatus() == "todo" {
			fmt.Println("ID:", value.Id, " Task:", value.GetDescription())
		}
	}

}

func ShowTasks() {

	var err error
	var file *os.File

	if file, err = os.Open("task.json"); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var task_list []task_struct.Task

	decoder := json.NewDecoder(file)

	if err = decoder.Decode(&task_list); err != nil {
		log.Fatal(err)
	}

	fmt.Println("All tasks list: ")

	for _, value := range task_list {
		fmt.Println("ID:", value.Id, " Task:", value.GetDescription())
	}

}

func SetStatus(id string, status int) {

	if status > 3 || status < 1 {
		fmt.Println("Incorrect task condition")
		return
	}

	var err error
	var file *os.File

	if file, err = os.Open("task.json"); err != nil {
		log.Fatal(err)
	}

	var tasks []task_struct.Task

	if err = json.NewDecoder(file).Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	file.Close()

	var flag bool

	for i := range tasks {
		if tasks[i].Id == id {
			flag = true
			tasks[i].UpdatedAt = time.Now().Format("02-01-2006 15.04.05")

			switch status {

			case 1:
				tasks[i].Status = "done"

			case 2:
				tasks[i].Status = "in-progress"

			case 3:
				tasks[i].Status = "todo"

			}
		}
	}

	if flag == false {
		fmt.Println("Incorrect id")
	}

	if file, err = os.OpenFile("task.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777); err != nil {
		log.Fatal(err)
	}

	if err = json.NewEncoder(file).Encode(tasks); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

}
