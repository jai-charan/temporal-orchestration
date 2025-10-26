package mocks

import (
	"time"
	"go.temporal.io/sdk/client"
  "context"
    "log"
)


func MockEmailSignalForTemporal (c client.Client ,userID string){
	go func() {
		time.Sleep(10 * time.Second)
		err := c.SignalWorkflow(context.Background(), userID, "", "email_sent", "signal data")
		if err != nil {
			log.Println("Error sending 'email_sent' signal:", err)
		}
	}()



}

func MockSubscribeSignalForTemporal (c client.Client,userID string){
	go func() {
		time.Sleep(20 * time.Second)
		err := c.SignalWorkflow(context.Background(), userID, "", "subscribed", "signal data")
		if err != nil {
			log.Println("Error sending 'subscribed' signal:", err)
		}
	}()

}