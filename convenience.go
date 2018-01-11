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
func FromJSONFile(filename string, instance interface{}, signals ...os.Signal) *DefaultListener {
	reader := NewJSONReader(filename, instance)
	return FromReader(reader, signals...)
}

func FromReader(reader Reader, signals ...os.Signal) *DefaultListener {
	if listener, err := New(reader, WithSignal(signals...)); err != nil {
		log.Fatalln("[ERROR] Unable to read configuration:", err)
		return nil
	} else {
		go listener.Listen()
		return listener
	}
}
