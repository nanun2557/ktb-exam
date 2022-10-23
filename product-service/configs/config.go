package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	App   AppConfiguration
	MySql MySqlConfiguration
	Log   LogConfiguration
}

type AppConfiguration struct {
	Env   string
	Port  int
	Debug bool
}

type MySqlConfiguration struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
}

type LogConfiguration struct {
	Level  int
	Format string
}

func LoadConfig() Configuration {
	viper.SetConfigName("config")     // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./configs/") //// optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var configuration Configuration
	err = viper.Unmarshal(&configuration)
	if err != nil {
		panic(fmt.Errorf("unable to decode struct: %w", err))
	}

	// fmt.Println(configuration)
	return configuration
}
