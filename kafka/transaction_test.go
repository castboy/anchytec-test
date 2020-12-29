package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/prometheus/common/log"
	"strconv"
	"testing"
	"time"
)

// 未完结。。。。。。
func Test_TransactionBatchProduce(t *testing.T) {
	broker := sarama.NewBroker("localhost:9092")
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0

	err := broker.Open(config)
	if err != nil {
		log.Fatal(err)
	}

	transactionalID := "first_transaction_test"

	crdReq := &sarama.FindCoordinatorRequest{
		Version:         2,
		CoordinatorKey:  transactionalID,
		CoordinatorType: sarama.CoordinatorTransaction,
	}

	crdRes, err := broker.FindCoordinator(crdReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("FindCoordinator res: ", crdRes.Version, crdRes.Err, crdRes.ThrottleTime, *crdRes.ErrMsg)
	fmt.Println("new broker: ", *crdRes.Coordinator)

	//
	initPdRes, err := broker.InitProducerID(&sarama.InitProducerIDRequest{TransactionalID: &transactionalID, TransactionTimeout: time.Millisecond})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("InitProducerID res: ", initPdRes.Err, initPdRes.ProducerEpoch, initPdRes.ProducerID, initPdRes.ThrottleTime)

	//
	ptrTxnReq := &sarama.AddPartitionsToTxnRequest{
		TransactionalID: transactionalID,
		ProducerID:      initPdRes.ProducerID,
		ProducerEpoch:   initPdRes.ProducerEpoch,
		TopicPartitions: map[string][]int32{"batch1": []int32{0}},
	}

	ptrTxnRes, err := broker.AddPartitionsToTxn(ptrTxnReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AddPartitionsToTxn res: ", ptrTxnRes.Errors["batch1"][0].Err)

	pdReq := &sarama.ProduceRequest{
		//TransactionalID: &transactionalID,
		RequiredAcks: sarama.WaitForAll,
		Version:      3, // not change
	}

	const record_length = 10
	const batch_topic = "batch1"

	records := make([]*sarama.Record, record_length)

	for i := 0; i < record_length; i++ {
		records[i] = &sarama.Record{OffsetDelta: int64(i), Value: []byte("value" + strconv.Itoa(i))}
	}

	batch := &sarama.RecordBatch{
		FirstOffset:     0,                 // not change
		LastOffsetDelta: record_length - 1, // not change
		Version:         2,                 // not change
		ProducerID:      initPdRes.ProducerID,
		ProducerEpoch:   initPdRes.ProducerEpoch,
		FirstSequence:   0,
		Records:         records,
	}

	pdReq.AddBatch(batch_topic, 0, batch)

	pdRes, err := broker.Produce(pdReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", pdRes.Blocks["batch1"][0].Err)

	oftTxnReq := &sarama.AddOffsetsToTxnRequest{
		TransactionalID: transactionalID,
		ProducerID:      initPdRes.ProducerID,
		ProducerEpoch:   initPdRes.ProducerEpoch,
		GroupID:         "112",
	}

	oftTxnRes, err := broker.AddOffsetsToTxn(oftTxnReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("AddOffsetsToTxn res: ", oftTxnRes.Err)

	//oftCmtReq := &sarama.TxnOffsetCommitRequest{
	//	TransactionalID: transactionalID,
	//	GroupID:         "111",
	//	ProducerID:      initPdRes.ProducerID,
	//	ProducerEpoch:   initPdRes.ProducerEpoch,
	//	Topics: map[string][]*sarama.PartitionOffsetMetadata{
	//		"batch1": []*sarama.PartitionOffsetMetadata{&sarama.PartitionOffsetMetadata{Partition: 0, Offset: 10}},
	//	},
	//}
	//oftCmtRes, err := broker.TxnOffsetCommit(oftCmtReq)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("TxnOffsetCommit res: ", oftCmtRes.Topics["batch1"][0].Err)

	endTxnReq := &sarama.EndTxnRequest{
		TransactionalID:   transactionalID,
		ProducerID:        initPdRes.ProducerID,
		ProducerEpoch:     initPdRes.ProducerEpoch,
		TransactionResult: false,
	}

	endTxnRes, err := broker.EndTxn(endTxnReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("EndTxn res: ", endTxnRes.Err)
}
