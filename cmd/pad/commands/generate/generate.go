package generate

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	layoutVar        = "layout"
	layoutVarShort   = "l"
	destPathVar      = "destPath"
	destPathVarShort = "o"
	serveVar         = "serve"
	serveVarShort    = "s"
	portVar          = "port"
	portVarShort     = "p"
)

var (
	layout   string
	destPath string
	serve    bool
	port     int
)

var Command = &cobra.Command{
	Use:   "generate",
	Short: "Generate templated email html",
	Long:  "Use this command to generate templated for specific or multiple layouts.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: update
	},
}

func init() {
	Command.PersistentFlags().StringVarP(&layout, layoutVar, layoutVarShort, "",
		"Specify name of the layout to generate templated email for. Generate for all layouts if not provided.")
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
