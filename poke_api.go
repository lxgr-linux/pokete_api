package main
 
import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
)
 
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Pokete-API!")
    fmt.Println("Endpoint Hit: homePage")
}
 
func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/poketes", returnAllPokes)
    http.HandleFunc("/attacks", returnAllAttacks)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func returnAllPokes(w http.ResponseWriter,r *http.Request){
    returnAllJson(w, "pokes")
}

func returnAllAttacks(w http.ResponseWriter,r *http.Request){
    returnAllJson(w, "attacks")
}

func returnAllJson(w http.ResponseWriter, name string) {
    fmt.Println("Endpoint Hit: returnAllJson", name)
    pokes, _ := exec.Command("./get_json.py", name).Output()
    fmt.Fprintf(w, "%s", pokes)
}
 
func main() {
    handleRequests()
}
