package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	stdexec "os/exec"

	"github.com/nobe4/clias/internal/aliases"
	"github.com/nobe4/clias/internal/config"
	"github.com/nobe4/clias/internal/exec"
	"github.com/nobe4/clias/internal/version"
)

const usage = `Usage: %s [clias-flags] [binary [flags]]
Version %s

clias-flags:
`

func main() {
	ctx := context.TODO()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0], version.String())
		flag.CommandLine.PrintDefaults()
	}

	listF := flag.Bool("list", false, "list defined aliases")
	debugF := flag.Bool("debug", false, "print debug information")
	configF := flag.String("config", config.Dir(), "custom config file")

	flag.Parse()

	if *debugF {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	slog.DebugContext(ctx, "config file", "path", *configF)

	c, err := os.Open(*configF)
	if err != nil {
		slog.ErrorContext(ctx, "failed to open the config file", "path", *configF, "err", err)
		os.Exit(1)
	}

	al, err := aliases.Parse(c)
	if err != nil {
		slog.ErrorContext(ctx, "failed to parse the aliases", "err", err)
		os.Exit(1)
	}

	if *listF {
		al.List(os.Stdout)

		return
	}

	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	binary, args := args[0], args[1:]
	if _, err := stdexec.LookPath(binary); err != nil {
		slog.ErrorContext(ctx, "binary not found", "binary", binary)

		return
	}

	if a := al.Find(binary, args); a != nil {
		slog.DebugContext(ctx, "alias found", "binary", binary, "args", args, "alias", a)

		args = a
	} else {
		slog.DebugContext(ctx, "alias not found, passing all args", "binary", binary, "args", args)
	}

	os.Exit(exec.Exec(ctx, binary, args))
}
