package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/CaliDog/certstream-go"
)

const workerCount = 10

func compareStrings(domain string, cert *CertStreamMessage, keywords []StringCompare) {
	for _, sc := range keywords {
		if strings.Contains(domain, sc.str) {
			sc.callback(cert)
		}
	}
}

func processStream() {
	keywordsEnv := os.Getenv("KEYWORDS")
	if keywordsEnv == "" {
		fmt.Println("Error: KEYWORDS environment variable is not set")
		return
	}

	var keywordStringsList []string
	err := json.Unmarshal([]byte(keywordsEnv), &keywordStringsList)
	if err != nil {
		fmt.Printf("Error: Could not parse KEYWORDS: %v\n", err)
		return
	}

	matches := 0

	arbitraryStringCompares := make([]StringCompare, len(keywordStringsList))
	for i, s := range keywordStringsList {
		arbitraryStringCompares[i] = StringCompare{s, func(cert *CertStreamMessage) {
			// certBytes, _ := json.MarshalIndent(cert, "", "  ")
			// fmt.Printf("Match found:\n%s\n", string(certBytes))
			fmt.Printf("Match found:\n%v\n", matches)
			matches++
		}}
	}

	stream, errStream := certstream.CertStreamEventStream(false)

	// Create a buffered channel for the jobs and a channel to signal that all workers are done
	jobs := make(chan *CertStreamMessage, workerCount*2)
	done := make(chan bool)

	// Start the workers
	for i := 0; i < workerCount; i++ {
		go func() {
			for cert := range jobs {
				domain := cert.Data.LeafCert.Subject.CN

				// CN is nil sometimes, if so then get a domain from all_domains for comparison
				if domain == "" && len(cert.Data.LeafCert.AllDomains) > 0 {
					domain = cert.Data.LeafCert.AllDomains[0]
				}

				if domain != "" {
					compareStrings(domain, cert, arbitraryStringCompares)
				}
			}
			done <- true
		}()
	}

	// Process the certificate stream.
	for {
		select {
		case jq := <-stream:
			jq_object, err := jq.Object()
			if err != nil {
				fmt.Printf("Error extracting jq: %v\n", err)
			}

			jq_bytes, err := json.Marshal(jq_object)
			if err != nil {
				fmt.Printf("Error marshaling jq object: %v\n", err)
			}

			cert := CertStreamMessage{}
			err = json.Unmarshal(jq_bytes, &cert)
			if err != nil {
				fmt.Printf("Error unmarshaling certificate JSON: %v\n", err)
				continue
			}

			// Send the certificate to one of the workers.
			jobs <- &cert

		case err := <-errStream:
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	go processStream()

	http.HandleFunc("/health", healthCheckHandler)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
