package main

import (
	"fmt"
	"time"
	
	"github.com/rubrikinc/rubrik-polaris-sdk-for-go/rubrikpolaris"

)

const (

	timeFormat = "2006-01-02T15:04:05.000Z"
)



func main() {



	

	polaris := rubrikpolaris.Connect("rubrik-se", "drew.russell@outlook.com", "Welcome1!")
	// polaris := rubrikpolaris.Connect("rubrik-se", "drew.russell@rubrik.com", "XQHBFn7D7xjTdm")


	

	timeOfLastCheck := time.Now().UTC().Add(time.Minute * -120)

	radarEvent, err := polaris.GetAllPolarisEvents(timeOfLastCheck.Format(timeFormat)) // nolint
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(radarEvent)

	// go func() {
	// 	for {
	// 		log.Println("******************** Looking for new Polaris events ********************")
	// 		radarEvent, err := polaris.GetAllPolarisEvents(timeOfLastCheck.Format(timeFormat)) // nolint
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		for _, event := range radarEvent.Data.ActivitySeriesConnection.Edges { // nolint
	// 			log.Println("Found new event")

	// 			// structedField := util.CreateSyslogStructuredField(event.Node)

	// 			// Send the event message(s) to the Syslog server
	// 			for _, v := range event.Node.ActivityConnection.Nodes {
	// 				log.Println(v.Message)

	// 				// err = util.SendToSyslog(fmt.Sprintf("%s %s", structedField, v.Message), flag["log_server"], flag["log_port"], flag["log_network"]) // nolint
	// 				// if err != nil {
	// 				// 	log.Fatal(err)
	// 				// }
	// 			}

	// 			if event.Node.LastActivityStatus != "Success" || event.Node.LastActivityStatus != "Failure" {
	// 				log.Println("Event has completed")
	// 				// If the event has not completed, save it for follow up
	// 				inProgressEvents[event.Node.ActivitySeriesID] = event.Node.Cluster.ID // nolint

	// 			}
	// 		}
	// 		timeOfLastCheck = time.Now().UTC()
	// 		time.Sleep(time.Duration(newEventWaitPeriod) * time.Minute)
	// 	}
	// }()

	// for {
	// 	select {}
	// }
	
	
}
