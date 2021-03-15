package parser

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strings"
)

type Parser struct {
	raw            io.Reader
	CurrentVersion *Version
	CurrentSection *Section
}

type Changelog struct {
	Title    string
	Versions []*Version
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
	cl.Versions = append(cl.Versions, p.CurrentVersion)

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

		v := versionFactory()

		if p.CurrentVersion != nil {
			cl.Versions = append(cl.Versions, p.CurrentVersion)
		}

		lineSlice := strings.SplitN(line, "-", 2)

		re := regexp.MustCompile(`(?:\[)(.*)(?:\])`)
		match := re.FindStringSubmatch(lineSlice[0])

		v.Name = strings.TrimSpace(match[1])
		v.Date = strings.TrimSpace(lineSlice[1])

		p.CurrentVersion = v

		return
	}

	matched, _ = regexp.MatchString(`(?:^###)(?:\s)`, line)

	if matched {

		if p.CurrentVersion == nil {
			log.Fatal("Iterator became decoupled from version object")
		}

		name := strings.TrimSpace(strings.TrimPrefix(line, "###"))
		section := sectionFactory(name)

		if p.CurrentSection != nil {
			p.CurrentVersion.Body = append(p.CurrentVersion.Body, p.CurrentSection)
		}

		p.CurrentSection = section
	}

	matched, _ = regexp.MatchString(`(?:^-)(?:\s)`, line)

	if matched {

		if p.CurrentSection == nil {
			log.Fatal("Iterator became decoupled from section object")
		}

		content := strings.TrimSpace(strings.TrimPrefix(line, "-"))

		p.CurrentSection.Content = append(p.CurrentSection.Content, content)

	}

}

func versionFactory() *Version {
	return &Version{}
}

func sectionFactory(name string) *Section {
	return &Section{Name: name}
}
