package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type TomlConfig struct {
	AppName        string
	MySQL          MySQLConfig
	Log            LogConfig
	StaticPath     PathConfig
	MsgChannelType MsgChannelType
}

// mysql related configuration
type MySQLConfig struct {
	Host        string
	Name        string
	Password    string
	Port        int
	TablePrefix string
	User        string
}

// Log saving address
type LogConfig struct {
	Path  string
	Level string
}

// Related address information, such as static file address
type PathConfig struct {
	FilePath string
}

// Message queue type and information queue related information
// GoChannel uses Go's default Channel to make messages for a stand -alone machine
// kafka is used as a message queue with Kafka, which can be distributed to expand message chat programs
type MsgChannelType struct {
	ChannelType string
	KafkaHosts  string
	KafkaTopic  string
}

var c TomlConfig

func init() {
	// Set the file name
	viper.SetConfigName("config")
	// Set the file type
	viper.SetConfigType("toml")
	// Set the file path, you can find multiple VIPERs in order in the order of settings
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	viper.Unmarshal(&c)
}
func GetConfig() TomlConfig {
	return c
}
