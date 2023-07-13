package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cyan-store/hook/log"
)

func Serve(port int) error {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        routes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Info.Printf("Running on port :%d", port)
	return server.ListenAndServe()
}
