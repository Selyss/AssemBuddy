package chtsht

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func DisplaySyscall(syscallJSON []byte) {
	var data map[string]interface{}
	if err := json.Unmarshal(syscallJSON, &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ARCH", "NR", "SYSCALL", "RAX", "rdi", "rsi", "rdx", "r10", "r8", "r9"})

	// Extract and add data to the table
	table.Append([]string{
		fmt.Sprintf("%v", data["arch"]),
		fmt.Sprintf("%v", data["nr"]),
		fmt.Sprintf("%v", data["name"]),
		fmt.Sprintf("%v", data["return"]),
		fmt.Sprintf("%v", data["arg0"]),
		fmt.Sprintf("%v", data["arg1"]),

		fmt.Sprintf("%v", data["arg2"]),
		fmt.Sprintf("%v", data["arg3"]),
		fmt.Sprintf("%v", data["arg4"]),
		fmt.Sprintf("%v", data["arg5"]),
	})

	table.Render()
}
