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
