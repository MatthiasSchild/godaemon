package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MatthiasSchild/godaemon"
)

var daemonizer = godaemon.New(godaemon.Options{
	Name:    "service",
	Service: service,
})

func service() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pid := os.Getpid()
		_, _ = fmt.Fprintf(w, "Service is running with PID %d!\n", pid)
	})

	log.Println("Server started and is listening")
	return http.ListenAndServe(":8080", nil)
}

func startService() {
	fmt.Println("Start service")
	err := daemonizer.Start()
	if err != nil {
		fmt.Println("Failed to start service:", err)
	}
}

func stopService() {
	fmt.Println("Stop service")
	err := daemonizer.Stop()
	if err != nil {
		fmt.Println("Failed to stop service:", err)
	}
}

func main() {
	isService, err := daemonizer.IsService()
	if err != nil {
		fmt.Println("Failed to start service:", err)
		os.Exit(1)
	}
	if isService {
		return
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "start":
			startService()
			return
		case "stop":
			stopService()
			return
		}
	}

	fmt.Println("usage: go run ./example <start|stop>")
}
