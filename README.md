![banner](/assets/Banner.png)

# Learning Kafka

>This is a repository to save my learning progress in kafka. I'll be working with Go while learning the Kafka technologies. **Remember to remove the docker containers after finishing each project**.

Go to project:

>* <a href="p01">Project 01</a>
>* <a href="p02">Project 02</a>

<h2 id="p01">## Project 01</h2>

>* Here we'll learn the **basic kafka structure**.
>* We'll create simple kafka producer and consumer using Go.

*Some important commands:*

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

run these commands inside "kafka-container" container.

```bash
# create topic
kafka-topics --create --bootstrap-server localhost:9092 --topic first --partitions 3

# console consumer (for early tests) (to see all messages add --from-beginning)
kafka-console-consumer --bootstrap-server localhost:9092 --topic first --group --first-consumers

# console producer (for early tests)
kafka-console-producer --bootstrap-server localhost:9092 --topic first
```

### go

run these commands inside "go-kafka" container.

```bash
# run consumer
go run cmd/consumer/main.go
# run producer
go run cmd/producer/main.go
```

<h2 id="p02">Project 02</h2>

>* Producing and consuming e-mails with kafka.
>* We'll use the go smtp package to **send e-mails**.

* In this project there's a struct called *mail* that has all the attributes and methods we'll use while working with the e-mails.
* You will have to create a mailconf.yaml file with your email config (smtp config, email and password)
* You will be using the same kafka and docker configuration from project 01.

*Some important commands:*

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
# name it "mailconf.yaml"
# put it in in the directory /02
# use your real config
host:     "smtp.domain.com"
port:     "111"
id:       "yourmail@domain.com"
password: "1234password"
```
