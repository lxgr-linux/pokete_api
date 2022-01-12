package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lxgr-linux/pokete_api/server"
)

func showHelp() {
	fmt.Printf(`Downloads data and starts server
Usage: %s [ARG1] [ARG2] ...
Options:
	--no-download    : Does not download the data
	--update         : Just downloads the data
	--port <port>    : The port the server will run on, default is 8000
	--help           : Shows this dialog`, os.Args[0])
}

func main() {
	skip := false
	port := "8000"
	doDownload := true
	startServer := true

	for idx, i := range os.Args[1:] {
		if skip {
			skip = false
		} else {
			switch i {
			case "--help":
				showHelp()
				os.Exit(0)
			case "--no-download":
				doDownload = false
			case "--update":
				startServer = false
			case "--port":
				port = os.Args[idx+2]
				skip = true
			default:
				showHelp()
				os.Exit(1)
			}
		}
	}
	if doDownload {
		cmd := exec.Command("./download.sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	if startServer {
		fmt.Println(":: Starting...")
		server.NewServer(port).HandleRequests()
	}
}
