package main
 
import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "encoding/json"

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
    fmt.Println("Endpoint Hit: homePage")
}
 
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/{cat}", returnJson)
    myRouter.HandleFunc("/{cat}/{name}", returnJson)
    log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func handleNotFound(w http.ResponseWriter){
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "Error 404 - Page not found") 
}

func returnJson(w http.ResponseWriter, r *http.Request) {
    cat, _ := mux.Vars(r)["cat"]
    if ! itemInList(cat, []string{"types", "poketes", "attacks"}) {
	handleNotFound(w)
	return
    }
    fmt.Println("Endpoint Hit: returnJson", cat)
    data, _ := exec.Command("./get_json.py", cat).Output()
    key, exists := mux.Vars(r)["name"] 
    if ! exists {
        fmt.Fprintf(w, "%s", data)
    } else {
	var rawData map[string]interface{}
	json.Unmarshal([]byte(data), &rawData)
	item, item_exists := rawData[key]
	if ! item_exists {
	    handleNotFound(w)
	    return
	}
	res, _ := json.Marshal(item)
	fmt.Fprintf(w, "%s", res)
    }
}
 
func main() {
    handleRequests()
}
