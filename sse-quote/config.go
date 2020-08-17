package ssequote

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DebugMode   bool
	Redis       RedisConfig
	ServiceName string
	Rest        RestConfig
	Cloud       CloudConfig
	Socket      SocketConfig
	GinMode     string
}

type SocketConfig struct {
	Port int
}

type CloudConfig struct {
	Url string
	Key string
}

type RestConfig struct {
	CacheInterval int
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

type ConfigInterface interface {
	GetHost() string
}

func NewConfig() *Config {
	return &Config{
		DebugMode: getEnvAsBool("DEBUG_MODE", false),
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "127.0.0.1"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
		},
		ServiceName: getEnv("SERVICE_NAME", ""),
		Rest: RestConfig{
			CacheInterval: getEnvAsInt("REST_CACHE_INTERVAL", 60),
		},
		Cloud: CloudConfig{
			Url: getEnv("SSE_CLOUD_URL", "https://sandbox-sse.iexapis.com/stable/"),
			Key: getEnv("CURL_CLOUD_PKEY", ""),
		},
		Socket: SocketConfig{
			Port: getEnvAsInt("SOCKET_PORT", 9090),
		},
		GinMode: getEnv("GIN_MODE", "debug"),
	}
}

func (r *RedisConfig) GetHost() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

func (r *CloudConfig) GetHost() string {
	return r.Url
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
