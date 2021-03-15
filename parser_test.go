package parser

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, _ := os.Open("testdata/changelog.md")
	parser, _ := New(f)
	changelog := parser.Parse()

	t.Run("Check Versions", func(t *testing.T) {
		checkVersion(t, changelog)
	})

	t.Run("Check Versions", func(t *testing.T) {
		checkDates(t, changelog)
	})

}

func checkDates(t *testing.T, c *Changelog) {

	var testDates = []string{
		"2017-06-20",
		"2015-12-03",
		"2015-10-06",
		"2015-10-06",
		"2015-02-17",
		"2015-02-16",
		"2014-12-12",
		"2014-08-09",
		"2014-08-09",
		"2014-08-09",
		"2014-07-10",
		"2014-05-31",
	}

	for k, v := range c.Versions {
		if v.Date != testDates[k] {
			t.Fail()
		}
	}

	if len(c.Versions) != 12 {
		t.Fail()
	}

}

func checkVersion(t *testing.T, c *Changelog) {

	var testVersions = []string{
		"1.0.0",
		"0.3.0",
		"0.2.0",
		"0.1.0",
		"0.0.8",
		"0.0.7",
		"0.0.6",
		"0.0.5",
		"0.0.4",
		"0.0.3",
		"0.0.2",
		"0.0.1",
	}

	for k, v := range c.Versions {
		if v.Name != testVersions[k] {
			t.Fail()
		}
	}

	if len(c.Versions) != 12 {
		t.Fail()
	}

}
