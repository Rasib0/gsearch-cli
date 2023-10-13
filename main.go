package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Define a struct to hold the name and URL
type SearchEngine struct {
	Name string
	URL  string
}

func main() {
	version := "0.1.0"
	numArgs := len(os.Args)

	searchURLs := map[string]SearchEngine{
		"!g": {Name: "Google", URL: "https://www.google.com/search?q="},
		"!y": {Name: "YouTube", URL: "https://www.youtube.com/results?search_query="},
		"!d": {Name: "DuckDuckGo", URL: "https://duckduckgo.com/?q="},
		"!p": {Name: "Phind", URL: "https://phind.com/search?q="},
	}

	var query string
	searchName := "!g"

	switch numArgs {
	case 2:
		query = os.Args[1]

		if query == "-h" || query == "--help" {
			printUsage(&searchURLs)
			os.Exit(0)
		}

		if query == "--version" {
			fmt.Printf("gsearch-cli version %s\n", version)
			os.Exit(0)
		}

	case 3:
		searchName = os.Args[1]
		query = os.Args[2]
	default:
		printUsage(&searchURLs)
		os.Exit(1)
	}

	searchEngine, exists := searchURLs[searchName]

	if !exists {
		fmt.Println("Invalid search flag")
		os.Exit(1)
	}

	searchAndOpen(query, &searchEngine)
}

// Print usage message
func printUsage(searchURLs *map[string]SearchEngine) {
	message := `Please provide a search query as argument (default search engine is Google)

Usage: ./gs <flag> <query> or ./gs <query>

Available search flags:`

	fmt.Println(message)

	for key, engine := range *searchURLs {
		fmt.Printf("%s (%s)\n", key, engine.Name)
	}
}

// Search and open the browser
func searchAndOpen(query string, searchEngine *SearchEngine) {
	query = strings.ReplaceAll(query, " ", "+")
	searchURL := searchEngine.URL + query

	fmt.Printf("Searching %s for %s...\n", searchEngine.Name, query)

	runtimeOS := runtime.GOOS
	var cmd *exec.Cmd

	switch runtimeOS {
	case "linux":
		cmd = exec.Command("xdg-open", searchURL)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", searchURL)
	case "darwin":
		cmd = exec.Command("open", searchURL)
	case "freebsd":
		cmd = exec.Command("xdg-open", searchURL)
	default:
		fmt.Printf("Unsupported OS: %s\n", runtimeOS)
		os.Exit(1)
	}

	// Start the browser or display an error message
	if err := cmd.Start(); err != nil {
		fmt.Println("Error opening browser")
		os.Exit(1)
	}
}
