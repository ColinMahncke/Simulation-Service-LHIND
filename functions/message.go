package functions

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

func WriteMessage(jsonData []byte, conn *kafka.Conn) {
	_, err := conn.WriteMessages(kafka.Message{Value: jsonData})
	if err != nil {
		fmt.Println("failed to send message:", err)
	}
}
