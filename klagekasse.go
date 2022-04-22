package main

import (
	"fmt"
	helpers "github.com/gadelkareem/go-helpers"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Get arguments
	from := os.Args[1]

	// Create a random inquiry ID
	idLen := 7
	maxId := math.Pow(10, float64(idLen))
	id := rand.Intn(int(maxId))
	strId := string(id)[:idLen]

	// Expand with leading zeroes if not long enough
	if len(strId) < idLen {
		strId = strings.Repeat("0", idLen-len(strId)) + strId
	}

	// Send an acknowledgement email
	subject := fmt.Sprintf("[%s] Your inquiry has been received", strId)
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
