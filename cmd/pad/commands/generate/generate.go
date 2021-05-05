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
	serveVar          = "serve"
	serveVarShort     = "s"
	portVar           = "port"
	portVarShort      = "p"
)

var (
	layout    string
	layoutDir string
	destPath  string
	serve     bool
	port      int
)

var Command = &cobra.Command{
	Use:   "generate",
	Short: "Generate templated email html",
	Long:  "Use this command to generate templated for specific or multiple layouts.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		g := generate.NewGenerator(
			defaultStringValueOrNil(layoutDir),
			defaultStringValueOrNil(layout),
			defaultStringValueOrNil(destPath))
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
		"Specify name of the layout to generate templated email for. Generate for all layouts if not provided.")
	viper.BindPFlag(layoutVar, Command.PersistentFlags().Lookup(layoutVar))

	Command.PersistentFlags().StringVarP(&layoutDir, layoutDirVar, layoutDirVarShort, "",
		"Specify layout director.")
	viper.BindPFlag(layoutVar, Command.PersistentFlags().Lookup(layoutVar))

	Command.PersistentFlags().StringVarP(&destPath, destPathVar, destPathVarShort, "",
		"Specify the path to write generated html to.")
	viper.BindPFlag(destPathVar, Command.PersistentFlags().Lookup(destPathVar))

	Command.PersistentFlags().BoolVarP(&serve, serveVar, serveVarShort, false,
		"Serve and preview.")
	viper.BindPFlag(serveVar, Command.PersistentFlags().Lookup(serveVar))

	Command.PersistentFlags().IntVarP(&port, portVar, portVarShort, 8080,
		"Specify the port to serve to.")
	viper.BindPFlag(portVar, Command.PersistentFlags().Lookup(portVar))
}
