# log-daemon

Used to send a folder of log files to the log-collector webservice.

Usage:
  log-daemon [OPTIONS]

Application Options:
  -p, --path-to-logs= Path to the folder containing the logs
  -t, --topic=        The Kafka topic to store the logs under
  -k, --key-prefix=   A prefix added in front of each filename-key in kafka.
  -c, --clear-topic   A flag that when set will cause the passed topic to be cleared before any logs are stored into that topic (false)
  -r, --recurse       If daemon encounters a folder, restart the search inside that folder and so on (false)

Help Options:
  -h, --help          Show this help message


