package main

import (
	"context"
	"io"
	"os"
	"os/signal"

	"github.com/yael-castro/passgen/internal/cli"
)

func main() {
	notifyCh := make(chan os.Signal)

	signal.Notify(notifyCh, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-notifyCh
		cancel()
	}()

	flags := cli.Flags{}

	cli.InitFlags(&flags)
	cli.ValidateFlags(flags)

	output := io.Discard

	if flags.Verbose {
		output = os.Stdout
	}

	cli.RunContext(ctx, output, flags)
}
