package main

import (
    "L0/internal/cache"
    "L0/internal/config"
    "L0/internal/delivery/http"
    "L0/internal/repository/kafka"
    "L0/internal/repository/postgres"
    "L0/internal/usecase"
    "log"
    "os"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Println("Failed to load config: ", err)
        os.Exit(1)
    }

    db, err := postgres.NewPostgresConnection(cfg.DatabaseURL)
    if err != nil {
        log.Println("Failed to connect to database: ", err)
        os.Exit(1)
    }

    kafkaClient, err := kafka.NewKafkaClient(cfg.KafkaBroker)
    if err != nil {
        log.Println("Failed to connect to Kafka: ", err)
        os.Exit(1)
    }

    memoryCache := cache.NewMemoryCache()
    useCase := usecase.NewOrderUseCase(db, memoryCache)

    go kafkaClient.StartConsumer(db)

    handler := http.NewHandler(useCase)
    handler.StartServer(cfg.HTTPPort)
}
