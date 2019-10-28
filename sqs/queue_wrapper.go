package sqs

import "github.com/pjmyburg/virago/sqs/queue"

type queueWrapper struct {
	url   string
	queue *queue.Queue
	in    chan []message
	rcv   chan ReceiveRequest
}

type ReceiveRequest struct {
	Count int
	Reply chan []message
}

func (wq *queueWrapper) Run() {
	go func() {
		for {
			select {
			case msgs := <-wq.in:
				for _, msg := range msgs {
					wq.queue.Enqueue(msg)
				}
			case rcv := <-wq.rcv:
				var msgs []message
				count := rcv.Count
				if count > wq.queue.Depth() {
					count = wq.queue.Depth()
				}

				for i := 0; i < count; i++ {
					msg := wq.queue.Dequeue()
					msgs = append(msgs, msg.(message))
				}

				rcv.Reply <- msgs
			}
		}
	}()
}
