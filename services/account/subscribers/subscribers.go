package subscribers

import (
	"github.com/lileio/pubsub"
)

type AccountServiceSubscriber struct{}

func (s *AccountServiceSubscriber) Setup(c *pubsub.Client) {
	// https://godoc.org/github.com/lileio/pubsub#Client.On
	// c.On(pubsub.HandlerOptions{
	// 	Topic:    "some_topic",
	// 	Name:     "service_name",
	// 	Handler:  s.SomeMethod,
	// 	Deadline: 10 * time.Second,
	// 	Concurrency: 10,
	// 	AutoAck:  true,
	// })
}

// func (s *AccountServiceSubscriber) SomeMethod(ctx context.Context, req *proto.Message, _ *pubsub.Msg) error {
// 	return nil
// }
