package exec

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/exec"
)

func Exec(ctx context.Context, binary string, args []string) int {
	cmd := exec.CommandContext(ctx, binary, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err == nil {
		return 0
	}

	var exitE *exec.ExitError
	if errors.As(err, &exitE) {
		return exitE.ExitCode()
	}

	slog.ErrorContext(ctx, "command failed with a non-exit error", "err", err)

	return 1
}
