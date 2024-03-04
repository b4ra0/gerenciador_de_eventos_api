package models

import (
	"time"
)

type Evento struct {
	Id        int
	Nome      string    `binding:"required"`
	Descricao string    `binding:"required"`
	Local     string    `binding:"required"`
	Datetime  time.Time `binding:"required"`
	UserID    int
}

var eventos = []Evento{}

func (e Evento) Save() {
	//salvar no banco de dados
	eventos = append(eventos, e)
}

func GetAllEventos() []Evento {
	return eventos
}
