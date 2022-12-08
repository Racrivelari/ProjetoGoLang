package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //mysql connection driver
)

// Abertura de conexao com o bd
func Conectar() (*sql.DB, error) {
	// stringConexao := "root:1q2w3e4r5t@/deposito?charset=utf8&parseTime=True&loc=Local"

	stringConexao := "admin:cursogolang@tcp(deposito-db.cmfdgmbsgaqu.sa-east-1.rds.amazonaws.com:3306)/deposito"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil

}
