package formats

import (
	//"fmt"
	"strconv"
)

type irssiFormat struct {
}

func (*irssiFormat) NormalRegex() string {
	return `^(\d+):\d+[^<*^!-]+<[@%+~&[ ]+]?([^>]+)> (.*)`
}

func (*irssiFormat) ActionRegex() string {
	return `^(\d+):\d+[^ ]+ +\* (\S+) (.*)`
}

func (*irssiFormat) OtherRegex() string {
	return `^(\d+):(\d+)[^-]+-\!- (\S+) (\S+) (\S+) (\S+) (\S+)(.*)`
}

func (ip *irssiFormat) ProcessNormal(matches []string) (int, string, string) {
	hour, _ := strconv.Atoi(matches[1])
	nick := matches[2]
	text := matches[3]
	return hour, nick, text
}

func (ip *irssiFormat) ProcessAction(matches []string) (int, string, string) {
	hour, _ := strconv.Atoi(matches[1])
	nick := matches[2]
	text := matches[3]
	return hour, nick, text
}

func (ip *irssiFormat) ProcessOther(matches []string) (int, string, map[string]string) {
	hour, _ := strconv.Atoi(matches[1])
	nick := matches[3]
	return hour, nick, nil
}

func NewIrssiFormatter() Formatter {
	return &irssiFormat{}
}
