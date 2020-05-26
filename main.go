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
	app.Name = "Backgrounder CLI"
	app.Usage = "Set the background for your machine automagically."

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "tutorialledge.net",
		},
	}
	app.Commands = []*cli.Command{
		&cli.Command{
			Name:  "ns",
			Usage: "Looks up the nameservers for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
