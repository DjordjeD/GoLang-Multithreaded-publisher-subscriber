package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type PubSub struct {
	mu     sync.Mutex               // mutual ex
	topics map[string][]chan string // map with string keys used for array of string channels
}

var wg sync.WaitGroup // waiting for the group

// TODO: creates and returns a new channel on a given topic, updating the PubSub struct
func (ps *PubSub) subscribe(topic string) chan string {
	ps.mu.Lock()
	newChan := make(chan string)
	ps.topics[topic] = append(ps.topics[topic], newChan)
	ps.mu.Unlock()
	return newChan

}

// TODO: writes the given message on all the channels associated with the given topic
func (ps PubSub) publish(topic string, msg string) {
	ps.mu.Lock()

	// temp := ps.topics[topic]
	// temp[0] <- msg
	for topicIter, message := range ps.topics {
		if topicIter == topic {
			i := 0
			for range message {
				//fmt.Println("upisa sam na channel ")
				message[i] <- msg
				i++
			}
		}
	}
	ps.mu.Unlock()
}

// TODO: sends messages taken from a given array of message, one at a time and at random intervals,
//to all topic subscribers
func publisher(ps PubSub, topic string, msgs []string) {

	//ps.publish(topic, msgs[0])
	//time.Sleep(time.Second * 2)
	//ps.mu.Lock()
	for _, msg := range msgs {
		num := rand.Intn(2)
		//go subscriber(ps, "John", topic)
		ps.publish(topic, msg)
		//fmt.Println("publishovao sam ovo ")
		time.Sleep(time.Second * time.Duration(num))
	}
	//ps.mu.Unlock()
	time.Sleep(time.Second * 2)
	wg.Done()
}

// TODO: reads and displays all messages received from a particular topic
func subscriber(ps PubSub, name string, topic string) {
	// x := ps.topics[topic]
	// fmt.Println(<-x[0])
	//ps.mu.Lock()
	currChan := ps.subscribe(topic)
	//ps.mu.Unlock()
	for {

		channels := ps.topics[topic]
		//i := 0
		for _, tempChan := range channels {
			if tempChan == currChan {
				fmt.Println(name + " received message: " + <-currChan)
			}
		}

	}

}

func main() {

	// TODO: create the ps struct
	ps := PubSub{mu: sync.Mutex{}, topics: make(map[string][]chan string)}

	// TODO: create the arrays of messages to be sent on each topic
	bees := []string{"bees are pollinators",
		"bees produce honey",
		"all worker bees are female",
		"bees have 5 eyes",
		"bees fly about 20mph",
	}
	colorado := []string{"Colfax in Denver is the longest continuous street in America,",
		"Colorado has the highest mean altitude of all the states",
		"Colorado state flower is the columbine",
	}

	go subscriber(ps, "john", "bees")
	go subscriber(ps, "mary", "bees")
	go subscriber(ps, "mary", "colorado")

	time.Sleep(time.Second * 3)
	//go publisher(ps, "colorado", colorado)

	wg.Add(2)
	go publisher(ps, "bees", bees)
	go publisher(ps, "colorado", colorado)

	wg.Wait()
	//fmt.Scan()
}
