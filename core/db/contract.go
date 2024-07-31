package db

import (
	"cryptocatalysts/core/entities"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tidwall/buntdb"
)

func ConSave(key string, data *entities.Contract) {
	db, err := buntdb.Open(LEDGERDB)
	if err != nil {
		log.Printf("Error rendering chain.db: %s\n", err)
		return
	}
	defer db.Close()

	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error converting data to json: %s\n", err)
		return
	}
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, string(dataJson), nil)
		return err
	})
	if err != nil {
		log.Printf("Save: %s\n", err.Error())
	}
}

func ConUpdate(address string, contract entities.Contract) {
	db, err := buntdb.Open(LEDGERDB)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	dataJson, err := json.Marshal(contract)
	if err != nil {
		log.Println("Error converting data to json")
	}
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(address, string(dataJson), nil)
		return err
	})
	if err != nil {
		log.Println("Error updating conctract")
	}
}

func ConRetrieve(key string) entities.Contract {
	db, err := buntdb.Open(LEDGERDB)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	var contract entities.Contract
	err = db.View(func(tx *buntdb.Tx) error {
		data, err := tx.Get(key)
		if err != nil {
			return nil
		}
		err = json.Unmarshal([]byte(data), &contract)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return nil
		}
		return nil
	})
	if err != nil {
		log.Printf("Retrieve: %s\n", err.Error())
	}
	return contract
}

func ConList() ([]entities.Contract, bool) {
	db, err := buntdb.Open(LEDGERDB)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	contracts := []entities.Contract{}
	err = db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			//fmt.Printf("key: %s, value: %s\n", key, value)
			contract :=  entities.Contract{}
			json.Unmarshal([]byte(value), &contract)
			contracts = append(contracts, contract)
			return true // continue iteration
		})
		return err
	})
	if err != nil {
		return nil, false
	}
	return contracts, true
}
