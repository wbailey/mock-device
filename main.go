package main

import "fmt"
import "github.com/wbailey/emitters"
import "github.com/urfave/cli"
import "time"
import "os"

//import "flag"

func main() {
	app := cli.NewApp()
	app.Name = "mock-device"
	app.Usage = "simulate a device taking and reporting various measurements"
	app.Action = func(c *cli.Context) error {
		e := emitters.ConstantEmitter{1.0, 1}

		for i := 0; i < 10; i++ {
			fmt.Println(i, e.Emit())
			time.Sleep(e.UseFreq())
		}

		return nil
	}

	app.Run(os.Args)
}
