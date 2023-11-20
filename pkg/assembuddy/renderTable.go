package assembuddy

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const (
	purple = lipgloss.Color("#EE73F0")
	orange = lipgloss.Color("#FEA96A")
	blue   = lipgloss.Color("#23B0FF")
	red    = lipgloss.Color("#FF6367")
	green  = lipgloss.Color("#78E3A1")
	gray   = lipgloss.Color("#20242c")
	white  = lipgloss.Color("#ffffff")
)

func RenderTable(arch string, tableData []Syscall) {
	// TODO: add the ARG0 (x0) stuff to headers
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		CellStyle = re.NewStyle().Padding(0, 1).Align(lipgloss.Center)
		RowStyle  = CellStyle.Copy().Foreground(white)
	)

	var (
		renderOrange = re.NewStyle().Foreground(orange).Bold(true).Align(lipgloss.Center)
		renderBlue   = re.NewStyle().Foreground(blue).Bold(true).Align(lipgloss.Center)
		renderRed    = re.NewStyle().Foreground(red).Bold(true).Align(lipgloss.Center)
		renderPurple = re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		renderGreen  = re.NewStyle().Foreground(green).Bold(true).Align(lipgloss.Center)
	)

	rows := [][]string{}
	for _, syscall := range tableData {
		rows = append(rows, []string{syscall.Arch, fmt.Sprint(syscall.Nr), syscall.Name, syscall.ReturnValue, syscall.Arg0, syscall.Arg1, syscall.Arg2, syscall.Arg3, syscall.Arg4, syscall.Arg5})
	}

	t := table.New().
		Border(lipgloss.ThickBorder()).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style

			if row == 0 {
				switch arch {
				case "x64":
					return renderOrange
				case "x86":
					return renderBlue
				case "arm64":
					return renderRed
				case "arm":
					return renderPurple
				default:
					return renderGreen
				}
			}

			style = RowStyle
			return style
		}).
		Rows(rows...)
	getHeaders(arch, t)
	fmt.Println(t)
}

func getHeaders(arch string, t *table.Table) {
	switch arch {
	case "x64":
		t.Headers("NR", "SYSCALL", "RAX", "rdi", "rsi", "rdx", "r10", "r8", "r9")
		t.BorderStyle(lipgloss.NewStyle().Foreground(orange))
	case "x86":
		t.Headers("NR", "SYSCALL", "eax", "ebx", "ecx", "edx", "esi", "edi", "ebp")
		t.BorderStyle(lipgloss.NewStyle().Foreground(blue))
	case "arm64":
		t.Headers("NR", "SYSCALL", "x8", "x0", "x1", "x2", "x3", "x4", "x5")
		t.BorderStyle(lipgloss.NewStyle().Foreground(red))
	case "arm":
		t.Headers("NR", "SYSCALL", "r7", "r0", "r1", "r2", "r3", "r4", "r5")
		t.BorderStyle(lipgloss.NewStyle().Foreground(purple))
	case "":
		t.Headers("ARCH", "NR", "NAME", "RETURN", "ARG0", "ARG1", "ARG2", "ARG3", "ARG4", "ARG5")
		t.BorderStyle(lipgloss.NewStyle().Foreground(green))
	}
}
