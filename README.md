Logger implements a shared logging infrastructure. It add several
features on top of the standard logging package:

* The usual log levels, from trace up to panic.
* Logs are created for named services, each of which can log at its own
  log level.
* Eacch log has a group number that can be used to group related log entries.
* Each log entry has a globally unique sequence number that can be
  used to reconstruct the order in which log entries were generated
  even in the face of out-of-order log recording.
* Logs can have callback functions registered for them to facilitate
  sending log entries as events.

