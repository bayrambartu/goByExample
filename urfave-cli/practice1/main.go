package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "MyApp",
		Usage: "Demo with global and subcommand flags",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Enable verbose output (global flag)",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "hello",
				Usage: "Say hello to someone",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "Name to greet",
					},
				},
				Action: func(ctx *cli.Context) error {
					name := ctx.String("name")
					if name == "" {
						name = "stranger"
					}
					fmt.Printf("Hello, %s!\n", name)

					if ctx.App.Metadata["verbose"] == true {
						fmt.Println("Verbose mode is enabled (from global flag)")
					}
					return nil
				},
			},
		},
		Before: func(ctx *cli.Context) error {
			ctx.App.Metadata = map[string]interface{}{
				"verbose": ctx.Bool("verbose"),
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
