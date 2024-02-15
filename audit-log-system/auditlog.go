package main

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch"
)

type AuditLog struct {
	Actor   string      `json:"actor"`
	Action  string      `json:"action"`
	Module  string      `json:"module"`
	When    time.Time   `json:"when"`
	Details interface{} `json:"details"`
}

func logAuditEvent(actor string, action string, module string, details interface{}) {
	// Log the audit event
	entry := AuditLog{
		Actor:   actor,
		Action:  action,
		Module:  module,
		When:    time.Now(),
		Details: details,
	}

	logEntry, _ := json.Marshal(entry)
	file, _ := os.OpenFile("audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	file.WriteString(string(logEntry) + "\n")
}

func sendToElasticSearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	file, err := os.Open("audit.log")
	if err != nil {
		log.Fatal("Error in Opening audit logs: %s", err)
	}
	defer close(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry AuditLog
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			log.Fatalf("Error in unmarshalling the log entry: %s", err)
		}

		body, err := json.Marshal(entry)
		if err != nil {
			log.Fatalf("Error in marshalling the log entry: %s", err)
		}

		req := esapi.IndexRequest{
			Index:        "audit-log",
			DocumentID:   "1",
			Body:         strings.NewReader(string(body)),
			Refresh:      "true",
			DocumentType: "_doc",
		}

		_, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatal("Error in sending the log entry to ElasticSearch: %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error in reading the audit logs: %s", err)
	}

}

func main() {
	logAuditEvent("Yash_Pimple", "create", "User Profile", "User created with username: Yash_Pimple")
}
