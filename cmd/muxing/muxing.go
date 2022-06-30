package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
/*
Endpoints expected by tests:

| METHOD | REQUEST                               | RESPONSE                      |
|--------|---------------------------------------|-------------------------------|
| GET    | `/name/{PARAM}`                       | body: `Hello, PARAM!`         |
| GET    | `/bad`                                | Status: `500`                 |
| POST   | `/data` + Body `PARAM`                | body: `I got message:\nPARAM` |
| POST   | `/headers`+ Headers{"a":"2", "b":"3"} | Header `"a+b": "5"`           |

If not defined in table:
Request will be:
* No body set
* No headers set
Response expected to have
* Status: 200 OK
* Empty body
*/
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", handleName).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handleData).Methods(http.MethodPost)
	router.HandleFunc("/headers", handleHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func handleName(w http.ResponseWriter, r *http.Request) {
	//code
	param := mux.Vars(r)["PARAM"]
	data := []byte("Hello, " + param + "!")
	//w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleBad(w http.ResponseWriter, r *http.Request) {
	//code
	w.WriteHeader(http.StatusInternalServerError)
}

//error fixed
func handleData(w http.ResponseWriter, r *http.Request) {
	//code
	var param []byte
	//if count, err := r.Body.Read(param); err != nil {
	param, err := io.ReadAll(r.Body)
	if err != nil {
		//if _, err := r.Body.Read(param); err != nil {
		log.Fatal(err)
	}
	//if count == 0 {
	//	param = []byte("")
	//}
	data := []byte("I got message:\n") // + PARAM body
	data = append(data, param...)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//error
func handleHeaders(w http.ResponseWriter, r *http.Request) {
	//code
	var a1, b1, result int
	var err error
	a := r.Header.Get("a")
	b := r.Header.Get("b")
	if a1, err = strconv.Atoi(a); err != nil {
		log.Fatal(err)
	}
	if b1, err = strconv.Atoi(b); err != nil {
		log.Fatal(err)
	}
	result = a1 + b1
	w.Header().Add("a+b", string(rune(result)))
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(""))
}
