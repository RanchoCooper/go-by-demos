package main

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

type MessagePayload struct {
	UniqueID    string       `json:"unique_id"`
	Recipient   string       `json:"recipient"`
	MessageType string       `json:"message_type"` // refer to models.MessageType
	SendTime    string       `json:"send_time"`    // reserved to support scheduled sent, default value is `immediately`
	Params      ParamPayload `json:"params"`
}

type ParamPayload struct {
	InviterEmail        string `json:"inviter_email"`         // for invite message
	AlertMessage        string `json:"alert_message"`         // for alert message
	NotifyMessage       string `json:"notify_message"`        // for notify message
	SubgraphName        string `json:"subgraph_name"`         // for notify message
	SubgraphStatus      string `json:"subgraph_status"`       // for notify message
	SubgraphProjectName string `json:"subgraph_project_name"` // for notify message
}

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	uniqueID := uuid.NewString()
	// payload, _ := json.Marshal(MessagePayload{
	// 	UniqueID:    uniqueID,
	// 	Recipient:   "rancho@ormi.xyz",
	// 	MessageType: "welcome",
	// 	SendTime:    "immediately",
	// })
	payload, _ := json.Marshal(MessagePayload{
		UniqueID:    uniqueID,
		Recipient:   "rancho@ormi.xyz",
		MessageType: "invite",
		SendTime:    "immediately",
		Params:      ParamPayload{InviterEmail: "rancho@ormi.xyz"},
	})

	task := asynq.NewTask("message", payload)

	client.Enqueue(task)
	// ------------------------------------------------------
	// Example 1: Enqueue task to be processed immediately.
	//            Use (*Client).Enqueue method.
	// ------------------------------------------------------

	// task, err := tasks.NewEmailDeliveryTask(42, "some:template:id")
	// if err != nil {
	// 	log.Fatalf("could not create task: %v", err)
	// }
	// info, err := client.Enqueue(task)
	// if err != nil {
	// 	log.Fatalf("could not enqueue task: %v", err)
	// }
	// log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	//
	// // ------------------------------------------------------------
	// // Example 2: Schedule task to be processed in the future.
	// //            Use ProcessIn or ProcessAt option.
	// // ------------------------------------------------------------
	//
	// info, err = client.Enqueue(task, asynq.ProcessIn(24*time.Hour))
	// if err != nil {
	// 	log.Fatalf("could not schedule task: %v", err)
	// }
	// log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	//
	// // ----------------------------------------------------------------------------
	// // Example 3: Set other options to tune task processing behavior.
	// //            Options include MaxRetry, Queue, Timeout, Deadline, Unique etc.
	// // ----------------------------------------------------------------------------
	//
	// task, err = tasks.NewImageResizeTask("https://example.com/myassets/image.jpg")
	// if err != nil {
	// 	log.Fatalf("could not create task: %v", err)
	// }
	// info, err = client.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(3*time.Minute))
	// if err != nil {
	// 	log.Fatalf("could not enqueue task: %v", err)
	// }
	// log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
