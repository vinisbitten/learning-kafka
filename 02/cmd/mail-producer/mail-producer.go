package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vinisbitten/learning-kafka/02/cmd/mail"
)

// DeliveryReport check errors
func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada:", ev.TopicPartition)
			}
		}
	}
}

func main() {
	deliveryChan := make(chan kafka.Event)

	producer, err := NewKafkaProducer()
	if err != nil {
		log.Println("error producer:", err.Error())
	}

	body := "Seu aluguel de janeiro tá pago"
	sbjct := "Pagamentos"
	rc := mail.Receiver{
		Id: "guilherme@conceitho.com",
	}
	sd := mail.Sender{
		Id:       "suporte@conceitho.com",
		Password: "123456a.",
	}
	smtp := mail.SmtpServerConf{
		Host: "mail.conceitho.com",
		Port: "465",
	}

	mailMessage := newMail(body, sbjct, rc, sd, smtp)

	kafkaMailMessage := []byte(mailMessage)

	// key might be []byte("email")
	Publish(kafkaMailMessage, "mail", producer, nil, deliveryChan)

	go DeliveryReport(deliveryChan)

	producer.Flush(5000)
}

func newMail(body string, subject string, receiver mail.Receiver, sender mail.Sender, smtp mail.SmtpServerConf) (mailMessage []byte) {
	mailSchema := mail.Mail{
		Body:     body,
		Subject:  subject,
		Receiver: receiver,
		Sender:   sender,
		Smtp:     smtp,
	}
	mailMessage = []byte(fmt.Sprintf("%v", mailSchema))
	return
}

// NewKafkaProducer implements a High-level Apache Kafka Producer instance
func NewKafkaProducer() (producer *kafka.Producer, err error) {
	// map containing configuration
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-container:9092",
		// sets a limit on the time to wait for an acknowledgment of the success or failure to deliver a message
		"delivery.timeout.ms": "0",
		// denotes the number of brokers that must receive the record before we consider the write as successful
		"acks": "all",
		// determines whether the producer may write duplicates of a retried message
		"enable.idempotence": "true",
	}
	producer, err = kafka.NewProducer(configMap)
	return
}

// Publish publishes message to apache kafka
func Publish(msg []byte, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) {
	message := &kafka.Message{
		Value: msg,
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		log.Println(err.Error())
	}
}
