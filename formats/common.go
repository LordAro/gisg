package formats

import ()

type Formatter interface {
	ProcessNormal(matches []string) (int, string, string)
	ProcessAction(matches []string) (int, string, string)
	ProcessOther(matches []string) (int, string, map[string]string)

	NormalRegex() string
	ActionRegex() string
	OtherRegex() string
}

var formatMap = map[string]func() Formatter{
	"irssi": NewIrssiFormatter,
}

// Returns nil in the case of
func NewFormatter(format string) Formatter {
	if fun, ok := formatMap[format]; ok {
		return fun()
	}
	return nil
}
