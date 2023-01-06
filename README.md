# Learning Kafka

This is a repository to save my kafka progress. I'll be working with Go while learning the Kafka technologies.Remember to remove the docker containers after finishing each project.

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

### kafka (inside kafka-container container)

```bash
# create topic
kafka-topics --create --bootstrap-server localhost:9092 --topic first --partitions 3

# console consumer (for early tests) (to see all messages add --from-beginning)
kafka-console-consumer --bootstrap-server localhost:9092 --topic first --group --first-consumers

# console producer (for early tests)
kafka-console-producer --bootstrap-server localhost:9092 --topic first
```

### go (inside go-kafka container)

```bash
# run consumer
go run cmd/consumer/main.go
# run producer
go run cmd/producer/main.go
```

## Project 02

In this project I used a Kafka to send emails.
The mail package has all the functions to consume kafka messages.
You will have to create a mailConfig.yaml file with your email config (smtp config, email and password).
You will be using the same command from project 01 to run the docker containers and kafka config.

### go

```bash
# run consumer (inside consumer container)
go run cmd/mail-consumer/mail-consumer.go
# run producer (inside producer container)
go run cmd/mail-producer/mail-producer.go
```

### yaml

```yaml
# create a file like this
# name it "mailConfig.yaml"
# put it in in the directory /02
# use your real config
host:     "smtp.domain.com"
port:     "111"
id:       "yourmail@domain.com"
password: "1234password"
```
