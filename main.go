package main

import (
	"flag"
	"fmt"
	"muzucode/goroutines/client"
	"muzucode/goroutines/server"
	"os"
)

func main() {

	modePtr := flag.String("mode", "", "Specify 'client' or 'server' mode")
	addressPtr := flag.String("address", "localhost:9800", "Specify the server address in 'client' mode")

	flag.Parse()

	switch *modePtr {
	case "client":
		client.RunClient(*addressPtr)
	case "server":
		server.RunServer(*addressPtr)
	default:
		fmt.Println("Invalid mode specified. Use '-mode client' or '-mode server'")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
