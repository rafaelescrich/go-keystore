package main

import (
	"os"
	"runtime"

	"github.com/abiosoft/ishell"
	"github.com/rafaelescrich/go-keystore/database"
)

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
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			c.Println("Insert password to generate master key:")

			// prompt for input
			c.Print("Password: ")
			password := c.ReadPassword()

			// ciphering.
			// do something with username and password
			c.Println("Your input is: " + password + ".")

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "encryptfile",
		Help: "creates new encrypted file",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			c.Println("Insert filename to encrypt:")

			// prompt for input
			c.Print("Filename: ")
			filename := c.ReadLine()

			c.Println("Choose a key from your keystore:")

			// prints keys
			// if none return error

			// choose key
			// creates new encrypted file
			// deletes original file

			// ciphering.
			// do something with username and password
			c.Println("Your input is: " + filename + ".")

		},
	})

	// run shell
	shell.Run()
}
