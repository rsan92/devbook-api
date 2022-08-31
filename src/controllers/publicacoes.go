package controllers

import (
	"encoding/json"
	"errors"
	"io"
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

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorios := repositorios.NovoRepositorioPublicacoes(db)

	resultado, err := repositorios.Buscar(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, resultado)

}

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

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

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

	repositorio := repositorios.NovoRepositorioPublicacoes(db)

	publicacaoSalvaNoBanco, err := repositorio.BuscarPorID(publicacaoId)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if usuarioID != publicacaoSalvaNoBanco.AutorID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar uma publicacao que não é a sua"))
		return
	}

	var publicacaoCorpo modelos.Publicacoes

	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = json.Unmarshal(corpoRequest, &publicacaoCorpo); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = repositorio.DeletarPublicacao(publicacaoId); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

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

	repositorio := repositorios.NovoRepositorioPublicacoes(db)
	publicacaoSalvaNoBanco, err := repositorio.BuscarPorID(publicacaoId)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if usuarioID != publicacaoSalvaNoBanco.AutorID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar uma publicacao que não é a sua"))
		return
	}

	var publicacaoCorpo modelos.Publicacoes

	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = json.Unmarshal(corpoRequest, &publicacaoCorpo); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = repositorio.AtualizarPublicacao(publicacaoId, publicacaoCorpo); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)

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

	repositorio := repositorios.NovoRepositorioPublicacoes(db)

	publicacoes, err := repositorio.BuscarPorUsuario(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)

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

	repositorio := repositorios.NovoRepositorioPublicacoes(db)

	err = repositorio.Curtir(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)

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

	repositorio := repositorios.NovoRepositorioPublicacoes(db)

	err = repositorio.Descurtir(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}
