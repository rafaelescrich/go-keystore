package main

import (
	"github.com/rafaelescrich/go-keystore/cfg"
	"github.com/rafaelescrich/go-keystore/cmd"
	"github.com/rafaelescrich/go-keystore/controller"
)

func main() {

	cfg.SetMaxProcs()
	controller.InitDB()
	cmd.Run()
}
