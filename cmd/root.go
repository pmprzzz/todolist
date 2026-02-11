package cmd

import (
	"CLI-TodoList/functions"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Cli application to track daily tasks",
	Long:  `My first pet project - cli task tracker for daily usage, its can add, delete, update and track your daily tasks`,
}

func RootCmd() *cobra.Command {
	return rootCmd
}

var add = &cobra.Command{
	Use:   "add",
	Short: "Add task by description",
	Long:  `Use this function to add your new task to file, enter the description of task, for example: "Buy groceries". `,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		functions.TaskAdder(args[0])
	},
}

var delete = &cobra.Command{
	Use:   "delete",
	Short: "Delete task by id",
	Long:  `Use this function to delete your new task from file, enter the id of task, for example: "5".`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cleanAll, _ := cmd.Flags().GetBool("all")
		if cleanAll {
			functions.Clean()
		} else if len(args) == 0 {
			fmt.Println("Error: need id or -a flag")
		} else {
			functions.TaskEraser(args[0])
		}
	},
}

var show = &cobra.Command{
	Use:   "list",
	Short: "Use this function to see the list of all your tasks",
	Long:  `Use this functions Use this function to see the list of all your tasks, not depends from tasks status`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			functions.ShowTasks()
		} else {

			status := args[0]

			switch status {
			case "done":
				functions.ShowTasksDone()
			case "todo":
				functions.ShowTasksTodo()
			case "in-progress":
				functions.ShowTasksInProgress()
			default:
				fmt.Println("Incorrect argument")
			}
		}

	},
}

var mark_in_progress = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Use this function to mark your task as in-progress by id",
	Long:  `Use this function ti mark your task as in-progress by id, for example: mark-in-progress 5`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		functions.SetStatus(args[0], 2)
	},
}

var mark_done = &cobra.Command{
	Use:   "mark-done",
	Short: "Use this function to mark your task as done by id",
	Long:  `Use this function ti mark your task as done by id, for example: mark-done 5`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		functions.SetStatus(args[0], 1)
	},
}

var mark_todo = &cobra.Command{
	Use:   "mark-todo",
	Short: "Use this function to mark your task as in-progress by id",
	Long:  `Use this function ti mark your task as todo by id, for example: mark-todo 5`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		functions.SetStatus(args[0], 3)
	},
}

func Init() *cobra.Command {

	functions.CreateIfNotEx()

	rootCmd.AddCommand(add, delete, show, mark_done, mark_in_progress, mark_todo)

	delete.Flags().Bool("all", false, "Delete all tasks")

	return rootCmd
}
