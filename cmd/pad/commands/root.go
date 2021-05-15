package commands

import (
	"fmt"
	"os"

	"github.com/apiobuild/post-it-pad/cmd/pad/commands/generate"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "pad",
	Short: "Pad is generates templated email html",
	Long:  `Use Post-it Pad to generate templated email html with specific layout or as a development kit to create new email layout. Check out documentation is at https://apiobuild.com/docs/.`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

var version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Pad",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: update
		// fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

func Execute() {
	root.AddCommand(version)
	root.AddCommand(generate.Command)
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
