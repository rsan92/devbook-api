package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacoes struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadoEm,omitempty"`
}

func (p *Publicacoes) Preparar() error {
	if err := p.validar(); err != nil {
		return err
	}
	p.formatar()
	return nil
}

func (p *Publicacoes) validar() error {
	if p.Titulo == "" {
		return errors.New("o titulo é obrigário e não pode estar em branco")
	}

	if p.Conteudo == "" {
		return errors.New("o conteudo é obrigário e não pode estar em branco")
	}
	return nil
}

func (p *Publicacoes) formatar() {
	p.Titulo = strings.TrimSpace(p.Titulo)
	p.Conteudo = strings.TrimSpace(p.Conteudo)
}
