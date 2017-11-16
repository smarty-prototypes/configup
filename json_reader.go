package confighup

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
)

type JSONReader struct {
	path       string
	configType reflect.Type
}

func NewJSONReader(path string, instance interface{}) *JSONReader {
	return &JSONReader{path: path, configType: reflect.TypeOf(instance)}
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

	config := reflect.New(this.configType.Elem()).Interface()
	if err = json.Unmarshal(raw, config); err != nil {
		return nil, err
	}

	return config, nil
}
