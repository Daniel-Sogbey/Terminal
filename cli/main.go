package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	debug = flag.Bool("debug", false, "Log out all debug information")
	// apiKey = flag.String("apiKey", "", "Sets the api key")

	usage = `Specify a command to execute:
	- search-repos: Search for github repos
	- search-users: search for users on github`
)

func main() {
	flag.Parse()
	fmt.Printf("Flag args : %v", flag.Args())
	// fmt.Printf("API KEY : %s", *apiKey)

	if len(flag.Args()) < 1 {
		fmt.Println(usage)
		os.Exit(1)
	}

	command := flag.Args()[0]
	err := executeCommand(command, flag.Args()[1:])
	if err != nil {
		fmt.Println("Err: ", err)
		os.Exit(1)
	}

	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}

func executeCommand(command string, args []string) error {
	printDebug("Execute command : " + command)

	switch command {
	case "search-repos":
		fmt.Println("Search repos command")
		return executeSearchRepos(args)
	case "search-users":
		fmt.Println("Search users command")
		return nil
	default:
		return fmt.Errorf("Invald command: '%s' \n\n %s \n", command, usage)
	}
}

func executeSearchRepos(args []string) error {

	if len(args) == 0 {
		return errors.New("Provide a search term for searching repositories : search-repos <search-term>")
	}

	searchTerm := args[0]

	printDebug(fmt.Sprintf("[search-repos] Search Term: %s ", searchTerm))

	repos, err := findRepos(searchTerm)

	if err != nil {
		return err
	}

	fmt.Println(strings.Join(repos, ","))

	return nil

}

func findRepos(term string) ([]string, error) {
	url := "https://api.github.com/search/repositories"

	type repoName struct {
		FullName string `json:"full_name"`
	}

	type searchResult struct {
		TotalCount        int        `json:"total_count"`
		IncompleteResults bool       `json:"incomplete_results"`
		Items             []repoName `json:"items"`
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		printDebug(fmt.Sprintf("%v", err))
		return nil, errors.New("Failed to connect to github")
	}

	//set search term as query parameter and encode query param
	query := req.URL.Query()
	query.Set("q", term)
	req.URL.RawQuery = query.Encode()

	//make http request to github server
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		printDebug(fmt.Sprintf("%v", err))
		return nil, errors.New("Failed to connect to github")
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, errors.New("Failed to connect to github")
	}

	// resBody, _ := io.ReadAll(req.Body)
	bs := make([]byte, 0)

	_, _ = res.Body.Read(bs)

	// fmt.Printf("Response : %v", string(bs))

	// io.Reader.Read(req.Body, bs)

	// fmt.Println(string(bs))

	results := searchResult{}

	err = json.NewDecoder(res.Body).Decode(&results)

	if err != nil {
		printDebug(fmt.Sprintf("ERROR : %v", err))
		return nil, err
	}

	// fmt.Printf("Result> %v", results)

	repos := make([]string, 0)

	for _, r := range results.Items {
		repos = append(repos, r.FullName)
	}

	return repos, nil
}

func printDebug(msg string) {
	if *debug {
		fmt.Printf("[DEBUG]: %s\n", msg)
	}
}
