package repositorios

import (
	"database/sql"
	"errors"
	"fmt"

	"api/src/modelos"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db: db}
}

func (u usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statment, err := u.db.Prepare("update usuarios set senha = ? where id = ?")
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err = statment.Exec(senha, usuarioID); err != nil {
		return err
	}

	return nil
}

func (u usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, err := u.db.Query("select senha from usuarios where id = ?", usuarioID)
	if err != nil {
		return "", err
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.Senha); err != nil {
			return "", err
		}
		return usuario.Senha, nil
	}
	return "", errors.New("usuario não encontrado")
}

func (u usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, err := u.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(userID), nil
}

func (u usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	linhas, err := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if err := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err

		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (u usuarios) BuscarPorID(id uint64) (modelos.Usuario, error) {
	linhas, err := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id=?",
		id,
	)

	if err != nil {
		return modelos.Usuario{}, err
	}

	if linhas.Next() {
		var usuario modelos.Usuario
		if err := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return modelos.Usuario{}, err
		}
		return usuario, nil
	}
	return modelos.Usuario{}, nil
}

func (u usuarios) AtualizarUsuario(id uint64, usuario modelos.Usuario) error {
	statament, err := u.db.Prepare("update usuarios set nome=?, nick=?, email=? where id=?")
	if err != nil {
		return err
	}
	defer statament.Close()

	if _, err := statament.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); err != nil {
		return err
	}

	return nil
}

func (u usuarios) DeletarUsuario(id uint64) error {
	statement, err := u.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (u usuarios) BuscarPorEmail(email string) (result modelos.Usuario, err error) {
	linha, err := u.db.Query("select id, senha from usuarios where email = ?", email)

	if err != nil {
		return result, err
	}

	defer linha.Close()

	if linha.Next() {
		if err := linha.Scan(&result.ID, &result.Senha); err != nil {
			return result, err
		}
	}

	return result, nil
}

func (u usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, err := u.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

func (u usuarios) DeixarDeSeguir(usuarioID, seguidorID uint64) error {
	statement, err := u.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

func (u usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, err := u.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criadoEm
	from usuarios u inner join seguidores s
	on u.id = s.seguidor_id
	where s.usuario_id = ?
	`, usuarioID)

	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (u usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, err := u.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criadoEm
	from usuarios u inner join seguidores s
	on u.id = s.usuario_id
	where s.seguidor_id = ?
	`, usuarioID)

	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
