package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	driver   = "postgres"
	host     = "192.168.0.189"
	port     = "5432"
	user     = "postgres"
	password = "senha"
	dbname   = "gerenciador_de_eventos"
)

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

var DB *sql.DB
var err error

func Conectar() {
	fmt.Println(DataSourceName)
	DB, err = sql.Open(driver, DataSourceName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	criarTabelas()
}

func criarTabelas() {
	criarTabelaUsuarios := `
	CREATE TABLE IF NOT EXISTS users (
	    id SERIAL,
	    email TEXT NOT NULL UNIQUE ,
	    password TEXT NOT NULL
	)
`
	_, err := DB.Exec(criarTabelaUsuarios)

	if err != nil {
		panic(err)
	}
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
	_, err = DB.Exec(criarTabelaEventos)

	if err != nil {
		fmt.Println(err)
		panic("Não foi possível criar tabela de eventos")
	} else {
		fmt.Println("Tabelas criadas com sucesso!")
	}
}
