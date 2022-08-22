package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	publi := modelos.Publicacoes{}

	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = json.Unmarshal(corpoRequest, &publi); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	publi.AutorID = usuarioID

	if err = publi.Preparar(); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioPublicacoes(db)

	publi.ID, err = repositorio.Criar(publi)

	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, publi)

}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorios := repositorios.NovoRepositorioPublicacoes(db)

	publicacao, err := repositorios.BuscarPorID(publicacaoId)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacao)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request)   {}
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {}
