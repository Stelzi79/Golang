package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func main() {
	flags := []cli.Flag{
		altsrc.NewIntFlag(&cli.IntFlag{Name: "test"}),
		&cli.StringFlag{Name: "load"},
	}

	app := &cli.App{
		Action: func(*cli.Context) error {
			fmt.Println("--test value.*default: 0")
			return nil
		},
		Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("load")),
		Flags:  flags,
	}
	// var Reset = "\033[0m"
	// var Red = "\033[31m"
	// var Green = "\033[32m"
	// var Yellow = "\033[33m"
	// var Blue = "\033[34m"
	// var Magenta = "\033[35m"
	// var Cyan = "\033[36m"
	// var Gray = "\033[37m"
	// var White = "\033[97m"
	// println(Red + "This is 🚧 Red" + Reset)
	// println(Green + "This is Green" + Reset)
	// println(Yellow + "This is Yellow" + Reset)
	// println(Blue + "This is Blue" + Reset)
	// println(Magenta + "This is Purple" + Reset)
	// println(Cyan + "This is Cyan" + Reset)
	// println(Gray + "This is Gray" + Reset)
	// println(White + "This is White" + Reset)

	println(color.RedString("This is Red"))
	println(color.GreenString("This is Green"))
	println(color.YellowString("This is Yellow"))
	println(color.BlueString("This is Blue"))
	println(color.MagentaString("This is Purple"))
	println(color.CyanString("This is Cyan"))


	app.Run(os.Args)
}
