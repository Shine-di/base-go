package kafka

import (
	"cortex3/conf"
	"github.com/bsm/sarama-cluster"
	"time"
	"github.com/Shopify/sarama"
	"strings"
	"github.com/labstack/gommon/log"
)

func Consumer() {
	groupID := conf.Yaml.Conf.Kafka.Group
	topicList := conf.Yaml.Conf.Kafka.Input
	urls := conf.Yaml.Conf.Kafka.Host
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始

	consumer, err := cluster.NewConsumer(strings.Split(urls, ","), groupID, strings.Split(topicList, ","), config)
	if err != nil {
		log.Error(err)
		return
	}

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				consumer.MarkOffset(msg, "")
			}
		case err := <-consumer.Errors():
			log.Error("err :%s\n", err.Error())
		case ntf := <-consumer.Notifications():
			log.Info("Rebalanced: %+v\n", ntf)
		}
	}
}

