package usuarios

import (
    "encoding/json"
    "fmt"
    db "github.com/florgm/webplanner_api/src/api/db"
    "github.com/florgm/webplanner_api/src/api/domain/usuarios"
)

//ParseLoginUsuario esto es una funcion
func ParseLoginUsuario(data []byte) (*usuarios.LoginUsuario, error) {
    var usuario usuarios.LoginUsuario
    if err := json.Unmarshal(data, &usuario); err != nil {
        return nil, err
    }

    return &usuario, nil
}

//ParseUsuario esto es una funcion
func ParseUsuario(data []byte) (*usuarios.Usuarios, error) {
    var usuario usuarios.Usuarios
    if err := json.Unmarshal(data, &usuario); err != nil {
        return nil, err
    }

    return &usuario, nil
}

//CreateUsuario funcion
func CreateUsuario(usuario *usuarios.Usuarios) error {
	stmt, err := db.Init().Prepare("insert into usuarios (nombre, usuario, password) values(?,?,?);")

    if err != nil {
		fmt.Print(err.Error())
		return err
    }

    _, err = stmt.Exec(usuario.Nombre, usuario.Usuario, usuario.Password)

    defer stmt.Close()
    return err
}

//UpdateUsuario funcion
func UpdateUsuario(user int64, usuario *usuarios.Usuarios) error {
	stmt, err := db.Init().Prepare("update usuarios set nombre=?, password=? where id_usuario=?;")

    if err != nil {
        fmt.Print(err.Error())
    }

    _, err = stmt.Exec(usuario.Nombre, usuario.Password, user)

    defer stmt.Close()
    return err
}

//Login funcion
func Login(usuario *usuarios.LoginUsuario) (*usuarios.Usuarios, error) {
    var user usuarios.Usuarios
    stmt, err := db.Init().Prepare("select * from usuarios where usuario = ? and password = ?;")

    if err != nil {
		fmt.Print(err.Error())
        return nil, err
    }

	result := stmt.QueryRow(usuario.Usuario, usuario.Password)
	err = result.Scan(
		&user.IDUsuario,
        &user.Nombre,
        &user.Usuario,
		&user.Password)
		
	if err != nil {
		return nil, err
	}
	
	defer stmt.Close()
    return &user, nil
}

//Logout funcion
func Logout(user int64) error {
    stmt, err := db.Init().Prepare("delete from sessions where user = ?;")

    if err != nil {
        fmt.Print(err.Error())
    }
    _, err = stmt.Exec(user)

    defer stmt.Close()
    return err
}
