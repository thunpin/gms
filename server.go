package gms

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thunpin/gerrors"
)

func ToJSON(context *gin.Context, value interface{}, err error) {
	result := BuildAPIResult(value, gerrors.New(err))
	context.JSON(result.Code, result)
}

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
