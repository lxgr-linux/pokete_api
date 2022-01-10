package main
 
import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "encoding/json"

    "github.com/gorilla/mux"
)
 
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Pokete-API!")
    fmt.Println("Endpoint Hit: homePage")
}
 
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/poketes", returnAllPokes)
    myRouter.HandleFunc("/attacks", returnAllAttacks)
    myRouter.HandleFunc("/poketes/{name}", returnSinglePoke)
    myRouter.HandleFunc("/attacks/{name}", returnSingleAttack)
    log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func returnSinglePoke(w http.ResponseWriter, r *http.Request){
    returnAllJson(w, "pokes", mux.Vars(r)["name"])
}

func returnSingleAttack(w http.ResponseWriter, r *http.Request){
    returnAllJson(w, "attacks", mux.Vars(r)["name"])
}

func returnAllPokes(w http.ResponseWriter, r *http.Request){
    returnAllJson(w, "pokes", "")
}

func returnAllAttacks(w http.ResponseWriter, r *http.Request){
    returnAllJson(w, "attacks", "")
}

func returnAllJson(w http.ResponseWriter, name string, key string) {
    fmt.Println("Endpoint Hit: returnAllJson", name, key)
    data, _ := exec.Command("./get_json.py", name).Output()
    if key == "" {
        fmt.Fprintf(w, "%s", data)
    } else {
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	res, _ := json.Marshal(result[key])
	fmt.Fprintf(w, "%s", res)
    }
}
 
func main() {
    handleRequests()
}
