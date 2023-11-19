package assembuddy

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func RenderNameTable(tableData []Syscall) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ARCH", "NR", "SYSCALL", "RAX", "rdi", "rsi", "rdx", "r10", "r8", "r9"})
}

func RenderArchTable(tableData []Syscall) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NR", "SYSCALL", "RAX", "rdi", "rsi", "rdx", "r10", "r8", "r9"})
}
