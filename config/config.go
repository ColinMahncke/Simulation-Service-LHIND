package config

import (
	"family/functions"
	"fmt"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

//loads the enviroment variables

func Load() error {
	err := loadEnv()
	if err != nil {
		return err
	}
	err = setSeed()
	return err

}
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")

		return err
	}
	return nil
}

func setSeed() error {
	envTime := functions.Getenv("time", "")
	if envTime == "" {
		envTime = time.Now().Format(time.RFC3339)
	}
	seedtime, err := time.Parse(time.RFC3339, envTime) //????
	if err != nil {
		fmt.Println("fail to load time", err)
		return err
	}
	rand.Seed(seedtime.UTC().UnixNano())
	return nil

}
