package config

import (
	"os"
)

type AwsConfig struct {
	Region    string
	AccessKey string
	SecretKey string
}

type DynamoTableConfig struct {
	User    string
	Blog    string
	UserGSI string
	BlogGSI string
}

type Config struct {
	Env            string
	AwsConfig      AwsConfig
	ServerPort     string
	DynamoEndpoint string
	DynamoTable    DynamoTableConfig
}

var instance *Config

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Init() {
	if instance == nil {
		instance = &Config{
			Env: getEnvWithDefault("ENV", "dev"),
			AwsConfig: AwsConfig{
				Region:    getEnvWithDefault("AWS_REGION", ""),
				AccessKey: getEnvWithDefault("AWS_ACCESS_KEY_ID", ""),
				SecretKey: getEnvWithDefault("AWS_SECRET_ACCESS_KEY", ""),
			},
			ServerPort:     getEnvWithDefault("SERVER_PORT", "3000"),
			DynamoEndpoint: getEnvWithDefault("DYNAMO_ENDPOINT", ""),
			DynamoTable: DynamoTableConfig{
				User: getEnvWithDefault("USER_TABLE", "Users"),
				Blog: getEnvWithDefault("BLOG_TABLE", "Blogs"),
				UserGSI: getEnvWithDefault("USER_TABLE", "UsersGsi"),
				BlogGSI: getEnvWithDefault("BLOG_TABLE", "BlogsGsi"),
			},
		}
	}
}

func GetInstance() *Config {
	return instance
}

func GetEnv() string {
	return instance.Env
}

func GetAwsConfig() AwsConfig {
	return instance.AwsConfig
}

func GetServerPort() string {
	return instance.ServerPort
}

func GetDynamoEndpoint() string {
	return instance.DynamoEndpoint
}

func GetDynamoTable() DynamoTableConfig {
	return instance.DynamoTable
}

func GetDynamoUserTable() string {
	return instance.DynamoTable.User
}

func GetDynamoBlogTable() string {
	return instance.DynamoTable.Blog
}

func GetDynamoUserGsiTable() string {
	return instance.DynamoTable.UserGSI
}

func GetDynamoBlogGsiTable() string {
	return instance.DynamoTable.BlogGSI
}
