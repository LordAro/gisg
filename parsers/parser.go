package parsers

import (
	"bufio"
	"os"
	"regexp"
)

type Parser interface {
	ProcessNormal(matches []string)
	ProcessAction(matches []string)
	ProcessOther(matches []string)

	NormalRegex() string
	ActionRegex() string
	OtherRegex() string
}

var parserMap = map[string]func() Parser{
	"irssi": NewIrssiParser,
}

// Returns nil in the case of
func NewParser(format string) Parser {
	if fun, ok := parserMap[format]; ok {
		return fun()
	}
	return nil
}

func GetLines(path string) ([]string, error) {
	inFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	lines := []string{}

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func ParseFile(parser Parser, infile string) error {
	lines, err := GetLines(infile)
	if err != nil {
		return err
	}

	normalregex := regexp.MustCompile(parser.NormalRegex())
	actionregex := regexp.MustCompile(parser.ActionRegex())
	otherregex := regexp.MustCompile(parser.OtherRegex())
	for _, line := range lines {
		if matches := normalregex.FindStringSubmatch(line); matches != nil {
			parser.ProcessNormal(matches)
		} else if matches := actionregex.FindStringSubmatch(line); matches != nil {
			parser.ProcessAction(matches)
		} else if matches := otherregex.FindStringSubmatch(line); matches != nil {
			parser.ProcessOther(matches)
		}
		// Ignore everything else
	}
	return nil
}
