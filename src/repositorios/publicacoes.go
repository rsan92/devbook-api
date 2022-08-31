package repositorios

import (
	"database/sql"

	"api/src/erros"
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

	return modelos.Publicacoes{}, erros.ErrPublicacaoNotFound{}
}

func (p Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacoes, error) {
	linhas, err := p.db.Query(`
		select distinct p.*, u.nick
		from publicacoes p 
			inner join usuarios u on p.autor_id = u.id 
			inner join seguidores s on p.autor_id = s.usuario_id
		where u.id = ?
		or s.seguidor_id = ?
		order by 1 desc
	`, usuarioID, usuarioID)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	resultado := []modelos.Publicacoes{}

	for linhas.Next() {
		publicacao := modelos.Publicacoes{}
		if err := linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); err != nil {
			return nil, err
		}
		resultado = append(resultado, publicacao)
	}

	return resultado, nil
}

func (p Publicacoes) AtualizarPublicacao(publicacaoID uint64, publicacao modelos.Publicacoes) error {
	statement, err := p.db.Prepare(`
		update publicacoes set titulo = ?, conteudo = ? where id = ?
	`)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); err != nil {
		return err
	}

	return nil
}

func (p Publicacoes) DeletarPublicacao(publicacaoID uint64) error {
	statement, err := p.db.Prepare(`
		delete from publicacoes where id = ?
	`)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}

func (p Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacoes, error) {
	linhas, err := p.db.Query(`
		select p.*, u.nick
		from publicacoes p join usuarios u on u.id = p.autor_id
		where autor_id = ?
	`, usuarioID)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	resultado := []modelos.Publicacoes{}

	for linhas.Next() {
		publicacao := modelos.Publicacoes{}
		if err := linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); err != nil {
			return nil, err
		}
		resultado = append(resultado, publicacao)
	}

	return resultado, nil
}

func (p Publicacoes) Curtir(publicacaoID uint64) error {
	statement, err := p.db.Prepare(`update publicacoes set curtidas = curtidas + 1 where id = ?`)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}

func (p Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, err := p.db.Prepare(`update publicacoes set curtidas = 
		CASE 
			WHEN 
				curtidas > 0 THEN curtidas - 1
			ELSE curtidas 
			END
		where id = ?`)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}
