package main

import (
	goard "github.com/dnl-jst/goard-daemon/v1"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	goard.Run()
}
