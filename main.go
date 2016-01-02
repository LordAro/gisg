package main

import (
	"io/ioutil"
	"log"

	"github.com/LordAro/gisg/parsers"
	"github.com/docopt/docopt-go"
	"gopkg.in/yaml.v2"
)

//  pisg [-ch channel] [-l logfile] [-o outputfile] [-ma maintainer]
//       [-f format] [-ne network] [-d logdir] [-mo moduledir] [-s] [-v] [-h]

func ParseArgs() map[string]interface{} {
	usage := `Go IRC Statistics Generator.

Usage:
  gisg -c CFG [-s]

Options:
  -c --config-file=CFG  Configuration file [default: gisg.yaml].
  -s  --silent          Suppress output (except error messages).
  -v  --version         Show version.
  -h  --help            Output this message and exit.`

	arguments, _ := docopt.Parse(usage, nil, true, "GISG 0.1", false)
	return arguments
}

type logConfig struct {
	Network    string
	Format     string
	InputFile  string
	OutputFile string
}

type yamlConfig struct {
	Maintainer string
	Channels   map[string]logConfig
}

func main() {
	args := ParseArgs()

	cfgData, err := ioutil.ReadFile(args["--config-file"].(string))
	if err != nil {
		log.Fatalln(err.Error())
	}

	var cfg yamlConfig
	if err = yaml.Unmarshal(cfgData, &cfg); err != nil {
		log.Fatalln(err.Error())
	}

	for chanName, chanData := range cfg.Channels {
		parser := parsers.NewParser(chanData.Format)
		if parser == nil {
			log.Println("Unknown format: " + chanData.Format)
			continue
		}

		log.Println(chanName + ":")
		err = parsers.ParseFile(parser, chanData.InputFile)
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
}
