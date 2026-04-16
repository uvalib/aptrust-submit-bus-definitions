//
// Events related to explicit commands
//

package uvaaptsbus

//
// event names
//

var EventCommandSubmissionInitiate = "command.submission.initiate" // command to reprocess a submission
var EventCommandSubmissionPurge = "command.submission.purge"       // command to purge a submission's assets
var EventCommandBagInitiate = "command.bag.initiate"               // command to reprocess a bag
var EventCommandBagSubmit = "command.bag.submit"                   // command to submit a bag (to APT)
var EventCommandBagStatus = "command.bag.status"                   // command to request a bag status (from APT)
var EventCommandBagPurge = "command.bag.purge"                     // command to purge a bag's assets

//
// end of file
//
