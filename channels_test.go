package tricks

import (
	"testing"
	"time"
	"fmt"
)

// Cannot send or receive from nil channel but it does not panic and blocks forever
// https://golang.org/ref/spec#Send_statements `A send on a nil channel blocks forever.`
// https://golang.org/ref/spec#Receive_operator `Receiving from a nil channel blocks forever.`
func TestNilChannel(t *testing.T) {
	var ch chan int

	go func() {
		fmt.Println(<-ch)
		t.Error("received from channel")
	}()
	go func() {
		ch <- 1
		t.Error("sent to channel")
	}()

	time.Sleep(10 * time.Millisecond)
}

// Close nil channel panics
// https://golang.org/ref/spec#Close `Closing the nil channel also causes a run-time panic.`
func TestCloseNilChannel(t *testing.T) {
	var ch chan int

	shouldPanic(t, func() { close(ch) }, "close of nil channel")
}

// Receive on closed channel always returns empty value and
// https://golang.org/ref/spec#Receive_operator `A receive operation on a closed channel can always proceed immediately, yielding the element type's zero value after any previously sent values have been received.`
func TestReceiveClosedChan(t *testing.T) {
	ch := make(chan int)

	close(ch)

	for i := 0; i < 5; i++ {
		i, ok := <-ch
		if i != 0 || ok {
			t.Errorf("unexpected response i = %d, ok = %t", i, ok)
		}
	}
}

// Send on closed channel panics
// https://golang.org/ref/spec#Send_statements `A send on a closed channel proceeds by causing a run-time panic.`
func TestSendClosedChan(t *testing.T) {
	ch := make(chan int)
	close(ch)

	shouldPanic(t, func() { ch <- 1 }, "send on closed channel")
}

// Close on closed channel panics
// https://golang.org/ref/spec#Close `Sending to or closing a closed channel causes a run-time panic.`
func TestCloseClosedChan(t *testing.T) {
	ch := make(chan int)
	close(ch)

	shouldPanic(t, func() { close(ch) }, "close of closed channel")
}
