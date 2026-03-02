//
// Events related to scheduling
//

package uvaaptsbus

//
// event names
//

var EventCommandSubmissionInitiate = "command.submission.initiate" // command to reprocess a submission
var EventCommandBagInitiate = "command.bag.initiate"               // command to reprocess a bag
var EventCommandBagSubmit = "command.bag.submit"                   // command to submit a bag (to APT)
var EventCommandBagStatus = "command.bag.status"                   // command to request a bag status (from APT)

//
// end of file
//
