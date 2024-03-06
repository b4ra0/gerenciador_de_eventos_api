package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	driver   = "postgres"
	host     = "192.168.15.122"
	port     = "5432"
	user     = "postgres"
	password = "senha"
	dbname   = "gerenciador_de_eventos"
)

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

var db *sql.DB
var err error

func Connectar() {
	fmt.Println(DataSourceName)
	db, err = sql.Open(driver, DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	criarTabelas()
}

func criarTabelas() {
	criarTabelaEventos := `
	CREATE TABLE IF NOT EXISTS eventos (
	    id SERIAL,
	    nome TEXT NOT NULL,
	    descricao TEXT NOT NULL,
	    local TEXT NOT NULL,
	    dateTime TIMESTAMP NOT NULL,
	    user_id INTEGER
	    )
`
	_, err := db.Exec(criarTabelaEventos)

	if err != nil {
		fmt.Println(err)
		panic("Não foi possível criar tabela de eventos")
	} else {
		fmt.Println("Tabelas criadas com sucesso!")
	}
}