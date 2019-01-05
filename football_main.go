package main
import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "api/handlers"
)

func main(){
  router := mux.NewRouter()
  router.HandleFunc("/handlers/game_info", handlers.GameHandler).Methods("GET")
  router.HandleFunc("/handlers/game_run", handlers.GameRunHandler).Methods("GET")
  router.HandleFunc("/", ServeIndex).Methods("GET")
  // fmt.Println(router)
  fmt.Println("serving")
  log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
  fmt.Println("serving")

}

func ServeIndex(w http.ResponseWriter, r *http.Request){
  fmt.Println(w, "Index")
}
