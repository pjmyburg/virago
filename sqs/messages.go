package sqs

// ResponseMetaData is used in all AWS response messages
type ResponseMetaData struct {
	RequestID string `xml:"RequestId"`
}

// ListQueuesResult is part of the ListQueuesResponse
type ListQueuesResult struct {
	QueueURL []string `xml:"QueueUrl"`
}

// ListQueuesResponse represents the corresponding AWS XML structure
type ListQueuesResponse struct {
	Result   ListQueuesResult `xml:"ListQueuesResult"`
	MetaData ResponseMetaData `xml:"ResponseMetadata"`
}
