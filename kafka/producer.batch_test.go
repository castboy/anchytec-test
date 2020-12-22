package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/prometheus/common/log"
	"strconv"
	"testing"
	"time"
)

// ProduceID、producer epoch、sequence 等信息都是为了实现幂等producer之用
//
// 每次添加新的消息batch时，都会检查它的epoch 和 sequence，校检步骤如下：
// 1)检查消息batch的produce epoch,如果该消息barch的epoch比上条消息小，则会报错。
// 2)检查消息batch的sequence，分为两种情况:
//    1. 如果消息batch的epoch比上条消息大，它的序列号必须从0开始。否则就会出错。
//    2. 如果消息batch的epoch相等，它的起始序列号和上条消息batch的结束序列号，必须是连续递增的，否则就会出错。

func Test_BatchProduce(t *testing.T) {
	broker := sarama.NewBroker("localhost:9092")
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0

	err := broker.Open(config)
	if err != nil {
		log.Fatal(err)
	}

	pdReq := &sarama.ProduceRequest{
		RequiredAcks: sarama.WaitForAll,
		Version:      3, // not change
	}

	const record_length = 10
	const batch_topic = "batch1"

	produceID := time.Now().UnixNano()

	records := make([]*sarama.Record, record_length)

	for i := 0; i < record_length; i++ {
		records[i] = &sarama.Record{OffsetDelta: int64(i), Value: []byte("value" + strconv.Itoa(i))}
	}

	batch := &sarama.RecordBatch{
		FirstOffset:     0,                 // not change
		LastOffsetDelta: record_length - 1, // not change
		Version:         2,                 // not change
		ProducerID:      produceID,
		ProducerEpoch:   0,
		FirstSequence:   0,
		Records:         records,
	}

	pdReq.AddBatch(batch_topic, 0, batch)

	res, err := broker.Produce(pdReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", res.Blocks["batch1"][0])

	// if producerID and ProducerEpoch are both same with last run, FirstSequence must be continuous with last sequence.
	batch.FirstSequence = record_length

	pdReq.AddBatch(batch_topic, 0, batch)

	res, err = broker.Produce(pdReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", res.Blocks["batch1"][0])
}
