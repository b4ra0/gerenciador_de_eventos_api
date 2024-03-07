package models

import (
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

func GetAllEventos() ([]Evento, error) {
	query := `SELECT * FROM eventos`
	row, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		var evento Evento
		err := row.Scan(&evento.Id, &evento.Nome, &evento.Descricao, &evento.Local, &evento.Datetime, &evento.UserID)

		if err != nil {
			return nil, err
		}
		eventos = append(eventos, evento)
	}
	return eventos, err
}

func GetEventoById(id int64) (*Evento, error) {
	query := `SELECT * FROM eventos WHERE ID = $1`
	row := db.DB.QueryRow(query, id)

	var evento Evento
	err := row.Scan(&evento.Id, &evento.Nome, &evento.Descricao, &evento.Local, &evento.Datetime, &evento.UserID)
	if err != nil {
		return nil, err
	}
	return &evento, nil
}
