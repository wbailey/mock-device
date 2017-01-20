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
	app.Usage = "simulate an IoT device"

	c := emitters.ConstantEmitter{1.0, 1}

	for i := 0; i < 10; i++ {
		fmt.Println(i, c.Emit())
		time.Sleep(c.UseFreq())
	}

	app.Run(os.Args)
}
