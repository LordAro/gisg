package parsers

import (
	//"fmt"
	"strconv"
)

type irssiParser struct {
	hourCount map[int]int
	nickCount map[string]int
}

func (*irssiParser) NormalRegex() string {
	return `^(\d+):\d+[^<*^!-]+<[@%+~& ]?([^>]+)> (.*)`
}

func (*irssiParser) ActionRegex() string {
	return `^(\d+):\d+[^ ]+ +\* (\S+) (.*)`
}

func (*irssiParser) OtherRegex() string {
	return `^(\d+):(\d+)[^-]+-\!- (\S+) (\S+) (\S+) (\S+) (\S+)(.*)`
}

func (ip *irssiParser) ProcessNormal(matches []string) {
	hour, _ := strconv.Atoi(matches[1])
	nick := matches[2]
	//text := matches[3]
	ip.hourCount[hour]++
	ip.nickCount[nick]++

}

func (ip *irssiParser) ProcessAction(matches []string) {
	//hour, _ := strconv.Atoi(matches[1])
	//nick := matches[2]
	//text := matches[3]
}

func (ip *irssiParser) ProcessOther(matches []string) {
	//fmt.Println(matches)
}

func NewIrssiParser() Parser {
	return &irssiParser{
		hourCount: make(map[int]int),
		nickCount: make(map[string]int),
	}
}
