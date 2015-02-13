package main

// reference: https://github.com/rabbitmq/rabbitmq-tutorials/blob/master/go/rpc_client.go

import (
    "fmt"
    "log"

    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func echo(msgs string) string {
    return msgs
}

func main() {
    conn, err := amqp.Dial("amqp://rpc:rpc@targethostname:5672//rpc")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "rpc_queue", // name
        false,       // durable
        false,       // delete when usused
        false,        // exclusive
        false,       // noWait
        nil,         // arguments
    )
    failOnError(err, "Failed to declare a queue")

    err = ch.Qos(
        1,     // prefetch count
        0,     // prefetch size
        false, // global
    )
    failOnError(err, "Failed to set QoS")

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consume
        false,  // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            n := string(d.Body)

            log.Printf(" [.] msgs(%s)", n)
            response := echo(n)

            err = ch.Publish(
                "",        // exchange
                d.ReplyTo, // routing key
                false,     // mandatory
                false,     // immediate
                amqp.Publishing{
                    ContentType:   "text/plain",
                    CorrelationId: d.CorrelationId,
                    Body:          []byte(response),
                })
            failOnError(err, "Failed to publish a message")

            d.Ack(false)
        }
    }()

    log.Printf(" [*] Awaiting RPC requests")
    <-forever
}
