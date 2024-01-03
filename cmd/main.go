package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/4ra1n/y4-lang/cli"
	"github.com/4ra1n/y4-lang/color"
)

func init() {
	// VERSION
	cli.Version = "v0.0.1"
	// BUILD TIME
	cli.BuildTime = "2024/01/01"
}

// Y4-Lang
// y4-lang is a script language based on golang
func main() {
	// CHECK COLOR
	if !color.IsSupported() {
		color.DisableColor()
	}
	// PRINT LOGO
	cli.PrintLogo()
	// START
	go cli.Start()
	// WAIT
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Println("stop")
}
