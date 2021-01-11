// Package d2thread is a package graciously taken from https://github.com/faiface/mainthread
package d2thread

import (
	"errors"
)

// CallQueueCap is the capacity of the call queue. This means how many calls to CallNonBlock will not
// block until some call finishes.
//
// The default value is 16 and should be good for 99% usecases.
var (
	callQueue chan func() //nolint:gochecknoglobals // necessary evil for now
)

func checkRun() {
	if callQueue == nil {
		panic(errors.New("mainthread: did not call Run"))
	}
}

// Run enables mainthread package functionality. To use mainthread package, put your main function
// code into the run function (the argument to Run) and simply call Run from the real main function.
//
// Run returns when run (argument) function finishes.
func Run(run func()) {
	var CallQueueCap = 16

	callQueue = make(chan func(), CallQueueCap)

	done := make(chan struct{})

	go func() {
		run()
		done <- struct{}{}
	}()

	for {
		select {
		case f := <-callQueue:
			f()
		case <-done:
			return
		}
	}
}

// CallNonBlock queues function f on the main thread and returns immediately. Does not wait until f
// finishes.
func CallNonBlock(f func()) {
	checkRun()
	callQueue <- f
}

// Call queues function f on the main thread and blocks until the function f finishes.
func Call(f func()) {
	checkRun()

	done := make(chan struct{})
	callQueue <- func() {
		f()
		done <- struct{}{}
	}
	<-done
}

// CallErr queues function f on the main thread and returns an error returned by f.
func CallErr(f func() error) error {
	checkRun()

	errChan := make(chan error)
	callQueue <- func() {
		errChan <- f()
	}

	return <-errChan
}

// CallVal queues function f on the main thread and returns a value returned by f.
func CallVal(f func() interface{}) interface{} {
	checkRun()

	respChan := make(chan interface{})
	callQueue <- func() {
		respChan <- f()
	}

	return <-respChan
}
