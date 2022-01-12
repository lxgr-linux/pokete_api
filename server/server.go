package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func itemInList(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Pokete-API!")
	log.Println("Endpoint Hit: homePage")
}

// Server that configures the webservers behaviour
type Server struct {
	port string
}

// NewServer contructs a server object that configures the webservers behaviour
func NewServer(port string) Server {
	return Server{port}
}

// HandleRequests handles requests
func (s Server) HandleRequests() {
	log.Println("Starting server on port", s.port)
	fmt.Println(":: Starting server...")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/{cat}", returnJSON)
	myRouter.HandleFunc("/{cat}/{name}", returnJSON)
	log.Fatal(http.ListenAndServe(":"+s.port, myRouter))
}

func handleNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Error 404 - Page not found")
}

func returnJSON(w http.ResponseWriter, r *http.Request) {
	cat, _ := mux.Vars(r)["cat"]
	if !itemInList(cat, []string{"types", "poketes", "attacks"}) {
		handleNotFound(w)
		return
	}
	log.Println("Endpoint Hit: returnJson", cat)
	data, err := exec.Command("./get_json.py", cat).Output()
	if err != nil {
		log.Fatal("Error in get_json: ", err)
	}
	key, exists := mux.Vars(r)["name"]
	if !exists {
		fmt.Fprintf(w, "%s", data)
	} else {
		var rawData map[string]interface{}
		json.Unmarshal([]byte(data), &rawData)
		item, itemExists := rawData[key]
		if !itemExists {
			handleNotFound(w)
			return
		}
		res, _ := json.Marshal(item)
		fmt.Fprintf(w, "%s", res)
	}
}
