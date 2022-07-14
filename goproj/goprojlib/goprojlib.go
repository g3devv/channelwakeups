package goprojlib

import (
	"fmt"
	"time"
)

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
