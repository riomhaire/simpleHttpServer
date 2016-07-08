package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"riomhaire/simpleHttpServer"

	"gopkg.in/yaml.v2"
)

func main() {
	// Process flags
	configFile := flag.String("m", "mapping.yml", "configuration containing uri's to files")
	fmt.Println("Using Config File: ", *configFile)

	var config simpleHttpServer.Configuration
	source, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Printf("Cannot read config file because : %v\n", err)
		os.Exit(-1)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		fmt.Printf("Cannot parse config file because : %v\n", err)
		os.Exit(-2)
	}

	if config.Port == 0 {
		config.Port = 8080
	}

	for _, mapping := range config.Mappings {
		fmt.Printf("Value: %v -> %v\n", mapping.URI, mapping.File)
	}

	fmt.Printf("Starting Server On Port: %v\n", config.Port)
	simpleHttpServer.Start(&config)
}
