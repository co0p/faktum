package main

import (
	"fmt"
	"log"
	"strings"

	prose "gopkg.in/jdkato/prose.v2"
)

type Entry struct {
	Text string
	Tags []string
}

func (e *Entry) IsArgument() bool {

	startsWithIf := false
	hasCommaThen := false
	previousTag := ""
	for k, v := range e.Tags {
		if k > 0 {
			previousTag = e.Tags[k-1]
		}
		if k == 0 && v == "IN" {
			startsWithIf = true
		}

		if previousTag == "," && v == "RB" {
			hasCommaThen = true
		}
	}

	return startsWithIf || hasCommaThen
}

func main() {

	doc, err := prose.NewDocument("I like cats. Peter is a man, therefore he drinks beer. If peter drinks beer, others will follow.")
	if err != nil {
		log.Fatal(err)
	}

	entries := []Entry{}

	sents := doc.Sentences()
	for _, sent := range sents {
		doc, _ := prose.NewDocument(sent.Text)
		entry := Entry{Text: strings.TrimSpace(sent.Text), Tags: []string{}}
		for _, tok := range doc.Tokens() {
			entry.Tags = append(entry.Tags, tok.Tag)
		}

		entries = append(entries, entry)
	}

	for _, k := range entries {
		fmt.Printf("%s (%v) = Argument? %v\n", k.Text, k.Tags, k.IsArgument())
	}

}
