package main

import (
	helpers "github.com/gadelkareem/go-helpers"
	"os"
)

func main() {
	// Get arguments
	from := os.Args[1]

	// Send an acknowledgement email
	response := "Your complaint has been received"
	smtpServer := "localhost:25"
	replyFrom := "noreply@trok.no"
	to := []string{
		from,
	}

	helpers.SendMail(smtpServer, replyFrom, response, response, to)
}
