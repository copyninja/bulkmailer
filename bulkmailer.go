package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/scorredoira/email"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"
)

var addressFile string
var contentFile string
var subjectLine string
var fromAddress string
var smtpServer string
var user string
var password string

func init() {
	flag.StringVar(&fromAddress, "from", "",
		"From address for the mail")
	flag.StringVar(&subjectLine, "subject", "",
		"Subject for the mail")
	flag.StringVar(&addressFile, "addresses", "",
		"File path containing mail addresses")
	flag.StringVar(&contentFile, "content", "",
		"File path containing mail content to be sent")
	flag.StringVar(&smtpServer, "server", "",
		"SMTP server for sending mails")
	flag.StringVar(&user, "username", "",
		"Username for SMTP server login")
	flag.StringVar(&password, "password", "",
		"Password for SMTP server login")
}

type attachments []string

func (a *attachments) String() string {
	return fmt.Sprint(*a)
}

func (a *attachments) Set(value string) error {
	if len(*a) > 0 {
		return errors.New("Flag is already set")
	}

	for _, at := range strings.Split(value, ",") {
		if _, err := os.Stat(at); err != nil {
			return err
		}

		*a = append(*a, at)
	}

	return nil
}

var files attachments

func init() {
	flag.Var(&files, "attachment",
		"comma-separated list of attachments (file path)")
}

func validateArgs() error {
	if len(subjectLine) == 0 {
		return errors.New("Please provide a subject line for mail")
	}

	if len(addressFile) == 0 {
		return errors.New(`Please provide file containing mail \
		addresses`)
	}

	if len(contentFile) == 0 {
		return errors.New(`Please provide file containing mail \
		body`)
	}

	if len(fromAddress) == 0 {
		return errors.New(`Please provide valid From address`)
	}

	if len(smtpServer) == 0 {
		return errors.New(`Please provide valid SMTP server to \
		use`)
	}

	if _, err := os.Stat(addressFile); err != nil {
		return err
	}

	if _, err := os.Stat(contentFile); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := validateArgs(); err != nil {
		log.Fatalln(err)
	}

	tmp, err := ioutil.ReadFile(addressFile)
	if err != nil {
		log.Fatalln(err)
	}

	var addresses []Address
	for _, addr := range bytes.Split(tmp, []byte{'\n'}) {
		if len(bytes.TrimSpace(addr)) > 0 {
			addresses = append(addresses,
				NewAddress(string(addr)))
		}
	}

	content, err := ioutil.ReadFile(contentFile)
	if err != nil {
		log.Fatalln(err)
	}

	auth := smtp.PlainAuth("", user, password, smtpServer)
	for _, mail := range addresses {
		body := strings.Replace(string(content), "REPLACE_ME",
			mail.FirstName, 1)
		m := email.NewMessage(subjectLine, body)
		m.From = fromAddress
		m.To = []string{fmt.Sprint(mail)}

		if err := email.Send(smtpServer, auth, m); err != nil {
			log.Fatalln(err)
		}

		log.Printf("Mail sent to %s successfully\n",
			strings.Join([]string{mail.FirstName,
				mail.LastName}, " "))
	}
}
