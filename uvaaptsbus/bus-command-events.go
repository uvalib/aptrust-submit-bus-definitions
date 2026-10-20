//
// Events related to explicit commands
//

package uvaaptsbus

//
// event names
//

// reprocess a submission
var EventCommandSubmissionInitiate = "command.submission.initiate"

// purge a submission's assets
var EventCommandSubmissionPurge = "command.submission.purge"

// reprocess a bag
var EventCommandBagInitiate = "command.bag.initiate"

// submit a bag (to APT)
var EventCommandBagSubmit = "command.bag.submit"

// request a bag status (from APT)
var EventCommandBagStatus = "command.bag.status"

// request a bag purge
var EventCommandBagPurge = "command.bag.purge"

// resend the email requesting submission approval
var EventCommandSubmissionEmailApprove = "command.submission.email.approve" // command to purge a bag's assets

//
// end of file
//
