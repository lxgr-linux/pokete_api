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
    myRouter.HandleFunc("/poketes", PoketeData{"pokes"}.returnJson)
    myRouter.HandleFunc("/attacks", PoketeData{"attacks"}.returnJson)
    myRouter.HandleFunc("/types", PoketeData{"types"}.returnJson)
    myRouter.HandleFunc("/poketes/{name}", PoketeData{"pokes"}.returnJson)
    myRouter.HandleFunc("/attacks/{name}", PoketeData{"attacks"}.returnJson)
    myRouter.HandleFunc("/types/{name}", PoketeData{"types"}.returnJson)
    log.Fatal(http.ListenAndServe(":8000", myRouter))
}

type PoketeData struct {
    name string
}

func (self PoketeData)returnJson(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnJson", self.name)
    data, _ := exec.Command("./get_json.py", self.name).Output()
    key, exists := mux.Vars(r)["name"] 
    if ! exists {
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
