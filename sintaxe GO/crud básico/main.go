package main

import (
	"crud/servidor"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TIPOS DE REQUISIÇÃO DE CRUD
// CREATE	- 	POST
// READ 	-	GET
// UPDATE 	- 	PUT
// DELETE 	- 	DELETE

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuarios).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":5000", router))

}
