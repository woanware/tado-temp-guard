package main

import (
	"fmt"
	"log"
	"strconv"

	fsnotify "github.com/fsnotify/fsnotify"
	viper "github.com/spf13/viper"
)

// ##### Structs ##############################################################

// Config holds configuration data for the application
type Config struct {
	UserName              string
	Password              string
	MaxTemperatureCelsius float64
	IntervalMinutes       int
}

// ##### Methods ##############################################################

// LoadConfig loads the configuration data from the "bgpm" config file
func initialiseConfiguration() {

	configReader = viper.New()
	configReader.SetConfigName("tado-temp-guard")
	configReader.AddConfigPath(".")
	err := configReader.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s \n", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		reloadConfig()
	})
}

//
func parseConfiguration() *Config {

	config := new(Config)

	config.UserName = configReader.GetString("user_name")
	config.Password = configReader.GetString("password")
	config.IntervalMinutes = configReader.GetInt("interval_minutes")

	tempMax, err := strconv.ParseFloat(configReader.GetString("max_temperature"), 32)
	if err != nil {
		fmt.Printf("Invalid max temperature: %v", configReader.GetString("max_temperature"))
		return nil
	}
	config.MaxTemperatureCelsius = tempMax

	return config
}
