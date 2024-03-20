package base

import "flag"

const (
	Kafka_Cataory_Reader  = "reader"
	Kafka_Category_Writer = "writer"
	Kafka_Category_Booth  = "both"
)

var (
	avgKafkaDevice = flag.String("kafka_device", "", "kafka_device")
)
