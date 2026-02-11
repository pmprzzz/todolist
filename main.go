package main

import (
	"CLI-TodoList/cmd"
	//"CLI-TodoList/administration"
	"log"
)

func main() {

	//Choose how you are going to use this application

	//1.

	//administration.Administration()

	//2.

	rootCmd := cmd.Init()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
