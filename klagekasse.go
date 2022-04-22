package main

import (
	helpers "github.com/gadelkareem/go-helpers"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Get arguments
	from := os.Args[1]

	// Send an acknowledgement email
	subject := "Your complaint has been received"
	smtpServer := "localhost:25"
	replyFrom := "noreply@trok.no"
	to := []string{
		from,
	}

	// Get the acknowledgement email from the PHP file
	response, err := exec.Command("php", "/usr/lib/klagekasse/acknowledgement.php").Output()
	if err != nil {
		log.Fatal(err)
	}

	helpers.SendMail(smtpServer, replyFrom, subject, string(response), to)
}
