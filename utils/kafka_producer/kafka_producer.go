package kafkaproducer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	config_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/config"
	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	interface_kafkaproducer "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/kafka_producer/interface"
)

type KafakProducer struct {
	Config config_postNrelSvc.KafkaConfigs
}

func NewKafkaProducer(config config_postNrelSvc.KafkaConfigs) interface_kafkaproducer.IKafkaProducer {
	return &KafakProducer{Config: config}
}

func (k *KafakProducer) KafkaNotificationProducer(message *requestmodels_posnrel.KafkaNotificationTopicModel) error {

	fmt.Println("---------------to kafkaProducer:", *message)
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{k.Config.KafkaPort}, configs)
	if err != nil {
		log.Println("---------kafka producer err--------", err)
		return err
	}

	msgJson, _ := marshalStructJson(message)

	msg := &sarama.ProducerMessage{Topic: k.Config.KafkaTopicNotification,
		Key:   sarama.StringEncoder(message.UserID),
		Value: sarama.StringEncoder(*msgJson)}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("\nerr sending message to kafkaproducer on partition: %s,error: %v", k.Config.KafkaTopicNotification, err)
	}
	log.Printf("[producer] partition id: %d; offset:%d, value: %v\n", partition, offset, msg)
	return nil
}

func marshalStructJson(msgModel interface{}) (*[]byte, error) {
	data, err := json.Marshal(msgModel)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
