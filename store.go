package dog

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type DB struct {
	filepath string
	mapping  map[string]interface{}
}

func Open(file string) (*DB, error) {
	db := &DB{
		filepath: file,
		mapping:  make(map[string]interface{}),
	}

	// Pull if file exists
	if _, err := os.Stat(file); err == nil {
		if err3 := db.pull(); err3 != nil {
			log.Printf("Error while starting DB pull: %v", err3)
		}
	} else {
		f, err := os.Create(file)
		defer f.Close()
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func (d *DB) pull() error {
	log.Print("Pulling dog")
	f, err := os.Open(d.filepath)
	defer f.Close()
	if err == nil {
		err = gob.NewDecoder(f).Decode(&d.mapping)
	}
	return err
}

func (d *DB) push() error {
	for _, value := range map[string]interface{}(d.mapping) {
		gob.Register(value)
	}

	log.Print("Flushing dog")
	f, err := os.Create(d.filepath)
	defer f.Close()
	if err == nil {
		err = gob.NewEncoder(f).Encode(d.mapping)
	}
	return err
}

func (d *DB) Get(key string) (interface{}, error) {
	if value, ok := d.mapping[key]; !ok {
		return value, nil
	} else {
		return nil, fmt.Errorf("No such key [%s]", key)
	}
}

func (d *DB) Set(key string, value interface{}) {
	d.mapping[key] = value
}

func (d *DB) Close() error {
	return d.push()
}
