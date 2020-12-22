package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"testing"
	"time"
)

func Test_BatchConsume(t *testing.T) {
	broker := sarama.NewBroker("localhost:9092")
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	err := broker.Open(config)
	if err != nil {
		log.Fatal(err)
	}

	fetReq := sarama.FetchRequest{
		MaxWaitTime: 10e8,
		MinBytes:    0,
		MaxBytes:    0,
		Version:     3,
	}

	fetReq.AddBlock("batch1", 0, 198184, 0)

	res, err := broker.Fetch(&fetReq)
	if err != nil {
		fmt.Println(err)
	}

	if res == nil {
		fmt.Println("res is nil")
	}

	if res.Blocks["batch1"] == nil {
		time.Sleep(time.Second)
		fmt.Println("batch1 nil")
	}

	rdsSet := res.Blocks["batch1"][0].RecordsSet

	fmt.Println(res.Blocks["batch1"][0].HighWaterMarkOffset, len(rdsSet[0].MsgSet.Messages))

	for _, v := range rdsSet[0].MsgSet.Messages {
		fmt.Println(string(v.Msg.Key), string(v.Msg.Value))
	}
}
