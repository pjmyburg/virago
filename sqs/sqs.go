package sqs

import (
	"fmt"
	"github.com/pjmyburg/virago/config"
)

// SQS represents an instance of the SQS mock
type SQS struct {
	queues map[string]queue
}

type queue struct {
	url      string
	in       chan message
	inFlight chan message
}

type message struct {
	body string
}

// New creates a new SQS instance and creates queues using the supplied config on startup
func New(conf *config.Config) *SQS {
	queues := make(map[string]queue)
	for _, confQ := range conf.Queues {
		queues[confQ.Name] = queue{
			url:      fmt.Sprintf("https://%s.queue.amazonaws.com/%s/%s", config.Region, config.AccountID, confQ.Name),
			in:       nil,
			inFlight: nil,
		}
	}

	return &SQS{
		queues: queues,
	}
}
