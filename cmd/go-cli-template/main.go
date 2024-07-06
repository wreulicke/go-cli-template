package main

import (
	"debug/buildinfo"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func mainInternal() error {
	//nolint:wrapcheck
	return NewApp().Execute()
}

func main() {
	if err := mainInternal(); err != nil {
		log.Fatal(err)
	}
}

func NewApp() *cobra.Command {
	c := cobra.Command{
		Use:   "go-cli-template",
		Short: "go-cli-template",
		RunE: func(cmd *cobra.Command, args []string) error {
			// nop
			return nil
		},
	}

	c.AddCommand(
		NewVersionCommand(),
	)
	return &c
}

func NewVersionCommand() *cobra.Command {
	var detail bool
	c := &cobra.Command{
		Use:   "version",
		Short: "show version",
		RunE: func(cmd *cobra.Command, args []string) error {
			w := cmd.OutOrStdout()
			info, err := buildinfo.ReadFile(os.Args[0])
			if err != nil {
				return fmt.Errorf("cannot read buildinfo: %w", err)
			}

			fmt.Fprintf(w, "go version: %s\n", info.GoVersion)
			fmt.Fprintf(w, "path: %s\n", info.Path)
			fmt.Fprintf(w, "mod: %s\n", info.Main.Path)
			fmt.Fprintf(w, "module version: %s\n", info.Main.Version)
			if detail {
				fmt.Fprintln(w)
				fmt.Fprintln(w, info)
			}
			return nil
		},
	}
	c.Flags().BoolVarP(&detail, "detail", "d", false, "show details")
	return c
}
