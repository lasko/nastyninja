package main

import (
	"fmt"
	"os"

	"github.com/caarlos0/spin"
	"github.com/lasko/nastyninja"
	"github.com/urfave/cli"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "nastyninja"
	app.Version = version
	app.Author = "Brandon Height (lasko@nastyninja.net)"
	app.Usage = "This is a set of microservices written in Go"
	app.Action = func(c *cli.Context) error {
		spin := spin.New("\033[36m %s Working...\033[m")
		spin.Start()
		err := nastyninja.Foo()
		spin.Stop()
		if err != nil {
			return err
		}
		fmt.Println("Done!")
		return nil
	}
	_ = app.Run(os.Args)
}
