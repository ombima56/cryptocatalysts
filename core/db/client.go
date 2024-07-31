package db

import (
	"encoding/json"
	"log"

	"github.com/tidwall/buntdb"

	"cryptocatalysts/core/entities"
)

func CSave(key string, client *entities.Client) {
	db, err := buntdb.Open(CLIENTDB)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	dataJson, err := json.Marshal(client)
	if err != nil {
		log.Printf("Error converting data to json")
	}
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, string(dataJson), nil)
		return err
	})
	if err != nil {
		log.Printf("Save: %s\n", err.Error())
		return
	}
}

func CRetrieve(key string) (*entities.Client, bool) {
	db, err := buntdb.Open(CLIENTDB)
	if err != nil {
		return nil, false
	}
	defer db.Close()

	var client entities.Client
	err = db.View(func(tx *buntdb.Tx) error {
		data, err := tx.Get(key)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(data), &client)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, false
	}
	return &client, true
}
