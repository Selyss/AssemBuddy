package assembuddy

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func DisplaySyscall(syscallJSON []byte) {
	var data map[string]interface{}
	if err := json.Unmarshal(syscallJSON, &data); err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ARCH", "NR", "SYSCALL", "RAX", "rdi", "rsi", "rdx", "r10", "r8", "r9"})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	)

	arch := data["arch"]
	// I am truely sorry for the following code. I just want colors.

	if arch == "x64" {
		table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor})
	} else if arch == "x86" {
	} else if arch == "arm64" {
		table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor})
	} else if arch == "arm" {
		table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor})
	}

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
