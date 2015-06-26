# log-daemon

Used to send a folder of log files to the log-collector webservice.

USAGE:

	go run log-daemon.go PATH_TO_FOLDER TOPIC_NAME

where topic name is the name of the topic in kafka you want to add these logs to
and the key in kafka will be the file name
