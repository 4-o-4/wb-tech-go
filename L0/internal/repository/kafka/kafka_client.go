package kafka

import (
    "L0/internal/domain"
    "L0/internal/repository/postgres"
    "context"
    "encoding/json"
    "github.com/segmentio/kafka-go"
    "log"
)

type KafkaClient struct {
    reader *kafka.Reader
}

func NewKafkaClient(broker string) (*KafkaClient, error) {
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{broker},
        Topic:   "orders",
    })
    return &KafkaClient{reader: reader}, nil
}

func (kc *KafkaClient) StartConsumer(db *postgres.OrderRepository) {
    for {
        msg, err := kc.reader.ReadMessage(context.Background())
        if err != nil {
            log.Println("Error reading message: ", err)
            continue
        }
        var order domain.Order
        if err := json.Unmarshal(msg.Value, &order); err != nil {
            log.Println("Error unmarshalling message: ", err)
            continue
        }
        if err := db.SaveOrder(&order); err != nil {
            log.Println("Error saving order to database: ", err)
        }
    }
}
