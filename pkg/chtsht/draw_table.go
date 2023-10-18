package chtsht

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// Define a struct to represent your JSON data structure.
type Item struct {
	Arch string `json:"arch"`
	Nr   int    `json:"nr"`
	Name string `json:"name"`
	Refs string `json:"refs"`

	Return string `json:"return"`
	Rdi    string `json:"arg0"`
	Rsi    string `json:"arg1"`
	Rdx    string `json:"arg2"`
	R10    string `json:"arg3"`
	R8     string `json:"arg4"`
	R9     string `json:"arg5"`
}

func DrawRawTable(dataJSON string) {
	var jsonData []map[string]interface{}
	err := json.Unmarshal([]byte(dataJSON), &jsonData)
	if err != nil {

		fmt.Println("Error parsing JSON:", err)
	}

	var data [][]string
	for _, item := range jsonData {
		dataRow := []string{
			fmt.Sprintf("%v", item["nr"]),
			fmt.Sprintf("%v", item["name"]),
			fmt.Sprintf("%v", item["return"]),
			fmt.Sprintf("%v", item["arg0"]),
			fmt.Sprintf("%v", item["arg1"]),
			fmt.Sprintf("%v", item["arg2"]),
			fmt.Sprintf("%v", item["arg3"]),
			fmt.Sprintf("%v", item["arg4"]),
			fmt.Sprintf("%v", item["arg5"]),
		}
		data = append(data, dataRow)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NR", "SYSCALL", "RAX", "rdi", "rsi", "rdx", "r10", "r8", "r9"})
	table.SetAutoWrapText(true)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(0)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetReflowDuringAutoWrap(false)
	table.SetBorder(false)

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor})

	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func DrawTable(jsonData string) {
	// Parse JSON data into a slice of struct.
	var items []Item
	err := json.Unmarshal([]byte(jsonData), &items)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Create a new table.
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NR", "SYSCALL NAME", "references", "RAX", "ARG0 (rdi)", "ARG1 (rsi)", "ARG2 (rdx)", "ARG3 (r10)", "ARG4 (r8)", "ARG5 (r9)"})

	table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	table.ClearFooter()
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetCenterSeparator("│")
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetNoWhiteSpace(true)

	// Render the table.

	// Iterate through the data and add rows to the table.
	table.Render()
}
