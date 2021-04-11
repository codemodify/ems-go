package commands

import (
	"os"
	"path/filepath"

	"github.com/codemodify/ems-go/simulation"
	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
	logging "github.com/codemodify/systemkit-logging"
)

// AppRootCmdFlags - these are the main's app params specified in the CLI
type AppRootCmdFlags struct {
	DataFilePath string `flagName:"data" flagDefault:"data.json" flagDescription:"path to .json data file"`
}

var appRootCmd = &clicmdflags.Command{
	Name:        filepath.Base(os.Args[0]),
	Description: "Runs the simulation in Go",
	Examples: []string{
		filepath.Base(os.Args[0]) + " --data DATA_FILE_PATH",
	},
	Handler: func(command *clicmdflags.Command) {

		// get all the flags from CLI
		flags, ok := command.Flags.(AppRootCmdFlags)
		if !ok {
			logging.Panic("Can't parse args")
			panic("Can't parse args")
		}

		// run with the data file specified
		simulation.RunSimulation(flags.DataFilePath)
	},
	Flags: AppRootCmdFlags{},
}

// Execute - this is a convenience call
func Execute() error {
	return appRootCmd.Execute()
}
