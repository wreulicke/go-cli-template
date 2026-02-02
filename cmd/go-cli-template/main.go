package main

// patchr:cut-start
//go:generate go run github.com/wreulicke/patchr/cmd/patchr@v0.0.4 ../../
// patchr:cut-end
import (
	"debug/buildinfo"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func mainInternal() error {
	//nolint:wrapcheck
	return NewApp().Execute()
}

func main() {
	err := mainInternal()
	if err != nil {
		log.Fatal(err)
	}
}

func NewApp() *cobra.Command {
	c := cobra.Command{
		// patchr:template-start
		// Use:   "{{ input "name" }}",
		// Short: "{{ input "name" }}",
		// patchr:template-end
		// patchr:cut-start
		Use:   "go-cli-template",
		Short: "go-cli-template",
		// patchr:cut-end
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
			return version(cmd.OutOrStdout(), detail)
		},
	}
	c.Flags().BoolVarP(&detail, "detail", "d", false, "show details")

	return c
}

func version(w io.Writer, detail bool) error {
	path, err := os.Executable()
	if err != nil {
		return fmt.Errorf("cannot get executable path: %w", err)
	}

	info, err := buildinfo.ReadFile(path)
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
}
