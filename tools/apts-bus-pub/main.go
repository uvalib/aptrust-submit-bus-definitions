package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/uvalib/apts-bus-definitions/uvaaptsbus"
)

// main entry point
func main() {

	var eventBus string
	var eventSource string
	var eventName string
	var debug bool
	var logger *log.Logger
	var clientId string
	var submissionId string
	var bagId string
	var bulk string

	flag.StringVar(&eventBus, "bus", "", "Event bus name")
	flag.StringVar(&eventSource, "source", "", "Event source name")
	flag.StringVar(&eventName, "event", "", "The name of the event")
	flag.StringVar(&clientId, "cid", "", "The event client identifier")
	flag.StringVar(&submissionId, "sid", "", "The event submission identifier (optional)")
	flag.StringVar(&bagId, "bid", "", "The event bag identifier (optional)")
	flag.StringVar(&bulk, "bulk", "", "File of events to publish (name/client/submission/bag)")
	flag.BoolVar(&debug, "debug", false, "Log debug information")
	flag.Parse()

	if debug == true {
		logger = log.Default()
	}

	// validate required parameters
	if len(eventBus) == 0 ||
		len(eventSource) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// validate optional parameters
	if len(bulk) == 0 && (len(eventName) == 0 ||
		len(clientId) == 0) {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// create the bus client
	cfg := uvaaptsbus.UvaBusConfig{
		Source:  eventSource,
		BusName: eventBus,
		Log:     logger,
	}
	bus, err := uvaaptsbus.NewUvaBus(cfg)
	if err != nil {
		log.Fatalf("ERROR: creating event bus client (%s)", err.Error())
	}
	fmt.Printf("Using: %s@%s\n", eventSource, eventBus)

	count := 0
	if len(bulk) != 0 {

		var bytesRead []byte
		bytesRead, err = os.ReadFile(bulk)
		if err == nil {
			lines := strings.Split(string(bytesRead), "\n")
			for _, line := range lines {
				attrs := strings.Split(line, "/")
				if len(attrs) == 4 {
					err = publishEvent(bus, attrs[0], attrs[1], attrs[2], attrs[3])
					count++
					if err != nil {
						break
					}
				} else {
					if len(line) != 0 {
						fmt.Printf("WARNING: ignoring: %s\n", line)
					}
				}
			}
		}
	} else {
		err = publishEvent(bus, eventName, clientId, submissionId, bagId)
		count++
	}

	if err != nil {
		fmt.Printf("Terminating with error (%s)\n", err.Error())
		os.Exit(1)
	} else {
		fmt.Printf("Terminating normally, %d event(s) published\n", count)
	}
}

func publishEvent(bus uvaaptsbus.UvaBus, name string, clientId string, submissionId string, bagId string) error {

	// if we have a submissionId and/or bagId, it is a workflow event
	wf := uvaaptsbus.UvaWorkflowEvent{
		SubmissionId: submissionId,
		BagId:        bagId,
	}

	jsonStr, _ := wf.Serialize()
	ev := uvaaptsbus.UvaBusEvent{
		EventName: name,
		ClientId:  clientId,
		Detail:    jsonStr,
	}
	err := bus.PublishEvent(&ev)
	if err != nil {
		return fmt.Errorf("publishing event (%s)", err.Error())
	}

	fmt.Printf("INFO: published: %s\n", ev.String())
	return nil
}

//
// end of file
//
