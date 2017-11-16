package confighup

type Watcher struct {
	storage  Storage
	listener Listener
}

func (this *Watcher) Load() interface{} {
	return this.storage.Load()
}

func (this *Watcher) Listen() {
	this.listener.Listen()
}
