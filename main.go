package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type message struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    uint64 `json:"data"`
}

type errorMessage struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    uint64 `json:"data"`
}

//handleRoot handles calls to the root endpoint
func handleRoot(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome home!")
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	z := rand.NewZipf(rnd, 1.2, 4, 60)
	n := z.Uint64()
	if n == 0 {
		n = 1
	} else if n > 30 {
		//error out due to timeouts
		time.Sleep(time.Duration(2) * time.Second)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errorMessage{500, "Internal Server Error", "Request timed out", n})
		return
	}
	time.Sleep(time.Duration(n) * time.Millisecond * 100)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(message{200, true, "Request Successful", n})
}

func main() {
	//set the port to run server at
	portPtr := flag.String("port", "8080", "port number to run at")
	flag.Parse()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/request", handleRoot).Methods("GET")
	log.Println("Listening on port " + *portPtr)
	log.Fatal(http.ListenAndServe(":"+*portPtr, router))

}
