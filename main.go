package main

import (
	"github.com/lxgr-linux/pokete_api/server"
)

func main() {
	server.NewServer("8000").HandleRequests()
}
