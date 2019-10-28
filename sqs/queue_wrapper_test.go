package sqs

import (
	"fmt"
	"github.com/pjmyburg/virago/sqs/queue"
	"testing"
)

func TestQueueWrapper(t *testing.T) {
	in := make(chan []message)
	rcv := make(chan ReceiveRequest)
	qw := queueWrapper{
		url:   "queueURL",
		queue: queue.New(),
		in:    in,
		rcv:   rcv,
	}
	qw.Run()

	// Sending to the wrapper's in channel shouldn't block
	sendCount := 5
	msgsPerSend := 2
	for i := 0; i < sendCount; i++ {
		var msgs []message
		for j := 0; j < msgsPerSend; j++ {
			msg := message{body: fmt.Sprintf("%d:%d", i, j)}
			msgs = append(msgs, msg)
		}
		in <- msgs
	}

	requestCount := 6
	reply := make(chan []message, requestCount)
	request := ReceiveRequest{
		Count: requestCount,
		Reply: reply,
	}

	outstanding := sendCount * msgsPerSend
	for outstanding > 0 {
		rcv <- request
		msgs := <-reply
		if outstanding >= requestCount {
			if len(msgs) != requestCount {
				t.Fatalf("Expected %d messages to be received, but got %d", requestCount, len(msgs))
			}
		} else {
			if len(msgs) != outstanding {
				t.Fatalf("Expected %d messages to be received, but got %d", outstanding, len(msgs))
			}
		}
		outstanding = outstanding - requestCount
	}

}
