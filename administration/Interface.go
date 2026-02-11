package administration

import (
	"CLI-TodoList/functions"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Administration() {
	if _, err := os.Stat("task.json"); os.IsNotExist(err) {
		os.WriteFile("task.json", []byte("[]"), 0777)
	}

	for {
		fmt.Println("\n" + "TaskList")

		fmt.Println("1. Show all tasks list\n" +
			"2. Show done tasks list\n" +
			"3. Show in-progress task list\n" +
			"4. Show todo tasks list\n" +
			"5. Add task\n" +
			"6. Update task\n" +
			"7. Delete task\n" +
			"8. Exit")

		fmt.Println("Choose option: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, _ := strconv.Atoi(scanner.Text())

		switch option {

		case 1:
			functions.ShowTasks()

		case 2:
			functions.ShowTasksDone()

		case 3:
			functions.ShowTasksInProgress()

		case 4:
			functions.ShowTasksTodo()

		case 5:
			fmt.Println("Enter task description: ")
			scanner.Scan()
			task := scanner.Text()

			functions.TaskAdder(task)

		case 6:
			var option_update int
			functions.ShowTasks()

			fmt.Println("Enter task Id: ")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			id := scanner.Text()

			fmt.Println("Choose option:" + " 1. Mark as done " + " 2. Mark as in-progress " + " 3. Mark as todo")

			scanner.Scan()
			option_update, _ = strconv.Atoi(scanner.Text())

			switch option_update {

			case 1:
				functions.SetStatus(id, option_update)
				fmt.Println("Success!")

			case 2:
				functions.SetStatus(id, option_update)
				fmt.Println("Success!")

			case 3:
				functions.SetStatus(id, option_update)
				fmt.Println("Success!")

			default:
				fmt.Println("Incorrect condition!")

			}
		case 7:
			fmt.Println("Enter task Id: ")
			scanner.Scan()
			id := scanner.Text()

			functions.TaskEraser(id)

		case 8:
			return

		default:
			fmt.Println("Incorrect input")

		}
	}
}
