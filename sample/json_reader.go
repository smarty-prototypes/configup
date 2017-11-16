package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type JSONReader struct {
	path string
}

func NewJSONReader(path string) *JSONReader {
	return &JSONReader{path: path}
}

func (this *JSONReader) Read() (interface{}, error) {
	file, err := os.Open(this.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	config := ConfigFile{}
	if err = json.Unmarshal(raw, &config); err != nil {
		return nil, err
	}

	return config, nil
}
