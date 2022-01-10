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
    log.Fatal(http.ListenAndServe(":8000", nil))
}
 
func returnAllPokes(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllPokes")
    pokes, _ := exec.Command("./get_json.py").Output()
    fmt.Fprintf(w, "%s", pokes)
}
 
func main() {
    handleRequests()
}
