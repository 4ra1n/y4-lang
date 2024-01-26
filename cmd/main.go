package main

import (
	"context"
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
	cli.BuildTime = "2024/01/26"
}

// Y4-Lang
// y4-lang is a script language based on golang
func main() {
	// check color
	if !color.IsSupported() {
		color.DisableColor()
	}

	// new context
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	// start
	go cli.Start(cancel)

	// wait
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-sigChan:
			if !cli.QuietFlag {
				fmt.Println("ctrl+c stop")
			}
			return
		case <-ctx.Done():
			if !cli.QuietFlag {
				fmt.Println("y4-lang run finish")
			}
			return
		}
	}
}
