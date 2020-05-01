package worker

import (
	"fmt"
	"os"
	"os/signal"
	"pando/worker/jobs"
	"strconv"
	"syscall"

	"github.com/gofort/dispatcher"
)

// Start starts worker
func Start() {
	amqpUser := os.Getenv("PANDO_AMQP_USER")
	amqpPass := os.Getenv("PANDO_AMQP_PASS")
	amqpHost := os.Getenv("PANDO_AMQP_HOST")
	amqpPort := os.Getenv("PANDO_AMQP_PORT")
	amqpReconnectForever, _ := strconv.ParseBool(os.Getenv("PANDO_AMQP_RECONNECT_FOREVER"))
	amqpReconnectRetries, _ := strconv.ParseInt(os.Getenv("PANDO_AMQP_RECONNECT_RETRIES"), 10, 64)
	amqpReconnectInterval, _ := strconv.ParseInt(os.Getenv("PANDO_AMQP_RECONNECT_INTERVAL"), 10, 64)
	amqpDebugMode, _ := strconv.ParseBool(os.Getenv("PANDO_AMQP_RECONNECT_DEBUGMODE"))

	workerName := os.Getenv("PANDO_WORKER_NAME")
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

	workerConfig := dispatcher.WorkerConfig{
		Queue: qName,
		Name:  workerName,
	}

	tasks := DefineWorkerTasks()

	worker, err := server.NewWorker(&workerConfig, tasks)
	if err != nil {
		panic(err)
	}

	if err := worker.Start(server); err != nil {
		panic(err)
	}

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// When you push CTRL+C close worker gracefully
	<-sig

	server.Close()
}

// DefineWorkerTasks lists worker tasks
func DefineWorkerTasks() map[string]dispatcher.TaskConfig {
	tasks := make(map[string]dispatcher.TaskConfig)

	tasks["job_example"] = dispatcher.TaskConfig{
		Function: func() {
			jobs.JobTaskExample()
		},
	}

	return tasks
}
