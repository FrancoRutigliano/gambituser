package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/FrancoRutigliano/gambituser/awsgo"
	"github.com/FrancoRutigliano/gambituser/bd"
	"github.com/FrancoRutigliano/gambituser/models"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

// Llamamos a la lambda desde main mediante lambda.start
func main() {
	lambda.Start(EjecutoLambda)
}

//Funcion que ejecuta lambda y parametros con los que vamos a trabajar dentro de la lambda

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros. deber enviar 'SecretName'")
		err := errors.New("error en los parametros debe enviar SecretName")

		return event, err
	}

	var datos models.SignUp

	//El evento viene con coleccion de valores (atributos de user) y yo tengo que iterar hasta que, encuentre el valor deseado

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}

	err := bd.ReadSecret()
	// Esto es si no puedo leer el secreto no me conecto a la bd, por ende hay que abortar
	if err != nil {
		fmt.Println("Error al leer el Secret " + err.Error())
		return event, err
	}

	//En la variable error vamos a grabar el resultado de lo que devuelve signup y le pasamos los datos en la variable llamada datos
	err = bd.SignUp(datos)
	//Si aca devuelve de manera correcta err debería ser igual a nil, es decir sin errores y en caso de errores se devuelven
	return event, err
}

func ValidoParametros() bool {
	var traeParametro bool
	// Esto va a traerme en caso de encontrarla la variable de entorno de SecretName
	_, traeParametro = os.LookupEnv("SecretName")

	return traeParametro
}
