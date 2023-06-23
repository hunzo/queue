package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"golang-consumer/config"
	"golang-consumer/graylog"
	"golang-consumer/models"
	"golang-consumer/rabbitmq"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

func main() {

	queue, err := rabbitmq.New()
	if err != nil {
		log.Fatal(err)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	QUEUE_NAME := config.New().QueueName

	fmt.Printf("queue_name: %s\n", QUEUE_NAME)
	fmt.Printf("graylog_server: %s\n", config.New().GraylogSrv)

	forever := make(chan bool)

	err = queue.StartCosumer(QUEUE_NAME, "consumer_1", CallBack)

	if err != nil {
		fmt.Printf("error %s\v", err)
		log.Fatal(err)
	}
	<-forever

}

func CallBack(d amqp.Delivery) {

	if d.Body == nil {
		fmt.Printf("No message")
	}

	// fmt.Printf("routing: %s ", d.RoutingKey)
	// fmt.Printf("timestamp: %s\n", time.Now())

	var ReqPayload models.GraylogPayload
	if err := json.Unmarshal(d.Body, &ReqPayload); err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	// Logging to graylog
	// graylog.LogToGraylog(ReqPayload, res)

	graylog.LogToGraylog(ReqPayload)

	fmt.Printf("%+v", ReqPayload)

	fmt.Printf("\n----------------------------\n")

	// fmt.Printf("body: %v\n", ReqPayload)

	d.Ack(false)

}
