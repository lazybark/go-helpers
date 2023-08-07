package main

import (
	"fmt"
	"time"

	"github.com/lazybark/go-helpers/npt"
)

func main() {
	//Calling Now() will create a new NPT from current moment in time:
	t := npt.Now()
	fmt.Println("Now it's:", t.Time())

	//ToNow() will set internals of NPT to current moment:
	//Time() will return time.Time object from NPT internals
	time.Sleep(2 * time.Second)
	t.ToNow()
	fmt.Println("Now it's:", t.Time())

	//FromTime() will set NPT to specified time value
	t.FromTime(time.Now().Add(time.Hour))
	fmt.Println("And now it's:", t.Time())

	//Add() will add specified duration to NPT
	t.Add(time.Hour)
	fmt.Println("And now it's:", t.Time())

}
