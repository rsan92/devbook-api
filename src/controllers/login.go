package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario modelos.Usuario

	if err := json.Unmarshal(corpoRequest, &usuario); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusServiceUnavailable, err)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarioSalvo, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		respostas.Erro(w, http.StatusServiceUnavailable, err)
		return
	}

	if err = seguranca.VerificarSenha(usuario.Senha, usuarioSalvo.Senha); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := autenticacao.CriarToken(usuarioSalvo.ID)
	if err != nil {
		respostas.Erro(w, http.StatusServiceUnavailable, err)
		return
	}

	respostas.JSON(w, http.StatusOK, token)
}
