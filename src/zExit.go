package src

import (
	"fmt"
	"os"
	"syscall"
)

func zExit() {
	os.Exit(0)
}

// listen sig channel
func GracefullyExit(sigC chan os.Signal) {
	sig := <-sigC
	if sig == syscall.SIGINT || sig == syscall.SIGSTOP || sig == syscall.SIGTERM || sig == syscall.SIGHUP {
		fmt.Println("gracefully exiting...")
		// do data persistence
		//.....
		zExit()
	}

}
