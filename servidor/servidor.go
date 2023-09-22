package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID 		int 	`json:"id"`
	Nome 	string 	`json:"nome"`
	Email 	string	`json:"email"`
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body) // Lê o corpo da requisição
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil { // Converte o JSON para struct
		w.Write([]byte("Erro ao converter o usuário para struct!"))
		log.Fatal(erro)
		return
	}
	fmt.Println(usuario)

	db, erro := banco.Conectar() // Conecta ao banco de dados
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)") // Prepara a query
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		log.Fatal(erro)
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email) // Executa a query
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId() // Pega o ID do usuário inserido
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated) // Retorna o status 201 Created
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}

// BuscarUsuarios busca todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios") // Executa a query
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() { // Percorre todas as linhas retornadas
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil { // Pega os valores de cada coluna
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK) // Retorna o status 200 OK
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil { // Converte o slice de usuários para JSON
		w.Write([]byte("Erro ao converter os usuários para JSON!"))
		return
	}
}

// BuscarUsuario busca um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r) // Pega os parâmetros da requisição
	
	ID, erro := strconv.ParseInt(parametros["id"], 10, 32) // Converte o ID de string para int
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID) // Executa a query
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}

	var usuario usuario
	if linha.Next() { // Pega a primeira linha retornada
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil { // Pega os valores de cada coluna
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil { // Converte o usuário para JSON
		w.Write([]byte("Erro ao converter o usuário para JSON!"))
		return
	}
}

// AtualizarUsuario altera um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r) // Pega os parâmetros da requisição

	ID, erro := strconv.ParseInt(parametros["id"], 10, 32) // Converte o ID de string para int
	if erro != erro {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}
	
	corpoRequisicao, erro := io.ReadAll(r.Body) // Lê o corpo da requisição
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil { // Converte o JSON para struct
		w.Write([]byte("Erro ao converter o usuário para struct!"))
		return
	}

	db, erro := banco.Conectar() // Conecta ao banco de dados
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?") // Prepara a query
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil { // Executa a query
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent) // Retorna o status 204 No Content
}

// DeletarUsuario remove um usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r) // Pega os parâmetros da requisição

	ID, erro := strconv.ParseInt(parametros["id"], 10, 32) // Converte o ID de string para int
	if erro != erro {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}

	db, erro := banco.Conectar() // Conecta ao banco de dados
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?") // Prepara a query
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil { // Executa a query
		w.Write([]byte("Erro ao deletar o usuário!"))
		return
	}
	
	w.WriteHeader(http.StatusNoContent) // Retorna o status 204 No Content

}