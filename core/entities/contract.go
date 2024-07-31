package entities

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"cryptocatalysts/core/utils"
)

type Contract struct {
	Address string `json:"address"`
	Duration   int           `json:"duration"`
	Start      time.Time     `json:"start"`
	End        time.Time     `json:"end"`
	Cost       int        `json:"cost"`
	Milestones []*Milestone `json:"milestones"`
	Owner      string        `json:"owner"`
	Awarded    string        `json:"awarded"`
	Oversight  []*Client     `json:"oversight"`
	Parties    []string      `json:"parties"`
	Signatures []string      `json:"signatures"`
}
type Milestone struct {
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Expectations string `json:"expectations"`
	Completed    string `json:"completed"`
}


func NewContract(client, recipient string, amount int, duration int) (*Contract, bool) {
	address, err := utils.GenerateRandomHash(64)
	if err != nil {
		log.Printf("Unable to generate random key")
		return nil, false
	}

	jsonTmpl, err := os.Open("./templates/contract.json")
	if err != nil {
		log.Fatalln("Unable to read json template file")
	}
	var contract Contract
	json.NewEncoder(jsonTmpl).Encode(contract)

	contract.Duration = duration
	contract.Start = time.Now()
	contract.End = utils.Time(duration);
	contract.Cost = amount
	contract.Owner = client
	contract.Awarded = recipient
	contract.Address = address
	
	return &contract, true
}
