package assembuddy

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func RenderTable(arch string, tableData []Syscall) {
	// TODO: add the ARG0 (x0) stuff to headers
	//
	// HACK: hard coding the header len as 9 for now
	//
	// INFO: I actually need the col stuff because the library needs it
	table := tablewriter.NewWriter(os.Stdout)
	switch arch {
	case "x64":
		table.SetHeader([]string{"NR", "SYSCALL", "RAX", "rdi", "rsi", "rdx", "r10", "r8", "r9"})
		col := tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor}
		table.SetHeaderColor(col, col, col, col, col, col, col, col, col)
		for _, syscall := range tableData {
			table.Append([]string{syscall.Name, syscall.ReturnValue, syscall.Arg0, syscall.Arg1, syscall.Arg2, syscall.Arg3, syscall.Arg4, syscall.Arg5})
		}
	case "x86":
		table.SetHeader([]string{"NR", "SYSCALL", "eax", "ebx", "ecx", "edx", "esi", "edi", "ebp"})
		col := tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor}
		table.SetHeaderColor(col, col, col, col, col, col, col, col, col)
		for _, syscall := range tableData {
			table.Append([]string{syscall.Name, syscall.ReturnValue, syscall.Arg0, syscall.Arg1, syscall.Arg2, syscall.Arg3, syscall.Arg4, syscall.Arg5})
		}
	case "arm64":
		table.SetHeader([]string{"NR", "SYSCALL", "x8", "x0", "x1", "x2", "x3", "x4", "x5"})
		col := tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor}
		table.SetHeaderColor(col, col, col, col, col, col, col, col, col)
		for _, syscall := range tableData {
			table.Append([]string{syscall.Name, syscall.ReturnValue, syscall.Arg0, syscall.Arg1, syscall.Arg2, syscall.Arg3, syscall.Arg4, syscall.Arg5})
		}
	case "arm":
		table.SetHeader([]string{"NR", "SYSCALL", "r7", "r0", "r1", "r2", "r3", "r4", "r5"})
		col := tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor}
		table.SetHeaderColor(col, col, col, col, col, col, col, col, col)
		for _, syscall := range tableData {
			table.Append([]string{syscall.Name, syscall.ReturnValue, syscall.Arg0, syscall.Arg1, syscall.Arg2, syscall.Arg3, syscall.Arg4, syscall.Arg5})
		}

	case "":
		table.SetHeader([]string{"ARCH", "NR", "NAME", "RETURN", "ARG0", "ARG1", "ARG2", "ARG3", "ARG4", "ARG5"})
		col := tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor}
		table.SetHeaderColor(col, col, col, col, col, col, col, col, col, col)

		for _, syscall := range tableData {
			table.Append([]string{syscall.Arch, fmt.Sprint(syscall.Nr), syscall.Name, syscall.ReturnValue, syscall.Arg0, syscall.Arg1, syscall.Arg2, syscall.Arg3, syscall.Arg4, syscall.Arg5})
		}
	}
	table.Render()
}
