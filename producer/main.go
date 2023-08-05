// inspired by: https://gist.github.com/RonnanSouza/a62e2fff0acd6e5d1271d0da0ef1e9b3

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

var (
	brokers = flag.String("bootstrap-servers", "localhost:29092,localhost:39092", "The URL of the Kafka broker to use")
	topic   = flag.String("topic", "test-topic-golang", "The name of the Kafka topic to use")
)

func main() {
	producer, err := setupProducer()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Kafka AsyncProducer up and running!")
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	numberOfMessages := produceMessages(producer, signals)
	log.Printf("Kafka AsyncProducer finished with %d messages produced.", numberOfMessages)
}

// setupProducer will create a AsyncProducer and returns it
func setupProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	return sarama.NewAsyncProducer(strings.Split(*brokers, ","), config)
}

// produceMessages will send 'testing 123 - <timestamp>' to KafkaTopic each second,
// until receive a os signal to stop e.g. control + c by the user in terminal
func produceMessages(producer sarama.AsyncProducer, signals chan os.Signal) (numberOfMessages int) {
	for {
		time.Sleep(100 * time.Millisecond)
		message := &sarama.ProducerMessage{
			Topic: *topic,
			Value: sarama.StringEncoder(
				fmt.Sprintf("%s - %d", "testing 123", time.Now().UnixNano()),
			),
		}
		select {
		case producer.Input() <- message:
			numberOfMessages++
			log.Println("New Message produced")
		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			return
		}
	}
}
