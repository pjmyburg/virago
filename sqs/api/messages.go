package api

import "encoding/xml"

// ResponseMetaData is used in all AWS response messages
type ResponseMetaData struct {
	RequestID string `xml:"RequestId"`
}

// GetQueueURLResult is part of the GetQueueUrlResponse message
type GetQueueURLResult struct {
	QueueURL string `xml:"QueueUrl"`
}

// GetQueueURLResponse represents the corresponding AWS XML structure
type GetQueueURLResponse struct {
	XMLName  xml.Name          `xml:"GetQueueUrlResponse"`
	Result   GetQueueURLResult `xml:"GetQueueUrlResult"`
	MetaData ResponseMetaData  `xml:"ResponseMetadata"`
}

// ListQueuesResult is part of the ListQueuesResponse message
type ListQueuesResult struct {
	QueueURL []string `xml:"QueueUrl"`
}

// ListQueuesResponse represents the corresponding AWS XML structure
type ListQueuesResponse struct {
	Result   ListQueuesResult `xml:"ListQueuesResult"`
	MetaData ResponseMetaData `xml:"ResponseMetadata"`
}

type SendMessageResult struct {
	MD5OfMessageAttributes string `xml:"MD5OfMessageAttributes"`
	MD5OfMessageBody       string `xml:"MD5OfMessageBody"`
	MessageID              string `xml:"MessageId"`
	SequenceNumber         string `xml:"SequenceNumber"`
}

type SendMessageResponse struct {
	Result   SendMessageResult `xml:"SendMessageResult"`
	Metadata ResponseMetaData  `xml:"ResponseMetadata"`
}

// ErrorResult is part of the ErrorResponse message
type ErrorResult struct {
	Type    string `xml:"Type,omitempty"`
	Code    string `xml:"Code,omitempty"`
	Message string `xml:"Message,omitempty"`
}

// ErrorResponse represents the AWS error XML structure
type ErrorResponse struct {
	Error     ErrorResult `xml:"Error"`
	RequestID string      `xml:"RequestId"`
}
