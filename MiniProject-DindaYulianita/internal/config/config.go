package config

import (
    "os"
)

type Config struct {
    AppPort   string
    MysqlDSN  string
    JwtSecret string
    UploadDir string
}

func LoadConfig() *Config {
    cfg := &Config{
        AppPort:   getEnv("APP_PORT", "8080"),
        MysqlDSN:  getEnv("MYSQL_DSN", "user:pass@tcp(127.0.0.1:3306)/mini_db?charset=utf8mb4&parseTime=True&loc=Local"),
        JwtSecret: getEnv("JWT_SECRET", "secret"),
        UploadDir: getEnv("UPLOAD_DIR", "./uploads"),
    }
    return cfg
}

func getEnv(k, d string) string {
    v := os.Getenv(k)
    if v == "" {
        return d
    }
    return v
}
