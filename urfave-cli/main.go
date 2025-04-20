/*
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/urfave/cli"
	"log"
	"os"
	"runtime"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "SecureToolCLI"
	app.Usage = "A simple CLI to kill processes and list mounted volumes"

	app.Commands = []cli.Command{
		{
			Name:        "kill",
			Usage:       "Kill a process by its ID or name",
			Description: "Terminate a running process",
			Action:      KillAction,
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:  "id",
					Usage: "Process ID to kill",
				},
				&cli.StringFlag{
					Name:  "name",
					Usage: "Process name to kill",
				},
			},
		},
		{
			Name:        "volumes",
			Usage:       "List mounted volumes",
			Description: "Shows details about mounted disk partitions",
			Action:      ActionVolumes,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func KillAction(c *cli.Context) error {
	if len(c.Args()) > 0 {
		return errors.New("no arguments expected, use flags instead")
	}
	if c.IsSet("id") && c.IsSet("name") {
		return errors.New("provide only one of --id or --name")
	}
	if !c.IsSet("id") && c.String("name") == "" {
		return errors.New("you must provide --id or --name")
	}
	if err := killProcess(c); err != nil {
		return err
	}
	fmt.Println("Process killed successfully.")
	return nil
}

func killProcess(c *cli.Context) error {
	if c.IsSet("id") {
		proc, err := process.NewProcess(int32(c.Uint("id")))
		if err != nil {
			return err
		}
		return proc.Kill()
	}

	processes, err := process.Processes()
	if err != nil {
		return err
	}

	var (
		errs  []string
		found bool
	)
	target := c.String("name")
	for _, p := range processes {
		name, _ := p.Name()
		if isEqualProcessName(name, target) {
			found = true
			if err := p.Kill(); err != nil {
				errs = append(errs, err.Error())
			}
		}
	}
	if !found {
		return errors.New("process not found")
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}
	return nil
}

func isEqualProcessName(p1, p2 string) bool {
	if runtime.GOOS == "linux" {
		return p1 == p2
	}
	return strings.EqualFold(p1, p2)
}

type Volume struct {
	Name       string
	Total      uint64
	Used       uint64
	Available  uint64
	UsePercent float64
	Mount      string
}

func ActionVolumes(c *cli.Context) error {
	stats, err := disk.Partitions(true)
	if err != nil {
		return err
	}

	var vols []Volume

	for _, stat := range stats {
		usage, err := disk.Usage(stat.Mountpoint)
		if err != nil {
			continue
		}
		vol := Volume{
			Name:       stat.Device,
			Total:      usage.Total,
			Used:       usage.Used,
			Available:  usage.Free,
			UsePercent: usage.UsedPercent,
			Mount:      stat.Mountpoint,
		}
		vols = append(vols, vol)
	}

	out, err := json.MarshalIndent(vols, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}
*/

/*

package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {

	app := cli.App{
		Action: func(ctx *cli.Context) error {
			fmt.Println("hello world")
			return nil
		},
	}

	app.Run(os.Args) // starts a cli app



}
*/

package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "MyApp",
		Usage: "A sample app with a flag",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "enable verbose output",
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.Bool("verbose") {
				fmt.Println("Verbose mode enabled.")
			} else {
				fmt.Println("Verbose mode disabled.")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
