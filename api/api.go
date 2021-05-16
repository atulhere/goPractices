package api

import (
	"fmt"
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

func TestApis() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")

	http.ListenAndServe(":3030", router)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	log.Print("Server started on port 3030")
	fmt.Println("hello World")
}
