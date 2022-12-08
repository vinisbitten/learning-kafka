package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
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
	// create delivery channel
	deliveryChan := make(chan kafka.Event)

	producer, err := NewKafkaProducer()
	if err != nil {
		log.Println(err.Error())
	}

	Publish([]byte("transferiu"), "first", producer, []byte("transferencia"), deliveryChan)

	// routine to check errors and show message in terminal
	go DeliveryReport(deliveryChan)

	// wait for message to complete delivery
	producer.Flush(5000)
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
