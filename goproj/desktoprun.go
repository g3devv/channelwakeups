package main

import (
	"fmt"
	"goprojlib/goprojlib"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("There isn't much to see on desktop, it's just the convenience of being able to run the project directly on the computer")
	// goprojlib.RunChannels()
	goprojlib.RunChannelConciseExample(true)
	// goprojlib.RunChannelsMinimizing()

	keepProcessActive()
}

func keepProcessActive() {
	quitChannel := make(chan os.Signal, 1) // https://stackoverflow.com/a/65127173/2161301
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	//time for cleanup before exit
	fmt.Println("Bye!")
}
