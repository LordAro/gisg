package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/LordAro/gisg/formats"
	"github.com/docopt/docopt-go"
	"gopkg.in/yaml.v2"
)

func ParseArgs() map[string]interface{} {
	usage := `Go IRC Statistics Generator.

Usage:
  gisg [-c CFG -s]

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
		formatter := formats.NewFormatter(chanData.Format)
		if formatter == nil {
			log.Println("Unknown format: " + chanData.Format)
			continue
		}

		fmt.Println(chanName + ":")
		channel := NewChannel(chanName, chanData.Network, chanData.InputFile)
		err = channel.Process(formatter)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		//for _, u := range channel.Users {
		//	fmt.Println(u)
		//}
		//fmt.Println(channel.HoursActive())

		data, err := channel.HTML(cfg.Maintainer)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		f, err := os.Create(chanData.OutputFile)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		_, err = f.Write(data.Bytes())
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
}
