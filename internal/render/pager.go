package render

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/term"
)

func UsePager(output string, format string, noPager bool) bool {
	lines := 1 + strings.Count(output, "\n")
	isTTY := term.IsTerminal(int(os.Stdout.Fd()))
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 0
	}
	return shouldUsePager(lines, height, format, noPager, isTTY)
}

func shouldUsePager(lines int, height int, format string, noPager bool, isTTY bool) bool {
	if format == "json" || noPager {
		return false
	}
	if !isTTY {
		return false
	}
	if height <= 0 {
		return false
	}
	return lines > height
}

func OutputWithPager(output string, format string, noPager bool) error {
	if !UsePager(output, format, noPager) {
		fmt.Fprint(os.Stdout, output)
		return nil
	}

	pager := os.Getenv("PAGER")
	if pager == "" {
		pager = "less"
	}

	cmd := exec.Command(pager)
	cmd.Stdin = strings.NewReader(output)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stdout, output)
		return nil
	}
	return nil
}
