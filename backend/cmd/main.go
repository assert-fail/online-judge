package main

import (
	"backend/config"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("❌ Unable to read config file: ", err)
	}
}
