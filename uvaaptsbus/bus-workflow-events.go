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

var EventSubmissionRegister = "workflow.submission.register" // submission registered
var EventSubmissionInitiate = "workflow.submission.initiate" // submission initiated
var EventSubmissionValidate = "workflow.submission.validate" // submission validate
var EventSubmissionComplete = "workflow.submission.complete" // submission complete
var EventSubmissionApproved = "workflow.submission.approved" // submission approved
var EventSubmissionRejected = "workflow.submission.rejected" // submission rejected
var EventBagInitiate = "workflow.bag.initiate"               // bag initiate
var EventBagValidate = "workflow.bag.validate"               // bag initiate
var EventBagComplete = "workflow.bag.complete"               // bag complete
var EventBagSubmit = "workflow.bag.submit"                   // bag submit (to APT)
var EventBagAccept = "workflow.bag.accept"                   // bag accepted (by APT)

//
// corresponding schema for these events
//

type UvaWorkflowEvent struct {
	SubmissionId string `json:"submission_id"` // submission identifier
	BagId        string `json:"bag_id"`        // bag identifier (optional)
}

// standard behavior
func (impl UvaWorkflowEvent) String() string {
	return fmt.Sprintf("(%s/%s)>", impl.SubmissionId, impl.BagId)
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
