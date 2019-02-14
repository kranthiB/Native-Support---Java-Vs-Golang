package main

import (
	"fmt"
	"log"

	"posx_mq"
)

const maxTickNum = 10

func main() {
	oflag := posix_mq.O_RDONLY
	mq, err := posix_mq.NewMessageQueue("/test_queue", oflag, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer mq.Close()

	fmt.Println("Start receiving messages")

	count := 0
	for {
		count++

		msg, _, err := mq.Receive()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(string(msg))

		if count >= maxTickNum {
			break
		}
	}
}
