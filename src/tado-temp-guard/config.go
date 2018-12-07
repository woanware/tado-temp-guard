package main

import (
	"fmt"
	"os"
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

	config = new(Config)

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v \n", err)
		os.Exit(-1)
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		loadConfiguration()
	})
}

//
func loadConfiguration() {

	config.UserName = viper.GetString("user_name")
	config.Password = viper.GetString("password")
	config.IntervalMinutes = viper.GetInt("interval_minutes")

	tempMax, err := strconv.ParseFloat(viper.GetString("max_temperature"), 32)
	if err != nil {
		fmt.Printf("Invalid max temperature: %v (defaulting to 21)", viper.GetString("max_temperature"))
		os.Exit(-1)
	}
	config.MaxTemperatureCelsius = tempMax
}
