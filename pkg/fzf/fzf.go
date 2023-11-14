package fzf

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mattn/go-isatty"
)

func InteractiveChoice(command string) (string, error) {
	cmd := exec.Command("fzf", "--ansi", "--no-preview")
	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = &out

	cmd.Env = append(os.Environ(), fmt.Sprintf("FZF_DEFAULT_COMMAND=%s", command))
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return "", err
		}
	}
	choice := strings.TrimSpace(out.String())
	parts := strings.Split(choice, "\t")
	choice = strings.TrimSpace(parts[0])
	if choice == "" {
		return "", errors.New("you did not choose any of the options")
	}
	return choice, nil
}

// IsInteractiveMode determines if we can do choosing with fzf.
func IsInteractiveMode(stdout *os.File) bool {
	value := os.Getenv("uptimectl_IGNORE_FZF")
	return value == "" && isTerminal(stdout) && fzfInstalled()
}

// isTerminal determines if given fd is a TTY.
func isTerminal(fd *os.File) bool {
	return isatty.IsTerminal(fd.Fd())
}

// fzfInstalled determines if fzf(1) is in PATH.
func fzfInstalled() bool {
	v, _ := exec.LookPath("fzf")
	if v != "" {
		return true
	}
	return false
}
