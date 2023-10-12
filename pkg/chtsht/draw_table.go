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
	for _, item := range items {
		table.Append([]string{fmt.Sprint(item.Nr), item.Name, item.Refs, item.Return, item.Rdi, item.Rsi, item.Rdx, item.R10, item.R8, item.R9})
	}
	table.Render()
}
