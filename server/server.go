package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pjmyburg/virago/sqs/api"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
)

// New creates a new instance of the Virago HTTP server
func New(sqsAPI *api.API) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", handler(sqsAPI)).Methods("GET", "POST")
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/{account}", handler(sqsAPI)).Methods("GET", "POST")
	r.HandleFunc("/queue/{queueName}", handler(sqsAPI)).Methods("GET", "POST")
	r.HandleFunc("/{account}/{queueName}", handler(sqsAPI)).Methods("GET", "POST")
	return r
}

func health(w http.ResponseWriter, req *http.Request) {
	log.Debug("health")
	w.WriteHeader(200)
	if _, err := fmt.Fprint(w, "OK"); err != nil {
		log.Errorf("Error responding to HTTP request: %s", err)
	}
}

func handler(sqsAPI *api.API) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if log.IsLevelEnabled(log.DebugLevel) {
			requestDump, err := httputil.DumpRequest(req, true)
			if err != nil {
				log.Debug("Failed to dump request: ", err)
			}
			log.Debug(string(requestDump))
		}
		action := req.FormValue("Action")

		fn, err := routeAction(action, sqsAPI)
		if err != nil {
			log.Errorf("Error in handler: %s", err)
			w.WriteHeader(400)
			if _, err := fmt.Fprint(w, "Bad Request"); err != nil {
				log.Errorf("Error responding to HTTP request: %s", err)
			}
			return
		}

		fn.ServeHTTP(w, req)
	}
}

func routeAction(action string, sqsAPI *api.API) (http.HandlerFunc, error) {
	switch action {
	case "ListQueues":
		return sqsAPI.ListQueues, nil
	case "AddPermission":
		return sqsAPI.AddPermission, nil
	case "ChangeMessageVisibility":
		return sqsAPI.ChangeMessageVisibility, nil
	case "ChangeMessageVisibilityBatch":
		return sqsAPI.ChangeMessageVisibilityBatch, nil
	case "CreateQueue":
		return sqsAPI.CreateQueue, nil
	case "DeleteMessage":
		return sqsAPI.DeleteMessage, nil
	case "DeleteMessageBatch":
		return sqsAPI.DeleteMessageBatch, nil
	case "DeleteQueue":
		return sqsAPI.DeleteQueue, nil
	case "GetQueueAttributes":
		return sqsAPI.GetQueueAttributes, nil
	case "GetQueueUrl":
		return sqsAPI.GetQueueURL, nil
	case "ListDeadLetterSourceQueues":
		return sqsAPI.ListDeadLetterSourceQueues, nil
	case "ListQueueTags":
		return sqsAPI.ListQueueTags, nil
	case "PurgeQueue":
		return sqsAPI.PurgeQueue, nil
	case "ReceiveMessage":
		return sqsAPI.ReceiveMessage, nil
	case "RemovePermission":
		return sqsAPI.RemovePermission, nil
	case "SendMessage":
		return sqsAPI.SendMessage, nil
	case "SendMessageBatch":
		return sqsAPI.SendMessageBatch, nil
	case "SetQueueAttributes":
		return sqsAPI.SetQueueAttributes, nil
	case "TagQueue":
		return sqsAPI.TagQueue, nil
	case "UntagQueue":
		return sqsAPI.UntagQueue, nil

	default:
		return nil, fmt.Errorf("invalid action: %s", action)
	}
}
