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
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		CellStyle = re.NewStyle().Align(lipgloss.Center).Padding(0, 1)
		RowStyle  = CellStyle.Copy().Foreground(white).MaxWidth(12).PaddingBottom(1)
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
		if arch == "" {
			rows = append(rows, []string{syscall.Arch, fmt.Sprint(syscall.Nr), syscall.Name, syscall.ReturnValue, syscall.Arg0, syscall.Arg1, syscall.Arg2, syscall.Arg3, syscall.Arg4, syscall.Arg5})
		} else {
			rows = append(rows, []string{fmt.Sprint(syscall.Nr), syscall.Name, syscall.ReturnValue, syscall.Arg0, syscall.Arg1, syscall.Arg2, syscall.Arg3, syscall.Arg4, syscall.Arg5})
		}
	}

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderRow(true).
		StyleFunc(func(row, _ int) lipgloss.Style {
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
		t.Headers("NR", "SYSCALL", "RAX", "ARG0 (rdi)", "ARG1 (rsi)", "ARG2 (rdx)", "ARG3 (r10)", "ARG4 (r8)", "ARG5 (r9)")
		t.BorderStyle(lipgloss.NewStyle().Foreground(orange))
	case "x86":
		t.Headers("NR", "SYSCALL", "eax", "ARG0 (ebx)", "ARG1 (ecx)", "ARG2 (edx)", "ARG3 (esi)", "ARG4 (edi)", "ARG5 (ebp)")
		t.BorderStyle(lipgloss.NewStyle().Foreground(blue))
	case "arm64":
		t.Headers("NR", "SYSCALL", "x8", "ARG0 (x0)", "ARG1 (x1)", "ARG2 (x2)", "ARG3 (x3)", "ARG4 (x4)", "ARG5 (x5)")
		t.BorderStyle(lipgloss.NewStyle().Foreground(red))
	case "arm":
		t.Headers("NR", "SYSCALL", "r7", "ARG0 (r0)", "ARG1 (r1)", "ARG2 (r2)", "ARG3 (r3)", "ARG4 (r4)", "ARG5 (r5)")
		t.BorderStyle(lipgloss.NewStyle().Foreground(purple))
	case "":
		t.Headers("ARCH", "NR", "NAME", "RETURN", "ARG0", "ARG1", "ARG2", "ARG3", "ARG4", "ARG5")
		t.BorderStyle(lipgloss.NewStyle().Foreground(green))
	}
}
