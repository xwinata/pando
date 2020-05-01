package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pando/server/helper"
	"strconv"

	"github.com/gofort/dispatcher"
	"github.com/labstack/echo/v4"
)

// TestQueueAtask test sending task to worker
func TestQueueAtask(c echo.Context) error {
	defer c.Request().Body.Close()

	QueueTask("job_example", []dispatcher.TaskArgument{})

	return helper.ReturnJSONresp(c, http.StatusOK, "0000", "Success", "done")
}

// QueueTask adds a task into queue
func QueueTask(tName string, args []dispatcher.TaskArgument) {
	amqpUser := os.Getenv("PANDO_AMQP_USER")
	amqpPass := os.Getenv("PANDO_AMQP_PASS")
	amqpHost := os.Getenv("PANDO_AMQP_HOST")
	amqpPort := os.Getenv("PANDO_AMQP_PORT")
	amqpReconnectForever, _ := strconv.ParseBool(os.Getenv("PANDO_AMQP_RECONNECT_FOREVER"))
	amqpReconnectRetries, _ := strconv.ParseInt(os.Getenv("PANDO_AMQP_RECONNECT_RETRIES"), 10, 64)
	amqpReconnectInterval, _ := strconv.ParseInt(os.Getenv("PANDO_AMQP_RECONNECT_INTERVAL"), 10, 64)
	amqpDebugMode, _ := strconv.ParseBool(os.Getenv("PANDO_AMQP_RECONNECT_DEBUGMODE"))
	qName := os.Getenv("PANDO_QUEUE_NAME")

	serverConfig := dispatcher.ServerConfig{
		AMQPConnectionString:        fmt.Sprintf("amqp://%s:%s@%s:%s", amqpUser, amqpPass, amqpHost, amqpPort), // "amqp://user:pass@0.0.0.0:5672/"
		ReconnectionRetriesForever:  amqpReconnectForever,
		ReconnectionRetries:         int(amqpReconnectRetries),
		ReconnectionIntervalSeconds: amqpReconnectInterval,
		DebugMode:                   amqpDebugMode, // enables extended logging
		Exchange:                    os.Getenv("PANDO_AMQP_EXCHANGE_NAME"),
		InitQueues: []dispatcher.Queue{ // creates queues and binding keys if they are not created already
			{
				Name:        qName,
				BindingKeys: []string{qName}, // BindingKeys: []string{BindName},
			},
		},
		DefaultRoutingKey: qName, // DefaultRoutingKey: BindName, // default routing key which is used for publishing messages
	}

	server, _, err := dispatcher.NewServer(&serverConfig)
	if err != nil {
		panic(err)
	}

	tasks := dispatcher.Task{
		Name: tName,
		Args: args,
	}

	if err = server.Publish(&tasks); err != nil {
		log.Println(err)
	}

	server.Close()
}
