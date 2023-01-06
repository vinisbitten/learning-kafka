package main

import (
	"encoding/json"
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

	myMail, err := mail.NewMail()
	if err != nil {
		log.Fatal("Erro ao ler variáveis de ambiente", err.Error())
	}

	myMail.AddReceiver("yourmail@yourdomain.com")
	myMail.Message = mail.Message{
		To:      myMail.To(),
		Body:    "Você ganhou 1 milhão de reais",
		Subject: "Loteria",
	}

	kafkaMailMessage := JsonMail(*myMail)

	fmt.Println(string(kafkaMailMessage))

	// key might be []byte("email")
	Publish(kafkaMailMessage, "mail", producer, nil, deliveryChan)

	go DeliveryReport(deliveryChan)

	producer.Flush(5000)
}

// JsonMail turns mail struct into json
func JsonMail(mail mail.Mail) (producerMessage []byte) {
	producerMessage, err := json.Marshal(mail.Message)
	if err != nil {
		log.Println("Erro ao codificar json:", err.Error())
		return
	}
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
