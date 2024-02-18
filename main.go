package main

import (
	"dal/cmd"
	"github.com/spf13/cast"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const Version = "v1.0"

func main() {

	app := &cli.App{
		Name:        "dal",
		Usage:       "dal",
		Description: "dal language for dsl",
		Version:     Version,
		Commands: []*cli.Command{
			{
				Name:   "server",
				Usage:  "dal server --path=<path> --port=<port>",
				Action: cmd.ServerHandler.Server,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "path",
						Usage:       "dsl file directory",
						DefaultText: cmd.ServerHandler.DefaultPath,
					},
					&cli.IntFlag{
						Name:        "port",
						Usage:       "http server port",
						DefaultText: cast.ToString(cmd.ServerHandler.DefaultPort),
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
