package main

import "fmt"
import "github.com/wbailey/emitters"
import "time"

//import "flag"

func main() {
	c := emitters.ConstantEmitter{1.0, 1}

	for i := 0; i < 10; i++ {
		fmt.Println(i, c.Emit())
		time.Sleep(c.UseFreq())
	}
}
