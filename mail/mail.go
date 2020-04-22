package mail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
)

/* Estructura del destinatario */
type Dest struct {
	Name  string
	Email string
}

//Funcion para enviar un email por medio de un template
func SendMail(name string, email string, mtemplate string) {
	/* Destinatario */
	from := mail.Address{"Links app", "josealberto1107@hotmail.com"}
	/* Remitente */
	to := mail.Address{name, email}
	/* Asunto */
	subject := "Recover password from Links app"
	/* Estructura del destinatario */
	dest := Dest{Name: to.Address}

	/* Encabezado del correo */
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	/* Especificar el tipo de archivo a enviar en este caso un html */
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	message := ""

	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	/* t, err := template.ParseFiles(path.Join("/templates", mtemplate)) */
	t, err := template.ParseFiles("correo.html")
	controlError(err)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	controlError(err)

	message += buf.String()

	/* Servidor a usar para el envio */
	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	/* Autenticacion plana */

	auth := smtp.PlainAuth("", "ljose297@gmail.com", "Nicromano11", host)

	tslConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tslConfig)
	controlError(err)

	client, err := smtp.NewClient(conn, host)
	controlError(err)
	err = client.Auth(auth)
	controlError(err)

	/* Quien va a ser el que envia y quien recibe */

	err = client.Mail(from.Address)
	controlError(err)
	err = client.Rcpt(from.Address)
	controlError(err)

	w, err := client.Data()
	controlError(err)

	_, err = w.Write([]byte(message))
	controlError(err)

	err = w.Close()
	controlError(err)

	client.Quit()

}

func controlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
