package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vinisbitten/learning-kafka/02/cmd/mail"
)

func main() {
	consumer, err := NewKafkaConsumer()
	if err != nil {
		log.Println("error consumer:", err.Error())
	}

	topics := []string{"mail"}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Println("Erro ao inscrever-se em topico", err.Error())
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			// IMPLEMENT EMAIL SENDING
			// fmt.Println(string(msg.Value), msg.TopicPartition)
			mailSchema := WorkKafkaMessage(msg)

		} else {
			log.Print("Mensagem com erro:", err.Error())
		}
	}
}

func NewKafkaConsumer() (consumer *kafka.Consumer, err error) {
	// map containing configuration
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-container:9092",
		// consumer id
		"client.id": "go-mail-consumer",
		// kafka group id
		"group.id": "mail-consumers",

		// to consume all messages
		// "auto.offset.reset": "earliest",
	}
	consumer, err = kafka.NewConsumer(configMap)
	return
}

func WorkKafkaMessage(msg *kafka.Message) (mailSchema mail.Mail){
	slpitedMsg := strings.Split(string(msg.Value), ",")
	mailMessage := make()
}

// func newMail(body string, subject string, receiver mail.Receiver, sender mail.Sender, smtp mail.SmtpServerConf) (mailMessage []byte) {
// 	mailSchema := mail.Mail{
// 		Body:     body,
// 		Subject:  subject,
// 		Receiver: receiver,
// 		Sender:   sender,
// 		Smtp:     smtp,
// 	}
// 	mailMessage = []byte(fmt.Sprintf("%#v", mailSchema))
// 	return
// }
