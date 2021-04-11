package helpers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/codemodify/ems-go/contracts"
)

func LoadDataFile(dataFilePath string) (contracts.DataFile, error) {
	dataFileRaw, err := ioutil.ReadFile(dataFilePath)
	if err != nil {
		return contracts.DataFile{}, err
	}

	var dataFile contracts.DataFile
	err = json.Unmarshal(dataFileRaw, &dataFile)
	if err != nil {
		return contracts.DataFile{}, err
	}

	return dataFile, nil
}

func ObjectToString(o interface{}, indent bool) string {
	var objectAsBytes []byte
	if indent {
		objectAsBytes, _ = json.Marshal(o)
	} else {
		objectAsBytes, _ = json.MarshalIndent(o, "", "\t")
	}

	return string(objectAsBytes)
}
