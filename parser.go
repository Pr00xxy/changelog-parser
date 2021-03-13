package parser

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strings"
)

type Parser struct {
	raw io.Reader
}

type Changelog struct {
	Title          string
	CurrentVersion *Version
	CurrentSection *Section
	Versions       []*Version
}

type Version struct {
	Name string
	Date string
	Body []*Section
}

type Section struct {
	Name    string
	Content []string
}

func New(content io.Reader) (*Parser, error) {
	p := Parser{raw: content}
	return &p, nil
}

func (p Parser) Parse() *Changelog {
	scanner := bufio.NewScanner(p.raw)

	scanner.Split(bufio.ScanLines)

	cl := &Changelog{}

	for scanner.Scan() {
		p.handleLine(cl, scanner.Text())
	}

	// push last version to main data
	cl.Versions = append(cl.Versions, cl.CurrentVersion)

	return cl
}

func (p *Parser) handleLine(cl *Changelog, line string) {

	matched, _ := regexp.MatchString(`(?:^#)(?:\s)([A-z]*)`, line)

	if matched {
		cl.Title = strings.TrimSpace(strings.TrimPrefix(line, "#"))
		return
	}

	matched, _ = regexp.MatchString(`(?:^##)(?:\s)`, line)

	if matched {

		v := p.VersionFactory()

		if cl.CurrentVersion != nil {
			cl.Versions = append(cl.Versions, cl.CurrentVersion)
		}

		lineSlice := strings.Split(line, "-")

		re := regexp.MustCompile(`(?:\[)(.*)(?:\])`)
		match := re.FindStringSubmatch(lineSlice[0])

		v.Name = strings.TrimSpace(match[1])
		v.Date = strings.TrimSpace(lineSlice[1])

		cl.CurrentVersion = v

		return
	}

	matched, _ = regexp.MatchString(`(?:^###)(?:\s)`, line)

	if matched {

		if cl.CurrentVersion == nil {
			log.Fatal("Iterator became decoupled from version object")
		}

		name := strings.TrimSpace(strings.TrimPrefix(line, "###"))
		section := &Section{
			Name: name,
		}

		if cl.CurrentSection != nil {
			cl.CurrentVersion.Body = append(cl.CurrentVersion.Body, cl.CurrentSection)
		}

		cl.CurrentSection = section
	}

	matched, _ = regexp.MatchString(`(?:^-)(?:\s)`, line)

	if matched {

		if cl.CurrentSection == nil {
			log.Fatal("Iterator became decoupled from section object")
		}

		content := strings.TrimSpace(strings.TrimPrefix(line, "-"))

		cl.CurrentSection.Content = append(cl.CurrentSection.Content, content)

	}

}

func (p Parser) VersionFactory() *Version {
	return &Version{}
}
