package main

import (
	"github.com/rafaelescrich/go-keystore/cfg"
	"github.com/rafaelescrich/go-keystore/cmd"
	"github.com/rafaelescrich/go-keystore/database"
)

func main() {

	cfg.SetMaxProcs()
	database.InitDB()
	cmd.Run()
}
