# Learning Kafka

This is a repository to save my kafka progress. I'll be working with Go while learning the Kafka technologies. *Remember to remove the docker containers* after finishing a project.

## Project 01

This my first project working with Kafka.
A simple kafka producer and consumer using Go.

Some important docker commands:

### bash

```bash
# to start the project
sudo service docker start
docker-compose up -d

# execute docker containers
docker exec -it go-kafka bash
docker exec -it kafka-container bash
```

### kafka

```bash
# create topic
kafka-topics --create --bootstrap-server localhost:9092 --topic first --partitions 3

# console consumer (for early tests) (to see all messages add --from-beginning)
kafka-console-consumer --bootstrap-server localhost:9092 --topic first --group --first-consumers

# console producer (for early tests)
kafka-console-producer --bootstrap-server localhost:9092 --topic first
```

### go

```bash
# run consumer
go run cmd/consumer/main.go
# run producer
go run cmd/producer/main.go
```

## Project 02

This is a bit more complicated project using go routines to publish and send emails.