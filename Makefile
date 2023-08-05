KAFKA_VERSION = kafka_2.13-3.5.0
TOPIC = test-topic-golang

default:
	wget https://dlcdn.apache.org/kafka/3.5.0/${KAFKA_VERSION}.tgz
	tar -xzf ${KAFKA_VERSION}.tgz

create-topic:
	./${KAFKA_VERSION}/bin/kafka-topics.sh \
		--create \
		--topic ${TOPIC} \
		--bootstrap-server localhost:29092,localhost:39092
