package main

import (
	"encoding/json"
	"family/functions"
	"fmt"
	s "strings"
	"time"

	"github.com/segmentio/kafka-go"
)

func generateEvents[Measurement any](function func(float64) int, sensorEnvName string, generateMeasurement func(string, int, time.Time, string) *Measurement, conn *kafka.Conn) {
	var event *Measurement
	events := make([]Measurement, 0)

	numbersensors := functions.Getenv(sensorEnvName, "kamera")
	fmt.Println(s.Split(numbersensors, ","))
	for _, sensor := range s.Split(numbersensors, ",") {
		currentValue := function(functions.GetTimeToHours())
		event = generateMeasurement(sensor, currentValue, time.Now(), "person")
		if event != nil {
			fmt.Println("Besucher:", currentValue)
			//connect the score of the different sensors
			events = append(events, *event)
		}

	}

	jsonData, err := json.Marshal(&events)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	functions.WriteMessage(jsonData, conn)
	fmt.Println(string(jsonData))
}
