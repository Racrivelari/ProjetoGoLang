package main

import (
	// "deposito/handler"
	"fmt"
	// "log"
	"net/http"

	// "github.com/gorilla/mux"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("dist/spa/"))
	http.Handle("/webui/", http.StripPrefix("/webui/", fs))

	http.ListenAndServe(":5050", nil)

	//create - post
	//read - get
	//update - put
	//delete - delete

	// router := mux.NewRouter()

	// router.HandleFunc("/products", handler.CreateProduct).Methods("POST")

	// router.HandleFunc("/products", handler.GetAllProducts).Methods("GET")

	// router.HandleFunc("/products/{id}", handler.GetById).Methods("GET")

	// router.HandleFunc("/product/{id}", handler.DeleteById).Methods("DELETE")

	// router.HandleFunc("/product/{id}", handler.UpdateById).Methods("PUT")

	// fmt.Println("Escutando na porta 5000")
	// log.Fatal(http.ListenAndServe(":5000", router))
}
