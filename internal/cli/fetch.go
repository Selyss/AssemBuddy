package cli

import (
	"fmt"
	"strings"

	"github.com/Selyss/AssemBuddy/internal/fetch"
	"github.com/spf13/cobra"
)

func newFetchCommand() *cobra.Command {
	var outDir string
	var overwrite bool
	var arches string
	var apiBase string

	cmd := &cobra.Command{
		Use:   "fetch-data",
		Short: "Fetch and generate syscall dataset files",
		RunE: func(cmd *cobra.Command, args []string) error {
			archList := []string{}
			if strings.TrimSpace(arches) != "" {
				for _, part := range strings.Split(arches, ",") {
					if trimmed := strings.TrimSpace(part); trimmed != "" {
						archList = append(archList, trimmed)
					}
				}
			}

			_, err := fetch.GenerateDataset(fetch.Options{
				OutDir:    outDir,
				Overwrite: overwrite,
				Arches:    archList,
				APIBase:   apiBase,
			})
			if err != nil {
				return internalError(err)
			}
			fmt.Fprintln(cmd.OutOrStdout(), "dataset generated")
			return nil
		},
	}

	cmd.Flags().StringVar(&outDir, "out", "data", "Output directory")
	cmd.Flags().BoolVar(&overwrite, "overwrite", false, "Overwrite existing files")
	cmd.Flags().StringVar(&arches, "arches", "", "Comma-separated architectures")
	cmd.Flags().StringVar(&apiBase, "api-base", "https://api.syscall.sh/v1", "API base URL")

	return cmd
}
