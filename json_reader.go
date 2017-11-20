package configup

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

	// TODO: use this if the instance interface contains a pointer
	// otherwise, use another algorithm to create an instance of a struct
	item := this.instance()
	if err = json.Unmarshal(raw, item); err != nil {
		return nil, err
	}

	return item, nil
}

func (this *JSONReader) instance() interface{} {
	return reflect.New(this.configType.Elem()).Interface()
}
