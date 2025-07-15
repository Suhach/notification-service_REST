package kafkaNP

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"testAPI/pkg/logger"
)

type NotificaitonProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewNotificationProducer(brokers []string, topic string) (*NotificaitonProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		logger.Log.Warn("Failed to start Sarama producer", zap.Error(err))
		return nil, err
	}
	return &NotificaitonProducer{
		producer: producer,
		topic:    topic,
	}, nil
}

func (np *NotificaitonProducer) SendUUID(uuid string) error {
	msg := &sarama.ProducerMessage{
		Topic: np.topic,
		Key:   sarama.StringEncoder("notification_uuid"),
		Value: sarama.StringEncoder(uuid),
	}

	partition, offset, err := np.producer.SendMessage(msg)
	if err != nil {
		zap.L().Warn("Kafka send failed", zap.Error(err))
		return err
	}

	zap.L().Info("Kafka message sent",
		zap.String("uuid", uuid),
		zap.Int32("partition", partition),
		zap.Int64("offset", offset),
	)

	return nil
}
