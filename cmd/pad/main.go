package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/apiobuild/post-it-pad/cmd/pad/commands"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	commands.Execute()
}
