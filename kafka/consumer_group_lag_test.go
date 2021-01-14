package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"testing"
)

func Test_ConsumerGroupLag(t *testing.T) {
	broker := sarama.NewBroker("localhost:9092")
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	err := broker.Open(config)
	if err != nil {
		log.Fatal(err)
	}

	req := &sarama.OffsetFetchRequest{
		ConsumerGroup: "my-group",
		Version:       1,
	}

	req.AddPartition("my-test-2", 0)

	res, err := broker.FetchOffset(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Blocks["my-test-2"][0].Offset)
	//

	req2 := &sarama.OffsetRequest{
		Version: 1,
	}

	req2.AddBlock("my-test-2", 0, -1, 0)

	res2, err := broker.GetAvailableOffsets(req2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res2.GetBlock("my-test-2", 0).Offset)
}
