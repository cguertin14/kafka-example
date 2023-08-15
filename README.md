# kafka-example

Kafka example project with consumer &amp; producer. 
Quickstart guide available [here](https://kafka.apache.org/quickstart).

## Setup

Run `make` to install Kafka tools to interact with the cluster.

## Deployment

Run `docker-compose up` to launch the Kafka stack.

## Libraries

[Sarama](https://github.com/IBM/sarama) is preferred over [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) because of its pure Go implementation and lack of dependencies on the `librdkafka` C library. A good comparison of the two libraries is made [here](https://www.reddit.com/r/golang/comments/olrd34/comment/h5h9xb3/?utm_source=share&utm_medium=web2x&context=3), which provides a great case for using Sarama.
