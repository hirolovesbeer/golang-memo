package main

// reference: https://github.com/rabbitmq/rabbitmq-tutorials/blob/master/go/rpc_client.go

import (
    "fmt"
    "log"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func randomString(l int) string {
    bytes := make([]byte, l)
    for i := 0; i < l; i++ {
        bytes[i] = byte(randInt(65, 90))
    }

    return string(bytes)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func echoRPC(n string) (res string, err error) {
    conn, err := amqp.Dial("amqp://rpc:rpc@targethostname:5672//rpc")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "",    // name
        false, // durable
        false, // delete when usused
        true,  // exclusive
        false, // noWait
        nil,   // arguments
    )
    failOnError(err, "Failed to declare a queue")

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consume
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    corrId := randomString(32)

    err = ch.Publish(
        "",          // exchange
        "rpc_queue", // routing key
        false,       // mandatory
        false,       // immediate
        amqp.Publishing {
            ContentType:    "text/plain",
            CorrelationId:  corrId,
            ReplyTo:        q.Name,
            Body:           []byte(n),
        })
    failOnError(err, "Failed to publish a message")

    for d := range msgs {
        if corrId == d.CorrelationId {
            res = string(d.Body)
            break
        }
    }

    return
}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    n := bodyFrom(os.Args)

    for i := 0; i < 1000; i++ {
        m := fmt.Sprintf("%s : %s", n, strconv.Itoa(i))

        log.Printf(" [x] Requesting echo(%s)", m)
        res, err := echoRPC(m)
        failOnError(err, "Failed to handle RPC request")

        log.Printf(" [.] Got %s", res)
    }
}

func bodyFrom(args []string) string {
    var s string
    if (len(args) < 2) || os.Args[1] == "" {
        s = "hoge"
    } else {
        s = strings.Join(args[1:], " ")
    }

    return s
}
