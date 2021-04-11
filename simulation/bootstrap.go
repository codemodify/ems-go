package simulation

import (
	"net/http"

	"github.com/codemodify/ems-go/helpers"
	httpServer "github.com/codemodify/systemkit-appserver-http"
	logging "github.com/codemodify/systemkit-logging"
	crashproof "github.com/remoteit/systemkit-crashproof"
)

var logID = "simulation"

const listenOn = ":63000"

func RunSimulation(dataFilePath string) {
	// load data file
	dataFile, err := helpers.LoadDataFile(dataFilePath)
	if err != nil {
		logging.Panicf("Can't read data file, %v", err)
		panic("Can't read data file")
	}

	logging.Debugf("%s: %s", logID, helpers.ObjectToString(dataFile, true))

	// run the http server
	crashproof.Go(func() {
		logging.Debugf("%s: listen on %s", logID, listenOn)
		server := httpServer.NewHTTPServer([]httpServer.HTTPHandler{
			httpServer.HTTPHandler{
				Route:   "/SimulationEndpoint",
				Verb:    "GET",
				Handler: simulationEndpointHandler,
			},
		})
		err := server.Run(listenOn, true)
		if err != nil {
			logging.Panic(err)
			panic(err)
		}
	})

	// run control loop
	// FIXME: just started
	o1 := essGetCurrentData(0, dataFile.ESS)
	o2 := meterGetCurrentData(0, dataFile.Meter)
	o3 := targetGetCurrentData(0, dataFile.Target)
	logging.Debugf("%s: o1 - %s", logID, helpers.ObjectToString(o1, true))
	logging.Debugf("%s: o2 - %s", logID, helpers.ObjectToString(o2, true))
	logging.Debugf("%s: o3 - %s", logID, helpers.ObjectToString(o3, true))
}

func simulationEndpointHandler(rw http.ResponseWriter, r *http.Request) {
	// FIXME: have to see what will be here
}
