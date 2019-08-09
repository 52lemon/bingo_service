package servers

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

// 默认配置
const NSQ_OPEN_STATUS = 0
const NSQ_HOST = "127.0.0.1:4150"
const NSQ_TOPIC_NAME = "test"
const NSQ_CHANNEL_NAME = "test-channel"

// 启动Nsq
func NsqRun() {
	if NSQ_OPEN_STATUS == 1 {
		Consumer()
	}
}

// nsq发布消息
func Producer(msgBody string) {

	// 新建生产者
	p, err := nsq.NewProducer(NSQ_HOST, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	// 发布消息
	if err := p.Publish(NSQ_TOPIC_NAME, []byte(msgBody)); err != nil {
		panic(err)
	}
}

// nsq订阅消息
type ConsumerT struct{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println(string(msg.Body))
	return nil
}

func Consumer() {

	// 新建一个消费者
	c, err := nsq.NewConsumer(NSQ_TOPIC_NAME, NSQ_CHANNEL_NAME, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	// 添加消息处理
	c.AddHandler(&ConsumerT{})

	// 建立连接
	if err := c.ConnectToNSQD(NSQ_HOST); err != nil {
		panic(err)
	}
}
