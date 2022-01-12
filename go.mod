module github.com/lxgr-linux/pokete_api

go 1.17

require github.com/lxgr-linux/pokete_api/server v0.1.0

require github.com/gorilla/mux v1.8.0 // indirect

replace github.com/lxgr-linux/pokete_api/server => ./server
