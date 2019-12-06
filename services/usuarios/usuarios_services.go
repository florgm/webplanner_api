package usuarios

import (
	"fmt"
	"../../db"
	"encoding/json"
	usuariosDomain "../../domain/usuarios"
)

//ParseLoginUsuario esto es una funcion
func ParseLoginUsuario(data []byte) (*usuariosDomain.LoginUsuario, error) {
    var usuario usuariosDomain.LoginUsuario
    if err := json.Unmarshal(data, &usuario); err != nil {
        return nil, err
	}
	
    return &usuario, nil
}

//ParseUsuario esto es una funcion
func ParseUsuario(data []byte) (*usuariosDomain.Usuarios, error) {
    var usuario usuariosDomain.Usuarios
    if err := json.Unmarshal(data, &usuario); err != nil {
        return nil, err
	}
	
    return &usuario, nil
}

//Login funcion
func Login(usuario *usuariosDomain.LoginUsuario) (*usuariosDomain.Usuarios, error) {
	var user usuariosDomain.Usuarios
	stmt, err := db.Init().Prepare("select * from usuarios where usuario = ? and password = ?;")

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	
	result, err := stmt.Query(usuario.Usuario, usuario.Password)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		err = result.Scan(
			&user.IDUsuario,
			&user.Nombre,
			&user.Usuario,
			&user.Password)
		if err != nil {
			fmt.Print(err.Error())
			return nil, err
		}
	}

	defer result.Close()
	defer stmt.Close()
	return &user, nil
}