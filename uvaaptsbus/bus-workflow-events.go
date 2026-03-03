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

var EventSubmissionRegister = "workflow.submission.register"     // submission registered
var EventSubmissionInitiate = "workflow.submission.initiate"     // submission initiated
var EventSubmissionValidate = "workflow.submission.validate"     // submission validate
var EventSubmissionReconcile = "workflow.submission.reconcile"   // submission reconciliation
var EventSubmissionApprove = "workflow.submission.approve"       // submission to approve
var EventSubmissionBag = "workflow.submission.bag"               // submission approved
var EventSubmissionComplete = "workflow.submission.complete"     // submission complete
var EventSubmissionIncomplete = "workflow.submission.incomplete" // submission incomplete
var EventSubmissionRejected = "workflow.submission.rejected"     // submission rejected

var EventBagInitiate = "workflow.bag.initiate"  // bag initiate
var EventBagBuilt = "workflow.bag.built"        // bag has been built
var EventBagSubmit = "workflow.bag.submit"      // bag submit (to APT)
var EventBagSubmitted = "workflow.bag.submited" // bag submitted (to APT)
var EventBagAccepted = "workflow.bag.accepted"  // bag accepted (by APT)
var EventBagRejected = "workflow.bag.rejected"  // bag rejected (by APT)

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
