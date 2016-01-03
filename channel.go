package main

import (
	"bufio"
	"bytes"
	"html/template"
	"os"
	"regexp"
	"time"

	"github.com/LordAro/gisg/formats"
)

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

type Channel struct {
	Name    string
	Network string

	Users []*User

	infile  string
	outfile string
}

func (c *Channel) lazyUser(name string) *User {
	for _, u := range c.Users {
		if u.Name == name {
			return u
		}
	}
	c.Users = append(c.Users, NewUser(name))
	return c.Users[len(c.Users)-1]
}

func (c *Channel) Process(f formats.Formatter) error {
	lines, err := GetLines(c.infile)
	if err != nil {
		return err
	}
	normalregex := regexp.MustCompile(f.NormalRegex())
	actionregex := regexp.MustCompile(f.ActionRegex())
	otherregex := regexp.MustCompile(f.OtherRegex())

	for _, line := range lines {
		var userRef *User
		if matches := normalregex.FindStringSubmatch(line); matches != nil {
			hour, nick, text := f.ProcessNormal(matches)
			userRef = c.lazyUser(nick)
			userRef.HoursActive[hour]++
			userRef.Messages = append(userRef.Messages, text)

		} else if matches := actionregex.FindStringSubmatch(line); matches != nil {
			hour, nick, text := f.ProcessAction(matches)
			userRef = c.lazyUser(nick)
			userRef.HoursActive[hour]++
			userRef.Actions = append(userRef.Actions, text)

		} else if matches := otherregex.FindStringSubmatch(line); matches != nil {
			_, nick, _ := f.ProcessOther(matches)
			userRef = c.lazyUser(nick)
		}
		// Ignore everything else
	}

	return nil
}

func (c *Channel) HoursActive() [24]int {
	var channelActive [24]int
	for _, u := range c.Users {
		for i := range u.HoursActive {
			channelActive[i] += u.HoursActive[i]
		}
	}
	return channelActive
}

func (c *Channel) HTML(maint string) (*bytes.Buffer, error) {
	t, err := template.ParseGlob("templates/*.tmpl")
	if err != nil {
		return nil, err
	}

	s := struct {
		*Channel
		Maintainer string
		GenTime    time.Time
	}{
		c,
		maint,
		time.Now().UTC(),
	}

	buf := new(bytes.Buffer)
	err = t.ExecuteTemplate(buf, "main.tmpl", s)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func NewChannel(name, network, infile string) *Channel {
	return &Channel{
		Name:    name,
		Network: network,

		infile: infile,
	}
}
