package main

import (
	"fmt"
	"os"

	callstack "github.com/codemodify/systemkit-callstack"
	crashproof "github.com/codemodify/systemkit-crashproof"

	logging "github.com/codemodify/systemkit-logging"
	loggingFormat "github.com/codemodify/systemkit-logging-formatters-timerfc3339nano"
	loggingMixMulti "github.com/codemodify/systemkit-logging-mixers-multi"
	loggingPersistConsole "github.com/codemodify/systemkit-logging-persisters-console"
	loggingPersistFile "github.com/codemodify/systemkit-logging-persisters-rollingfile"

	"github.com/codemodify/ems-go/commands"
)

func main() {

	// 1. setup crash handler for any nested concurrent code
	crashproof.ConcurrentCodeCrashCatcher = reportCrash
	crashproof.RunAppAndCatchCrashes(func() {
		// 2. setup logging
		var loggers = []logging.CoreLogger{
			loggingFormat.NewTimeRFC3339NanoFormatter(),
			loggingMixMulti.NewMultiLogger([]logging.CoreLogger{
				loggingPersistConsole.NewConsoleLogger(),
				loggingPersistFile.NewFileLoggerWithDefaultRotationDefaultName(),
			}),
		}

		logging.SetLogger(logging.NewPipe(loggers)).KeepOnlyLogs(logging.TypeDebug)

		// 3. run the app
		runApp()
	})
}

func reportCrash(err interface{}, callStack []callstack.Frame) {
	fmt.Fprintf(os.Stderr, "\n\nCRASH: %s\n\nstack: %v\n\n", err, callStack)
}

func runApp() {
	if err := commands.Execute(); err != nil {
		logging.Panic(err)
		panic("Can't read data file")
	}
}
