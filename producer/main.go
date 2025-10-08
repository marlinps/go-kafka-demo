package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"host.docker.internal:9092"}, // ✅ ganti ke 9092
		Topic:    "demo-topic",
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("Pesan ke-%d dari Golang Producer", i)
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("Key-%d", i)),
				Value: []byte(msg),
			},
		)
		if err != nil {
			log.Println("❌ Gagal kirim pesan:", err)
		} else {
			fmt.Println("✅ Terkirim:", msg)
		}
		time.Sleep(time.Second)
	}
}
