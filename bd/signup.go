package bd

import (
	"fmt"

	"github.com/FrancoRutigliano/gambituser/models"
	"github.com/FrancoRutigliano/gambituser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza registro")

	//Estan en el mismo paquete por eso no hace falta colocar el db. y le estamos diciendo que cuando se ejecute el signup se conecte la base de datos
	err := DbConnect()
	if err != nil {
		return err
	}
	// Instruccion que se va comentar al final de la funcion
	// En este caso debo asegurarme de cerrar la DB
	defer Db.Close()

	//Sentencia SQL
	//Resivimos un objeto llamado sig
	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"

	fmt.Println(sentencia)

	/*
		db.Exec se utiliza para ejecutar sentencias SQL en una base de datos. db.Exec se usa específicamente cuando se necesita ejecutar una consulta SQL que no devuelve filas de resultados y recibe como parametro un query string.
	*/

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//En caso de que ande todo bien devolvemos nil y imprimos lo siguiente para corroborar en cloudwatch
	fmt.Println("SignUp > Ejecucción Exitosa")
	return nil

}
