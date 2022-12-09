package main

import (
	"deposito/handler"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	// fs := http.FileServer(http.Dir("dist/spa"))
	fs := http.FileServer(http.Dir("dist/spa/"))
	

	fmt.Println("CHEGOU AQ")

	// create - post
	// read - get
	// update - put
	// delete - delete

	router := mux.NewRouter()

	router.HandleFunc("/products", handler.CreateProduct).Methods("POST")

	router.HandleFunc("/products", handler.GetAllProducts).Methods("GET")

	router.HandleFunc("/products/{id}", handler.GetById).Methods("GET")

	router.HandleFunc("/products/{id}", handler.DeleteById).Methods("DELETE")

	router.HandleFunc("/products/{id}", handler.UpdateById).Methods	("PUT")

	router.Handle("/webui", http.StripPrefix("/webui", fs))

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
