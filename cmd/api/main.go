package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phanikumarps/umc/umc"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/metadata", metadata).Methods("GET")
	router.HandleFunc("/{entity}/{id}", resolver).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}
func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}
func metadata(w http.ResponseWriter, r *http.Request) {
	resp := umc.Metadata()
	b, err := getBytes(resp)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(b))
}
func resolver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entity := vars["entity"]
	id := vars["id"]
	switch entity {
	case "account":
		resp := umc.Account(&id)
		j, err := json.Marshal(resp)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Hello, World!"))
	}
}
