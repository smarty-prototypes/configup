package configup

import (
	"log"
	"os"
)

// FromJSONFile receives a filename and an instance of the type of struct
// into which the contents of the specified file should be unmarshalled.
// The instance is only used to derive type information. Rather than fill
// the provided instance, the Storage will give back a fresh copy of the
// unmarshalled data.
func FromJSONFile(filename string, instance interface{}, signals ...os.Signal) Storage {
	reader := NewJSONReader(filename, instance)
	return FromReader(reader, signals...)
}

func FromReader(reader Reader, signals ...os.Signal) Storage {
	wireup := New(reader).WithSignal(signals...)
	if storage, err := wireup.Initialize(); err != nil {
		log.Fatalln("[ERROR] Unable to read configuration:", err)
		return nil
	} else {
		return storage
	}
}
