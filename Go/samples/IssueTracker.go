package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/sreekanthn/horizon/Go/samples/github"
)

//SearchIssues - Searches the issues based on the query from github.com
func SearchIssues(terms []string) (*github.IssuesSearchResult, error) {
	query := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(github.IssuesURL + "?q=" + query)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Error occured while fetching from HTTP")
	}

	var result github.IssuesSearchResult
	err1 := json.NewDecoder(resp.Body).Decode(&result)
	if err1 != nil {
		resp.Body.Close()
		return nil, err1
	}

	resp.Body.Close()
	return &result, nil
}

func main() {

	result, err2 := SearchIssues(os.Args[1:])
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}
