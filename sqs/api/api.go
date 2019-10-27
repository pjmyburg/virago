package api

import (
	"encoding/xml"
	"github.com/pjmyburg/virago/sqs"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// API implements the functions on the AWS SQS API
type API struct {
	sqs *sqs.SQS
}

// NewAPI creates a new instance of the SQS API using the supplied SQS instance
func NewAPI(sqs *sqs.SQS) *API {
	return &API{sqs}
}

// AddPermission adds a permission to a queue for a specific principal
func (s *API) AddPermission(w http.ResponseWriter, req *http.Request) {
	log.Debug("AddPermission")
	w.WriteHeader(http.StatusNotImplemented)
}

// ChangeMessageVisibility changes the visibility timeout of a specified message in a queue to a new value
func (s *API) ChangeMessageVisibility(w http.ResponseWriter, req *http.Request) {
	log.Debug("ChangeMessageVisibility")
	w.WriteHeader(http.StatusNotImplemented)
}

// ChangeMessageVisibilityBatch changes the visibility timeout of multiple messages
func (s *API) ChangeMessageVisibilityBatch(w http.ResponseWriter, req *http.Request) {
	log.Debug("ChangeMessageVisibilityBatch")
	w.WriteHeader(http.StatusNotImplemented)
}

// CreateQueue creates a new standard or FIFO queue
func (s *API) CreateQueue(w http.ResponseWriter, req *http.Request) {
	log.Debug("CreateQueue")
	w.WriteHeader(http.StatusNotImplemented)
}

// DeleteMessage deletes the specified message from the specified queue
func (s *API) DeleteMessage(w http.ResponseWriter, req *http.Request) {
	log.Debug("DeleteMessage")
	w.WriteHeader(http.StatusNotImplemented)
}

// DeleteMessageBatch deletes up to ten messages from the specified queue
func (s *API) DeleteMessageBatch(w http.ResponseWriter, req *http.Request) {
	log.Debug("DeleteMessageBatch")
	w.WriteHeader(http.StatusNotImplemented)
}

// DeleteQueue deletes the queue specified by the QueueUrl, regardless of the queue's contents
func (s *API) DeleteQueue(w http.ResponseWriter, req *http.Request) {
	log.Debug("DeleteQueue")
	w.WriteHeader(http.StatusNotImplemented)
}

// GetQueueAttributes gets attributes for the specified queue
func (s *API) GetQueueAttributes(w http.ResponseWriter, req *http.Request) {
	log.Debug("GetQueueAttributes")
	w.WriteHeader(http.StatusNotImplemented)
}

// GetQueueURL returns the URL of an existing Amazon SQS queue
func (s *API) GetQueueURL(w http.ResponseWriter, req *http.Request) {
	log.Debug("GetQueueURL")

	queueName := req.FormValue("QueueName")
	url, err := s.sqs.GetQueueURL(queueName)
	if err != nil {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusBadRequest)
		err := ErrorResponse{
			Error: ErrorResult{
				Type:    "Not Found",
				Code:    "AWS.SimpleQueueService.NonExistentQueue",
				Message: "The specified queue does not exist for this wsdl version.",
			},
			RequestID: "00000000-0000-0000-0000-000000000000",
		}
		enc := xml.NewEncoder(w)
		enc.Indent("  ", "    ")
		if err := enc.Encode(err); err != nil {
			log.Errorf("error: %s", err)
		}
		return
	}

	response := GetQueueURLResponse{
		Result:   GetQueueURLResult{url},
		MetaData: ResponseMetaData{"00000000-0000-0000-0000-000000000000"},
	}

	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")
	if err := enc.Encode(response); err != nil {
		log.Errorf("error: %s", err)
	}

}

// ListDeadLetterSourceQueues returns a list of your queues that have the RedrivePolicy queue attribute configured with a dead-letter queue
func (s *API) ListDeadLetterSourceQueues(w http.ResponseWriter, req *http.Request) {
	log.Debug("ListDeadLetterSourceQueues")
	w.WriteHeader(http.StatusNotImplemented)
}

// ListQueues returns a list of your queues
func (s *API) ListQueues(w http.ResponseWriter, req *http.Request) {
	log.Debug("ListQueues")

	queueNamePrefix := req.FormValue("QueueNamePrefix")
	queues := s.sqs.ListQueues(queueNamePrefix)

	response := ListQueuesResponse{
		Result:   ListQueuesResult{queues},
		MetaData: ResponseMetaData{"00000000-0000-0000-0000-000000000000"},
	}

	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")
	if err := enc.Encode(response); err != nil {
		log.Errorf("error: %s", err)
	}
}

// ListQueueTags list all cost allocation tags added to the specified Amazon SQS queue
func (s *API) ListQueueTags(w http.ResponseWriter, req *http.Request) {
	log.Debug("ListQueueTags")
	w.WriteHeader(http.StatusNotImplemented)
}

// PurgeQueue deletes the messages in a queue specified by the QueueURL parameter
func (s *API) PurgeQueue(w http.ResponseWriter, req *http.Request) {
	log.Debug("PurgeQueue")
	w.WriteHeader(http.StatusNotImplemented)
}

// ReceiveMessage retrieves one or more messages (up to 10), from the specified queue
func (s *API) ReceiveMessage(w http.ResponseWriter, req *http.Request) {
	log.Debug("ReceiveMessage")
	w.WriteHeader(http.StatusNotImplemented)
}

// RemovePermission revokes any permissions in the queue policy that matches the specified Label parameter
func (s *API) RemovePermission(w http.ResponseWriter, req *http.Request) {
	log.Debug("RemovePermission")
	w.WriteHeader(http.StatusNotImplemented)
}

// SendMessage delivers a message to the specified queue
func (s *API) SendMessage(w http.ResponseWriter, req *http.Request) {
	log.Debug("SendMessage")
	w.WriteHeader(http.StatusNotImplemented)
}

// SendMessageBatch delivers up to ten messages to the specified queue
func (s *API) SendMessageBatch(w http.ResponseWriter, req *http.Request) {
	log.Debug("SendMessageBatch")
	w.WriteHeader(http.StatusNotImplemented)
}

// SetQueueAttributes sets the value of one or more queue attributes
func (s *API) SetQueueAttributes(w http.ResponseWriter, req *http.Request) {
	log.Debug("SetQueueAttributes")
	w.WriteHeader(http.StatusNotImplemented)
}

// TagQueue add cost allocation tags to the specified Amazon SQS queue
func (s *API) TagQueue(w http.ResponseWriter, req *http.Request) {
	log.Debug("TagQueue")
	w.WriteHeader(http.StatusNotImplemented)
}

// UntagQueue remove cost allocation tags from the specified Amazon SQS queue
func (s *API) UntagQueue(w http.ResponseWriter, req *http.Request) {
	log.Debug("UntagQueue")
	w.WriteHeader(http.StatusNotImplemented)
}
