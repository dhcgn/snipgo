package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Snippet struct {
	Category string `json:"category"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}
// Snippet implements fmt.Stringer
var _ fmt.Stringer = (*Snippet)(nil)

func ReadSnippets(filename string) ([]Snippet, error) {
	snippetFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer snippetFile.Close()
	byteValue, err := ioutil.ReadAll(snippetFile)
	if err != nil {
		return nil, err
	}
	var snippets []Snippet
	if err = json.Unmarshal(byteValue, &snippets); err != nil {
		return nil, err
	}
	return snippets, nil
}

func (s Snippet) String() string { return fmt.Sprintf("%s: %s", s.Category, s.Title) }

type ByCategoryTitle []Snippet
// ByCategoryTitle implements sort.Interface
var _ sort.Interface = (*ByCategoryTitle)(nil)

// Implement sort interface
func (s ByCategoryTitle) Len() int      { return len(s) }
func (s ByCategoryTitle) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByCategoryTitle) Less(i, j int) bool {
	return (strings.ToLower(s[i].Category) < strings.ToLower(s[j].Category)) ||
		((strings.ToLower(s[i].Category) == strings.ToLower(s[j].Category)) &&
			(strings.ToLower(s[i].Title) < strings.ToLower(s[j].Title)))
}
