//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, stream_chan chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(stream_chan)
			return
		}
		stream_chan <- tweet
	}
}

func consumer(stream_chan chan *Tweet, done chan struct{}) {
	for t := range stream_chan {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	done <- struct{}{}
}

func main() {
	start := time.Now()
	stream_chan := make(chan *Tweet)
	done_chan := make(chan struct{})
	stream := GetMockStream()

	// Producer
	go producer(stream, stream_chan)

	// Consumer
	go consumer(stream_chan, done_chan)
	<-done_chan
	fmt.Printf("Process took %s\n", time.Since(start))
}
