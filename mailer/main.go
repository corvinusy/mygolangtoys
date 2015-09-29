package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	smtpServer := "smtp.gmail.com"
	smtpPort := ":587"
	auth := smtp.PlainAuth(
		"",
		"marat.noreply@gmail.com",
		"marat_noreply",
		smtpServer,
	)

	from := mail.Address{
		Name:    "Prototype License Server",
		Address: "marat.noreply@gmail.com",
	}
	to := mail.Address{
		Name:    "Client",
		Address: "mz5corvinus@gmail.com",
	}
	title := "Test subject"

	body := "This is test email body"

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpServer+smtpPort,
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		log.Fatal(err)
	}
}

/*package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"marat.noreply@gmail.com",
		"marat_noreply",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"marat.noreply@gmail.com",
		[]string{"mz5corvinus@gmail.com"},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
*/
