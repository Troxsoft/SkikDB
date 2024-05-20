package pkg

import (
	"encoding/json"
	"net"
	"os"
)

type Config struct {
	Name string `json:"name"`
}
type DBStorage struct {
	Datad   map[string]any `json:"data"`
	Configd Config         `json:"conf"`
}
type DB struct {
	Storage DBStorage `json:"DB"`
	Server  net.Listener
}

func NewDB(filepath string) (*DB, error) {
	d, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var storage DBStorage
	err = json.Unmarshal(d, &storage)
	if err != nil {
		return nil, err
	}
	return &DB{
		Storage: storage,
	}, nil
}
func (db *DB) Save() error {
	err := os.WriteFile(db.Storage.Configd.Name+".skl", []byte(jsonToStr(db.Storage)), 0644)
	if err != nil {
		return err
	}
	return nil
}
func CreateDB(name string) (*DB, error) {
	b, err := json.Marshal(DBStorage{
		Datad: make(map[string]any),
		Configd: Config{
			Name: name,
		},
	})
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(name+".skl", b, 0644)
	if err != nil {
		return nil, err
	}
	//os.WriteFile(name + ".skl",0777)
	return NewDB(name + ".skl")
}
