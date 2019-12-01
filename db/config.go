package db

import (
	_ "github.com/go-sql-driver/mysql"
  	"database/sql"
	"fmt"
)

//Init funcion que inicializa la conexion con la base de datos
func Init() (*sql.DB) {
  	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/calendario?parseTime=true")
  	checkErr(err)

	// defer db.Close()
	
	err = db.Ping()
	checkErr(err)
	fmt.Printf("Connection successfully")

	return db
}

func checkErr(err error) {
  if err != nil {
    fmt.Print(err.Error())
  }
}
