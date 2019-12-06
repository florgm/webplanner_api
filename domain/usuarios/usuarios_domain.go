package usuarios

//Usuarios esto es una estructura
type Usuarios struct {
	IDUsuario int64  `json:"id_usuario"`
	Nombre 	  string `json:"nombre"`
	Usuario	  string `json:"usuario"`
	Password  string `json:"password"`
}

//LoginUsuario estructura para validar el login
type LoginUsuario struct {
	Usuario	 string `json:"usuario"`
	Password string `json:"password"`
}