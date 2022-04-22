package main

import (
	"fmt"
	helpers "github.com/gadelkareem/go-helpers"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get arguments
	from := os.Args[1]

	// Create a random inquiry ID
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	idLen := 7
	maxId := math.Pow(10, float64(idLen))
	id := randGen.Intn(int(maxId))
	strId := strconv.Itoa(id)

	// Expand with leading zeroes if not long enough
	if len(strId) < idLen {
		strId = strings.Repeat("0", idLen-len(strId)) + strId
	}

	// Retract if too long
	if len(strId) > idLen {
		strId = strId[:idLen]
	}

	// Send an acknowledgement email
	subject := fmt.Sprintf("[%s] Your inquiry has been received", strId)
	smtpServer := "localhost:25"
	replyFrom := "noreply@trok.no"
	to := []string{
		from,
	}

	// Get the acknowledgement email from the PHP file
	response, err := exec.Command("php", "/usr/lib/klagekasse/acknowledgement.php", strId).Output()
	if err != nil {
		log.Fatal(err)
	}

	err = helpers.SendMail(smtpServer, replyFrom, subject, string(response), to)
	if err != nil {
		log.Fatal(err)
	}

	// Get the time to rejection email
	minTime := 1440
	maxTime := 7200
	replyTime := 0

	for replyTime < minTime {
		replyTime = randGen.Int() % maxTime
	}

	// Schedule the rejection email
	cmd := exec.Command("at", "now", "+", strconv.Itoa(replyTime), "minutes")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(stdin, fmt.Sprintf("php /usr/lib/klagekasse/rejection.php %s %s \"[%s] Your inquiry has been closed\" | /usr/sbin/sendmail %s", strId, from, strId, from))
	if err != nil {
		log.Fatal(err)
	}

	cmd.Start()

}
