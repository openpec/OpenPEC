package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //driver para o sql
	"github.com/gorilla/mux"
)

//Server é a estrutura principal do server
type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

//Connect conecta ao banco de dados
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@/openpec")

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
