package main

import (
	"encoding/json"
	"fmt"
)

type ConsoleOutputStrategy struct{}

func (console *ConsoleOutputStrategy) push(match *CertStreamMessage) error {
	matchBytes, _ := json.MarshalIndent(match, "", "  ")
	fmt.Printf("Match found:\n%s\n", string(matchBytes))

	return nil
}
