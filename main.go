package main

import (
	"fmt"
	"github.com/djungyvr/notr/cmd"
	"github.com/djungyvr/notr/models"
	"github.com/djungyvr/notr/notr"
	"os"
)

func initProgram() {
	// create the application directory
	// owner can read and write, everyone else can only read
	path := os.Getenv("HOME") + "/" + ".notr"
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create/connect the sqlite database
	err := models.Init(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	initProgram()
	defer models.Repo().Close()

	// Initialize the commands
	cmd.AddCommands(notr.RootCmd)
	if err := notr.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
