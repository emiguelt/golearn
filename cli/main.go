package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Test cli app"
	app.Usage = "test basic commands"

	flags := buidFlags()

	app.Commands = buildCommands(flags)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func buidFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name: "host",
		},
	}
}

func buildCommands(flags []cli.Flag) []cli.Command {
	return []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the name server for a given hostname",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}

				for _, server := range ns {

					fmt.Println(server.Host)
				}
				return nil
			},
		},
	}
}
