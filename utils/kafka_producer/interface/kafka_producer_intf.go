package interface_kafkaproducer

import requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"

type IKafkaProducer interface {
	KafkaNotificationProducer(message *requestmodels_posnrel.KafkaNotificationTopicModel) error
}
