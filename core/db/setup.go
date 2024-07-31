package db

import (
	"log"

	"github.com/tidwall/buntdb"
)

const LEDGERDB = "./database/ledger.db"
const CLIENTDB = "./database/client.db"

func Init() {
	db1, err := buntdb.Open(LEDGERDB)
	if err != nil {
		log.Printf("Error rendering %s: %s\n",LEDGERDB, err)
		return
	}
	defer db1.Close()
	db1.CreateIndex("contract", "*", buntdb.IndexJSON("contract.id"))

	db2, err := buntdb.Open(CLIENTDB)
	if err != nil {
		log.Printf("Error rendering %s: %s\n",CLIENTDB, err)
		return
	}
	defer db2.Close()
}
