package main

import (
	"os"
	"runtime"

	"github.com/abiosoft/ishell"
	"github.com/rafaelescrich/go-keystore/database"
)

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	database.InitDB()

	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("Go-keystore")
	shell.Println("version: 0.0.1")
	shell.Println("Author: github.com/rafaelescrich")
	shell.Println("If you need help please type help any time")

	shell.AddCmd(&ishell.Cmd{
		Name: "createkey",
		Help: "creates new key and saves on db",
		Func: func(c *ishell.Context) {
			// Prints with args
			// if c.Args < 1 {

			// } else {
			// 	if filesExists(c.Args) {
			// 		// TODO: Encrypt file and save key
			// 	}
			// }
		},
	})

	// run shell
	shell.Run()
}
