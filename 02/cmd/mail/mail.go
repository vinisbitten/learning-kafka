package mail

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/spf13/viper"
)

type Mail struct {
	sender    sender
	smtp      smtpConf
	Message   Message
	Receivers []string
}

type Message struct {
	To      string
	Body    string
	Subject string
}

type Receiver struct {
	id string
}

type sender struct {
	id       string
	password string
}

type smtpConf struct {
	host string
	port string
}

// AddReceiver adds a new receiver to the slice of receivers
func (m *Mail) AddReceiver(email string) {
	m.Receivers = append(m.Receivers, email)
}

// DecodeJson decodes the kafka response
func (m *Mail) DecodeJson(message []byte) {
	err := json.Unmarshal(message, &m.Message)
	if err != nil {
		log.Println("Error decoding json:", err.Error())
	}

	var receivers []string
	receivers = strings.Split(m.Message.To, ", ")
	m.Receivers = receivers
}

// NewMail gets the env information from de yaml file and updates the mail config
func NewMail() (mail *Mail, err error) {
	viper.AddConfigPath("/go/src")
	viper.SetConfigName("mailconf")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	mail = &Mail{
		smtp: smtpConf{
			host: viper.GetString("host"),
			port: viper.GetString("port"),
		},
		sender: sender{
			id:       viper.GetString("id"),
			password: viper.GetString("password"),
		},
	}
	return
}

// Send sends the e-mail using de smtp package
func (m *Mail) Send() {
	err := smtp.SendMail(m.smtp.address(), m.auth(), m.sender.id, m.Receivers, m.byte())
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Email sent successfully")
}

// To prints the receivers
func (m *Mail) To() string {
	var to string
	for _, receiver := range m.Receivers {
		to += receiver + ", "
	}
	to = to[:len(to)-2]
	return to
}

// auth returns a smtp.auth based on the mail credentials
func (m *Mail) auth() smtp.Auth {
	return smtp.PlainAuth("", m.sender.id, m.sender.password, m.smtp.host)
}

func (m *Mail) byte() []byte {
	return []byte("To: " + m.Message.To + "\r\n" +
		"Subject: " + m.Message.Subject + "\r\n" +
		"\r\n" +
		m.Message.Body + "\r\n")
}

func (s smtpConf) address() string {
	return s.host + ":" + s.port
}
