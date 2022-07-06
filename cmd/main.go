package main

import (
	"log"
	"os"

	"github.com/alash3al/scraply/cmd/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "scraply",
		Description: "if you know css, then you can scrap the world using this tool",
		Commands:    []*cli.Command{},
	}

	app.Commands = append(app.Commands, &cli.Command{
		Name:        "extract",
		Description: "extracts the required information from the specified url using the specified extractors",
		Action:      commands.Extractor(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Usage:    "the target url to extract the data from",
				Aliases:  []string{"u"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "user-agent",
				Usage:   "the User-Agent header value",
				Aliases: []string{"ua"},
				Value:   "scraply/fetch",
			},
			&cli.StringSliceFlag{
				Name:     "extract",
				Aliases:  []string{"x"},
				Required: true,
				Usage:    "the extractor(s) to be executed against the target url in the form of -x key=script -x key2=script2",
			},
			&cli.BoolFlag{
				Name:  "return-body",
				Usage: "whether to include body in the resulted information",
				Value: false,
			},
		},
	})

	app.Commands = append(app.Commands, &cli.Command{
		Name:        "serve",
		Description: "start the http server",
		Action:      commands.HTTPServer(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "address",
				Aliases: []string{"a"},
				Usage:   "the http server listen address",
				Value:   ":8010",
			},
			&cli.BoolFlag{
				Name:    "logging",
				Aliases: []string{"l"},
				Value:   true,
				Usage:   "whether to enable/disable logging (access logs)",
			},
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
