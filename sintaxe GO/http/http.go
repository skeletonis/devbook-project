package main

import (
	"log"
	"net/http"
)

//HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

//FUNCIONA COMO CLIENTE(FAZ A REQUISIÇÃO) - SERVIDOR(PROCESSA REQUISIÇÃO E ENVIA A RESPOSTA)

//REQUEST - RESPONSE

// MÉTODOS - GET, POST, PUT, DELETE

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}
