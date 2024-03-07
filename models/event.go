package models

import (
	"fmt"
	"gerenciador_de_eventos/db"
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

func (e Evento) Save() error {
	query := `INSERT INTO eventos(nome, descricao, local, dateTime, user_id) VALUES ($1,$2,$3,$4,$5)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	resultado, err := stmt.Exec(e.Nome, e.Descricao, e.Local, e.Datetime, e.UserID)
	if err != nil {
		return err
	}
	id, err := resultado.RowsAffected()
	if err != nil {
		return err
	}
	e.Id = int(id)
	return err
}

func GetAllEventos() {
	query := `SELECT * FROM eventos`
	stmt, err := db.DB.Prepare(query)
	resultado, err := stmt.Exec()
	if err != nil {
		panic(err)
	}
	fmt.Println(resultado)
}
