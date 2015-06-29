#Log-Daemon

* Usage:
  * log-daemon [OPTIONS]

* Application Options:
  * -f, --path-to-logs= Path to the folder containing the logs
  * -t, --topic=        The Kafka topic to store the logs under
  * -h, --host=         The host name of the server where the log processor is running, Must not end in a slash (http://localhost)
  * -p, --port=         The port that the log processor is listening on (8080)
  * -r, --recurse       If daemon encounters a folder, restart the search inside that folder and so on (false)
  * -k, --key-prefix=   A prefix added in front of each filename-key in kafka.
  * -c, --clear-topic   A flag that when set will cause the passed topic to be cleared before any logs are stored into that topic (false)

* Help Options:
  * -h, --help          Show this help message

