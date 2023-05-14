package main

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

func parse(extension string, fileData []byte) []Income {
	var incomes []Income
	var err error
	switch extension {
	case ".json":
		err = json.Unmarshal(fileData, &incomes)
	case ".yml":
	case ".yaml":
		err = yaml.Unmarshal(fileData, &incomes)
	}
	if err != nil {
		panic(err)
	}
	return incomes
}
