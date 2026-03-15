package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Buscando todos os usuario!"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Buscando um usuario!"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Atualizando um usuario!"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Deletando um usuario!"))
}
