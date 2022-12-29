package main

import (
	"context"
	"encoding/json"
	"log"
	"net"

	pb "cmdb-receiver/proto"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

var useProtoMarshal bool = false

type server struct {
	pb.UnimplementedReceiverServer
}

func sendKafkaMessage(message []byte) {
	// 连接kafka
	producer, err := sarama.NewSyncProducer([]string{"10.0.0.5:9092"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	msg := &sarama.ProducerMessage{Topic: "cmdb-receiver", Value: sarama.ByteEncoder(message)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}

func (s *server) Post(ctx context.Context, req *pb.Data) (*pb.Reply, error) {
	reply := &pb.Reply{
		Code: 0,
	}
	// 转发消息到kafka
	var bytes []byte
	var err error
	if useProtoMarshal {
		bytes, err = proto.Marshal(req)
	} else {
		bytes, err = json.Marshal(req)
	}
	if err != nil {
		log.Println(err)
		reply.Code = 1
	}
	sendKafkaMessage(bytes)
	// 返回回复
	return reply, nil
}

func main() {
	// 启动rpc
	rpcs := grpc.NewServer()
	pb.RegisterReceiverServer(rpcs, &server{})
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	if err = rpcs.Serve(listen); err != nil {
		panic(err)
	}
}
