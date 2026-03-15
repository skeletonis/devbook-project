package servidor

import (
	"crud/bd"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/* Cliente
   │
   │ POST /usuarios
   │ JSON (nome,email)
   ▼
Servidor Go
   │
   ├─ lê body
   ├─ converte JSON → struct
   ├─ conecta no banco
   ├─ executa INSERT
   ├─ pega ID criado
   ▼
Resposta
201 Created
Usuário inserido com sucesso! ID: X */

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	//O servidor lê o corpo de requisição do HTTP
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	//Adapta o JSON para o struct por meio da função json.Unmarshal
	var user usuario
	if erro = json.Unmarshal(corpoRequisicao, &user); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
	}

	//Aqui eu faço a conexão com o db. Por meio disso eu posso realizar as outras funções do CRUD
	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	//statement prepara o SQL. Isso serve para contornar tentativas de SQL Injection.
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}

	defer statement.Close()
	//Aqui os ? são substituídos por:
	insercao, erro := statement.Exec(user.Nome, user.Email)
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	// STATUS CODES
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	//Como é uma consulta não virá nenhum corpo na requisição.
	//Aqui eu faço a conexão com o db. Por meio disso eu posso realizar as outras funções do CRUD
	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}

	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	// STATUS CODES
	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	//Como é uma consulta não virá nenhum corpo na requisição.
	//Aqui eu faço a conexão com o db. Por meio disso eu posso realizar as outras funções do CRUD

	paramentros := mux.Vars(r)

	ID, erro := strconv.ParseUint(paramentros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro em converter o parâmetro para inteiro"))
	}

	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados."))
		return
	}

	defer db.Close()

	//statement prepara o SQL. Isso serve para contornar tentativas de SQL Injection.
	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}

	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}

	// STATUS CODES
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON"))
		return
	}
}

func AtualizarUsuarios(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Falha ao converter o usuário para struct"))
		return
	}

	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dado!"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome=?, email=? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}

	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados."))
		return
	}

	defer db.Close()

	//statement prepara o SQL. Isso serve para contornar tentativas de SQL Injection.
	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuários"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
