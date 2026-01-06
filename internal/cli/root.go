package cli

import (
	"fmt"
	"strings"

	"github.com/Selyss/AssemBuddy/internal/store"
	"github.com/spf13/cobra"
)

type Config struct {
	Format      string
	NoPager     bool
	Color       string
	Width       int
	DataVersion bool
}

type App struct {
	Store  *store.Store
	Config Config
}

func NewRootCommand(version string) *cobra.Command {
	app := &App{}

	cmd := &cobra.Command{
		Use:           "assembuddy",
		Short:         "Offline Linux syscall lookup",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if app.Store == nil {
				loaded, err := store.Load()
				if err != nil {
					return internalError(err)
				}
				app.Store = loaded
			}
			if app.Config.DataVersion {
				fmt.Fprintln(cmd.OutOrStdout(), app.Store.Meta.DatasetVersion)
				return ExitError{Code: 0}
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&app.Config.Format, "format", "table", "Output format: table or json")
	cmd.PersistentFlags().BoolVar(&app.Config.NoPager, "no-pager", false, "Disable pager output")
	cmd.PersistentFlags().StringVar(&app.Config.Color, "color", "auto", "Color output: auto, always, never")
	cmd.PersistentFlags().IntVar(&app.Config.Width, "width", 0, "Max cell width for table output")
	cmd.PersistentFlags().BoolVar(&app.Config.DataVersion, "data-version", false, "Print embedded dataset version and exit")

	cmd.AddCommand(newQueryCommand(app))
	cmd.AddCommand(newListCommand(app))
	cmd.AddCommand(newArchesCommand(app))
	cmd.AddCommand(newVersionCommand(app, version))
	cmd.AddCommand(newFetchCommand())

	cmd.SetHelpCommand(&cobra.Command{Hidden: true})

	cmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		return usageError(err)
	})

	cmd.PersistentPreRunE = chainPreRun(cmd.PersistentPreRunE, func(cmd *cobra.Command, args []string) error {
		if err := validateFormat(app.Config.Format); err != nil {
			return usageError(err)
		}
		if err := validateColor(app.Config.Color); err != nil {
			return usageError(err)
		}
		return nil
	})

	return cmd
}

func chainPreRun(existing func(*cobra.Command, []string) error, next func(*cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if existing != nil {
			if err := existing(cmd, args); err != nil {
				return err
			}
		}
		return next(cmd, args)
	}
}

func validateFormat(format string) error {
	format = strings.ToLower(format)
	if format != "table" && format != "json" {
		return fmt.Errorf("invalid format: %s", format)
	}
	return nil
}

func validateColor(color string) error {
	color = strings.ToLower(color)
	switch color {
	case "auto", "always", "never":
		return nil
	default:
		return fmt.Errorf("invalid color: %s", color)
	}
}
