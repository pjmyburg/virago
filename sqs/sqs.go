package sqs

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pjmyburg/virago/config"
	"github.com/pjmyburg/virago/sqs/queue"
	"strings"
)

// SQS represents an instance of the SQS mock
type SQS struct {
	queues map[string]queueWrapper
}

type message struct {
	id           string
	body         string
	attribute    map[string]string
	sysAttribute map[string]string
	delaySeconds int
	deDupID      string
	groupID      string
}

// New creates a new SQS instance and creates queues using the supplied config on startup
func New(conf *config.Config) *SQS {
	queues := make(map[string]queueWrapper)
	for _, confQ := range conf.Queues {
		qw := queueWrapper{
			url:   fmt.Sprintf("https://%s.queue.amazonaws.com/%s/%s", config.Region, config.AccountID, confQ.Name),
			queue: queue.New(),
			in:    make(chan []message),
			rcv:   make(chan ReceiveRequest),
		}
		qw.Run()
		queues[confQ.Name] = qw
	}

	return &SQS{
		queues: queues,
	}
}

func (s *SQS) GetQueueURL(name string) (string, error) {
	q, ok := s.queues[name]
	if !ok {
		return "", fmt.Errorf("queue not found")
	}

	return q.url, nil
}

func (s *SQS) ListQueues(prefix string) []string {
	var queues []string

	for k, v := range s.queues {
		if strings.HasPrefix(k, prefix) {
			queues = append(queues, v.url)
		}
	}

	return queues
}

func (s *SQS) SendMessage(queueURL string, body string, groupID string) (string, error) {
	name := queueName(queueURL)
	q, ok := s.queues[name]
	if !ok {
		return "", fmt.Errorf("queue not found")
	}

	msg := message{
		id:      uuid.New().String(),
		body:    body,
		groupID: groupID,
	}
	q.in <- []message{msg}

	return msg.id, nil
}

func queueName(queueURL string) string {
	segments := strings.Split(queueURL, "/")
	return segments[len(segments)-1]
}
