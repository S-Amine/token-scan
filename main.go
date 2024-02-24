package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/s-Amine/token-scan/scanners/goplus"
	"github.com/s-Amine/token-scan/scanners/ishoneypot"
	"github.com/s-Amine/token-scan/scanners/multiscan"
	"github.com/s-Amine/token-scan/scanners/quickintel"
)

func main() {
	// Define command-line flags
	mode := flag.String("mode", "", "Mode of operation: multiscan, goplus, ishoneypot, or quickIntel")
	tokenHash := flag.String("token", "", "Token hash to scan")
	flag.Parse()

	if *mode == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *tokenHash == "" {
		fmt.Println("Error: Token hash is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var result interface{}
	var err error

	switch *mode {
	case "multiscan":
		result = multiscan.Scan(*tokenHash)
	case "goplus":
		result, err = goplus.Scan(*tokenHash)
	case "ishoneypot":
		result, err = ishoneypot.Scan(*tokenHash)
	case "quickIntel":
		result, err = quickintel.Scan(*tokenHash)
	default:
		fmt.Println("Error: Invalid mode specified")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error occurred during %s scan: %v\n", *mode, err)
		os.Exit(1)
	}

	printJSON(result)
}

// printJSON prints the provided data structure as JSON to stdout
func printJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
