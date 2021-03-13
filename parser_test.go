package parser

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestParse(t *testing.T) {

	f, err := os.Open("testdata/changelog.md")
	defer f.Close()

	if err != nil {
		t.Fail()
	}

	parse, _ := New(f)

	Changelog := parse.Parse()

	if Changelog.Title != "Changelog" {
		t.Fail()
	}

	for _, version := range Changelog.Versions {
		if version.Name == "" {
			t.Fail()
		}
	}

	b, _ := json.Marshal(Changelog)
	log.Print(string(b))

}
