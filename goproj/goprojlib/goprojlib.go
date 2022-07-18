package goprojlib

import (
	"fmt"
	"time"
)

var ch chan (string)

/*
   I was able to break down this issue to the simplest common demoninator.
   Basically all we are doing here is to write to (and read from) a channel every 4ms. This is enough to
   trigger the wakeups_resource handling in iOS on an iPhone SE (2016) with iOS 14.3.

   My tests triggered this behavior with `multi` set to false, so one single channel seems to be enough to cause that.

*/

func RunChannelConciseExample(multi bool) {

	ch = make(chan string)
	go launchChannelReads()
	go launchChannelWrites()

	if multi { // not needed, `wakeup_ressource` also triggers when only one channel is in use.
		ch2 = make(chan string)
		go launchChannelReads2()
		go launchChannelWrites2()
	}

}

func launchChannelReads() {
	for elem := range ch {
		fmt.Println(elem)
	}
}

func launchChannelWrites() {

	for i := 0; i != 30000000; i++ {
		ch <- "one123"
		time.Sleep(time.Millisecond * 4)
	}

}

var ch2 chan (string)

func launchChannelReads2() {
	for elem := range ch2 {
		fmt.Println(elem)
	}
}

func launchChannelWrites2() {
	for i := 0; i != 30000000; i++ {
		ch2 <- "one1234"
		time.Sleep(time.Millisecond * 4)
	}
}

/*

	WIP of "optimizing" RunChannels()
	Changes:
	- fixed memory leak (issue still persists)

*/

func RunChannelsMinimizing() {

	for i := 0; i != 3000; i++ {
		if i%100 == 0 {
			fmt.Println("Reached iteration ", i)
		}

		go blockUnblockMinimizing(i, 1)
		go blockUnblockMinimizing(i, 2)
		blockUnblockMinimizing(i, 3)
	}

}

func blockUnblockMinimizing(mainID int, idd int) {

	for i := 0; i != 3; i++ {

		msg := make(chan string, 1)
		for j := 0; j != 30; j++ {
			msg <- fmt.Sprintf("Channel write mainID:%d  callID:%d i:%d j:%d", mainID, idd, i, j)

			for k := 0; k != 3; k++ {
				go func() {
					fmt.Println(<-msg)
				}()
			}

			/*
				The sleep below is to prevent the app to exceed the CPU time limit in iOS.

				We don't want to fry the CPU, it's not about putting useless load onto the CPU but
				rather about creating, stopping, and waking up threads. The sleep takes some weight
				off the CPU.
			*/

			time.Sleep(time.Millisecond * 1)
		}

	}
}

/*

I don't know golang very well. Basically all this func does is to create some channels, with and
without subroutines so we can see some context switches occuring in Xcode's profiler.

*/

func RunChannels() {

	for i := 0; i != 3000; i++ {
		if i%100 == 0 {
			fmt.Println("Reached iteration ", i)
		}

		go blockUnblock(i, 1)
		go blockUnblock(i, 2)
		blockUnblock(i, 3)
	}

}

func blockUnblock(mainID int, idd int) {

	for i := 0; i != 3; i++ {

		msg := make(chan string, 1)
		for j := 0; j != 30; j++ {
			msg <- fmt.Sprintf("Channel write mainID:%d  callID:%d i:%d j:%d", mainID, idd, i, j)

			for k := 0; k != 3; k++ {
				go func() {
					fmt.Println(<-msg)
				}()
			}

			/*
				The sleep below is to prevent the app to exceed the CPU time limit in iOS.

				We don't want to fry the CPU, it's not about putting useless load onto the CPU but
				rather about creating, stopping, and waking up threads. The sleep takes some weight
				off the CPU.
			*/

			time.Sleep(time.Millisecond * 1)
		}

	}
}
