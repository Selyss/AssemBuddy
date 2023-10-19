package chtsht

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func DisplaySyscall(syscallJSON []byte) error {
	jsonData := string(syscallJSON)
	var syscalls []Syscall
	fmt.Println(jsonData)
	err := json.Unmarshal([]byte(jsonData), &syscalls)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NR", "SYSCALL", "RAX", "rdi", "rsi", "rdx", "r10", "r8", "r9"})
	for _, syscall := range syscalls { // HACK: needed?
		table.Append([]string{syscall.Arch, string(syscall.Nr), syscall.Name, syscall.Refs, syscall.Return, syscall.Arg0, syscall.Arg1, syscall.Arg2, syscall.Arg3, syscall.Arg4, syscall.Arg5})
	}
	table.Render()
	return nil
}
