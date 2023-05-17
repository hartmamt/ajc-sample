package main

import (
	"fmt"

	"time"

	aj "github.com/choria-io/asyncjobs"
	//nats "github.com/nats-io/nats.go"
)

func main() {
	client, _ := aj.NewClient(aj.NatsContext("AJC"))

	// The deadline being an hour from now will result in a Schedule Task with a 1 hour deadline set
	task, _ := aj.NewTask("email:monthly", nil, aj.TaskDeadline(time.Now().Add(time.Hour)))

	// Create the schedule
	err := client.NewScheduledTask("EMAIL_MONTHLY_UPDATE", "@monthly", "EMAIL", task)

	if err != nil {
		panic(err)
	}
	// Load it
	st, _ := client.LoadScheduledTaskByName("EMAIL_MONTHLY_UPDATE")

	fmt.Println(st)
	// Remove it
	err = client.RemoveScheduledTask("EMAIL_MONTHLY_UPDATE")
}
