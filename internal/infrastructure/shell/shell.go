package shell

import (
	"android-cmd-server/internal/core/ports"
	"bytes"
	"context"
	"errors"
	"os/exec"
	"strings"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) RunCommand(ctx context.Context, name string, args []string, input []string) (*ports.Output, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if len(input) > 0 {
		joinedInput := strings.Join(input, "\n") + "\n"
		cmd.Stdin = strings.NewReader(joinedInput)
	}

	err := cmd.Run()

	if err != nil {
		return nil, errors.Join(err, errors.New(stderr.String()))
	}

	return &ports.Output{
		Stdout: strings.TrimSpace(stdout.String()),
		Stderr: strings.TrimSpace(stderr.String()),
	}, nil
}
