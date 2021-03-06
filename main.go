package main

import (
	"context"
	"family/data"
	"family/functions"
	"fmt"
	"strconv"
	"time"

	"family/config"

	"github.com/segmentio/kafka-go"
)

func main() {

	err := config.Load()
	if err != nil {
		return
	}
	host := functions.Getenv("host", "localhost:9093")

	connarea, err := kafka.DialLeader(context.Background(), "tcp", host, "area_counting_measurement_events", 0)
	defer connarea.Close()
	if err != nil {
		fmt.Println("failed to dial leader:", err)
		return
	}

	connline, err := kafka.DialLeader(context.Background(), "tcp", host, "line_counting_measurement_events", 0)
	defer connline.Close()
	if err != nil {
		fmt.Println("failed to dial leader:", err)
		return
	}
	messagecount := 0
	lastmessagecount := time.Now()
	starttime := time.Now()
	//generate the number of persons on different sensors
	for {

		messagecount += generateEvents(functions.GenerateAreaCount, "area_sensors", data.NewAreaCounting, connarea)

		messagecount += generateEvents(functions.GenerateLineCount, "line_sensors", data.NewLineCounting, connline)
		//buffertime between dataflow
		numberinterval, error := strconv.Atoi(functions.Getenv("interval", "2"))
		if error != nil {
			fmt.Println("error:", error)
			return
		}

		time.Sleep(time.Duration(numberinterval) * time.Millisecond)

		if lastmessagecount.Add(time.Second).Before(time.Now()) {
			fmt.Println("messages per second:", float64(messagecount)/float64(time.Now().Unix()-starttime.Unix()), "total messages:", messagecount)
			lastmessagecount = time.Now()

		}
	}
}
