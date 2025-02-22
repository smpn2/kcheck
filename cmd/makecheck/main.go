package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/littlecxm/kcheck/configs"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var (
	version, buildDate, commitID string
	srcDir, outFilename          string
	isIndent                     bool
)

func main() {
	configs.WorkDir, _ = os.Getwd()
	app := &cli.App{
		Name:    "makecheck",
		Usage:   "make check list for kcheck",
		Version: fmt.Sprintf("%s %s (built: %s)", version, commitID, buildDate),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "src",
				Aliases:     []string{"s"},
				Usage:       "source `DIR`",
				Destination: &srcDir,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Value:       "kcheck.json",
				Usage:       "output filename for check list `FILE`",
				Destination: &outFilename,
			},
			&cli.BoolFlag{
				Name:        "indent",
				Aliases:     []string{"ind"},
				Value:       false,
				Usage:       "enable indent for output list",
				Destination: &isIndent,
			},
		},
		Action: commandHandler,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
