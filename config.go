package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var config globalConfig

func loadConfig() {

	log.Println("Loading configuration...")

	// load the config
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/hostdb-collector-vrops")
	viper.AddConfigPath(".")

	// load env vars
	viper.SetEnvPrefix("hostdb_collector_vrops")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// read the config file, and handle any errors
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %s", err))
	}

	// unmarshal into our struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(fmt.Errorf("unable to decode into struct, %v", err))
	}

	// debug
	if config.Collector.Debug {
		log.Println(fmt.Sprintf("%v", os.Environ()))
		log.Println(fmt.Sprintf("%v", config))
	}

}
