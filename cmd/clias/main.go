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
	"github.com/nobe4/clias/internal/generators"
	"github.com/nobe4/clias/internal/version"
)

const usage = `Usage: %s [clias-flags] [binary [flags]]
Version %s

clias-flags:
`

const usageGenerators = `
Generators:
  - alias
  - comp-bash
  - comp-zsh
  - comp-clias-bash
  - comp-clias-zsh

Specify multiple by joining them with a comma ','.
`

// revive:disable:cognitive-complexity // This is fine.
func main() {
	ctx := context.TODO()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0], version.String())
		flag.CommandLine.PrintDefaults()
		fmt.Fprint(flag.CommandLine.Output(), usageGenerators)
	}

	generateF := flag.String("generate", "", "generate custom output, see Generators.")
	listF := flag.Bool("list", false, "list defined aliases")
	debugF := flag.Bool("debug", false, "print debug information")
	configF := flag.String("config", config.Dir(), "custom config file")

	flag.Parse()

	if *debugF {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	al, err := getAliases(ctx, *configF)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get aliases", "err", err)
		os.Exit(1)
	}

	if *generateF != "" {
		if err := generators.Generate(*generateF, al, os.Stdout); err != nil {
			slog.ErrorContext(ctx, "failed to generate", "err", err)
			os.Exit(1)
		}

		return
	}

	if *listF {
		al.List(os.Stdout)

		return
	}

	args := flag.Args()

	slog.DebugContext(ctx, "remaining args", "args", args)

	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	binary, args := args[0], args[1:]
	slog.DebugContext(ctx, "searching for binary", "binary", binary)

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
} // revive:enable:cognitive-complexity

func getAliases(ctx context.Context, path string) (aliases.Aliases, error) {
	slog.DebugContext(ctx, "config file", "path", path)

	c, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open the config file %q: %w", path, err)
	}

	al, err := aliases.Parse(c)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the aliases %w", err)
	}

	return al, nil
}
