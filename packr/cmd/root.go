package cmd

import (
	"fmt"
	"os"

	"github.com/gobuffalo/packr/costello/parser"
	"github.com/gobuffalo/packr/costello/store"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "packr",
	Short: "A brief description of your application",
	RunE: func(cmd *cobra.Command, args []string) error {
		pwd, err := os.Getwd()
		if err != nil {
			return errors.WithStack(err)
		}
		roots := append(args, pwd)
		for _, r := range roots {
			store.Clean(r)
		}
		p, err := parser.NewFromRoots(roots)
		if err != nil {
			return errors.WithStack(err)
		}
		boxes, err := p.Run()
		if err != nil {
			return errors.WithStack(err)
		}

		// reduce boxes - remove ones we don't want
		// MB: current assumption is we want all these
		// boxes, just adding a comment suggesting they're
		// might be a reason to exclude some

		fmt.Printf("Found %d boxes\n", len(boxes))

		// "pack" boxes
		d := store.NewDisk("", "")
		for _, b := range boxes {
			if err := d.Pack(b); err != nil {
				return errors.WithStack(err)
			}
		}
		return d.Close()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
}