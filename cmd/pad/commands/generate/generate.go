package generate

import (
	"github.com/apiobuild/post-it-pad/pkg/generate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	layoutVar         = "layout"
	layoutVarShort    = "n"
	layoutDirVar      = "layoutDir"
	layoutDirVarShort = "l"
	destPathVar       = "destPath"
	destPathVarShort  = "o"
	destDirVar        = "destDir"
	destDirVarShort   = "d"
	argsPathVar       = "args"
	argsPathVarShort  = "a"
	argsVar           = "json"
	argsVarShort      = "j"
)

var (
	layout    string
	layoutDir string
	destPath  string
	destDir   string
	argsPath  string
	argsJSON  string
)

// Command defines the generate command
var Command = &cobra.Command{
	Use:   "generate",
	Short: "Generate templated email html",
	Long:  "Use this command to generate templated for specific or multiple layouts.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		g := generate.NewGenerator(
			defaultStringValueOrNil(layoutDir),
			defaultStringValueOrNil(layout),
			defaultStringValueOrNil(destPath),
			defaultStringValueOrNil(destDir),
			defaultStringValueOrNil(argsPath),
			defaultStringValueOrNil(argsJSON),
		)
		err := g.Generate()
		if err != nil {
			panic(err)
		}
	},
}

func defaultStringValueOrNil(val string) *string {
	if val == "" {
		return nil
	}
	return &val
}

func init() {
	Command.PersistentFlags().StringVarP(&layout, layoutVar, layoutVarShort, "",
		"Name of the layout to generate templated email for. Generate for all layouts if not provided.")
	viper.BindPFlag(layoutVar, Command.PersistentFlags().Lookup(layoutVar))

	Command.PersistentFlags().StringVarP(&layoutDir, layoutDirVar, layoutDirVarShort, "",
		"Path to layout directory.")
	viper.BindPFlag(layoutVar, Command.PersistentFlags().Lookup(layoutDirVar))

	Command.PersistentFlags().StringVarP(&destPath, destPathVar, destPathVarShort, "",
		"Path to write generated html. Used if generate for specific layout.")
	viper.BindPFlag(destPathVar, Command.PersistentFlags().Lookup(destPathVar))

	Command.PersistentFlags().StringVarP(&destDir, destDirVar, destDirVarShort, "",
		"Dir to write generated html. Used if generate for all layouts.")
	viper.BindPFlag(destDirVar, Command.PersistentFlags().Lookup(destDirVar))

	Command.PersistentFlags().StringVarP(&argsPath, argsPathVar, argsPathVarShort, "",
		"Specify args json file path, default to example args.json if not provided.")
	viper.BindPFlag(argsPathVar, Command.PersistentFlags().Lookup(argsPathVar))

	Command.PersistentFlags().StringVarP(&argsJSON, argsVar, argsVarShort, "",
		"Specify args json directly.")
	viper.BindPFlag(argsVar, Command.PersistentFlags().Lookup(argsPathVar))
}
