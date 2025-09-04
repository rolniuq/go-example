package main

import "fmt"

// ex: subcribers and publisher
// when publisher publishs something, all subcribers can see
// or call pub/sub

type Subscriber struct {
	name string
}

func (s *Subscriber) sendNotification(msg string) {
	fmt.Printf("you are: %s and message received is %s\n", s.name, msg)
}

type Channel struct {
	subcribers []Subscriber
}

func newChannel() *Channel {
	return &Channel{}
}

func (c *Channel) subcribe(sub Subscriber) {
	c.subcribers = append(c.subcribers, sub)
}

func (c *Channel) notify(msg string) {
	for _, sub := range c.subcribers {
		sub.sendNotification(msg)
	}
}

func main() {
	sub1 := Subscriber{name: "sub1"}
	sub2 := Subscriber{name: "sub2"}
	sub3 := Subscriber{name: "sub3"}

	channel := newChannel()
	channel.subcribe(sub1)
	channel.subcribe(sub2)
	channel.subcribe(sub3)

	channel.notify("Hi everyone")
}
