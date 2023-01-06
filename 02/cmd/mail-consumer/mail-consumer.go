package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vinisbitten/learning-kafka/02/cmd/mail"
)

func main() {
	consumer, err := NewKafkaConsumer()
	if err != nil {
		log.Println("error consumer:", err.Error())
	}

	// "topics" must be an array of strings
	topics := []string{"mail"}

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Println("Erro ao inscrever-se em topico", err.Error())
	}

	//Reading message loop
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			//Getting the env settings
			myMail, err := mail.NewMail()
			if err != nil {
				log.Fatal("Erro ao ler variáveis de ambiente", err.Error())
			}
			//Creating a mail from kafka response
			myMail.DecodeJson(msg.Value)
			fmt.Printf("%#v", *myMail)
			//Sending mail
			myMail.Send()
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
