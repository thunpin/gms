package gms

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WaitStop(cError chan error, callback func(int)) {
	cSignal := make(chan os.Signal, 1)
	signal.Notify(cSignal, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-cError:
			if err != nil {
				callback(http.StatusInternalServerError)
			}
			os.Exit(http.StatusInternalServerError)
		case <-cSignal:
			log.Println("Exiting...")
			callback(http.StatusServiceUnavailable)
			os.Exit(0)
		}
		time.Sleep(1 * time.Millisecond)
	}
}
