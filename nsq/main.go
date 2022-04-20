package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

// ConsumerHandler 消费者处理者
type ConsumerHandler struct{}

// HandleMessage 处理消息
func (*ConsumerHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println(string(msg.Body), string(msg.NSQDAddress))
	return nil
}

// Producer 生产者
func Producer() {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		fmt.Println("NewProducer", err)
		panic(err)
	}

	i := 1
	for {
		if err := producer.Publish("test", []byte(fmt.Sprintf("Hello World %d", i))); err != nil {
			fmt.Println("Publish", err)
			panic(err)
		}

		time.Sleep(time.Second * 5)

		i++
	}
}

// ConsumerA 消费者
func ConsumerA() {
	consumer, err := nsq.NewConsumer("test", "chan_4151_2", nsq.NewConfig())
	if err != nil {
		fmt.Println("NewConsumer", err)
		panic(err)
	}

	consumer.AddHandler(&ConsumerHandler{})

	if err := consumer.ConnectToNSQLookupd("127.0.0.1:4161"); err != nil {
		fmt.Println("ConnectToNSQLookupd", err)
		panic(err)
	}
}

// ConsumerB 消费者
func ConsumerB() {
	consumer, err := nsq.NewConsumer("test", "test-channel-b", nsq.NewConfig())
	if err != nil {
		fmt.Println("NewConsumer", err)
		panic(err)
	}

	consumer.AddHandler(&ConsumerHandler{})

	if err := consumer.ConnectToNSQLookupd("127.0.0.1:4161"); err != nil {
		fmt.Println("ConnectToNSQLookupd", err)
		panic(err)
	}
}

func main() {
	ConsumerA()
	ConsumerB()
	Producer()
}
