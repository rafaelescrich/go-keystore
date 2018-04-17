package cfg

import (
	"os"
	"runtime"
)

// SetMaxProcs sets the maximum number of process to the number of cpu present
func SetMaxProcs() {
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}
