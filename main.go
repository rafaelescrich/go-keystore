package main

import (
	"os"
	"runtime"

	"github.com/rafaelescrich/go-keystore/cmd"
	"github.com/rafaelescrich/go-keystore/database"
)

func main() {

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	cmd.Execute()

	database.InitDB()
}
