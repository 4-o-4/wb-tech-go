package config

import "github.com/spf13/viper"

type Config struct {
    DatabaseURL string
    KafkaBroker string
    HTTPPort    string
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yml")
    viper.AddConfigPath("./configs")
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }
    config := &Config{
        DatabaseURL: viper.GetString("database_url"),
        KafkaBroker: viper.GetString("kafka_broker"),
        HTTPPort:    viper.GetString("http_port"),
    }
    return config, nil
}
