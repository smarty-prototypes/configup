# configup
--
    import "github.com/smartystreets/configup"

package configup provides a configuration auto-reload mechanism upon receipt of
SIGHUP (or any other signal you specify). Example usage:
github.com/smartystreets/configup/sample

## Usage

#### type DefaultListener

```go
type DefaultListener struct {
}
```


#### func  NewListener

```go
func NewListener(subscriber Subscriber, reader Reader) *DefaultListener
```

#### func (*DefaultListener) Close

```go
func (this *DefaultListener) Close() error
```

#### func (*DefaultListener) Listen

```go
func (this *DefaultListener) Listen()
```

#### type DefaultReader

```go
type DefaultReader struct {
}
```


#### func  NewReader

```go
func NewReader(inner Reader, storage Storage) *DefaultReader
```

#### func (*DefaultReader) Read

```go
func (this *DefaultReader) Read() (interface{}, error)
```

#### type DefaultSubscriber

```go
type DefaultSubscriber struct {
}
```


#### func  NewSubscriber

```go
func NewSubscriber() *DefaultSubscriber
```

#### func (*DefaultSubscriber) Subscribe

```go
func (this *DefaultSubscriber) Subscribe(signals ...os.Signal)
```

#### func (*DefaultSubscriber) Subscription

```go
func (this *DefaultSubscriber) Subscription() chan os.Signal
```

#### func (*DefaultSubscriber) Unsubscribe

```go
func (this *DefaultSubscriber) Unsubscribe()
```

#### type JSONReader

```go
type JSONReader struct {
}
```


#### func  NewJSONReader

```go
func NewJSONReader(path string, instance interface{}) *JSONReader
```

#### func (*JSONReader) Read

```go
func (this *JSONReader) Read() (interface{}, error)
```

#### type Listener

```go
type Listener interface {
	Listen()
	io.Closer
}
```


#### type Reader

```go
type Reader interface {
	Read() (interface{}, error)
}
```


#### type Storage

```go
type Storage interface {
	Store(interface{})
	Load() interface{}
}
```


#### func  FromJSONFile

```go
func FromJSONFile(filename string, instance interface{}) Storage
```
FromJSONFile receives a filename and an instance of the type of struct into
which the contents of the specified file should be unmarshalled. The instance is
only used to derive type information. Rather than fill the provided instance,
the Storage will give back a fresh copy of the unmarshalled data.

#### func  FromReader

```go
func FromReader(reader Reader) Storage
```

#### type Subscriber

```go
type Subscriber interface {
	Subscribe(...os.Signal)
	Unsubscribe()
	Subscription() chan os.Signal
}
```


#### type Wireup

```go
type Wireup struct {
}
```


#### func  New

```go
func New(reader Reader) *Wireup
```

#### func (*Wireup) Initialize

```go
func (this *Wireup) Initialize() (Storage, error)
```

#### func (*Wireup) WatchSignals

```go
func (this *Wireup) WatchSignals(signals ...os.Signal) *Wireup
```
