package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "Backgrounder CLI"
	app.Usage = "Set the background for your machine automagically."

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "topic",
			Value: "the topic of the pic you'd like",
		},
	}
	app.Commands = []*cli.Command{
		&cli.Command{
			Name:  "random",
			Usage: "Returns a random picture",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				unsplashResponse := unsplash(c.String("topic"))

				body, err := ioutil.ReadAll(unsplashResponse.Body)
				if err != nil {
					log.Fatal(err)
				}

				f, err := os.Create("/tmp/random.jpg")
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				n2, err := f.Write(body)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Wrote %v bytes into %v\n", n2, "/tmp/random.jpg")

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
