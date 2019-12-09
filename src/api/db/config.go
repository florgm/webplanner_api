package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Init funcion que inicializa la conexion con la base de datos
func Init() *sql.DB {
	db, err := sql.Open("mysql", "calendario:asd123**@tcp(157.245.129.202:3306)/calendario?parseTime=true")
	checkErr(err)

	// defer db.Close()

	err = db.Ping()
	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
