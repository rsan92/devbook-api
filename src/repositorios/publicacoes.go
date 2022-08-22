package repositorios

import (
	"database/sql"
	"errors"

	"api/src/modelos"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db: db}
}

func (p Publicacoes) Criar(publicacao modelos.Publicacoes) (uint64, error) {
	statement, err := p.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(id), nil

}

func (p Publicacoes) BuscarPorID(id uint64) (modelos.Publicacoes, error) {
	linha, err := p.db.Query(`
		select p.id, p.titulo, p.conteudo, p.autor_id, p.curtidas, p.criadaEm, u.nick
		from publicacoes p inner join usuarios u
		on p.autor_id = u.id
		where p.id = ?
	`, id)
	if err != nil {
		return modelos.Publicacoes{}, err
	}

	defer linha.Close()

	if linha.Next() {
		publicacao := modelos.Publicacoes{}
		if err := linha.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); err != nil {
			return modelos.Publicacoes{}, err
		}
		return publicacao, nil
	}

	return modelos.Publicacoes{}, errors.New("publicacao not found")
}
