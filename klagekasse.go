package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func SendMail(addr, from, subject, body string, to []string) error {
	r := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Mail(r.Replace(from)); err != nil {
		return err
	}
	for i := range to {
		to[i] = r.Replace(to[i])
		if err = c.Rcpt(to[i]); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	msg := "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func main() {
	// Get arguments
	from := os.Args[1]

	// Get the message from stdin
	message := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			message = append(message, text)
		} else {
			break
		}
	}

	// Get the subject and sender format name from headers
	var sender string
	var subject string

	for i := 0; i < len(message) && message[i] != ""; i++ {
		// Check if this is the from header
		if strings.ToLower(message[i][:5]) == "from:" {
			sender = strings.TrimSpace(message[i][5:])

			// Check if this is the subject header
		} else if strings.ToLower(message[i][:8]) == "subject:" {
			subject = strings.TrimSpace(message[i][8:])
		}
	}

	// Send an acknowledgement email
	response := "Your complaint has been received"
	smtpServer := "localhost:25"
	replyFrom := "noreply@trok.no"
	to := []string{
		from,
	}

	SendMail(smtpServer, replyFrom, response, response+sender+subject, to)
}
