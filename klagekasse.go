package main

import (
	"fmt"
	helpers "github.com/gadelkareem/go-helpers"
	"log"
	"math"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get arguments
	//from := os.Args[1]

	// Create a random inquiry ID
	idLen := 7
	maxId := math.Pow(10, float64(idLen))
	id := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(int(maxId))
	strId := strconv.Itoa(id)

	// Expand with leading zeroes if not long enough
	if len(strId) < idLen {
		strId = strings.Repeat("0", idLen-len(strId)) + strId
	}

	// Retract if too long
	if len(strId) > idLen {
		strId = strId[:idLen]
	}

	fmt.Println(strId)
	return
	from := ""

	// Send an acknowledgement email
	subject := fmt.Sprintf("[%s] Your inquiry has been received", strId)
	smtpServer := "localhost:25"
	replyFrom := "noreply@trok.no"
	to := []string{
		from,
	}

	// Get the acknowledgement email from the PHP file
	response, err := exec.Command("php", "/usr/lib/klagekasse/acknowledgement.php "+strId).Output()
	if err != nil {
		log.Fatal(err)
	}

	helpers.SendMail(smtpServer, replyFrom, subject, string(response), to)
}
