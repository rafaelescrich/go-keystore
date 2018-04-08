package main

import (
	"github.com/rafaelescrich/go-keystore/cmd"
	"github.com/rafaelescrich/go-keystore/database"
)

func main() {
	cmd.Execute()

	database.InitDB()
}
