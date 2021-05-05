package generate

import (
	"github.com/apiobuild/post-it-pad/pkg/generate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	layoutVar         = "layout"
	layoutVarShort    = "l"
	layoutDirVar      = "layoutDir"
	layoutDirVarShort = "d"
	destPathVar       = "destPath"
	destPathVarShort  = "o"
	argsPathVar       = "args"
	argsPathVarShort  = "a"
	argsVar           = "json"
	argsVarShort      = "j"
	serveVar          = "serve"
	serveVarShort     = "s"
	portVar           = "port"
	portVarShort      = "p"
)

var (
	layout    string
	layoutDir string
	destPath  string
	argsPath  string
	argsJSON  string
	serve     bool
	port      int
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
			defaultStringValueOrNil(argsPath),
			defaultStringValueOrNil(argsJSON),
		)
		g.Generate()
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
		"Path to layout director.")
	viper.BindPFlag(layoutVar, Command.PersistentFlags().Lookup(layoutDirVar))

	Command.PersistentFlags().StringVarP(&destPath, destPathVar, destPathVarShort, "",
		"Path to write generated html.")
	viper.BindPFlag(destPathVar, Command.PersistentFlags().Lookup(destPathVar))

	Command.PersistentFlags().StringVarP(&argsPath, argsPathVar, argsPathVarShort, "",
		"Specify args json file path, default to example args.json if not provided.")
	viper.BindPFlag(argsPathVar, Command.PersistentFlags().Lookup(argsPathVar))

	Command.PersistentFlags().StringVarP(&argsJSON, argsVar, argsVarShort, "",
		"Specify args json directly.")
	viper.BindPFlag(argsVar, Command.PersistentFlags().Lookup(argsPathVar))

	Command.PersistentFlags().BoolVarP(&serve, serveVar, serveVarShort, false,
		"Serve and preview.")
	viper.BindPFlag(serveVar, Command.PersistentFlags().Lookup(serveVar))

	Command.PersistentFlags().IntVarP(&port, portVar, portVarShort, 8080,
		"Port to serve to, default to 8080.")
	viper.BindPFlag(portVar, Command.PersistentFlags().Lookup(portVar))
}
