package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type KafkaMsg struct {
	Topic string
	Key   string
	Value string
}

var FolderTopic = "RouterLogs"

func sendMsg(msg KafkaMsg) {

	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Could not marshal json")
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/", bytes.NewBuffer(b))
	//req, err := http.NewRequest("GET", "http://example.com/?data", nil)
	if err != nil {
		fmt.Printf("FAILED !!! %s", err)
	}

	client := &http.Client{}

	_, err = client.Do(req)

	if err != nil {
		fmt.Printf("FAILED EVEN MORE !!! %s", err)

	}

}
func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("trouble opening file: %s", err)
		return nil
	}

	b := make([]byte, f.Size())
	n, err := file.Read(b)

	if int64(n) != f.Size() {
		fmt.Printf("Error reading entire file: %s", err)
	}

	var kmsg KafkaMsg = KafkaMsg{
		Topic: FolderTopic,
		Key:   f.Name(),
		Value: string(b),
	}

	sendMsg(kmsg)

	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	topicarg := flag.Arg(1)
	if topicarg != "" {
		FolderTopic = topicarg
	}
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
