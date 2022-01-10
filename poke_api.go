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
    myRouter := mx.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/poketes", returnAllPokes)
    myRouter.HandleFunc("/attacks", returnAllAttacks)
    myRouter.HandleFunc("/poketes/{name}", returnSinglePoke)
    myRouter.HandleFunc("/attacks/{name}", returnSingleAttack)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func returnSinglePoke(w http.ResponseWriter, r *http.Request){
    returnAllJson(w, r, "pokes", mux.Vars(r)["name"])
}

func returnSingleAttack(w http.ResponseWriter, r *http.Request){
}

func returnAllPokes(w http.ResponseWriter, r *http.Request){
    returnAllJson(w, r, "pokes", "")
}

func returnAllAttacks(w http.ResponseWriter, r *http.Request){
    returnAllJson(w, r, "attacks", "")
}

func returnAllJson(w http.ResponseWriter, name string, key string) {
    fmt.Println("Endpoint Hit: returnAllJson", name)
    pokes, _ := exec.Command("./get_json.py", name).Output()
    if key == "" {
        fmt.Fprintf(w, "%s", pokes)
    } else {
	var result map[string]interface{}
	json.Unmarshal([]byte(pokes), &result)
	res, err := json.Marshal(result[key])
	fmt.Fprintf(w, "%s", res)
    }
}
 
func main() {
    handleRequests()
}
