package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type LaboState struct {
	DatabaseFile string `yaml:"database_file"`
}

func getLaboState() (*LaboState, error) {
	lab, err := getLaboAppDir()
	if err != nil {
		return nil, err
	}

	if _, err = os.Open(lab); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(lab, 0755); err != nil {
				return nil, err
			}
			lDb := []byte(fmt.Sprintf("database_file: %s", path.Join(lab, "labo.db")))
			err = ioutil.WriteFile(path.Join(lab, ".labo"), []byte(lDb), 0644)
			if err != nil {
				return nil, err
			}
		}
	}

	data, err := ioutil.ReadFile(path.Join(lab, ".labo"))
	if err != nil {
		return nil, err
	}

	state := new(LaboState)
	if err = yaml.Unmarshal(data, state); err != nil {
		return nil, err
	}

	return state, nil

}
