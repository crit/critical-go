package file

import (
	"encoding/json"
	"github.com/crit/critical-go/storage"
)

type File struct {
	key   string
	store storage.Storage
}

func New(path string, name string) File {
	return File{
		store: storage.Folder(path),
		key:   name,
	}
}

func (f File) Get(key string) string {
	return f.data()[key]
}

func (f File) Put(key, value string) {
	data := f.data()

	data[key] = value

	f.setData(data)
}

func (f File) data() map[string]string {
	data := map[string]string{}

	raw := f.store.Get(f.key)

	json.Unmarshal(raw, &data)

	return data
}

func (f File) setData(data map[string]string) {
	raw, err := json.Marshal(data)

	if err != nil {
		return
	}

	f.store.Put(f.key, raw)
}
