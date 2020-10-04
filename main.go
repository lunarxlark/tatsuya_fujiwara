package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	if err := newApp().Run(os.Args); err != nil {
	}
}

var commands = []*cli.Command{
	{Name: "shout", Action: doShout},
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "tf"
	app.Usage = "respect tatsuya fujiwara"
	app.Version = "0.0.1"
	app.Authors = []*cli.Author{
		{Name: "lunarxlark", Email: "lunarxlark@gmail.com"},
	}
	app.Commands = commands

	return app
}

func doShout(ctx *cli.Context) error {
	args := ctx.Args().Slice()
	msg := strings.Join(args, " ")
	for _, c := range msg {
		fmt.Printf("%v゛", string([]rune{c}))
	}

	return nil
}