package graceshut

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func signals() os.Signal {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	return <-sigs
}

func Loop() {
	for {
		if sig := signals(); sig == syscall.SIGINT || sig == syscall.SIGTERM {
			fmt.Println("...Shutting app...")
			os.Exit(0)
		}
	}
}
