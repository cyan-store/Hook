package config

import (
	"flag"
	"os"

	"github.com/cyan-store/hook/log"
)

func getFlags() int {
	var port int

	flag.IntVar(&port, "port", port, "Webhook port")
	flag.Parse()

	if port == 0 {
		log.Error.Println("--port is required!")
		os.Exit(1)
	}

	return port
}
