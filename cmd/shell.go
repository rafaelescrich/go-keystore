package cmd

import (
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/rafaelescrich/go-keystore/controller"
	"github.com/rafaelescrich/go-keystore/keystore"
)

var shell *ishell.Shell

func logoASCII() string {
	return fmt.Sprintf(`
    
        ___         ___         ___         ___                 ___                 ___         ___         ___     
        /  /\       /  /\       /__/|       /  /\        ___    /  /\        ___    /  /\       /  /\       /  /\    
       /  /:/_     /  /::\     |  |:|      /  /:/_      /__/|  /  /:/_      /  /\  /  /::\     /  /::\     /  /:/_   
      /  /:/ /\   /  /:/\:\    |  |:|     /  /:/ /\    |  |:| /  /:/ /\    /  /:/ /  /:/\:\   /  /:/\:\   /  /:/ /\  
     /  /:/_/::\ /  /:/  \:\ __|  |:|    /  /:/ /:/_   |  |:|/  /:/ /::\  /  /:/ /  /:/  \:\ /  /:/~/:/  /  /:/ /:/_ 
    /__/:/__\/\:/__/:/ \__\:/__/\_|:|___/__/:/ /:/ /\__|__|:/__/:/ /:/\:\/  /::\/__/:/ \__\:/__/:/ /:/__/__/:/ /:/ /\
    \  \:\ /~~/:\  \:\ /  /:\  \:\/:::::\  \:\/:/ /:/__/::::\  \:\/:/~/:/__/:/\:\  \:\ /  /:\  \:\/:::::\  \:\/:/ /:/
     \  \:\  /:/ \  \:\  /:/ \  \::/~~~~ \  \::/ /:/   ~\~~\:\  \::/ /:/\__\/  \:\  \:\  /:/ \  \::/~~~~ \  \::/ /:/ 
      \  \:\/:/   \  \:\/:/   \  \:\      \  \:\/:/      \  \:\__\/ /:/      \  \:\  \:\/:/   \  \:\      \  \:\/:/  
       \  \::/     \  \::/     \  \:\      \  \::/        \__\/ /__/:/        \__\/\  \::/     \  \:\      \  \::/   
        \__\/       \__\/       \__\/       \__\/               \__\/               \__\/       \__\/       \__\/ 
    `)
}

func print(shell *ishell.Shell) {
	shell.Println(logoASCII())
	shell.Println("version: 0.0.1")
	shell.Println("Author: github.com/rafaelescrich")
	shell.Println("If you need help please type help any time")
}

func addCmd(shell *ishell.Shell) {

	shell.AddCmd(&ishell.Cmd{
		Name: "encryptfile",
		Help: "creates new encrypted file",
		Func: encryptFile(),
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "decryptfile",
		Help: "decrypts file",
		Func: decryptFile(),
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "deletekey",
		Help: "deletes key from db",
		Func: deleteKey(),
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "listkeys",
		Help: "list keys from db",
		Func: listKeys(),
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "insertpwd",
		Help: "insert password",
		Func: insertPwd(),
	})

}

func insertPwd() func(*ishell.Context) {
	return func(c *ishell.Context) {
		c.ShowPrompt(false)
		defer c.ShowPrompt(true)

		c.Println("Insert password to generate master key:")
		// prompt for input
		c.Print("Password: ")
		password := c.ReadPassword()

		err := controller.CreateMK(password)
		if err != nil {
			c.Println(err)
		}

		c.Println("Your password is: " + password)
	}
}

func deleteKey() func(*ishell.Context) {
	return func(c *ishell.Context) {
		c.ShowPrompt(false)
		defer c.ShowPrompt(true)

		if !keystore.MasterkeyExists() {
			c.Println("There is no master key!")
		} else {
			c.Println("Insert key to be deleted:")

			// TODO: list keys

			// prompt for input
			c.Print("Key number: ")
			key := c.ReadLine()

			// Are you sure:
			c.Print("Are you sure? (y/n)")
			yn := c.ReadLine()

			if yn != "y" {
				c.Println("Delete key was cancelled")
			} else {
				// TODO: delete key
				c.Println("Your input is: " + key + ".")
			}
		}

	}
}

func encryptFile() func(*ishell.Context) {
	return func(c *ishell.Context) {
		c.ShowPrompt(false)
		defer c.ShowPrompt(true)

		if !keystore.MasterkeyExists() {
			c.Println("There is no master key!")
		} else {
			c.Println("Insert filename to encrypt:")

			// prompt for input
			c.Print("Filename: ")
			filename := c.ReadLine()

			err := controller.EncryptFile(filename)
			if err != nil {
				c.Println(err)
				c.Print("Error while loading file from disk")
			} else {
				c.Print("File ciphered with success!")
			}
		}
	}

}
func decryptFile() func(*ishell.Context) {
	return func(c *ishell.Context) {
		c.ShowPrompt(false)
		defer c.ShowPrompt(true)

		if !keystore.MasterkeyExists() {
			c.Println("There is no master key!")
		} else {
			c.Println("Insert filename to decrypt:")

			// prompt for input
			c.Print("Filename: ")
			filename := c.ReadLine()

			err := controller.DecryptFile(filename)
			if err != nil {
				c.Println(err)
				c.Print("Error while loading file from disk")
			} else {
				c.Print("File deciphered with success!")
			}
		}
	}

}

func listKeys() func(*ishell.Context) {
	return func(c *ishell.Context) {

		c.Println("List of keys in DB:")

		// TODO: get all keys from db

		// TODO: print keys

	}
}

// Run the shell
func Run() {
	shl := ishell.New()
	addCmd(shl)
	print(shl)
	shl.Run()
}
