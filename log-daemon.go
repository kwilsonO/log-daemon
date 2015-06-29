package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var opts struct {
	Path          flags.Filename `short:"f" long:"path-to-logs" description:"Path to the folder containing the logs" required:"true"`
	Topic         string         `short:"t" long:"topic" description:"The Kafka topic to store the logs under" required:"true"`
	ProcessorHost string         `default:"http://localhost" short:"s" long:"host" description:"The host name of the server where the log processor is running, Must not end in a slash"`
	ProcessorPort string         `default:"8080" short:"p" long:"port" description:"The port that the log processor is listening on"`
	Recursive     bool           `default:"false" short:"r" long:"recurse" description:"If daemon encounters a folder, restart the search inside that folder and so on" required:"false"`
	KeyPrefix     string         `short:"k" long:"key-prefix" description:"A prefix added in front of each filename-key in kafka." required:"false"`
	ClearTopic    bool           `default:"false" short:"c" long:"clear-topic" description:"A flag that when set will cause the passed topic to be cleared before any logs are stored into that topic" required:"false"`
}

type KafkaMsg struct {
	Topic string
	Key   string
	Value string
}

var KEY_PREFIX = ""
var PROCESSOR_HOST = "http://localhost"
var PROCESSOR_PORT = "8080"
var FolderTopic = "RouterLogs"

func sendMsg(msg KafkaMsg) {

	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Could not marshal json\n")
		return
	}

	var reqDest = PROCESSOR_HOST + ":" + PROCESSOR_PORT + "/"
	req, err := http.NewRequest("POST", reqDest, bytes.NewBuffer(b))
	if err != nil {
		fmt.Printf("Failed creating request: %s\n\n", err)
	}

	client := &http.Client{}

	_, err = client.Do(req)

	if err != nil {
		fmt.Printf("Failed making request: %s\n", err)

	}

}
func visit(path string, f os.FileInfo, err error) error {

	var fMod os.FileMode
	fMod = f.Mode()
	if fMod.IsDir() {
		fmt.Printf("Entering Folder: %s\n", f.Name())
		return nil
	}
	fmt.Printf("Saving log: %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("trouble opening file: %s\n", err)
		return nil
	}

	b := make([]byte, f.Size())
	n, err := file.Read(b)

	if int64(n) != f.Size() {
		fmt.Printf("Error reading entire file: %s\n", err)
	}
	var finalkey = KEY_PREFIX + f.Name()
	var kmsg KafkaMsg = KafkaMsg{
		Topic: FolderTopic,
		Key:   finalkey,
		Value: string(b),
	}

	sendMsg(kmsg)

	return nil
}

func clearTopic(topicName string) {
	//todo delete all messages of certain typ
	fmt.Println("Clearing a topic is not yet supported at this time.\n")
}

func main() {

	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(1)
	}

	if opts.ClearTopic {
		clearTopic(opts.Topic)
	}

	if opts.KeyPrefix != "" {
		KEY_PREFIX = opts.KeyPrefix + "-"
	}

	if opts.ProcessorHost != "" {
		if !strings.HasSuffix(opts.ProcessorHost, "/") {
			PROCESSOR_HOST = opts.ProcessorHost
		} else {
			fmt.Println("Cannot have a hostname that ends with a slash, try again.")
			os.Exit(1)
		}
	}

	if opts.ProcessorPort != "" {
		PROCESSOR_PORT = opts.ProcessorPort
	}

	err = filepath.Walk(string(opts.Path), visit)
	if err != nil {
		fmt.Printf("Error saving log files: %s\n", err)
	}
}
