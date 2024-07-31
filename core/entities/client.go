package entities

import (
	"cryptocatalysts/core/utils"
)

type Client struct {
	Cid      string `json:"cid"`
	Cname    string `json:"cname"`
	Email    string `json:"email"`
	Location string `json:"location"`
	Account  string `json:"account"`
	Bio      string `json:"bio"`
}

func NewClient(cname, email, location, bio, account string) (*Client, bool) {
	id, err := utils.GenerateRandomHash(32)
	if err != nil {
		return nil, false
	}
	client := &Client{
		Cid:      id,
		Cname:    cname,
		Email:    email,
		Location: location,
		Bio:      bio,
		Account:  account,
	}
	return client, true
}
