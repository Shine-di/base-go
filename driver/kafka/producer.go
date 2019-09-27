package kafka

import (
	"cortex3/conf"
	"github.com/Shopify/sarama"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"strings"
)

func Producer(topic string, value interface{}) {
	urls := conf.Yaml.Conf.Kafka.Host
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic//topic没有的话会新建
	b,err := json.Marshal(value)
	if err != nil  {
		log.Error(err)
		return
	}
	msg.Value = sarama.StringEncoder(string(b))
	client, err := sarama.NewSyncProducer(strings.Split(urls, ","), config)

	if err != nil {
		log.Error("producer close err:", err)
		return
	}
	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		log.Error("send message failed,", err)
		return
	}
	log.Info("pid:%s offset:%s\n", pid, offset)
}