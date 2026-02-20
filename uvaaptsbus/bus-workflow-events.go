//
// Events related to workflow lifecycle
//

package uvaaptsbus

import (
	"encoding/json"
	"fmt"
)

//
// event names
//

var EventSubmissionCreate = "workflow.submission.create"     // work created
var EventSubmissionBegin = "workflow.submission.begin"       // work created
var EventSubmissionComplete = "workflow.submission.complete" // work published
var EventBagBegin = "workflow.bag.began"                     // work deleted
var EventBagComplete = "workflow.bag.complete"               // work published

//
// corresponding schema for these events
//

type UvaWorkflowEvent struct {
}

// standard behavior
func (impl UvaWorkflowEvent) String() string {
	return "<none>"
}

func (impl UvaWorkflowEvent) Serialize() ([]byte, error) {
	// serialize the event object
	buf, err := json.Marshal(impl)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, ErrEventSerialize)
	}
	return buf, nil
}

func MakeWorkflowEvent(buf []byte) (*UvaWorkflowEvent, error) {
	var event UvaWorkflowEvent
	err := json.Unmarshal(buf, &event)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, ErrEventDeserialize)
	}
	return &event, nil
}

//
// end of file
//
