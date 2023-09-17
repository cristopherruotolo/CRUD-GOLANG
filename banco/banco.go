package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	urlConexaoBanco := "mago:themago@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", urlConexaoBanco)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}