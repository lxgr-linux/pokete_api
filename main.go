package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/lxgr-linux/pokete_api/server"
)

func download() error {
	for _, i := range []string{"poketes", "attacks", "types"} {
		log.Println("Downloading", i)
		res, err := http.Get("https://raw.githubusercontent.com/lxgr-linux/pokete/master/pokete_data/" + i + ".py")
		defer res.Body.Close()
		if err != nil {
			return err
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		file, err := os.Create("./pokete_" + i + ".py")
		defer file.Close()
		if err != nil {
			return err
		}
		_, err = io.WriteString(file, string(body))
		if err != nil {
			return err
		}
	}
	return nil
}

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
		err := download()
		if err != nil {
			log.Fatal(err)
		}
	}
	if startServer {
		server.NewServer(port).HandleRequests()
	}
}
