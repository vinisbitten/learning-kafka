package mail

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/smtp"
	"strings"
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

type sender struct {
	id       string
	password string
}

type smtpConf struct {
	host string
	port string
}

type Receiver struct {
	id string
}

func (m *Mail) AddReceiver(email string) {
	m.Receivers = append(m.Receivers, email)
}

func (m *Mail) DecodeJson(message []byte) {
	err := json.Unmarshal(message, &m.Message)
	if err != nil {
		log.Println("Error decoding json:", err.Error())
	}

	var receivers []string
	receivers = strings.Split(m.Message.To, ", ")
	m.Receivers = receivers
}

func (m *Mail) Send() {
	err := smtp.SendMail(m.smtp.address(), m.auth(), m.sender.id, m.Receivers, m.byte())
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Email sent successfully")
}

func (m *Mail) auth() smtp.Auth {
	return smtp.PlainAuth("", m.sender.id, m.sender.password, m.smtp.host)
}

func (m *Mail) byte() []byte {
	return []byte("To: " + m.Message.To + "\r\n" +
		"Subject: " + m.Message.Subject + "\r\n" +
		"\r\n" +
		m.Message.Body + "\r\n")
}

func (m *Mail) To() string {
	var to string
	for _, receiver := range m.Receivers {
		to += receiver + ", "
	}
	to = to[:len(to)-2]
	return to
}

func (s smtpConf) address() string {
	return s.host + ":" + s.port
}

func NewMail() (mail *Mail, err error) {
	viper.AddConfigPath("/go/src")
	viper.SetConfigName("mailConfig")
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
